// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PagerDutyIntegration) DeepCopyInto(out *PagerDutyIntegration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PagerDutyIntegration.
func (in *PagerDutyIntegration) DeepCopy() *PagerDutyIntegration {
	if in == nil {
		return nil
	}
	out := new(PagerDutyIntegration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PagerDutyIntegration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PagerDutyIntegrationList) DeepCopyInto(out *PagerDutyIntegrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PagerDutyIntegration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PagerDutyIntegrationList.
func (in *PagerDutyIntegrationList) DeepCopy() *PagerDutyIntegrationList {
	if in == nil {
		return nil
	}
	out := new(PagerDutyIntegrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PagerDutyIntegrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PagerDutyIntegrationSpec) DeepCopyInto(out *PagerDutyIntegrationSpec) {
	*out = *in
	out.PagerdutyApiKeySecretRef = in.PagerdutyApiKeySecretRef
	in.ClusterDeploymentSelector.DeepCopyInto(&out.ClusterDeploymentSelector)
	out.TargetSecretRef = in.TargetSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PagerDutyIntegrationSpec.
func (in *PagerDutyIntegrationSpec) DeepCopy() *PagerDutyIntegrationSpec {
	if in == nil {
		return nil
	}
	out := new(PagerDutyIntegrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PagerDutyIntegrationStatus) DeepCopyInto(out *PagerDutyIntegrationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PagerDutyIntegrationStatus.
func (in *PagerDutyIntegrationStatus) DeepCopy() *PagerDutyIntegrationStatus {
	if in == nil {
		return nil
	}
	out := new(PagerDutyIntegrationStatus)
	in.DeepCopyInto(out)
	return out
}
