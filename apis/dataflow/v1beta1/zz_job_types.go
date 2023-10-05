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

type JobInitParameters struct {

	// List of experiments that should be used by the job. An example value is ["enable_stackdriver_agent_metrics"].
	AdditionalExperiments []*string `json:"additionalExperiments,omitempty" tf:"additional_experiments,omitempty"`

	// Enable/disable the use of Streaming Engine for the job. Note that Streaming Engine is enabled by default for pipelines developed against the Beam SDK for Python v2.21.0 or later when using Python 3.
	EnableStreamingEngine *bool `json:"enableStreamingEngine,omitempty" tf:"enable_streaming_engine,omitempty"`

	// The configuration for VM IPs.  Options are "WORKER_IP_PUBLIC" or "WORKER_IP_PRIVATE".
	IPConfiguration *string `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// The name for the Cloud KMS key for the job. Key format is: projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// User labels to be specified for the job. Keys and values should follow the restrictions
	// specified in the labeling restrictions page.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field effective_labels for all of the labels present on the resource.
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The machine type to use for the job.
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	MaxWorkers *float64 `json:"maxWorkers,omitempty" tf:"max_workers,omitempty"`

	// A unique name for the resource, required by Dataflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// One of "drain" or "cancel".  See above note.
	OnDelete *string `json:"onDelete,omitempty" tf:"on_delete,omitempty"`

	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region in which the created job should run.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The Service Account email used to create the job.
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// See above note.
	SkipWaitOnJobTermination *bool `json:"skipWaitOnJobTermination,omitempty" tf:"skip_wait_on_job_termination,omitempty"`

	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK". If the subnetwork is located in a Shared VPC network, you must use the complete URL. For example "googleapis.com/compute/v1/projects/PROJECT_ID/regions/REGION/subnetworks/SUBNET_NAME"
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// A writeable location on GCS for the Dataflow job to dump its temporary data.
	TempGcsLocation *string `json:"tempGcsLocation,omitempty" tf:"temp_gcs_location,omitempty"`

	// The GCS path to the Dataflow job template.
	TemplateGcsPath *string `json:"templateGcsPath,omitempty" tf:"template_gcs_path,omitempty"`

	// Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job. This field is not used outside of update.
	TransformNameMapping map[string]string `json:"transformNameMapping,omitempty" tf:"transform_name_mapping,omitempty"`

	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type JobObservation struct {

	// List of experiments that should be used by the job. An example value is ["enable_stackdriver_agent_metrics"].
	AdditionalExperiments []*string `json:"additionalExperiments,omitempty" tf:"additional_experiments,omitempty"`

	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// Enable/disable the use of Streaming Engine for the job. Note that Streaming Engine is enabled by default for pipelines developed against the Beam SDK for Python v2.21.0 or later when using Python 3.
	EnableStreamingEngine *bool `json:"enableStreamingEngine,omitempty" tf:"enable_streaming_engine,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The configuration for VM IPs.  Options are "WORKER_IP_PUBLIC" or "WORKER_IP_PRIVATE".
	IPConfiguration *string `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// The unique ID of this job.
	JobID *string `json:"jobId,omitempty" tf:"job_id,omitempty"`

	// The name for the Cloud KMS key for the job. Key format is: projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// User labels to be specified for the job. Keys and values should follow the restrictions
	// specified in the labeling restrictions page.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field effective_labels for all of the labels present on the resource.
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The machine type to use for the job.
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	MaxWorkers *float64 `json:"maxWorkers,omitempty" tf:"max_workers,omitempty"`

	// A unique name for the resource, required by Dataflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// One of "drain" or "cancel".  See above note.
	OnDelete *string `json:"onDelete,omitempty" tf:"on_delete,omitempty"`

	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region in which the created job should run.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The Service Account email used to create the job.
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// See above note.
	SkipWaitOnJobTermination *bool `json:"skipWaitOnJobTermination,omitempty" tf:"skip_wait_on_job_termination,omitempty"`

	// The current state of the resource, selected from the JobState enum
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK". If the subnetwork is located in a Shared VPC network, you must use the complete URL. For example "googleapis.com/compute/v1/projects/PROJECT_ID/regions/REGION/subnetworks/SUBNET_NAME"
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// A writeable location on GCS for the Dataflow job to dump its temporary data.
	TempGcsLocation *string `json:"tempGcsLocation,omitempty" tf:"temp_gcs_location,omitempty"`

	// The GCS path to the Dataflow job template.
	TemplateGcsPath *string `json:"templateGcsPath,omitempty" tf:"template_gcs_path,omitempty"`

	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`

	// Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job. This field is not used outside of update.
	TransformNameMapping map[string]string `json:"transformNameMapping,omitempty" tf:"transform_name_mapping,omitempty"`

	// The type of this job, selected from the JobType enum
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type JobParameters struct {

	// List of experiments that should be used by the job. An example value is ["enable_stackdriver_agent_metrics"].
	// +kubebuilder:validation:Optional
	AdditionalExperiments []*string `json:"additionalExperiments,omitempty" tf:"additional_experiments,omitempty"`

	// Enable/disable the use of Streaming Engine for the job. Note that Streaming Engine is enabled by default for pipelines developed against the Beam SDK for Python v2.21.0 or later when using Python 3.
	// +kubebuilder:validation:Optional
	EnableStreamingEngine *bool `json:"enableStreamingEngine,omitempty" tf:"enable_streaming_engine,omitempty"`

	// The configuration for VM IPs.  Options are "WORKER_IP_PUBLIC" or "WORKER_IP_PRIVATE".
	// +kubebuilder:validation:Optional
	IPConfiguration *string `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// The name for the Cloud KMS key for the job. Key format is: projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// User labels to be specified for the job. Keys and values should follow the restrictions
	// specified in the labeling restrictions page.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field effective_labels for all of the labels present on the resource.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The machine type to use for the job.
	// +kubebuilder:validation:Optional
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	// +kubebuilder:validation:Optional
	MaxWorkers *float64 `json:"maxWorkers,omitempty" tf:"max_workers,omitempty"`

	// A unique name for the resource, required by Dataflow.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// One of "drain" or "cancel".  See above note.
	// +kubebuilder:validation:Optional
	OnDelete *string `json:"onDelete,omitempty" tf:"on_delete,omitempty"`

	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	// +kubebuilder:validation:Optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region in which the created job should run.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The Service Account email used to create the job.
	// +kubebuilder:validation:Optional
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// See above note.
	// +kubebuilder:validation:Optional
	SkipWaitOnJobTermination *bool `json:"skipWaitOnJobTermination,omitempty" tf:"skip_wait_on_job_termination,omitempty"`

	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK". If the subnetwork is located in a Shared VPC network, you must use the complete URL. For example "googleapis.com/compute/v1/projects/PROJECT_ID/regions/REGION/subnetworks/SUBNET_NAME"
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// A writeable location on GCS for the Dataflow job to dump its temporary data.
	// +kubebuilder:validation:Optional
	TempGcsLocation *string `json:"tempGcsLocation,omitempty" tf:"temp_gcs_location,omitempty"`

	// The GCS path to the Dataflow job template.
	// +kubebuilder:validation:Optional
	TemplateGcsPath *string `json:"templateGcsPath,omitempty" tf:"template_gcs_path,omitempty"`

	// Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job. This field is not used outside of update.
	// +kubebuilder:validation:Optional
	TransformNameMapping map[string]string `json:"transformNameMapping,omitempty" tf:"transform_name_mapping,omitempty"`

	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// JobSpec defines the desired state of Job
type JobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JobParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider JobInitParameters `json:"initProvider,omitempty"`
}

// JobStatus defines the observed state of Job.
type JobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Job is the Schema for the Jobs API. Creates a job in Dataflow according to a provided config file.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Job struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tempGcsLocation) || (has(self.initProvider) && has(self.initProvider.tempGcsLocation))",message="spec.forProvider.tempGcsLocation is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.templateGcsPath) || (has(self.initProvider) && has(self.initProvider.templateGcsPath))",message="spec.forProvider.templateGcsPath is a required parameter"
	Spec   JobSpec   `json:"spec"`
	Status JobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobList contains a list of Jobs
type JobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Job `json:"items"`
}

// Repository type metadata.
var (
	Job_Kind             = "Job"
	Job_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Job_Kind}.String()
	Job_KindAPIVersion   = Job_Kind + "." + CRDGroupVersion.String()
	Job_GroupVersionKind = CRDGroupVersion.WithKind(Job_Kind)
)

func init() {
	SchemeBuilder.Register(&Job{}, &JobList{})
}
