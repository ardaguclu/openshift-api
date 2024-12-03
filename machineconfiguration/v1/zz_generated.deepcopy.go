//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	configv1 "github.com/openshift/api/config/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildInputs) DeepCopyInto(out *BuildInputs) {
	*out = *in
	if in.BaseImagePullSecret != nil {
		in, out := &in.BaseImagePullSecret, &out.BaseImagePullSecret
		*out = new(ImageSecretObjectReference)
		**out = **in
	}
	out.ImageBuilder = in.ImageBuilder
	out.RenderedImagePushSecret = in.RenderedImagePushSecret
	if in.Containerfile != nil {
		in, out := &in.Containerfile, &out.Containerfile
		*out = make([]MachineOSContainerfile, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildInputs.
func (in *BuildInputs) DeepCopy() *BuildInputs {
	if in == nil {
		return nil
	}
	out := new(BuildInputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildOutputs) DeepCopyInto(out *BuildOutputs) {
	*out = *in
	if in.CurrentImagePullSecret != nil {
		in, out := &in.CurrentImagePullSecret, &out.CurrentImagePullSecret
		*out = new(ImageSecretObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildOutputs.
func (in *BuildOutputs) DeepCopy() *BuildOutputs {
	if in == nil {
		return nil
	}
	out := new(BuildOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertExpiry) DeepCopyInto(out *CertExpiry) {
	*out = *in
	if in.Expiry != nil {
		in, out := &in.Expiry, &out.Expiry
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertExpiry.
func (in *CertExpiry) DeepCopy() *CertExpiry {
	if in == nil {
		return nil
	}
	out := new(CertExpiry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeConfig) DeepCopyInto(out *ContainerRuntimeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeConfig.
func (in *ContainerRuntimeConfig) DeepCopy() *ContainerRuntimeConfig {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerRuntimeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeConfigCondition) DeepCopyInto(out *ContainerRuntimeConfigCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeConfigCondition.
func (in *ContainerRuntimeConfigCondition) DeepCopy() *ContainerRuntimeConfigCondition {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfigCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeConfigList) DeepCopyInto(out *ContainerRuntimeConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContainerRuntimeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeConfigList.
func (in *ContainerRuntimeConfigList) DeepCopy() *ContainerRuntimeConfigList {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerRuntimeConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeConfigSpec) DeepCopyInto(out *ContainerRuntimeConfigSpec) {
	*out = *in
	if in.MachineConfigPoolSelector != nil {
		in, out := &in.MachineConfigPoolSelector, &out.MachineConfigPoolSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerRuntimeConfig != nil {
		in, out := &in.ContainerRuntimeConfig, &out.ContainerRuntimeConfig
		*out = new(ContainerRuntimeConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeConfigSpec.
func (in *ContainerRuntimeConfigSpec) DeepCopy() *ContainerRuntimeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeConfigStatus) DeepCopyInto(out *ContainerRuntimeConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ContainerRuntimeConfigCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeConfigStatus.
func (in *ContainerRuntimeConfigStatus) DeepCopy() *ContainerRuntimeConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeConfiguration) DeepCopyInto(out *ContainerRuntimeConfiguration) {
	*out = *in
	if in.PidsLimit != nil {
		in, out := &in.PidsLimit, &out.PidsLimit
		*out = new(int64)
		**out = **in
	}
	if in.LogSizeMax != nil {
		in, out := &in.LogSizeMax, &out.LogSizeMax
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.OverlaySize != nil {
		in, out := &in.OverlaySize, &out.OverlaySize
		x := (*in).DeepCopy()
		*out = &x
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeConfiguration.
func (in *ContainerRuntimeConfiguration) DeepCopy() *ContainerRuntimeConfiguration {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerCertificate) DeepCopyInto(out *ControllerCertificate) {
	*out = *in
	if in.NotBefore != nil {
		in, out := &in.NotBefore, &out.NotBefore
		*out = (*in).DeepCopy()
	}
	if in.NotAfter != nil {
		in, out := &in.NotAfter, &out.NotAfter
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerCertificate.
func (in *ControllerCertificate) DeepCopy() *ControllerCertificate {
	if in == nil {
		return nil
	}
	out := new(ControllerCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfig) DeepCopyInto(out *ControllerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfig.
func (in *ControllerConfig) DeepCopy() *ControllerConfig {
	if in == nil {
		return nil
	}
	out := new(ControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfigList) DeepCopyInto(out *ControllerConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfigList.
func (in *ControllerConfigList) DeepCopy() *ControllerConfigList {
	if in == nil {
		return nil
	}
	out := new(ControllerConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfigSpec) DeepCopyInto(out *ControllerConfigSpec) {
	*out = *in
	if in.KubeAPIServerServingCAData != nil {
		in, out := &in.KubeAPIServerServingCAData, &out.KubeAPIServerServingCAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.RootCAData != nil {
		in, out := &in.RootCAData, &out.RootCAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CloudProviderCAData != nil {
		in, out := &in.CloudProviderCAData, &out.CloudProviderCAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalTrustBundle != nil {
		in, out := &in.AdditionalTrustBundle, &out.AdditionalTrustBundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.ImageRegistryBundleUserData != nil {
		in, out := &in.ImageRegistryBundleUserData, &out.ImageRegistryBundleUserData
		*out = make([]ImageRegistryBundle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImageRegistryBundleData != nil {
		in, out := &in.ImageRegistryBundleData, &out.ImageRegistryBundleData
		*out = make([]ImageRegistryBundle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.InternalRegistryPullSecret != nil {
		in, out := &in.InternalRegistryPullSecret, &out.InternalRegistryPullSecret
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(configv1.ProxyStatus)
		**out = **in
	}
	if in.Infra != nil {
		in, out := &in.Infra, &out.Infra
		*out = new(configv1.Infrastructure)
		(*in).DeepCopyInto(*out)
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = new(configv1.DNS)
		(*in).DeepCopyInto(*out)
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(NetworkInfo)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfigSpec.
func (in *ControllerConfigSpec) DeepCopy() *ControllerConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfigStatus) DeepCopyInto(out *ControllerConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ControllerConfigStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ControllerCertificates != nil {
		in, out := &in.ControllerCertificates, &out.ControllerCertificates
		*out = make([]ControllerCertificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfigStatus.
func (in *ControllerConfigStatus) DeepCopy() *ControllerConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ControllerConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfigStatusCondition) DeepCopyInto(out *ControllerConfigStatusCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfigStatusCondition.
func (in *ControllerConfigStatusCondition) DeepCopy() *ControllerConfigStatusCondition {
	if in == nil {
		return nil
	}
	out := new(ControllerConfigStatusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistryBundle) DeepCopyInto(out *ImageRegistryBundle) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryBundle.
func (in *ImageRegistryBundle) DeepCopy() *ImageRegistryBundle {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryBundle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSecretObjectReference) DeepCopyInto(out *ImageSecretObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSecretObjectReference.
func (in *ImageSecretObjectReference) DeepCopy() *ImageSecretObjectReference {
	if in == nil {
		return nil
	}
	out := new(ImageSecretObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfig) DeepCopyInto(out *KubeletConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfig.
func (in *KubeletConfig) DeepCopy() *KubeletConfig {
	if in == nil {
		return nil
	}
	out := new(KubeletConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeletConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfigCondition) DeepCopyInto(out *KubeletConfigCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfigCondition.
func (in *KubeletConfigCondition) DeepCopy() *KubeletConfigCondition {
	if in == nil {
		return nil
	}
	out := new(KubeletConfigCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfigList) DeepCopyInto(out *KubeletConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeletConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfigList.
func (in *KubeletConfigList) DeepCopy() *KubeletConfigList {
	if in == nil {
		return nil
	}
	out := new(KubeletConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeletConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfigSpec) DeepCopyInto(out *KubeletConfigSpec) {
	*out = *in
	if in.AutoSizingReserved != nil {
		in, out := &in.AutoSizingReserved, &out.AutoSizingReserved
		*out = new(bool)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(int32)
		**out = **in
	}
	if in.MachineConfigPoolSelector != nil {
		in, out := &in.MachineConfigPoolSelector, &out.MachineConfigPoolSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeletConfig != nil {
		in, out := &in.KubeletConfig, &out.KubeletConfig
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSSecurityProfile != nil {
		in, out := &in.TLSSecurityProfile, &out.TLSSecurityProfile
		*out = new(configv1.TLSSecurityProfile)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfigSpec.
func (in *KubeletConfigSpec) DeepCopy() *KubeletConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KubeletConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfigStatus) DeepCopyInto(out *KubeletConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]KubeletConfigCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfigStatus.
func (in *KubeletConfigStatus) DeepCopy() *KubeletConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KubeletConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfig) DeepCopyInto(out *MachineConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfig.
func (in *MachineConfig) DeepCopy() *MachineConfig {
	if in == nil {
		return nil
	}
	out := new(MachineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigList) DeepCopyInto(out *MachineConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigList.
func (in *MachineConfigList) DeepCopy() *MachineConfigList {
	if in == nil {
		return nil
	}
	out := new(MachineConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigPool) DeepCopyInto(out *MachineConfigPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigPool.
func (in *MachineConfigPool) DeepCopy() *MachineConfigPool {
	if in == nil {
		return nil
	}
	out := new(MachineConfigPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineConfigPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigPoolCondition) DeepCopyInto(out *MachineConfigPoolCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigPoolCondition.
func (in *MachineConfigPoolCondition) DeepCopy() *MachineConfigPoolCondition {
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigPoolList) DeepCopyInto(out *MachineConfigPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineConfigPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigPoolList.
func (in *MachineConfigPoolList) DeepCopy() *MachineConfigPoolList {
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineConfigPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigPoolReference) DeepCopyInto(out *MachineConfigPoolReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigPoolReference.
func (in *MachineConfigPoolReference) DeepCopy() *MachineConfigPoolReference {
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigPoolSpec) DeepCopyInto(out *MachineConfigPoolSpec) {
	*out = *in
	if in.MachineConfigSelector != nil {
		in, out := &in.MachineConfigSelector, &out.MachineConfigSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	in.Configuration.DeepCopyInto(&out.Configuration)
	if in.PinnedImageSets != nil {
		in, out := &in.PinnedImageSets, &out.PinnedImageSets
		*out = make([]PinnedImageSetRef, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigPoolSpec.
func (in *MachineConfigPoolSpec) DeepCopy() *MachineConfigPoolSpec {
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigPoolStatus) DeepCopyInto(out *MachineConfigPoolStatus) {
	*out = *in
	in.Configuration.DeepCopyInto(&out.Configuration)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MachineConfigPoolCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CertExpirys != nil {
		in, out := &in.CertExpirys, &out.CertExpirys
		*out = make([]CertExpiry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PoolSynchronizersStatus != nil {
		in, out := &in.PoolSynchronizersStatus, &out.PoolSynchronizersStatus
		*out = make([]PoolSynchronizerStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigPoolStatus.
func (in *MachineConfigPoolStatus) DeepCopy() *MachineConfigPoolStatus {
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigPoolStatusConfiguration) DeepCopyInto(out *MachineConfigPoolStatusConfiguration) {
	*out = *in
	out.ObjectReference = in.ObjectReference
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigPoolStatusConfiguration.
func (in *MachineConfigPoolStatusConfiguration) DeepCopy() *MachineConfigPoolStatusConfiguration {
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolStatusConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigSpec) DeepCopyInto(out *MachineConfigSpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	if in.KernelArguments != nil {
		in, out := &in.KernelArguments, &out.KernelArguments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigSpec.
func (in *MachineConfigSpec) DeepCopy() *MachineConfigSpec {
	if in == nil {
		return nil
	}
	out := new(MachineConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineOSBuild) DeepCopyInto(out *MachineOSBuild) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(MachineOSBuildStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineOSBuild.
func (in *MachineOSBuild) DeepCopy() *MachineOSBuild {
	if in == nil {
		return nil
	}
	out := new(MachineOSBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineOSBuild) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineOSBuildList) DeepCopyInto(out *MachineOSBuildList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineOSBuild, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineOSBuildList.
func (in *MachineOSBuildList) DeepCopy() *MachineOSBuildList {
	if in == nil {
		return nil
	}
	out := new(MachineOSBuildList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineOSBuildList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineOSBuildSpec) DeepCopyInto(out *MachineOSBuildSpec) {
	*out = *in
	out.DesiredConfig = in.DesiredConfig
	out.MachineOSConfig = in.MachineOSConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineOSBuildSpec.
func (in *MachineOSBuildSpec) DeepCopy() *MachineOSBuildSpec {
	if in == nil {
		return nil
	}
	out := new(MachineOSBuildSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineOSBuildStatus) DeepCopyInto(out *MachineOSBuildStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Builder.DeepCopyInto(&out.Builder)
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]ObjectReference, len(*in))
		copy(*out, *in)
	}
	in.BuildStart.DeepCopyInto(&out.BuildStart)
	if in.BuildEnd != nil {
		in, out := &in.BuildEnd, &out.BuildEnd
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineOSBuildStatus.
func (in *MachineOSBuildStatus) DeepCopy() *MachineOSBuildStatus {
	if in == nil {
		return nil
	}
	out := new(MachineOSBuildStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineOSBuilderReference) DeepCopyInto(out *MachineOSBuilderReference) {
	*out = *in
	if in.ImageBuilderRef != nil {
		in, out := &in.ImageBuilderRef, &out.ImageBuilderRef
		*out = new(ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineOSBuilderReference.
func (in *MachineOSBuilderReference) DeepCopy() *MachineOSBuilderReference {
	if in == nil {
		return nil
	}
	out := new(MachineOSBuilderReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineOSConfig) DeepCopyInto(out *MachineOSConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(MachineOSConfigStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineOSConfig.
func (in *MachineOSConfig) DeepCopy() *MachineOSConfig {
	if in == nil {
		return nil
	}
	out := new(MachineOSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineOSConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineOSConfigList) DeepCopyInto(out *MachineOSConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineOSConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineOSConfigList.
func (in *MachineOSConfigList) DeepCopy() *MachineOSConfigList {
	if in == nil {
		return nil
	}
	out := new(MachineOSConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineOSConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineOSConfigReference) DeepCopyInto(out *MachineOSConfigReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineOSConfigReference.
func (in *MachineOSConfigReference) DeepCopy() *MachineOSConfigReference {
	if in == nil {
		return nil
	}
	out := new(MachineOSConfigReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineOSConfigSpec) DeepCopyInto(out *MachineOSConfigSpec) {
	*out = *in
	out.MachineConfigPool = in.MachineConfigPool
	in.BuildInputs.DeepCopyInto(&out.BuildInputs)
	if in.BuildOutputs != nil {
		in, out := &in.BuildOutputs, &out.BuildOutputs
		*out = new(BuildOutputs)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineOSConfigSpec.
func (in *MachineOSConfigSpec) DeepCopy() *MachineOSConfigSpec {
	if in == nil {
		return nil
	}
	out := new(MachineOSConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineOSConfigStatus) DeepCopyInto(out *MachineOSConfigStatus) {
	*out = *in
	if in.MachineOSBuild != nil {
		in, out := &in.MachineOSBuild, &out.MachineOSBuild
		*out = new(ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineOSConfigStatus.
func (in *MachineOSConfigStatus) DeepCopy() *MachineOSConfigStatus {
	if in == nil {
		return nil
	}
	out := new(MachineOSConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineOSContainerfile) DeepCopyInto(out *MachineOSContainerfile) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineOSContainerfile.
func (in *MachineOSContainerfile) DeepCopy() *MachineOSContainerfile {
	if in == nil {
		return nil
	}
	out := new(MachineOSContainerfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineOSImageBuilder) DeepCopyInto(out *MachineOSImageBuilder) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineOSImageBuilder.
func (in *MachineOSImageBuilder) DeepCopy() *MachineOSImageBuilder {
	if in == nil {
		return nil
	}
	out := new(MachineOSImageBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInfo) DeepCopyInto(out *NetworkInfo) {
	*out = *in
	if in.MTUMigration != nil {
		in, out := &in.MTUMigration, &out.MTUMigration
		*out = new(configv1.MTUMigration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInfo.
func (in *NetworkInfo) DeepCopy() *NetworkInfo {
	if in == nil {
		return nil
	}
	out := new(NetworkInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinnedImageSetRef) DeepCopyInto(out *PinnedImageSetRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinnedImageSetRef.
func (in *PinnedImageSetRef) DeepCopy() *PinnedImageSetRef {
	if in == nil {
		return nil
	}
	out := new(PinnedImageSetRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSynchronizerStatus) DeepCopyInto(out *PoolSynchronizerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSynchronizerStatus.
func (in *PoolSynchronizerStatus) DeepCopy() *PoolSynchronizerStatus {
	if in == nil {
		return nil
	}
	out := new(PoolSynchronizerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RenderedMachineConfigReference) DeepCopyInto(out *RenderedMachineConfigReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RenderedMachineConfigReference.
func (in *RenderedMachineConfigReference) DeepCopy() *RenderedMachineConfigReference {
	if in == nil {
		return nil
	}
	out := new(RenderedMachineConfigReference)
	in.DeepCopyInto(out)
	return out
}
