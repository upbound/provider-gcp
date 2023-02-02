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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AgentPoolObservation struct {

	// an identifier for the resource with format projects/{{project}}/agentPools/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the state of the AgentPool.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type AgentPoolParameters struct {

	// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	BandwidthLimit []BandwidthLimitParameters `json:"bandwidthLimit,omitempty" tf:"bandwidth_limit,omitempty"`

	// Specifies the client-specified AgentPool description.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type BandwidthLimitObservation struct {
}

type BandwidthLimitParameters struct {

	// Bandwidth rate in megabytes per second, distributed across all the agents in the pool.
	// +kubebuilder:validation:Required
	LimitMbps *string `json:"limitMbps" tf:"limit_mbps,omitempty"`
}

// AgentPoolSpec defines the desired state of AgentPool
type AgentPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AgentPoolParameters `json:"forProvider"`
}

// AgentPoolStatus defines the observed state of AgentPool.
type AgentPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AgentPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AgentPool is the Schema for the AgentPools API. Represents an On-Premises Agent pool.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AgentPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AgentPoolSpec   `json:"spec"`
	Status            AgentPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AgentPoolList contains a list of AgentPools
type AgentPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AgentPool `json:"items"`
}

// Repository type metadata.
var (
	AgentPool_Kind             = "AgentPool"
	AgentPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AgentPool_Kind}.String()
	AgentPool_KindAPIVersion   = AgentPool_Kind + "." + CRDGroupVersion.String()
	AgentPool_GroupVersionKind = CRDGroupVersion.WithKind(AgentPool_Kind)
)

func init() {
	SchemeBuilder.Register(&AgentPool{}, &AgentPoolList{})
}