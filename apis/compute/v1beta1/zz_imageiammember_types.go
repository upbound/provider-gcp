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

type ImageIAMMemberConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ImageIAMMemberConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ImageIAMMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title" tf:"title,omitempty"`
}

type ImageIAMMemberInitParameters struct {
	Condition []ImageIAMMemberConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +crossplane:generate:reference:type=Image
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Reference to a Image to populate image.
	// +kubebuilder:validation:Optional
	ImageRef *v1.Reference `json:"imageRef,omitempty" tf:"-"`

	// Selector for a Image to populate image.
	// +kubebuilder:validation:Optional
	ImageSelector *v1.Selector `json:"imageSelector,omitempty" tf:"-"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type ImageIAMMemberObservation struct {
	Condition []ImageIAMMemberConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type ImageIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ImageIAMMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +crossplane:generate:reference:type=Image
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Reference to a Image to populate image.
	// +kubebuilder:validation:Optional
	ImageRef *v1.Reference `json:"imageRef,omitempty" tf:"-"`

	// Selector for a Image to populate image.
	// +kubebuilder:validation:Optional
	ImageSelector *v1.Selector `json:"imageSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

// ImageIAMMemberSpec defines the desired state of ImageIAMMember
type ImageIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageIAMMemberParameters `json:"forProvider"`
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
	InitProvider ImageIAMMemberInitParameters `json:"initProvider,omitempty"`
}

// ImageIAMMemberStatus defines the observed state of ImageIAMMember.
type ImageIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ImageIAMMember is the Schema for the ImageIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ImageIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.member) || (has(self.initProvider) && has(self.initProvider.member))",message="spec.forProvider.member is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   ImageIAMMemberSpec   `json:"spec"`
	Status ImageIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageIAMMemberList contains a list of ImageIAMMembers
type ImageIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageIAMMember `json:"items"`
}

// Repository type metadata.
var (
	ImageIAMMember_Kind             = "ImageIAMMember"
	ImageIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImageIAMMember_Kind}.String()
	ImageIAMMember_KindAPIVersion   = ImageIAMMember_Kind + "." + CRDGroupVersion.String()
	ImageIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(ImageIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&ImageIAMMember{}, &ImageIAMMemberList{})
}
