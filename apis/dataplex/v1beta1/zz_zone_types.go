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

type DiscoverySpecCsvOptionsInitParameters struct {

	// Optional. The delimiter being used to separate values. This defaults to ','.
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// Optional. Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
	DisableTypeInference *bool `json:"disableTypeInference,omitempty" tf:"disable_type_inference,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// Optional. The number of rows to interpret as header rows that should be skipped when reading data rows.
	HeaderRows *float64 `json:"headerRows,omitempty" tf:"header_rows,omitempty"`
}

type DiscoverySpecCsvOptionsObservation struct {

	// Optional. The delimiter being used to separate values. This defaults to ','.
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// Optional. Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
	DisableTypeInference *bool `json:"disableTypeInference,omitempty" tf:"disable_type_inference,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// Optional. The number of rows to interpret as header rows that should be skipped when reading data rows.
	HeaderRows *float64 `json:"headerRows,omitempty" tf:"header_rows,omitempty"`
}

type DiscoverySpecCsvOptionsParameters struct {

	// Optional. The delimiter being used to separate values. This defaults to ','.
	// +kubebuilder:validation:Optional
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// Optional. Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
	// +kubebuilder:validation:Optional
	DisableTypeInference *bool `json:"disableTypeInference,omitempty" tf:"disable_type_inference,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// Optional. The number of rows to interpret as header rows that should be skipped when reading data rows.
	// +kubebuilder:validation:Optional
	HeaderRows *float64 `json:"headerRows,omitempty" tf:"header_rows,omitempty"`
}

type DiscoverySpecJSONOptionsInitParameters struct {

	// Optional. Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
	DisableTypeInference *bool `json:"disableTypeInference,omitempty" tf:"disable_type_inference,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`
}

type DiscoverySpecJSONOptionsObservation struct {

	// Optional. Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
	DisableTypeInference *bool `json:"disableTypeInference,omitempty" tf:"disable_type_inference,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`
}

type DiscoverySpecJSONOptionsParameters struct {

	// Optional. Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
	// +kubebuilder:validation:Optional
	DisableTypeInference *bool `json:"disableTypeInference,omitempty" tf:"disable_type_inference,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`
}

type ZoneAssetStatusInitParameters struct {
}

type ZoneAssetStatusObservation struct {
	ActiveAssets *float64 `json:"activeAssets,omitempty" tf:"active_assets,omitempty"`

	SecurityPolicyApplyingAssets *float64 `json:"securityPolicyApplyingAssets,omitempty" tf:"security_policy_applying_assets,omitempty"`

	// Output only. The time when the zone was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type ZoneAssetStatusParameters struct {
}

type ZoneDiscoverySpecInitParameters struct {

	// Optional. Configuration for CSV data.
	CsvOptions []DiscoverySpecCsvOptionsInitParameters `json:"csvOptions,omitempty" tf:"csv_options,omitempty"`

	// Required. Whether discovery is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Optional. The list of patterns to apply for selecting data to exclude during discovery. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	ExcludePatterns []*string `json:"excludePatterns,omitempty" tf:"exclude_patterns,omitempty"`

	// Optional. The list of patterns to apply for selecting data to include during discovery if only a subset of the data should considered. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	IncludePatterns []*string `json:"includePatterns,omitempty" tf:"include_patterns,omitempty"`

	// Optional. Configuration for Json data.
	JSONOptions []DiscoverySpecJSONOptionsInitParameters `json:"jsonOptions,omitempty" tf:"json_options,omitempty"`

	// Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron) for running discovery periodically. Successive discovery runs must be scheduled at least 60 minutes apart. The default value is to run discovery every 60 minutes. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone database. For example, "CRON_TZ=America/New_York 1 * * * *", or "TZ=America/New_York 1 * * * *".
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type ZoneDiscoverySpecObservation struct {

	// Optional. Configuration for CSV data.
	CsvOptions []DiscoverySpecCsvOptionsObservation `json:"csvOptions,omitempty" tf:"csv_options,omitempty"`

	// Required. Whether discovery is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Optional. The list of patterns to apply for selecting data to exclude during discovery. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	ExcludePatterns []*string `json:"excludePatterns,omitempty" tf:"exclude_patterns,omitempty"`

	// Optional. The list of patterns to apply for selecting data to include during discovery if only a subset of the data should considered. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	IncludePatterns []*string `json:"includePatterns,omitempty" tf:"include_patterns,omitempty"`

	// Optional. Configuration for Json data.
	JSONOptions []DiscoverySpecJSONOptionsObservation `json:"jsonOptions,omitempty" tf:"json_options,omitempty"`

	// Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron) for running discovery periodically. Successive discovery runs must be scheduled at least 60 minutes apart. The default value is to run discovery every 60 minutes. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone database. For example, "CRON_TZ=America/New_York 1 * * * *", or "TZ=America/New_York 1 * * * *".
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type ZoneDiscoverySpecParameters struct {

	// Optional. Configuration for CSV data.
	// +kubebuilder:validation:Optional
	CsvOptions []DiscoverySpecCsvOptionsParameters `json:"csvOptions,omitempty" tf:"csv_options,omitempty"`

	// Required. Whether discovery is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Optional. The list of patterns to apply for selecting data to exclude during discovery. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	// +kubebuilder:validation:Optional
	ExcludePatterns []*string `json:"excludePatterns,omitempty" tf:"exclude_patterns,omitempty"`

	// Optional. The list of patterns to apply for selecting data to include during discovery if only a subset of the data should considered. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	// +kubebuilder:validation:Optional
	IncludePatterns []*string `json:"includePatterns,omitempty" tf:"include_patterns,omitempty"`

	// Optional. Configuration for Json data.
	// +kubebuilder:validation:Optional
	JSONOptions []DiscoverySpecJSONOptionsParameters `json:"jsonOptions,omitempty" tf:"json_options,omitempty"`

	// Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron) for running discovery periodically. Successive discovery runs must be scheduled at least 60 minutes apart. The default value is to run discovery every 60 minutes. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone database. For example, "CRON_TZ=America/New_York 1 * * * *", or "TZ=America/New_York 1 * * * *".
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type ZoneInitParameters struct {

	// Optional. Description of the zone.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Required. Specification of the discovery feature applied to data in this zone.
	DiscoverySpec []ZoneDiscoverySpecInitParameters `json:"discoverySpec,omitempty" tf:"discovery_spec,omitempty"`

	// Optional. User friendly display name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Optional. User defined labels for the zone.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
	ResourceSpec []ZoneResourceSpecInitParameters `json:"resourceSpec,omitempty" tf:"resource_spec,omitempty"`

	// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ZoneObservation struct {

	// Output only. Aggregated status of the underlying assets of the zone.
	AssetStatus []ZoneAssetStatusObservation `json:"assetStatus,omitempty" tf:"asset_status,omitempty"`

	// Output only. The time when the zone was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Optional. Description of the zone.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Required. Specification of the discovery feature applied to data in this zone.
	DiscoverySpec []ZoneDiscoverySpecObservation `json:"discoverySpec,omitempty" tf:"discovery_spec,omitempty"`

	// Optional. User friendly display name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Optional. User defined labels for the zone.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The lake for the resource
	Lake *string `json:"lake,omitempty" tf:"lake,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
	ResourceSpec []ZoneResourceSpecObservation `json:"resourceSpec,omitempty" tf:"resource_spec,omitempty"`

	// Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// Output only. The time when the zone was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type ZoneParameters struct {

	// Optional. Description of the zone.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Required. Specification of the discovery feature applied to data in this zone.
	// +kubebuilder:validation:Optional
	DiscoverySpec []ZoneDiscoverySpecParameters `json:"discoverySpec,omitempty" tf:"discovery_spec,omitempty"`

	// Optional. User friendly display name.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Optional. User defined labels for the zone.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The lake for the resource
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/dataplex/v1beta1.Lake
	// +kubebuilder:validation:Optional
	Lake *string `json:"lake,omitempty" tf:"lake,omitempty"`

	// Reference to a Lake in dataplex to populate lake.
	// +kubebuilder:validation:Optional
	LakeRef *v1.Reference `json:"lakeRef,omitempty" tf:"-"`

	// Selector for a Lake in dataplex to populate lake.
	// +kubebuilder:validation:Optional
	LakeSelector *v1.Selector `json:"lakeSelector,omitempty" tf:"-"`

	// The location for the resource
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
	// +kubebuilder:validation:Optional
	ResourceSpec []ZoneResourceSpecParameters `json:"resourceSpec,omitempty" tf:"resource_spec,omitempty"`

	// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ZoneResourceSpecInitParameters struct {

	// Required. Immutable. The location type of the resources that are allowed to be attached to the assets within this zone. Possible values: LOCATION_TYPE_UNSPECIFIED, SINGLE_REGION, MULTI_REGION
	LocationType *string `json:"locationType,omitempty" tf:"location_type,omitempty"`
}

type ZoneResourceSpecObservation struct {

	// Required. Immutable. The location type of the resources that are allowed to be attached to the assets within this zone. Possible values: LOCATION_TYPE_UNSPECIFIED, SINGLE_REGION, MULTI_REGION
	LocationType *string `json:"locationType,omitempty" tf:"location_type,omitempty"`
}

type ZoneResourceSpecParameters struct {

	// Required. Immutable. The location type of the resources that are allowed to be attached to the assets within this zone. Possible values: LOCATION_TYPE_UNSPECIFIED, SINGLE_REGION, MULTI_REGION
	// +kubebuilder:validation:Optional
	LocationType *string `json:"locationType" tf:"location_type,omitempty"`
}

// ZoneSpec defines the desired state of Zone
type ZoneSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ZoneParameters `json:"forProvider"`
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
	InitProvider ZoneInitParameters `json:"initProvider,omitempty"`
}

// ZoneStatus defines the observed state of Zone.
type ZoneStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Zone is the Schema for the Zones API. The Dataplex Zone resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Zone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.discoverySpec) || (has(self.initProvider) && has(self.initProvider.discoverySpec))",message="spec.forProvider.discoverySpec is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceSpec) || (has(self.initProvider) && has(self.initProvider.resourceSpec))",message="spec.forProvider.resourceSpec is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   ZoneSpec   `json:"spec"`
	Status ZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneList contains a list of Zones
type ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Zone `json:"items"`
}

// Repository type metadata.
var (
	Zone_Kind             = "Zone"
	Zone_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Zone_Kind}.String()
	Zone_KindAPIVersion   = Zone_Kind + "." + CRDGroupVersion.String()
	Zone_GroupVersionKind = CRDGroupVersion.WithKind(Zone_Kind)
)

func init() {
	SchemeBuilder.Register(&Zone{}, &ZoneList{})
}
