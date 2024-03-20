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

type EntitiesInitParameters struct {

	// A collection of value synonyms. For example, if the entity type is vegetable, and value is scallions, a synonym could be green onions.
	// For KIND_LIST entity types: This collection must contain exactly one synonym equal to value.
	Synonyms []*string `json:"synonyms,omitempty" tf:"synonyms,omitempty"`

	// The primary value associated with this entity entry. For example, if the entity type is vegetable, the value could be scallions.
	// For KIND_MAP entity types: A canonical value to be used in place of synonyms.
	// For KIND_LIST entity types: A string that can contain references to other entity types (with or without aliases).
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EntitiesObservation struct {

	// A collection of value synonyms. For example, if the entity type is vegetable, and value is scallions, a synonym could be green onions.
	// For KIND_LIST entity types: This collection must contain exactly one synonym equal to value.
	Synonyms []*string `json:"synonyms,omitempty" tf:"synonyms,omitempty"`

	// The primary value associated with this entity entry. For example, if the entity type is vegetable, the value could be scallions.
	// For KIND_MAP entity types: A canonical value to be used in place of synonyms.
	// For KIND_LIST entity types: A string that can contain references to other entity types (with or without aliases).
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EntitiesParameters struct {

	// A collection of value synonyms. For example, if the entity type is vegetable, and value is scallions, a synonym could be green onions.
	// For KIND_LIST entity types: This collection must contain exactly one synonym equal to value.
	// +kubebuilder:validation:Optional
	Synonyms []*string `json:"synonyms,omitempty" tf:"synonyms,omitempty"`

	// The primary value associated with this entity entry. For example, if the entity type is vegetable, the value could be scallions.
	// For KIND_MAP entity types: A canonical value to be used in place of synonyms.
	// For KIND_LIST entity types: A string that can contain references to other entity types (with or without aliases).
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EntityTypeInitParameters struct {

	// Represents kinds of entities.
	AutoExpansionMode *string `json:"autoExpansionMode,omitempty" tf:"auto_expansion_mode,omitempty"`

	// The human-readable name of the entity type, unique within the agent.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction *bool `json:"enableFuzzyExtraction,omitempty" tf:"enable_fuzzy_extraction,omitempty"`

	// The collection of entity entries associated with the entity type.
	// Structure is documented below.
	Entities []EntitiesInitParameters `json:"entities,omitempty" tf:"entities,omitempty"`

	// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion.
	// If the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.
	// Structure is documented below.
	ExcludedPhrases []ExcludedPhrasesInitParameters `json:"excludedPhrases,omitempty" tf:"excluded_phrases,omitempty"`

	// Indicates whether the entity type can be automatically expanded.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The language of the following fields in entityType:
	// EntityType.entities.value
	// EntityType.entities.synonyms
	// EntityType.excluded_phrases.value
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// The agent to create a entity type for.
	// Format: projects//locations//agents/.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/dialogflowcx/v1beta1.Agent
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Reference to a Agent in dialogflowcx to populate parent.
	// +kubebuilder:validation:Optional
	ParentRef *v1.Reference `json:"parentRef,omitempty" tf:"-"`

	// Selector for a Agent in dialogflowcx to populate parent.
	// +kubebuilder:validation:Optional
	ParentSelector *v1.Selector `json:"parentSelector,omitempty" tf:"-"`

	// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
	Redact *bool `json:"redact,omitempty" tf:"redact,omitempty"`
}

type EntityTypeObservation struct {

	// Represents kinds of entities.
	AutoExpansionMode *string `json:"autoExpansionMode,omitempty" tf:"auto_expansion_mode,omitempty"`

	// The human-readable name of the entity type, unique within the agent.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction *bool `json:"enableFuzzyExtraction,omitempty" tf:"enable_fuzzy_extraction,omitempty"`

	// The collection of entity entries associated with the entity type.
	// Structure is documented below.
	Entities []EntitiesObservation `json:"entities,omitempty" tf:"entities,omitempty"`

	// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion.
	// If the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.
	// Structure is documented below.
	ExcludedPhrases []ExcludedPhrasesObservation `json:"excludedPhrases,omitempty" tf:"excluded_phrases,omitempty"`

	// an identifier for the resource with format {{parent}}/entityTypes/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates whether the entity type can be automatically expanded.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The language of the following fields in entityType:
	// EntityType.entities.value
	// EntityType.entities.synonyms
	// EntityType.excluded_phrases.value
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// The unique identifier of the entity type.
	// Format: projects//locations//agents//entityTypes/.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The agent to create a entity type for.
	// Format: projects//locations//agents/.
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
	Redact *bool `json:"redact,omitempty" tf:"redact,omitempty"`
}

type EntityTypeParameters struct {

	// Represents kinds of entities.
	// +kubebuilder:validation:Optional
	AutoExpansionMode *string `json:"autoExpansionMode,omitempty" tf:"auto_expansion_mode,omitempty"`

	// The human-readable name of the entity type, unique within the agent.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Enables fuzzy entity extraction during classification.
	// +kubebuilder:validation:Optional
	EnableFuzzyExtraction *bool `json:"enableFuzzyExtraction,omitempty" tf:"enable_fuzzy_extraction,omitempty"`

	// The collection of entity entries associated with the entity type.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Entities []EntitiesParameters `json:"entities,omitempty" tf:"entities,omitempty"`

	// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion.
	// If the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ExcludedPhrases []ExcludedPhrasesParameters `json:"excludedPhrases,omitempty" tf:"excluded_phrases,omitempty"`

	// Indicates whether the entity type can be automatically expanded.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// The language of the following fields in entityType:
	// EntityType.entities.value
	// EntityType.entities.synonyms
	// EntityType.excluded_phrases.value
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	// +kubebuilder:validation:Optional
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// The agent to create a entity type for.
	// Format: projects//locations//agents/.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/dialogflowcx/v1beta1.Agent
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Reference to a Agent in dialogflowcx to populate parent.
	// +kubebuilder:validation:Optional
	ParentRef *v1.Reference `json:"parentRef,omitempty" tf:"-"`

	// Selector for a Agent in dialogflowcx to populate parent.
	// +kubebuilder:validation:Optional
	ParentSelector *v1.Selector `json:"parentSelector,omitempty" tf:"-"`

	// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
	// +kubebuilder:validation:Optional
	Redact *bool `json:"redact,omitempty" tf:"redact,omitempty"`
}

type ExcludedPhrasesInitParameters struct {

	// The word or phrase to be excluded.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ExcludedPhrasesObservation struct {

	// The word or phrase to be excluded.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ExcludedPhrasesParameters struct {

	// The word or phrase to be excluded.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// EntityTypeSpec defines the desired state of EntityType
type EntityTypeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EntityTypeParameters `json:"forProvider"`
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
	InitProvider EntityTypeInitParameters `json:"initProvider,omitempty"`
}

// EntityTypeStatus defines the observed state of EntityType.
type EntityTypeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EntityTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EntityType is the Schema for the EntityTypes API. Entities are extracted from user input and represent parameters that are meaningful to your application.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type EntityType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.entities) || (has(self.initProvider) && has(self.initProvider.entities))",message="spec.forProvider.entities is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.kind) || (has(self.initProvider) && has(self.initProvider.kind))",message="spec.forProvider.kind is a required parameter"
	Spec   EntityTypeSpec   `json:"spec"`
	Status EntityTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EntityTypeList contains a list of EntityTypes
type EntityTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EntityType `json:"items"`
}

// Repository type metadata.
var (
	EntityType_Kind             = "EntityType"
	EntityType_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EntityType_Kind}.String()
	EntityType_KindAPIVersion   = EntityType_Kind + "." + CRDGroupVersion.String()
	EntityType_GroupVersionKind = CRDGroupVersion.WithKind(EntityType_Kind)
)

func init() {
	SchemeBuilder.Register(&EntityType{}, &EntityTypeList{})
}
