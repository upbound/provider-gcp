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

type GeoObservation struct {
}

type GeoParameters struct {

	// The location name defined in Google Cloud.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Same as rrdatas above.
	// +kubebuilder:validation:Required
	Rrdatas []*string `json:"rrdatas" tf:"rrdatas,omitempty"`
}

type RecordSetObservation struct {

	// an identifier for the resource with format projects/{{project}}/managedZones/{{zone}}/rrsets/{{name}}/{{type}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RecordSetParameters struct {

	// The name of the zone in which this record set will
	// reside.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/dns/v1beta1.ManagedZone
	// +kubebuilder:validation:Optional
	ManagedZone *string `json:"managedZone,omitempty" tf:"managed_zone,omitempty"`

	// Reference to a ManagedZone in dns to populate managedZone.
	// +kubebuilder:validation:Optional
	ManagedZoneRef *v1.Reference `json:"managedZoneRef,omitempty" tf:"-"`

	// Selector for a ManagedZone in dns to populate managedZone.
	// +kubebuilder:validation:Optional
	ManagedZoneSelector *v1.Selector `json:"managedZoneSelector,omitempty" tf:"-"`

	// The DNS name this record set will apply to.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The configuration for steering traffic based on query.
	// Now you can specify either Weighted Round Robin(WRR) type or Geolocation(GEO) type.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	RoutingPolicy []RoutingPolicyParameters `json:"routingPolicy,omitempty" tf:"routing_policy,omitempty"`

	// The string data for the records in this record set
	// whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding \" if you don't want your string to get split on spaces.g. "first255characters\" \"morecharacters").
	// +kubebuilder:validation:Optional
	Rrdatas []*string `json:"rrdatas,omitempty" tf:"rrdatas,omitempty"`

	// The time-to-live of this record set (seconds).
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The DNS record set type.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type RoutingPolicyObservation struct {
}

type RoutingPolicyParameters struct {

	// The configuration for Geolocation based routing policy.
	// Structure is document below.
	// +kubebuilder:validation:Optional
	Geo []GeoParameters `json:"geo,omitempty" tf:"geo,omitempty"`

	// The configuration for Weighted Round Robin based routing policy.
	// Structure is document below.
	// +kubebuilder:validation:Optional
	Wrr []WrrParameters `json:"wrr,omitempty" tf:"wrr,omitempty"`
}

type WrrObservation struct {
}

type WrrParameters struct {

	// Same as rrdatas above.
	// +kubebuilder:validation:Required
	Rrdatas []*string `json:"rrdatas" tf:"rrdatas,omitempty"`

	// The ratio of traffic routed to the target.
	// +kubebuilder:validation:Required
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

// RecordSetSpec defines the desired state of RecordSet
type RecordSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordSetParameters `json:"forProvider"`
}

// RecordSetStatus defines the observed state of RecordSet.
type RecordSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RecordSet is the Schema for the RecordSets API. Manages a set of DNS records within Google Cloud DNS.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RecordSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecordSetSpec   `json:"spec"`
	Status            RecordSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordSetList contains a list of RecordSets
type RecordSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RecordSet `json:"items"`
}

// Repository type metadata.
var (
	RecordSet_Kind             = "RecordSet"
	RecordSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RecordSet_Kind}.String()
	RecordSet_KindAPIVersion   = RecordSet_Kind + "." + CRDGroupVersion.String()
	RecordSet_GroupVersionKind = CRDGroupVersion.WithKind(RecordSet_Kind)
)

func init() {
	SchemeBuilder.Register(&RecordSet{}, &RecordSetList{})
}
