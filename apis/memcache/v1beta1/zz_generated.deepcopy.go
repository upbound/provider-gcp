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
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Instance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceInitParameters) DeepCopyInto(out *InstanceInitParameters) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
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
	if in.MaintenancePolicy != nil {
		in, out := &in.MaintenancePolicy, &out.MaintenancePolicy
		*out = make([]MaintenancePolicyInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemcacheParameters != nil {
		in, out := &in.MemcacheParameters, &out.MemcacheParameters
		*out = make([]MemcacheParametersInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemcacheVersion != nil {
		in, out := &in.MemcacheVersion, &out.MemcacheVersion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NodeConfig != nil {
		in, out := &in.NodeConfig, &out.NodeConfig
		*out = make([]NodeConfigInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(float64)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceInitParameters.
func (in *InstanceInitParameters) DeepCopy() *InstanceInitParameters {
	if in == nil {
		return nil
	}
	out := new(InstanceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceList) DeepCopyInto(out *InstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceList.
func (in *InstanceList) DeepCopy() *InstanceList {
	if in == nil {
		return nil
	}
	out := new(InstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceObservation) DeepCopyInto(out *InstanceObservation) {
	*out = *in
	if in.AuthorizedNetwork != nil {
		in, out := &in.AuthorizedNetwork, &out.AuthorizedNetwork
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.DiscoveryEndpoint != nil {
		in, out := &in.DiscoveryEndpoint, &out.DiscoveryEndpoint
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
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
	if in.MaintenancePolicy != nil {
		in, out := &in.MaintenancePolicy, &out.MaintenancePolicy
		*out = make([]MaintenancePolicyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaintenanceSchedule != nil {
		in, out := &in.MaintenanceSchedule, &out.MaintenanceSchedule
		*out = make([]MaintenanceScheduleObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemcacheFullVersion != nil {
		in, out := &in.MemcacheFullVersion, &out.MemcacheFullVersion
		*out = new(string)
		**out = **in
	}
	if in.MemcacheNodes != nil {
		in, out := &in.MemcacheNodes, &out.MemcacheNodes
		*out = make([]MemcacheNodesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemcacheParameters != nil {
		in, out := &in.MemcacheParameters, &out.MemcacheParameters
		*out = make([]MemcacheParametersObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemcacheVersion != nil {
		in, out := &in.MemcacheVersion, &out.MemcacheVersion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NodeConfig != nil {
		in, out := &in.NodeConfig, &out.NodeConfig
		*out = make([]NodeConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(float64)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
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
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceObservation.
func (in *InstanceObservation) DeepCopy() *InstanceObservation {
	if in == nil {
		return nil
	}
	out := new(InstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceParameters) DeepCopyInto(out *InstanceParameters) {
	*out = *in
	if in.AuthorizedNetwork != nil {
		in, out := &in.AuthorizedNetwork, &out.AuthorizedNetwork
		*out = new(string)
		**out = **in
	}
	if in.AuthorizedNetworkRef != nil {
		in, out := &in.AuthorizedNetworkRef, &out.AuthorizedNetworkRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AuthorizedNetworkSelector != nil {
		in, out := &in.AuthorizedNetworkSelector, &out.AuthorizedNetworkSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
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
	if in.MaintenancePolicy != nil {
		in, out := &in.MaintenancePolicy, &out.MaintenancePolicy
		*out = make([]MaintenancePolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemcacheParameters != nil {
		in, out := &in.MemcacheParameters, &out.MemcacheParameters
		*out = make([]MemcacheParametersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemcacheVersion != nil {
		in, out := &in.MemcacheVersion, &out.MemcacheVersion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NodeConfig != nil {
		in, out := &in.NodeConfig, &out.NodeConfig
		*out = make([]NodeConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(float64)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceParameters.
func (in *InstanceParameters) DeepCopy() *InstanceParameters {
	if in == nil {
		return nil
	}
	out := new(InstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenancePolicyInitParameters) DeepCopyInto(out *MaintenancePolicyInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.WeeklyMaintenanceWindow != nil {
		in, out := &in.WeeklyMaintenanceWindow, &out.WeeklyMaintenanceWindow
		*out = make([]WeeklyMaintenanceWindowInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenancePolicyInitParameters.
func (in *MaintenancePolicyInitParameters) DeepCopy() *MaintenancePolicyInitParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenancePolicyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenancePolicyObservation) DeepCopyInto(out *MaintenancePolicyObservation) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.WeeklyMaintenanceWindow != nil {
		in, out := &in.WeeklyMaintenanceWindow, &out.WeeklyMaintenanceWindow
		*out = make([]WeeklyMaintenanceWindowObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenancePolicyObservation.
func (in *MaintenancePolicyObservation) DeepCopy() *MaintenancePolicyObservation {
	if in == nil {
		return nil
	}
	out := new(MaintenancePolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenancePolicyParameters) DeepCopyInto(out *MaintenancePolicyParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.WeeklyMaintenanceWindow != nil {
		in, out := &in.WeeklyMaintenanceWindow, &out.WeeklyMaintenanceWindow
		*out = make([]WeeklyMaintenanceWindowParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenancePolicyParameters.
func (in *MaintenancePolicyParameters) DeepCopy() *MaintenancePolicyParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenancePolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceScheduleInitParameters) DeepCopyInto(out *MaintenanceScheduleInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceScheduleInitParameters.
func (in *MaintenanceScheduleInitParameters) DeepCopy() *MaintenanceScheduleInitParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenanceScheduleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceScheduleObservation) DeepCopyInto(out *MaintenanceScheduleObservation) {
	*out = *in
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.ScheduleDeadlineTime != nil {
		in, out := &in.ScheduleDeadlineTime, &out.ScheduleDeadlineTime
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceScheduleObservation.
func (in *MaintenanceScheduleObservation) DeepCopy() *MaintenanceScheduleObservation {
	if in == nil {
		return nil
	}
	out := new(MaintenanceScheduleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceScheduleParameters) DeepCopyInto(out *MaintenanceScheduleParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceScheduleParameters.
func (in *MaintenanceScheduleParameters) DeepCopy() *MaintenanceScheduleParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenanceScheduleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcacheNodesInitParameters) DeepCopyInto(out *MemcacheNodesInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcacheNodesInitParameters.
func (in *MemcacheNodesInitParameters) DeepCopy() *MemcacheNodesInitParameters {
	if in == nil {
		return nil
	}
	out := new(MemcacheNodesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcacheNodesObservation) DeepCopyInto(out *MemcacheNodesObservation) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.NodeID != nil {
		in, out := &in.NodeID, &out.NodeID
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcacheNodesObservation.
func (in *MemcacheNodesObservation) DeepCopy() *MemcacheNodesObservation {
	if in == nil {
		return nil
	}
	out := new(MemcacheNodesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcacheNodesParameters) DeepCopyInto(out *MemcacheNodesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcacheNodesParameters.
func (in *MemcacheNodesParameters) DeepCopy() *MemcacheNodesParameters {
	if in == nil {
		return nil
	}
	out := new(MemcacheNodesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcacheParametersInitParameters) DeepCopyInto(out *MemcacheParametersInitParameters) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcacheParametersInitParameters.
func (in *MemcacheParametersInitParameters) DeepCopy() *MemcacheParametersInitParameters {
	if in == nil {
		return nil
	}
	out := new(MemcacheParametersInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcacheParametersObservation) DeepCopyInto(out *MemcacheParametersObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcacheParametersObservation.
func (in *MemcacheParametersObservation) DeepCopy() *MemcacheParametersObservation {
	if in == nil {
		return nil
	}
	out := new(MemcacheParametersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcacheParametersParameters) DeepCopyInto(out *MemcacheParametersParameters) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcacheParametersParameters.
func (in *MemcacheParametersParameters) DeepCopy() *MemcacheParametersParameters {
	if in == nil {
		return nil
	}
	out := new(MemcacheParametersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfigInitParameters) DeepCopyInto(out *NodeConfigInitParameters) {
	*out = *in
	if in.CPUCount != nil {
		in, out := &in.CPUCount, &out.CPUCount
		*out = new(float64)
		**out = **in
	}
	if in.MemorySizeMb != nil {
		in, out := &in.MemorySizeMb, &out.MemorySizeMb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfigInitParameters.
func (in *NodeConfigInitParameters) DeepCopy() *NodeConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(NodeConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfigObservation) DeepCopyInto(out *NodeConfigObservation) {
	*out = *in
	if in.CPUCount != nil {
		in, out := &in.CPUCount, &out.CPUCount
		*out = new(float64)
		**out = **in
	}
	if in.MemorySizeMb != nil {
		in, out := &in.MemorySizeMb, &out.MemorySizeMb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfigObservation.
func (in *NodeConfigObservation) DeepCopy() *NodeConfigObservation {
	if in == nil {
		return nil
	}
	out := new(NodeConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfigParameters) DeepCopyInto(out *NodeConfigParameters) {
	*out = *in
	if in.CPUCount != nil {
		in, out := &in.CPUCount, &out.CPUCount
		*out = new(float64)
		**out = **in
	}
	if in.MemorySizeMb != nil {
		in, out := &in.MemorySizeMb, &out.MemorySizeMb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfigParameters.
func (in *NodeConfigParameters) DeepCopy() *NodeConfigParameters {
	if in == nil {
		return nil
	}
	out := new(NodeConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartTimeInitParameters) DeepCopyInto(out *StartTimeInitParameters) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = new(float64)
		**out = **in
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = new(float64)
		**out = **in
	}
	if in.Nanos != nil {
		in, out := &in.Nanos, &out.Nanos
		*out = new(float64)
		**out = **in
	}
	if in.Seconds != nil {
		in, out := &in.Seconds, &out.Seconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartTimeInitParameters.
func (in *StartTimeInitParameters) DeepCopy() *StartTimeInitParameters {
	if in == nil {
		return nil
	}
	out := new(StartTimeInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartTimeObservation) DeepCopyInto(out *StartTimeObservation) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = new(float64)
		**out = **in
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = new(float64)
		**out = **in
	}
	if in.Nanos != nil {
		in, out := &in.Nanos, &out.Nanos
		*out = new(float64)
		**out = **in
	}
	if in.Seconds != nil {
		in, out := &in.Seconds, &out.Seconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartTimeObservation.
func (in *StartTimeObservation) DeepCopy() *StartTimeObservation {
	if in == nil {
		return nil
	}
	out := new(StartTimeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartTimeParameters) DeepCopyInto(out *StartTimeParameters) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = new(float64)
		**out = **in
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = new(float64)
		**out = **in
	}
	if in.Nanos != nil {
		in, out := &in.Nanos, &out.Nanos
		*out = new(float64)
		**out = **in
	}
	if in.Seconds != nil {
		in, out := &in.Seconds, &out.Seconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartTimeParameters.
func (in *StartTimeParameters) DeepCopy() *StartTimeParameters {
	if in == nil {
		return nil
	}
	out := new(StartTimeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyMaintenanceWindowInitParameters) DeepCopyInto(out *WeeklyMaintenanceWindowInitParameters) {
	*out = *in
	if in.Day != nil {
		in, out := &in.Day, &out.Day
		*out = new(string)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = make([]StartTimeInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyMaintenanceWindowInitParameters.
func (in *WeeklyMaintenanceWindowInitParameters) DeepCopy() *WeeklyMaintenanceWindowInitParameters {
	if in == nil {
		return nil
	}
	out := new(WeeklyMaintenanceWindowInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyMaintenanceWindowObservation) DeepCopyInto(out *WeeklyMaintenanceWindowObservation) {
	*out = *in
	if in.Day != nil {
		in, out := &in.Day, &out.Day
		*out = new(string)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = make([]StartTimeObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyMaintenanceWindowObservation.
func (in *WeeklyMaintenanceWindowObservation) DeepCopy() *WeeklyMaintenanceWindowObservation {
	if in == nil {
		return nil
	}
	out := new(WeeklyMaintenanceWindowObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeeklyMaintenanceWindowParameters) DeepCopyInto(out *WeeklyMaintenanceWindowParameters) {
	*out = *in
	if in.Day != nil {
		in, out := &in.Day, &out.Day
		*out = new(string)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = make([]StartTimeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeeklyMaintenanceWindowParameters.
func (in *WeeklyMaintenanceWindowParameters) DeepCopy() *WeeklyMaintenanceWindowParameters {
	if in == nil {
		return nil
	}
	out := new(WeeklyMaintenanceWindowParameters)
	in.DeepCopyInto(out)
	return out
}
