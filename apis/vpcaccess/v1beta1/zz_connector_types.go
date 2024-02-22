// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type ConnectorInitParameters struct {

	// The range of internal addresses that follows RFC 4632 notation. Example: 10.132.0.0/28.
	IPCidrRange *string `json:"ipCidrRange,omitempty" tf:"ip_cidr_range,omitempty"`

	// Machine type of VM Instance underlying connector. Default is e2-micro
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// Maximum value of instances in autoscaling group underlying the connector.
	MaxInstances *float64 `json:"maxInstances,omitempty" tf:"max_instances,omitempty"`

	// Maximum throughput of the connector in Mbps, must be greater than min_throughput. Default is 300.
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`

	// Minimum value of instances in autoscaling group underlying the connector.
	MinInstances *float64 `json:"minInstances,omitempty" tf:"min_instances,omitempty"`

	// Minimum throughput of the connector in Mbps. Default and min is 200.
	MinThroughput *float64 `json:"minThroughput,omitempty" tf:"min_throughput,omitempty"`

	// Name or self_link of the VPC network. Required if ip_cidr_range is set.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The subnet in which to house the connector
	// Structure is documented below.
	Subnet []SubnetInitParameters `json:"subnet,omitempty" tf:"subnet,omitempty"`
}

type ConnectorObservation struct {

	// List of projects using the connector.
	ConnectedProjects []*string `json:"connectedProjects,omitempty" tf:"connected_projects,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{region}}/connectors/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The range of internal addresses that follows RFC 4632 notation. Example: 10.132.0.0/28.
	IPCidrRange *string `json:"ipCidrRange,omitempty" tf:"ip_cidr_range,omitempty"`

	// Machine type of VM Instance underlying connector. Default is e2-micro
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// Maximum value of instances in autoscaling group underlying the connector.
	MaxInstances *float64 `json:"maxInstances,omitempty" tf:"max_instances,omitempty"`

	// Maximum throughput of the connector in Mbps, must be greater than min_throughput. Default is 300.
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`

	// Minimum value of instances in autoscaling group underlying the connector.
	MinInstances *float64 `json:"minInstances,omitempty" tf:"min_instances,omitempty"`

	// Minimum throughput of the connector in Mbps. Default and min is 200.
	MinThroughput *float64 `json:"minThroughput,omitempty" tf:"min_throughput,omitempty"`

	// Name or self_link of the VPC network. Required if ip_cidr_range is set.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the VPC Access connector resides. If it is not provided, the provider region is used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The fully qualified name of this VPC connector
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// State of the VPC access connector.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The subnet in which to house the connector
	// Structure is documented below.
	Subnet []SubnetObservation `json:"subnet,omitempty" tf:"subnet,omitempty"`
}

type ConnectorParameters struct {

	// The range of internal addresses that follows RFC 4632 notation. Example: 10.132.0.0/28.
	// +kubebuilder:validation:Optional
	IPCidrRange *string `json:"ipCidrRange,omitempty" tf:"ip_cidr_range,omitempty"`

	// Machine type of VM Instance underlying connector. Default is e2-micro
	// +kubebuilder:validation:Optional
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// Maximum value of instances in autoscaling group underlying the connector.
	// +kubebuilder:validation:Optional
	MaxInstances *float64 `json:"maxInstances,omitempty" tf:"max_instances,omitempty"`

	// Maximum throughput of the connector in Mbps, must be greater than min_throughput. Default is 300.
	// +kubebuilder:validation:Optional
	MaxThroughput *float64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`

	// Minimum value of instances in autoscaling group underlying the connector.
	// +kubebuilder:validation:Optional
	MinInstances *float64 `json:"minInstances,omitempty" tf:"min_instances,omitempty"`

	// Minimum throughput of the connector in Mbps. Default and min is 200.
	// +kubebuilder:validation:Optional
	MinThroughput *float64 `json:"minThroughput,omitempty" tf:"min_throughput,omitempty"`

	// Name or self_link of the VPC network. Required if ip_cidr_range is set.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where the VPC Access connector resides. If it is not provided, the provider region is used.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The subnet in which to house the connector
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Subnet []SubnetParameters `json:"subnet,omitempty" tf:"subnet,omitempty"`
}

type SubnetInitParameters struct {

	// Subnet name (relative, not fully qualified). E.g. if the full subnet selfLink is
	// https://compute.googleapis.com/compute/v1/projects/{project}/regions/{region}/subnetworks/{subnetName} the correct input for this field would be {subnetName}"
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Subnetwork
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Subnetwork in compute to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Subnetwork in compute to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// Project in which the subnet exists. If not set, this project is assumed to be the project for which the connector create request was issued.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type SubnetObservation struct {

	// Subnet name (relative, not fully qualified). E.g. if the full subnet selfLink is
	// https://compute.googleapis.com/compute/v1/projects/{project}/regions/{region}/subnetworks/{subnetName} the correct input for this field would be {subnetName}"
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Project in which the subnet exists. If not set, this project is assumed to be the project for which the connector create request was issued.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type SubnetParameters struct {

	// Subnet name (relative, not fully qualified). E.g. if the full subnet selfLink is
	// https://compute.googleapis.com/compute/v1/projects/{project}/regions/{region}/subnetworks/{subnetName} the correct input for this field would be {subnetName}"
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Subnetwork
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Subnetwork in compute to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Subnetwork in compute to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// Project in which the subnet exists. If not set, this project is assumed to be the project for which the connector create request was issued.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

// ConnectorSpec defines the desired state of Connector
type ConnectorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectorParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ConnectorInitParameters `json:"initProvider,omitempty"`
}

// ConnectorStatus defines the observed state of Connector.
type ConnectorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Connector is the Schema for the Connectors API. Serverless VPC Access connector resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Connector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectorSpec   `json:"spec"`
	Status            ConnectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectorList contains a list of Connectors
type ConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connector `json:"items"`
}

// Repository type metadata.
var (
	Connector_Kind             = "Connector"
	Connector_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Connector_Kind}.String()
	Connector_KindAPIVersion   = Connector_Kind + "." + CRDGroupVersion.String()
	Connector_GroupVersionKind = CRDGroupVersion.WithKind(Connector_Kind)
)

func init() {
	SchemeBuilder.Register(&Connector{}, &ConnectorList{})
}
