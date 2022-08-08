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

type DNSSECConfigObservation struct {
}

type DNSSECConfigParameters struct {

	// Specifies parameters that will be used for generating initial DnsKeys
	// for this ManagedZone. If you provide a spec for keySigning or zoneSigning,
	// you must also provide one for the other.
	// default_key_specs can only be updated when the state is 'off'.
	// +kubebuilder:validation:Optional
	DefaultKeySpecs []DefaultKeySpecsParameters `json:"defaultKeySpecs,omitempty" tf:"default_key_specs,omitempty"`

	// Identifies what kind of resource this is
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the mechanism used to provide authenticated denial-of-existence responses.
	// non_existence can only be updated when the state is 'off'. Possible values: ["nsec", "nsec3"]
	// +kubebuilder:validation:Optional
	NonExistence *string `json:"nonExistence,omitempty" tf:"non_existence,omitempty"`

	// Specifies whether DNSSEC is enabled, and what mode it is in Possible values: ["off", "on", "transfer"]
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type DefaultKeySpecsObservation struct {
}

type DefaultKeySpecsParameters struct {

	// String mnemonic specifying the DNSSEC algorithm of this key Possible values: ["ecdsap256sha256", "ecdsap384sha384", "rsasha1", "rsasha256", "rsasha512"]
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Length of the keys in bits
	// +kubebuilder:validation:Optional
	KeyLength *float64 `json:"keyLength,omitempty" tf:"key_length,omitempty"`

	// Specifies whether this is a key signing key (KSK) or a zone
	// signing key (ZSK). Key signing keys have the Secure Entry
	// Point flag set and, when active, will only be used to sign
	// resource record sets of type DNSKEY. Zone signing keys do
	// not have the Secure Entry Point flag set and will be used
	// to sign all other types of resource record sets. Possible values: ["keySigning", "zoneSigning"]
	// +kubebuilder:validation:Optional
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// Identifies what kind of resource this is
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`
}

type ForwardingConfigObservation struct {
}

type ForwardingConfigParameters struct {

	// List of target name servers to forward to. Cloud DNS will
	// select the best available name server if more than
	// one target is given.
	// +kubebuilder:validation:Required
	TargetNameServers []TargetNameServersParameters `json:"targetNameServers" tf:"target_name_servers,omitempty"`
}

type ManagedZoneObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	NameServers []*string `json:"nameServers,omitempty" tf:"name_servers,omitempty"`
}

type ManagedZoneParameters struct {

	// The DNS name of this managed zone, for instance "example.com.".
	// +kubebuilder:validation:Required
	DNSName *string `json:"dnsName" tf:"dns_name,omitempty"`

	// DNSSEC configuration
	// +kubebuilder:validation:Optional
	DNSSECConfig []DNSSECConfigParameters `json:"dnssecConfig,omitempty" tf:"dnssec_config,omitempty"`

	// A textual description field. Defaults to 'Managed by Terraform'.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.
	// +kubebuilder:validation:Optional
	ForwardingConfig []ForwardingConfigParameters `json:"forwardingConfig,omitempty" tf:"forwarding_config,omitempty"`

	// A set of key/value label pairs to assign to this ManagedZone.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.
	// +kubebuilder:validation:Optional
	PeeringConfig []PeeringConfigParameters `json:"peeringConfig,omitempty" tf:"peering_config,omitempty"`

	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from.
	// +kubebuilder:validation:Optional
	PrivateVisibilityConfig []PrivateVisibilityConfigParameters `json:"privateVisibilityConfig,omitempty" tf:"private_visibility_config,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources. Default value: "public" Possible values: ["private", "public"]
	// +kubebuilder:validation:Optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

type NetworksObservation struct {
}

type NetworksParameters struct {

	// The id or fully qualified URL of the VPC network to bind to.
	// This should be formatted like 'projects/{project}/global/networks/{network}' or
	// 'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'
	// +kubebuilder:validation:Required
	NetworkURL *string `json:"networkUrl" tf:"network_url,omitempty"`
}

type PeeringConfigObservation struct {
}

type PeeringConfigParameters struct {

	// The network with which to peer.
	// +kubebuilder:validation:Required
	TargetNetwork []TargetNetworkParameters `json:"targetNetwork" tf:"target_network,omitempty"`
}

type PrivateVisibilityConfigObservation struct {
}

type PrivateVisibilityConfigParameters struct {

	// The list of VPC networks that can see this zone. Until the provider updates to use the Terraform 0.12 SDK in a future release, you
	// may experience issues with this resource while updating. If you've defined a 'networks' block and
	// add another 'networks' block while keeping the old block, Terraform will see an incorrect diff
	// and apply an incorrect update to the resource. If you encounter this issue, remove all 'networks'
	// blocks in an update and then apply another update adding all of them back simultaneously.
	// +kubebuilder:validation:Required
	Networks []NetworksParameters `json:"networks" tf:"networks,omitempty"`
}

type TargetNameServersObservation struct {
}

type TargetNameServersParameters struct {

	// Forwarding path for this TargetNameServer. If unset or 'default' Cloud DNS will make forwarding
	// decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
	// to the Internet. When set to 'private', Cloud DNS will always send queries through VPC for this target Possible values: ["default", "private"]
	// +kubebuilder:validation:Optional
	ForwardingPath *string `json:"forwardingPath,omitempty" tf:"forwarding_path,omitempty"`

	// IPv4 address of a target name server.
	// +kubebuilder:validation:Required
	IPv4Address *string `json:"ipv4Address" tf:"ipv4_address,omitempty"`
}

type TargetNetworkObservation struct {
}

type TargetNetworkParameters struct {

	// The id or fully qualified URL of the VPC network to forward queries to.
	// This should be formatted like 'projects/{project}/global/networks/{network}' or
	// 'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'
	// +kubebuilder:validation:Required
	NetworkURL *string `json:"networkUrl" tf:"network_url,omitempty"`
}

// ManagedZoneSpec defines the desired state of ManagedZone
type ManagedZoneSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedZoneParameters `json:"forProvider"`
}

// ManagedZoneStatus defines the observed state of ManagedZone.
type ManagedZoneStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedZone is the Schema for the ManagedZones API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ManagedZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedZoneSpec   `json:"spec"`
	Status            ManagedZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedZoneList contains a list of ManagedZones
type ManagedZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedZone `json:"items"`
}

// Repository type metadata.
var (
	ManagedZone_Kind             = "ManagedZone"
	ManagedZone_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedZone_Kind}.String()
	ManagedZone_KindAPIVersion   = ManagedZone_Kind + "." + CRDGroupVersion.String()
	ManagedZone_GroupVersionKind = CRDGroupVersion.WithKind(ManagedZone_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedZone{}, &ManagedZoneList{})
}