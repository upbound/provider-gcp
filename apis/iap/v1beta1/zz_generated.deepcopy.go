//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionObservation) DeepCopyInto(out *ConditionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionObservation.
func (in *ConditionObservation) DeepCopy() *ConditionObservation {
	if in == nil {
		return nil
	}
	out := new(ConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionParameters) DeepCopyInto(out *ConditionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionParameters.
func (in *ConditionParameters) DeepCopy() *ConditionParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebBackendServiceIAMMember) DeepCopyInto(out *WebBackendServiceIAMMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebBackendServiceIAMMember.
func (in *WebBackendServiceIAMMember) DeepCopy() *WebBackendServiceIAMMember {
	if in == nil {
		return nil
	}
	out := new(WebBackendServiceIAMMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebBackendServiceIAMMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebBackendServiceIAMMemberList) DeepCopyInto(out *WebBackendServiceIAMMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WebBackendServiceIAMMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebBackendServiceIAMMemberList.
func (in *WebBackendServiceIAMMemberList) DeepCopy() *WebBackendServiceIAMMemberList {
	if in == nil {
		return nil
	}
	out := new(WebBackendServiceIAMMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebBackendServiceIAMMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebBackendServiceIAMMemberObservation) DeepCopyInto(out *WebBackendServiceIAMMemberObservation) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebBackendServiceIAMMemberObservation.
func (in *WebBackendServiceIAMMemberObservation) DeepCopy() *WebBackendServiceIAMMemberObservation {
	if in == nil {
		return nil
	}
	out := new(WebBackendServiceIAMMemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebBackendServiceIAMMemberParameters) DeepCopyInto(out *WebBackendServiceIAMMemberParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]ConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.WebBackendService != nil {
		in, out := &in.WebBackendService, &out.WebBackendService
		*out = new(string)
		**out = **in
	}
	if in.WebBackendServiceRef != nil {
		in, out := &in.WebBackendServiceRef, &out.WebBackendServiceRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.WebBackendServiceSelector != nil {
		in, out := &in.WebBackendServiceSelector, &out.WebBackendServiceSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebBackendServiceIAMMemberParameters.
func (in *WebBackendServiceIAMMemberParameters) DeepCopy() *WebBackendServiceIAMMemberParameters {
	if in == nil {
		return nil
	}
	out := new(WebBackendServiceIAMMemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebBackendServiceIAMMemberSpec) DeepCopyInto(out *WebBackendServiceIAMMemberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebBackendServiceIAMMemberSpec.
func (in *WebBackendServiceIAMMemberSpec) DeepCopy() *WebBackendServiceIAMMemberSpec {
	if in == nil {
		return nil
	}
	out := new(WebBackendServiceIAMMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebBackendServiceIAMMemberStatus) DeepCopyInto(out *WebBackendServiceIAMMemberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebBackendServiceIAMMemberStatus.
func (in *WebBackendServiceIAMMemberStatus) DeepCopy() *WebBackendServiceIAMMemberStatus {
	if in == nil {
		return nil
	}
	out := new(WebBackendServiceIAMMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebIAMMember) DeepCopyInto(out *WebIAMMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebIAMMember.
func (in *WebIAMMember) DeepCopy() *WebIAMMember {
	if in == nil {
		return nil
	}
	out := new(WebIAMMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebIAMMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebIAMMemberConditionObservation) DeepCopyInto(out *WebIAMMemberConditionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebIAMMemberConditionObservation.
func (in *WebIAMMemberConditionObservation) DeepCopy() *WebIAMMemberConditionObservation {
	if in == nil {
		return nil
	}
	out := new(WebIAMMemberConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebIAMMemberConditionParameters) DeepCopyInto(out *WebIAMMemberConditionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebIAMMemberConditionParameters.
func (in *WebIAMMemberConditionParameters) DeepCopy() *WebIAMMemberConditionParameters {
	if in == nil {
		return nil
	}
	out := new(WebIAMMemberConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebIAMMemberList) DeepCopyInto(out *WebIAMMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WebIAMMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebIAMMemberList.
func (in *WebIAMMemberList) DeepCopy() *WebIAMMemberList {
	if in == nil {
		return nil
	}
	out := new(WebIAMMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebIAMMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebIAMMemberObservation) DeepCopyInto(out *WebIAMMemberObservation) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebIAMMemberObservation.
func (in *WebIAMMemberObservation) DeepCopy() *WebIAMMemberObservation {
	if in == nil {
		return nil
	}
	out := new(WebIAMMemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebIAMMemberParameters) DeepCopyInto(out *WebIAMMemberParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]WebIAMMemberConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebIAMMemberParameters.
func (in *WebIAMMemberParameters) DeepCopy() *WebIAMMemberParameters {
	if in == nil {
		return nil
	}
	out := new(WebIAMMemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebIAMMemberSpec) DeepCopyInto(out *WebIAMMemberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebIAMMemberSpec.
func (in *WebIAMMemberSpec) DeepCopy() *WebIAMMemberSpec {
	if in == nil {
		return nil
	}
	out := new(WebIAMMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebIAMMemberStatus) DeepCopyInto(out *WebIAMMemberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebIAMMemberStatus.
func (in *WebIAMMemberStatus) DeepCopy() *WebIAMMemberStatus {
	if in == nil {
		return nil
	}
	out := new(WebIAMMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeAppEngineIAMMember) DeepCopyInto(out *WebTypeAppEngineIAMMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeAppEngineIAMMember.
func (in *WebTypeAppEngineIAMMember) DeepCopy() *WebTypeAppEngineIAMMember {
	if in == nil {
		return nil
	}
	out := new(WebTypeAppEngineIAMMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebTypeAppEngineIAMMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeAppEngineIAMMemberConditionObservation) DeepCopyInto(out *WebTypeAppEngineIAMMemberConditionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeAppEngineIAMMemberConditionObservation.
func (in *WebTypeAppEngineIAMMemberConditionObservation) DeepCopy() *WebTypeAppEngineIAMMemberConditionObservation {
	if in == nil {
		return nil
	}
	out := new(WebTypeAppEngineIAMMemberConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeAppEngineIAMMemberConditionParameters) DeepCopyInto(out *WebTypeAppEngineIAMMemberConditionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeAppEngineIAMMemberConditionParameters.
func (in *WebTypeAppEngineIAMMemberConditionParameters) DeepCopy() *WebTypeAppEngineIAMMemberConditionParameters {
	if in == nil {
		return nil
	}
	out := new(WebTypeAppEngineIAMMemberConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeAppEngineIAMMemberList) DeepCopyInto(out *WebTypeAppEngineIAMMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WebTypeAppEngineIAMMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeAppEngineIAMMemberList.
func (in *WebTypeAppEngineIAMMemberList) DeepCopy() *WebTypeAppEngineIAMMemberList {
	if in == nil {
		return nil
	}
	out := new(WebTypeAppEngineIAMMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebTypeAppEngineIAMMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeAppEngineIAMMemberObservation) DeepCopyInto(out *WebTypeAppEngineIAMMemberObservation) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeAppEngineIAMMemberObservation.
func (in *WebTypeAppEngineIAMMemberObservation) DeepCopy() *WebTypeAppEngineIAMMemberObservation {
	if in == nil {
		return nil
	}
	out := new(WebTypeAppEngineIAMMemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeAppEngineIAMMemberParameters) DeepCopyInto(out *WebTypeAppEngineIAMMemberParameters) {
	*out = *in
	if in.AppID != nil {
		in, out := &in.AppID, &out.AppID
		*out = new(string)
		**out = **in
	}
	if in.AppIDRef != nil {
		in, out := &in.AppIDRef, &out.AppIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AppIDSelector != nil {
		in, out := &in.AppIDSelector, &out.AppIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]WebTypeAppEngineIAMMemberConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeAppEngineIAMMemberParameters.
func (in *WebTypeAppEngineIAMMemberParameters) DeepCopy() *WebTypeAppEngineIAMMemberParameters {
	if in == nil {
		return nil
	}
	out := new(WebTypeAppEngineIAMMemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeAppEngineIAMMemberSpec) DeepCopyInto(out *WebTypeAppEngineIAMMemberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeAppEngineIAMMemberSpec.
func (in *WebTypeAppEngineIAMMemberSpec) DeepCopy() *WebTypeAppEngineIAMMemberSpec {
	if in == nil {
		return nil
	}
	out := new(WebTypeAppEngineIAMMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeAppEngineIAMMemberStatus) DeepCopyInto(out *WebTypeAppEngineIAMMemberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeAppEngineIAMMemberStatus.
func (in *WebTypeAppEngineIAMMemberStatus) DeepCopy() *WebTypeAppEngineIAMMemberStatus {
	if in == nil {
		return nil
	}
	out := new(WebTypeAppEngineIAMMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeComputeIAMMember) DeepCopyInto(out *WebTypeComputeIAMMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeComputeIAMMember.
func (in *WebTypeComputeIAMMember) DeepCopy() *WebTypeComputeIAMMember {
	if in == nil {
		return nil
	}
	out := new(WebTypeComputeIAMMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebTypeComputeIAMMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeComputeIAMMemberConditionObservation) DeepCopyInto(out *WebTypeComputeIAMMemberConditionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeComputeIAMMemberConditionObservation.
func (in *WebTypeComputeIAMMemberConditionObservation) DeepCopy() *WebTypeComputeIAMMemberConditionObservation {
	if in == nil {
		return nil
	}
	out := new(WebTypeComputeIAMMemberConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeComputeIAMMemberConditionParameters) DeepCopyInto(out *WebTypeComputeIAMMemberConditionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeComputeIAMMemberConditionParameters.
func (in *WebTypeComputeIAMMemberConditionParameters) DeepCopy() *WebTypeComputeIAMMemberConditionParameters {
	if in == nil {
		return nil
	}
	out := new(WebTypeComputeIAMMemberConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeComputeIAMMemberList) DeepCopyInto(out *WebTypeComputeIAMMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WebTypeComputeIAMMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeComputeIAMMemberList.
func (in *WebTypeComputeIAMMemberList) DeepCopy() *WebTypeComputeIAMMemberList {
	if in == nil {
		return nil
	}
	out := new(WebTypeComputeIAMMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebTypeComputeIAMMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeComputeIAMMemberObservation) DeepCopyInto(out *WebTypeComputeIAMMemberObservation) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeComputeIAMMemberObservation.
func (in *WebTypeComputeIAMMemberObservation) DeepCopy() *WebTypeComputeIAMMemberObservation {
	if in == nil {
		return nil
	}
	out := new(WebTypeComputeIAMMemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeComputeIAMMemberParameters) DeepCopyInto(out *WebTypeComputeIAMMemberParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = make([]WebTypeComputeIAMMemberConditionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeComputeIAMMemberParameters.
func (in *WebTypeComputeIAMMemberParameters) DeepCopy() *WebTypeComputeIAMMemberParameters {
	if in == nil {
		return nil
	}
	out := new(WebTypeComputeIAMMemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeComputeIAMMemberSpec) DeepCopyInto(out *WebTypeComputeIAMMemberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeComputeIAMMemberSpec.
func (in *WebTypeComputeIAMMemberSpec) DeepCopy() *WebTypeComputeIAMMemberSpec {
	if in == nil {
		return nil
	}
	out := new(WebTypeComputeIAMMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTypeComputeIAMMemberStatus) DeepCopyInto(out *WebTypeComputeIAMMemberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTypeComputeIAMMemberStatus.
func (in *WebTypeComputeIAMMemberStatus) DeepCopy() *WebTypeComputeIAMMemberStatus {
	if in == nil {
		return nil
	}
	out := new(WebTypeComputeIAMMemberStatus)
	in.DeepCopyInto(out)
	return out
}
