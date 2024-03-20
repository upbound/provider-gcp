// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RegionSSLCertificateInitParameters struct {

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type RegionSSLCertificateObservation struct {

	// The unique identifier for the resource.
	CertificateID *float64 `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Expire time of the certificate in RFC3339 text format.
	ExpireTime *string `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/sslCertificates/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Region in which the created regional ssl certificate should reside.
	// If it is not provided, the provider region is used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type RegionSSLCertificateParameters struct {

	// The certificate in PEM format.
	// The certificate chain must be no greater than 5 certs long.
	// The chain must include at least one intermediate cert.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	CertificateSecretRef v1.SecretKeySelector `json:"certificateSecretRef" tf:"-"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The write-only private key in PEM format.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Region in which the created regional ssl certificate should reside.
	// If it is not provided, the provider region is used.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

// RegionSSLCertificateSpec defines the desired state of RegionSSLCertificate
type RegionSSLCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionSSLCertificateParameters `json:"forProvider"`
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
	InitProvider RegionSSLCertificateInitParameters `json:"initProvider,omitempty"`
}

// RegionSSLCertificateStatus defines the observed state of RegionSSLCertificate.
type RegionSSLCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionSSLCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RegionSSLCertificate is the Schema for the RegionSSLCertificates API. A RegionSslCertificate resource, used for HTTPS load balancing.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RegionSSLCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.certificateSecretRef)",message="spec.forProvider.certificateSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privateKeySecretRef)",message="spec.forProvider.privateKeySecretRef is a required parameter"
	Spec   RegionSSLCertificateSpec   `json:"spec"`
	Status RegionSSLCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionSSLCertificateList contains a list of RegionSSLCertificates
type RegionSSLCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionSSLCertificate `json:"items"`
}

// Repository type metadata.
var (
	RegionSSLCertificate_Kind             = "RegionSSLCertificate"
	RegionSSLCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionSSLCertificate_Kind}.String()
	RegionSSLCertificate_KindAPIVersion   = RegionSSLCertificate_Kind + "." + CRDGroupVersion.String()
	RegionSSLCertificate_GroupVersionKind = CRDGroupVersion.WithKind(RegionSSLCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionSSLCertificate{}, &RegionSSLCertificateList{})
}
