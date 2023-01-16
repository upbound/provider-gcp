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

type TagTemplateIAMPolicyObservation struct {

	// (Computed) The etag of the IAM policy.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TagTemplateIAMPolicyParameters struct {

	// The policy data generated by
	// a google_iam_policy data source.
	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Used to find the parent resource to bind the IAM policy to
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/datacatalog/v1beta1.TagTemplate
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	TagTemplate *string `json:"tagTemplate,omitempty" tf:"tag_template,omitempty"`

	// Reference to a TagTemplate in datacatalog to populate tagTemplate.
	// +kubebuilder:validation:Optional
	TagTemplateRef *v1.Reference `json:"tagTemplateRef,omitempty" tf:"-"`

	// Selector for a TagTemplate in datacatalog to populate tagTemplate.
	// +kubebuilder:validation:Optional
	TagTemplateSelector *v1.Selector `json:"tagTemplateSelector,omitempty" tf:"-"`
}

// TagTemplateIAMPolicySpec defines the desired state of TagTemplateIAMPolicy
type TagTemplateIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TagTemplateIAMPolicyParameters `json:"forProvider"`
}

// TagTemplateIAMPolicyStatus defines the observed state of TagTemplateIAMPolicy.
type TagTemplateIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TagTemplateIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TagTemplateIAMPolicy is the Schema for the TagTemplateIAMPolicys API. Collection of resources to manage IAM policy for Data catalog TagTemplate
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TagTemplateIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TagTemplateIAMPolicySpec   `json:"spec"`
	Status            TagTemplateIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TagTemplateIAMPolicyList contains a list of TagTemplateIAMPolicys
type TagTemplateIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TagTemplateIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	TagTemplateIAMPolicy_Kind             = "TagTemplateIAMPolicy"
	TagTemplateIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TagTemplateIAMPolicy_Kind}.String()
	TagTemplateIAMPolicy_KindAPIVersion   = TagTemplateIAMPolicy_Kind + "." + CRDGroupVersion.String()
	TagTemplateIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(TagTemplateIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&TagTemplateIAMPolicy{}, &TagTemplateIAMPolicyList{})
}
