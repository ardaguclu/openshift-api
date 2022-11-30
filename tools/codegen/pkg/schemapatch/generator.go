package schemapatch

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/openshift/api/tools/codegen/pkg/generation"
	"github.com/openshift/api/tools/codegen/pkg/utils"
	kyaml "sigs.k8s.io/yaml"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

// Options contains the configuration required for the schemapatch generator.
type Options struct {
	// ControllerGen is the path to a controller-gen binary to use for the generation.
	// When omitted, we will use the generator directly from the code.
	ControllerGen string

	// Disabled indicates whether the schemapatch generator is disabled or not.
	// This default to false as the schemapatch generator is enabled by default.
	Disabled bool

	// RequiredFeatureSets is used to filter the feature set manifests that
	// should be generated.
	// When omitted, any manifest with a feature set annotation will be ignored.
	RequiredFeatureSets []sets.String

	// Verify determines whether the generator should verify the content instead
	// of updating the generated file.
	Verify bool
}

// generator implements the generation.Generator interface.
// It is designed to generate schemapatch updates for a particular API group.
type generator struct {
	controllerGen       string
	disabled            bool
	requiredFeatureSets []sets.String
	verify              bool
}

// NewGenerator builds a new schemapatch generator.
func NewGenerator(opts Options) generation.Generator {
	return &generator{
		controllerGen:       opts.ControllerGen,
		disabled:            opts.Disabled,
		requiredFeatureSets: opts.RequiredFeatureSets,
		verify:              opts.Verify,
	}
}

// ApplyConfig creates returns a new generator based on the configuration passed.
// If the schemapatch configuration is empty, the existing generation is returned.
func (g *generator) ApplyConfig(config *generation.Config) generation.Generator {
	if config == nil || config.SchemaPatch == nil {
		return g
	}

	featureSets := []sets.String{}
	for _, featureSet := range config.SchemaPatch.RequiredFeatureSets {
		featureSets = append(featureSets, sets.NewString(strings.Split(featureSet, ",")...))
	}

	return NewGenerator(Options{
		ControllerGen:       g.controllerGen,
		Disabled:            config.SchemaPatch.Disabled,
		RequiredFeatureSets: featureSets,
		Verify:              g.verify,
	})
}

// Name returns the name of the generator.
func (g *generator) Name() string {
	return "schemapatch"
}

// GenGroup runs the schemapatch generator against the given group context.
func (g *generator) GenGroup(groupCtx generation.APIGroupContext) error {
	if g.disabled {
		klog.V(2).Infof("Skipping API schema generation for %s", groupCtx.Name)
		return nil
	}

	versionPaths := allVersionPaths(groupCtx.Versions)

	for _, version := range groupCtx.Versions {
		versionRequired, err := shouldProcessGroupVersion(version, g.requiredFeatureSets)
		if err != nil {
			return fmt.Errorf("could not determine if version %s is required: %w", version.Name, err)
		}

		if !versionRequired {
			continue
		}

		action := "Generating"
		if g.verify {
			action = "Verifying"
		}

		klog.V(1).Infof("%s API schema for for %s/%s", action, groupCtx.Name, version.Name)

		if err := g.genGroupVersion(groupCtx.Name, version, versionPaths); err != nil {
			return fmt.Errorf("could not run schemapatch generator for group/version %s/%s: %w", groupCtx.Name, version.Name, err)
		}
	}

	return nil
}

// genGroupVersion runs the schemapatch generator against a particular version of the API group.
func (g *generator) genGroupVersion(group string, version generation.APIVersionContext, versionPaths []string) error {
	generationContexts, err := loadSchemaPatchGenerationContextsForVersion(version, g.requiredFeatureSets)
	if err != nil {
		return fmt.Errorf("could not load generation contexts: %w", err)
	}

	for _, gc := range generationContexts {
		buf := bytes.NewBuffer(nil)

		if err := executeSchemaPatchForManifest(gc, buf, versionPaths, g.controllerGen); err != nil {
			return fmt.Errorf("could not execute schemapatch for manifest %s: %w", gc.manifestPath, err)
		}

		if err := executeYAMLPatchForManifest(gc, buf); err != nil {
			return fmt.Errorf("could not execute yaml patch for manifest %s: %w", gc.manifestPath, err)
		}

		manifestData, err := formatData(buf.Bytes())
		if err != nil {
			return fmt.Errorf("could not format data for manifest %s: %w", gc.manifestPath, err)
		}

		if g.verify {
			if !bytes.Equal(manifestData, gc.manifestData) {
				diff := utils.Diff(manifestData, gc.manifestData, gc.manifestPath)

				return fmt.Errorf("API schema for %s is out of date, please regenerate the API schema:\n%s", gc.manifestPath, diff)
			}

			return nil
		}

		if err := os.WriteFile(gc.manifestPath, manifestData, gc.manifestFileMode); err != nil {
			return fmt.Errorf("could not write manifest %s: %w", gc.manifestPath, err)
		}
	}

	return nil
}

// allVersionPaths creates a list of all version paths for the group.
func allVersionPaths(versions []generation.APIVersionContext) []string {
	out := []string{}

	for _, version := range versions {
		out = append(out, version.Path)
	}

	return out
}

// schemaPatchGenerationContext contains the context required to generate a schemapatch
// for a particular manifest.
type schemaPatchGenerationContext struct {
	manifestPath        string
	manifestFileMode    fs.FileMode
	manifestData        []byte
	patchPath           string
	requiredFeatureSets sets.String
}

// loadSchemaPatchGenerationContextsForVersion loads the generation contexts for all the manifests
// within a particular API group version.
// It finds all CRD manifests, their corresponding YAML patch manifest if available and the expected
// feature sets for the manifest.
func loadSchemaPatchGenerationContextsForVersion(version generation.APIVersionContext, requiredFeatureSets []sets.String) ([]schemaPatchGenerationContext, error) {
	errs := []error{}

	dirEntries, err := os.ReadDir(version.Path)
	if err != nil {
		return nil, fmt.Errorf("could not read file info for directory %s: %v", version.Path, err)
	}

	generationContexts := []schemaPatchGenerationContext{}

	for _, fileInfo := range dirEntries {
		// Ignore any file that isn't a yaml file.
		if fileInfo.IsDir() || filepath.Ext(fileInfo.Name()) != ".yaml" {
			continue
		}

		path := filepath.Join(version.Path, fileInfo.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			errs = append(errs, fmt.Errorf("could not read file %s: %v", path, err))
			continue
		}

		manifestInfo, err := os.Stat(path)
		if err != nil {
			errs = append(errs, fmt.Errorf("could not stat file %s: %v", path, err))
			continue
		}

		partialObject := &metav1.PartialObjectMetadata{}
		if err := kyaml.Unmarshal(data, partialObject); err != nil {
			errs = append(errs, fmt.Errorf("could not unmarshal YAML for type meta inspection: %v", err))
			continue
		}

		// Ignore any file that doesn't have a kind of CustomResourceDefinition or does not have the correct feature set annotation.
		if !isCustomResourceDefinition(partialObject) || !hasRequiredFeatureSet(partialObject, requiredFeatureSets) {
			continue
		}

		// The file is a CRD and has the correct feature set, build out the context.

		// Work out if there is a patch file for the CRD.
		patchPath := filepath.Join(version.Path, fmt.Sprintf("%s-patch", fileInfo.Name()))
		if _, err := os.Stat(patchPath); err != nil && os.IsNotExist(err) {
			// The patch file doesn't exist, clear the path.
			patchPath = ""
		} else if err != nil {
			errs = append(errs, fmt.Errorf("could not stat patch file %s: %w", patchPath, err))
			continue
		}

		generationContexts = append(generationContexts, schemaPatchGenerationContext{
			manifestPath:        filepath.Join(version.Path, fileInfo.Name()),
			manifestFileMode:    manifestInfo.Mode(),
			manifestData:        data,
			patchPath:           patchPath,
			requiredFeatureSets: getObjectFeatureSets(partialObject),
		})
	}

	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	return generationContexts, nil
}

// isCustomResourceDefinition returns true if the object is a CustomResourceDefinition.
// This is determined by the object having a Kind of CustomResourceDefinition and the
// correct APIVersion.
func isCustomResourceDefinition(partialObject *metav1.PartialObjectMetadata) bool {
	return partialObject.APIVersion == apiextensionsv1.SchemeGroupVersion.String() && partialObject.Kind == "CustomResourceDefinition"
}

// hasRequiredFeatureSet returns true if the object has the desired required feature set.
func hasRequiredFeatureSet(partialObject *metav1.PartialObjectMetadata, requiredFeatureSets []sets.String) bool {
	// Try an empty set in case no features were configured.
	// If this returns true then the object should be handled even if no
	// other requiredFeatureSets match.
	shouldHandle := mayHandleObject(partialObject, sets.NewString())

	for _, requiredFeatureSet := range requiredFeatureSets {
		if mayHandleObject(partialObject, requiredFeatureSet) {
			shouldHandle = true
			break
		}
	}

	return shouldHandle
}
