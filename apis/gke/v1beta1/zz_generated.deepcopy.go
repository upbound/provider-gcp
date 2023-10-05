//go:build !ignore_autogenerated

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
func (in *BackupBackupPlan) DeepCopyInto(out *BackupBackupPlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBackupPlan.
func (in *BackupBackupPlan) DeepCopy() *BackupBackupPlan {
	if in == nil {
		return nil
	}
	out := new(BackupBackupPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupBackupPlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupBackupPlanInitParameters) DeepCopyInto(out *BackupBackupPlanInitParameters) {
	*out = *in
	if in.BackupConfig != nil {
		in, out := &in.BackupConfig, &out.BackupConfig
		*out = make([]BackupConfigInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BackupSchedule != nil {
		in, out := &in.BackupSchedule, &out.BackupSchedule
		*out = make([]BackupScheduleInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Deactivated != nil {
		in, out := &in.Deactivated, &out.Deactivated
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.RetentionPolicy != nil {
		in, out := &in.RetentionPolicy, &out.RetentionPolicy
		*out = make([]RetentionPolicyInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBackupPlanInitParameters.
func (in *BackupBackupPlanInitParameters) DeepCopy() *BackupBackupPlanInitParameters {
	if in == nil {
		return nil
	}
	out := new(BackupBackupPlanInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupBackupPlanList) DeepCopyInto(out *BackupBackupPlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupBackupPlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBackupPlanList.
func (in *BackupBackupPlanList) DeepCopy() *BackupBackupPlanList {
	if in == nil {
		return nil
	}
	out := new(BackupBackupPlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupBackupPlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupBackupPlanObservation) DeepCopyInto(out *BackupBackupPlanObservation) {
	*out = *in
	if in.BackupConfig != nil {
		in, out := &in.BackupConfig, &out.BackupConfig
		*out = make([]BackupConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BackupSchedule != nil {
		in, out := &in.BackupSchedule, &out.BackupSchedule
		*out = make([]BackupScheduleObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(string)
		**out = **in
	}
	if in.Deactivated != nil {
		in, out := &in.Deactivated, &out.Deactivated
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EffectiveLabels != nil {
		in, out := &in.EffectiveLabels, &out.EffectiveLabels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
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
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ProtectedPodCount != nil {
		in, out := &in.ProtectedPodCount, &out.ProtectedPodCount
		*out = new(float64)
		**out = **in
	}
	if in.RetentionPolicy != nil {
		in, out := &in.RetentionPolicy, &out.RetentionPolicy
		*out = make([]RetentionPolicyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StateReason != nil {
		in, out := &in.StateReason, &out.StateReason
		*out = new(string)
		**out = **in
	}
	if in.TerraformLabels != nil {
		in, out := &in.TerraformLabels, &out.TerraformLabels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBackupPlanObservation.
func (in *BackupBackupPlanObservation) DeepCopy() *BackupBackupPlanObservation {
	if in == nil {
		return nil
	}
	out := new(BackupBackupPlanObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupBackupPlanParameters) DeepCopyInto(out *BackupBackupPlanParameters) {
	*out = *in
	if in.BackupConfig != nil {
		in, out := &in.BackupConfig, &out.BackupConfig
		*out = make([]BackupConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BackupSchedule != nil {
		in, out := &in.BackupSchedule, &out.BackupSchedule
		*out = make([]BackupScheduleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(string)
		**out = **in
	}
	if in.ClusterRef != nil {
		in, out := &in.ClusterRef, &out.ClusterRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterSelector != nil {
		in, out := &in.ClusterSelector, &out.ClusterSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Deactivated != nil {
		in, out := &in.Deactivated, &out.Deactivated
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.RetentionPolicy != nil {
		in, out := &in.RetentionPolicy, &out.RetentionPolicy
		*out = make([]RetentionPolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBackupPlanParameters.
func (in *BackupBackupPlanParameters) DeepCopy() *BackupBackupPlanParameters {
	if in == nil {
		return nil
	}
	out := new(BackupBackupPlanParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupBackupPlanSpec) DeepCopyInto(out *BackupBackupPlanSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBackupPlanSpec.
func (in *BackupBackupPlanSpec) DeepCopy() *BackupBackupPlanSpec {
	if in == nil {
		return nil
	}
	out := new(BackupBackupPlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupBackupPlanStatus) DeepCopyInto(out *BackupBackupPlanStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBackupPlanStatus.
func (in *BackupBackupPlanStatus) DeepCopy() *BackupBackupPlanStatus {
	if in == nil {
		return nil
	}
	out := new(BackupBackupPlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupConfigInitParameters) DeepCopyInto(out *BackupConfigInitParameters) {
	*out = *in
	if in.AllNamespaces != nil {
		in, out := &in.AllNamespaces, &out.AllNamespaces
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionKey != nil {
		in, out := &in.EncryptionKey, &out.EncryptionKey
		*out = make([]EncryptionKeyInitParameters, len(*in))
		copy(*out, *in)
	}
	if in.IncludeSecrets != nil {
		in, out := &in.IncludeSecrets, &out.IncludeSecrets
		*out = new(bool)
		**out = **in
	}
	if in.IncludeVolumeData != nil {
		in, out := &in.IncludeVolumeData, &out.IncludeVolumeData
		*out = new(bool)
		**out = **in
	}
	if in.SelectedApplications != nil {
		in, out := &in.SelectedApplications, &out.SelectedApplications
		*out = make([]SelectedApplicationsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SelectedNamespaces != nil {
		in, out := &in.SelectedNamespaces, &out.SelectedNamespaces
		*out = make([]SelectedNamespacesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupConfigInitParameters.
func (in *BackupConfigInitParameters) DeepCopy() *BackupConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(BackupConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupConfigObservation) DeepCopyInto(out *BackupConfigObservation) {
	*out = *in
	if in.AllNamespaces != nil {
		in, out := &in.AllNamespaces, &out.AllNamespaces
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionKey != nil {
		in, out := &in.EncryptionKey, &out.EncryptionKey
		*out = make([]EncryptionKeyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IncludeSecrets != nil {
		in, out := &in.IncludeSecrets, &out.IncludeSecrets
		*out = new(bool)
		**out = **in
	}
	if in.IncludeVolumeData != nil {
		in, out := &in.IncludeVolumeData, &out.IncludeVolumeData
		*out = new(bool)
		**out = **in
	}
	if in.SelectedApplications != nil {
		in, out := &in.SelectedApplications, &out.SelectedApplications
		*out = make([]SelectedApplicationsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SelectedNamespaces != nil {
		in, out := &in.SelectedNamespaces, &out.SelectedNamespaces
		*out = make([]SelectedNamespacesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupConfigObservation.
func (in *BackupConfigObservation) DeepCopy() *BackupConfigObservation {
	if in == nil {
		return nil
	}
	out := new(BackupConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupConfigParameters) DeepCopyInto(out *BackupConfigParameters) {
	*out = *in
	if in.AllNamespaces != nil {
		in, out := &in.AllNamespaces, &out.AllNamespaces
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionKey != nil {
		in, out := &in.EncryptionKey, &out.EncryptionKey
		*out = make([]EncryptionKeyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IncludeSecrets != nil {
		in, out := &in.IncludeSecrets, &out.IncludeSecrets
		*out = new(bool)
		**out = **in
	}
	if in.IncludeVolumeData != nil {
		in, out := &in.IncludeVolumeData, &out.IncludeVolumeData
		*out = new(bool)
		**out = **in
	}
	if in.SelectedApplications != nil {
		in, out := &in.SelectedApplications, &out.SelectedApplications
		*out = make([]SelectedApplicationsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SelectedNamespaces != nil {
		in, out := &in.SelectedNamespaces, &out.SelectedNamespaces
		*out = make([]SelectedNamespacesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupConfigParameters.
func (in *BackupConfigParameters) DeepCopy() *BackupConfigParameters {
	if in == nil {
		return nil
	}
	out := new(BackupConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupScheduleInitParameters) DeepCopyInto(out *BackupScheduleInitParameters) {
	*out = *in
	if in.CronSchedule != nil {
		in, out := &in.CronSchedule, &out.CronSchedule
		*out = new(string)
		**out = **in
	}
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupScheduleInitParameters.
func (in *BackupScheduleInitParameters) DeepCopy() *BackupScheduleInitParameters {
	if in == nil {
		return nil
	}
	out := new(BackupScheduleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupScheduleObservation) DeepCopyInto(out *BackupScheduleObservation) {
	*out = *in
	if in.CronSchedule != nil {
		in, out := &in.CronSchedule, &out.CronSchedule
		*out = new(string)
		**out = **in
	}
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupScheduleObservation.
func (in *BackupScheduleObservation) DeepCopy() *BackupScheduleObservation {
	if in == nil {
		return nil
	}
	out := new(BackupScheduleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupScheduleParameters) DeepCopyInto(out *BackupScheduleParameters) {
	*out = *in
	if in.CronSchedule != nil {
		in, out := &in.CronSchedule, &out.CronSchedule
		*out = new(string)
		**out = **in
	}
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupScheduleParameters.
func (in *BackupScheduleParameters) DeepCopy() *BackupScheduleParameters {
	if in == nil {
		return nil
	}
	out := new(BackupScheduleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionKeyInitParameters) DeepCopyInto(out *EncryptionKeyInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionKeyInitParameters.
func (in *EncryptionKeyInitParameters) DeepCopy() *EncryptionKeyInitParameters {
	if in == nil {
		return nil
	}
	out := new(EncryptionKeyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionKeyObservation) DeepCopyInto(out *EncryptionKeyObservation) {
	*out = *in
	if in.GCPKMSEncryptionKey != nil {
		in, out := &in.GCPKMSEncryptionKey, &out.GCPKMSEncryptionKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionKeyObservation.
func (in *EncryptionKeyObservation) DeepCopy() *EncryptionKeyObservation {
	if in == nil {
		return nil
	}
	out := new(EncryptionKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionKeyParameters) DeepCopyInto(out *EncryptionKeyParameters) {
	*out = *in
	if in.GCPKMSEncryptionKey != nil {
		in, out := &in.GCPKMSEncryptionKey, &out.GCPKMSEncryptionKey
		*out = new(string)
		**out = **in
	}
	if in.GCPKMSEncryptionKeyRef != nil {
		in, out := &in.GCPKMSEncryptionKeyRef, &out.GCPKMSEncryptionKeyRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.GCPKMSEncryptionKeySelector != nil {
		in, out := &in.GCPKMSEncryptionKeySelector, &out.GCPKMSEncryptionKeySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionKeyParameters.
func (in *EncryptionKeyParameters) DeepCopy() *EncryptionKeyParameters {
	if in == nil {
		return nil
	}
	out := new(EncryptionKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedNamesInitParameters) DeepCopyInto(out *NamespacedNamesInitParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedNamesInitParameters.
func (in *NamespacedNamesInitParameters) DeepCopy() *NamespacedNamesInitParameters {
	if in == nil {
		return nil
	}
	out := new(NamespacedNamesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedNamesObservation) DeepCopyInto(out *NamespacedNamesObservation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedNamesObservation.
func (in *NamespacedNamesObservation) DeepCopy() *NamespacedNamesObservation {
	if in == nil {
		return nil
	}
	out := new(NamespacedNamesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedNamesParameters) DeepCopyInto(out *NamespacedNamesParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedNamesParameters.
func (in *NamespacedNamesParameters) DeepCopy() *NamespacedNamesParameters {
	if in == nil {
		return nil
	}
	out := new(NamespacedNamesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPolicyInitParameters) DeepCopyInto(out *RetentionPolicyInitParameters) {
	*out = *in
	if in.BackupDeleteLockDays != nil {
		in, out := &in.BackupDeleteLockDays, &out.BackupDeleteLockDays
		*out = new(float64)
		**out = **in
	}
	if in.BackupRetainDays != nil {
		in, out := &in.BackupRetainDays, &out.BackupRetainDays
		*out = new(float64)
		**out = **in
	}
	if in.Locked != nil {
		in, out := &in.Locked, &out.Locked
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPolicyInitParameters.
func (in *RetentionPolicyInitParameters) DeepCopy() *RetentionPolicyInitParameters {
	if in == nil {
		return nil
	}
	out := new(RetentionPolicyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPolicyObservation) DeepCopyInto(out *RetentionPolicyObservation) {
	*out = *in
	if in.BackupDeleteLockDays != nil {
		in, out := &in.BackupDeleteLockDays, &out.BackupDeleteLockDays
		*out = new(float64)
		**out = **in
	}
	if in.BackupRetainDays != nil {
		in, out := &in.BackupRetainDays, &out.BackupRetainDays
		*out = new(float64)
		**out = **in
	}
	if in.Locked != nil {
		in, out := &in.Locked, &out.Locked
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPolicyObservation.
func (in *RetentionPolicyObservation) DeepCopy() *RetentionPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(RetentionPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPolicyParameters) DeepCopyInto(out *RetentionPolicyParameters) {
	*out = *in
	if in.BackupDeleteLockDays != nil {
		in, out := &in.BackupDeleteLockDays, &out.BackupDeleteLockDays
		*out = new(float64)
		**out = **in
	}
	if in.BackupRetainDays != nil {
		in, out := &in.BackupRetainDays, &out.BackupRetainDays
		*out = new(float64)
		**out = **in
	}
	if in.Locked != nil {
		in, out := &in.Locked, &out.Locked
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPolicyParameters.
func (in *RetentionPolicyParameters) DeepCopy() *RetentionPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(RetentionPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectedApplicationsInitParameters) DeepCopyInto(out *SelectedApplicationsInitParameters) {
	*out = *in
	if in.NamespacedNames != nil {
		in, out := &in.NamespacedNames, &out.NamespacedNames
		*out = make([]NamespacedNamesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectedApplicationsInitParameters.
func (in *SelectedApplicationsInitParameters) DeepCopy() *SelectedApplicationsInitParameters {
	if in == nil {
		return nil
	}
	out := new(SelectedApplicationsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectedApplicationsObservation) DeepCopyInto(out *SelectedApplicationsObservation) {
	*out = *in
	if in.NamespacedNames != nil {
		in, out := &in.NamespacedNames, &out.NamespacedNames
		*out = make([]NamespacedNamesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectedApplicationsObservation.
func (in *SelectedApplicationsObservation) DeepCopy() *SelectedApplicationsObservation {
	if in == nil {
		return nil
	}
	out := new(SelectedApplicationsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectedApplicationsParameters) DeepCopyInto(out *SelectedApplicationsParameters) {
	*out = *in
	if in.NamespacedNames != nil {
		in, out := &in.NamespacedNames, &out.NamespacedNames
		*out = make([]NamespacedNamesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectedApplicationsParameters.
func (in *SelectedApplicationsParameters) DeepCopy() *SelectedApplicationsParameters {
	if in == nil {
		return nil
	}
	out := new(SelectedApplicationsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectedNamespacesInitParameters) DeepCopyInto(out *SelectedNamespacesInitParameters) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectedNamespacesInitParameters.
func (in *SelectedNamespacesInitParameters) DeepCopy() *SelectedNamespacesInitParameters {
	if in == nil {
		return nil
	}
	out := new(SelectedNamespacesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectedNamespacesObservation) DeepCopyInto(out *SelectedNamespacesObservation) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectedNamespacesObservation.
func (in *SelectedNamespacesObservation) DeepCopy() *SelectedNamespacesObservation {
	if in == nil {
		return nil
	}
	out := new(SelectedNamespacesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectedNamespacesParameters) DeepCopyInto(out *SelectedNamespacesParameters) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectedNamespacesParameters.
func (in *SelectedNamespacesParameters) DeepCopy() *SelectedNamespacesParameters {
	if in == nil {
		return nil
	}
	out := new(SelectedNamespacesParameters)
	in.DeepCopyInto(out)
	return out
}
