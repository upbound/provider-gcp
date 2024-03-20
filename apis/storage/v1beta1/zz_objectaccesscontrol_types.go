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

type ObjectAccessControlInitParameters struct {

	// The name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.Bucket
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in storage to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in storage to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The entity holding the permission, in one of the following forms:
	Entity *string `json:"entity,omitempty" tf:"entity,omitempty"`

	// The name of the object to apply the access control to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.BucketObject
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("output_name",true)
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// Reference to a BucketObject in storage to populate object.
	// +kubebuilder:validation:Optional
	ObjectRef *v1.Reference `json:"objectRef,omitempty" tf:"-"`

	// Selector for a BucketObject in storage to populate object.
	// +kubebuilder:validation:Optional
	ObjectSelector *v1.Selector `json:"objectSelector,omitempty" tf:"-"`

	// The access permission for the entity.
	// Possible values are: OWNER, READER.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type ObjectAccessControlObservation struct {

	// The name of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The domain associated with the entity.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The email address associated with the entity.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The entity holding the permission, in one of the following forms:
	Entity *string `json:"entity,omitempty" tf:"entity,omitempty"`

	// The ID for the entity
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// The content generation of the object, if applied to an object.
	Generation *float64 `json:"generation,omitempty" tf:"generation,omitempty"`

	// an identifier for the resource with format {{bucket}}/{{object}}/{{entity}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the object to apply the access control to.
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// The project team associated with the entity
	// Structure is documented below.
	ProjectTeam []ObjectAccessControlProjectTeamObservation `json:"projectTeam,omitempty" tf:"project_team,omitempty"`

	// The access permission for the entity.
	// Possible values are: OWNER, READER.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type ObjectAccessControlParameters struct {

	// The name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in storage to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in storage to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The entity holding the permission, in one of the following forms:
	// +kubebuilder:validation:Optional
	Entity *string `json:"entity,omitempty" tf:"entity,omitempty"`

	// The name of the object to apply the access control to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.BucketObject
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("output_name",true)
	// +kubebuilder:validation:Optional
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// Reference to a BucketObject in storage to populate object.
	// +kubebuilder:validation:Optional
	ObjectRef *v1.Reference `json:"objectRef,omitempty" tf:"-"`

	// Selector for a BucketObject in storage to populate object.
	// +kubebuilder:validation:Optional
	ObjectSelector *v1.Selector `json:"objectSelector,omitempty" tf:"-"`

	// The access permission for the entity.
	// Possible values are: OWNER, READER.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type ObjectAccessControlProjectTeamInitParameters struct {
}

type ObjectAccessControlProjectTeamObservation struct {

	// The project team associated with the entity
	ProjectNumber *string `json:"projectNumber,omitempty" tf:"project_number,omitempty"`

	// The team.
	// Possible values are: editors, owners, viewers.
	Team *string `json:"team,omitempty" tf:"team,omitempty"`
}

type ObjectAccessControlProjectTeamParameters struct {
}

// ObjectAccessControlSpec defines the desired state of ObjectAccessControl
type ObjectAccessControlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectAccessControlParameters `json:"forProvider"`
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
	InitProvider ObjectAccessControlInitParameters `json:"initProvider,omitempty"`
}

// ObjectAccessControlStatus defines the observed state of ObjectAccessControl.
type ObjectAccessControlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectAccessControlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ObjectAccessControl is the Schema for the ObjectAccessControls API. The ObjectAccessControls resources represent the Access Control Lists (ACLs) for objects within Google Cloud Storage.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ObjectAccessControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.entity) || (has(self.initProvider) && has(self.initProvider.entity))",message="spec.forProvider.entity is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   ObjectAccessControlSpec   `json:"spec"`
	Status ObjectAccessControlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectAccessControlList contains a list of ObjectAccessControls
type ObjectAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectAccessControl `json:"items"`
}

// Repository type metadata.
var (
	ObjectAccessControl_Kind             = "ObjectAccessControl"
	ObjectAccessControl_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ObjectAccessControl_Kind}.String()
	ObjectAccessControl_KindAPIVersion   = ObjectAccessControl_Kind + "." + CRDGroupVersion.String()
	ObjectAccessControl_GroupVersionKind = CRDGroupVersion.WithKind(ObjectAccessControl_Kind)
)

func init() {
	SchemeBuilder.Register(&ObjectAccessControl{}, &ObjectAccessControlList{})
}
