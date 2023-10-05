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

type FilterLabelsInitParameters struct {

	// Name of the resource; provided by the client when the resource is created.
	// The name must be 1-63 characters long, and comply with
	// RFC1035.
	// Specifically, the name must be 1-63 characters long and match the regular
	// expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first
	// character must be a lowercase letter, and all following characters must
	// be a dash, lowercase letter, or digit, except the last character, which
	// cannot be a dash.
	// For Private Service Connect forwarding rules that forward traffic to Google
	// APIs, the forwarding rule name must be a 1-20 characters string with
	// lowercase letters and numbers and must start with a letter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value that the label must match. The value has a maximum
	// length of 1024 characters.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FilterLabelsObservation struct {

	// Name of the resource; provided by the client when the resource is created.
	// The name must be 1-63 characters long, and comply with
	// RFC1035.
	// Specifically, the name must be 1-63 characters long and match the regular
	// expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first
	// character must be a lowercase letter, and all following characters must
	// be a dash, lowercase letter, or digit, except the last character, which
	// cannot be a dash.
	// For Private Service Connect forwarding rules that forward traffic to Google
	// APIs, the forwarding rule name must be a 1-20 characters string with
	// lowercase letters and numbers and must start with a letter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value that the label must match. The value has a maximum
	// length of 1024 characters.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FilterLabelsParameters struct {

	// Name of the resource; provided by the client when the resource is created.
	// The name must be 1-63 characters long, and comply with
	// RFC1035.
	// Specifically, the name must be 1-63 characters long and match the regular
	// expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first
	// character must be a lowercase letter, and all following characters must
	// be a dash, lowercase letter, or digit, except the last character, which
	// cannot be a dash.
	// For Private Service Connect forwarding rules that forward traffic to Google
	// APIs, the forwarding rule name must be a 1-20 characters string with
	// lowercase letters and numbers and must start with a letter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value that the label must match. The value has a maximum
	// length of 1024 characters.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type GlobalForwardingRuleInitParameters struct {

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP protocol to which this rule applies.
	// For protocol forwarding, valid
	// options are TCP, UDP, ESP,
	// AH, SCTP, ICMP and
	// L3_DEFAULT.
	// The valid IP protocols are different for different load balancing products
	// as described in Load balancing
	// features.
	// Possible values are: TCP, UDP, ESP, AH, SCTP, ICMP.
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// The IP Version that will be used by this global forwarding rule.
	// Possible values are: IPV4, IPV6.
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Specifies the forwarding rule type.
	// For more information about forwarding rules, refer to
	// Forwarding rule concepts.
	// Default value is EXTERNAL.
	// Possible values are: EXTERNAL, EXTERNAL_MANAGED, INTERNAL_MANAGED, INTERNAL_SELF_MANAGED.
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty" tf:"load_balancing_scheme,omitempty"`

	// Opaque filter criteria used by Loadbalancer to restrict routing
	// configuration to a limited set xDS compliant clients. In their xDS
	// requests to Loadbalancer, xDS clients present node metadata. If a
	// match takes place, the relevant routing configuration is made available
	// to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set
	// to MATCH_ANY, at least one of the filterLabels must match the
	// corresponding label provided in the metadata. If its filterMatchCriteria
	// is set to MATCH_ALL, then all of its filterLabels must match with
	// corresponding labels in the provided metadata.
	// metadataFilters specified here can be overridden by those specified in
	// the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their
	// loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	// Structure is documented below.
	MetadataFilters []MetadataFiltersInitParameters `json:"metadataFilters,omitempty" tf:"metadata_filters,omitempty"`

	// This is used in PSC consumer ForwardingRule to control whether it should try to auto-generate a DNS zone or not. Non-PSC forwarding rules do not use this field.
	NoAutomateDNSZone *bool `json:"noAutomateDnsZone,omitempty" tf:"no_automate_dns_zone,omitempty"`

	// This field can only be used:
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each sourceIpRange entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
	SourceIPRanges []*string `json:"sourceIpRanges,omitempty" tf:"source_ip_ranges,omitempty"`
}

type GlobalForwardingRuleObservation struct {

	// [Output Only] The URL for the corresponding base Forwarding Rule. By base Forwarding Rule, we mean the Forwarding Rule that has the same IP address, protocol, and port settings with the current Forwarding Rule, but without sourceIPRanges specified. Always empty if the current Forwarding Rule does not have sourceIPRanges specified.
	BaseForwardingRule *string `json:"baseForwardingRule,omitempty" tf:"base_forwarding_rule,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// for all of the labels present on the resource.
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/forwardingRules/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP address for which this forwarding rule accepts traffic. When a client
	// sends traffic to this IP address, the forwarding rule directs the traffic
	// to the referenced target.
	// While creating a forwarding rule, specifying an IPAddress is
	// required under the following circumstances:
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The IP protocol to which this rule applies.
	// For protocol forwarding, valid
	// options are TCP, UDP, ESP,
	// AH, SCTP, ICMP and
	// L3_DEFAULT.
	// The valid IP protocols are different for different load balancing products
	// as described in Load balancing
	// features.
	// Possible values are: TCP, UDP, ESP, AH, SCTP, ICMP.
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// The IP Version that will be used by this global forwarding rule.
	// Possible values are: IPV4, IPV6.
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// The fingerprint used for optimistic locking of this resource.  Used
	// internally during updates.
	LabelFingerprint *string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`

	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Specifies the forwarding rule type.
	// For more information about forwarding rules, refer to
	// Forwarding rule concepts.
	// Default value is EXTERNAL.
	// Possible values are: EXTERNAL, EXTERNAL_MANAGED, INTERNAL_MANAGED, INTERNAL_SELF_MANAGED.
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty" tf:"load_balancing_scheme,omitempty"`

	// Opaque filter criteria used by Loadbalancer to restrict routing
	// configuration to a limited set xDS compliant clients. In their xDS
	// requests to Loadbalancer, xDS clients present node metadata. If a
	// match takes place, the relevant routing configuration is made available
	// to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set
	// to MATCH_ANY, at least one of the filterLabels must match the
	// corresponding label provided in the metadata. If its filterMatchCriteria
	// is set to MATCH_ALL, then all of its filterLabels must match with
	// corresponding labels in the provided metadata.
	// metadataFilters specified here can be overridden by those specified in
	// the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their
	// loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	// Structure is documented below.
	MetadataFilters []MetadataFiltersObservation `json:"metadataFilters,omitempty" tf:"metadata_filters,omitempty"`

	// This field is not used for external load balancing.
	// For Internal TCP/UDP Load Balancing, this field identifies the network that
	// the load balanced IP should belong to for this Forwarding Rule.
	// If the subnetwork is specified, the network of the subnetwork will be used.
	// If neither subnetwork nor this field is specified, the default network will
	// be used.
	// For Private Service Connect forwarding rules that forward traffic to Google
	// APIs, a network must be provided.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// This is used in PSC consumer ForwardingRule to control whether it should try to auto-generate a DNS zone or not. Non-PSC forwarding rules do not use this field.
	NoAutomateDNSZone *bool `json:"noAutomateDnsZone,omitempty" tf:"no_automate_dns_zone,omitempty"`

	// This field can only be used:
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The PSC connection id of the PSC Forwarding Rule.
	PscConnectionID *string `json:"pscConnectionId,omitempty" tf:"psc_connection_id,omitempty"`

	// The PSC connection status of the PSC Forwarding Rule. Possible values: STATUS_UNSPECIFIED, PENDING, ACCEPTED, REJECTED, CLOSED
	PscConnectionStatus *string `json:"pscConnectionStatus,omitempty" tf:"psc_connection_status,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each sourceIpRange entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
	SourceIPRanges []*string `json:"sourceIpRanges,omitempty" tf:"source_ip_ranges,omitempty"`

	// This field identifies the subnetwork that the load balanced IP should
	// belong to for this Forwarding Rule, used in internal load balancing and
	// network load balancing with IPv6.
	// If the network specified is in auto subnet mode, this field is optional.
	// However, a subnetwork must be specified if the network is in custom subnet
	// mode or when creating external forwarding rule with IPv6.
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// The URL of the target resource to receive the matched traffic.  For
	// regional forwarding rules, this target must be in the same region as the
	// forwarding rule. For global forwarding rules, this target must be a global
	// load balancing resource.
	// The forwarded traffic must be of a type appropriate to the target object.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`
}

type GlobalForwardingRuleParameters struct {

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// IP address for which this forwarding rule accepts traffic. When a client
	// sends traffic to this IP address, the forwarding rule directs the traffic
	// to the referenced target.
	// While creating a forwarding rule, specifying an IPAddress is
	// required under the following circumstances:
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.GlobalAddress
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Reference to a GlobalAddress in compute to populate ipAddress.
	// +kubebuilder:validation:Optional
	IPAddressRef *v1.Reference `json:"ipAddressRef,omitempty" tf:"-"`

	// Selector for a GlobalAddress in compute to populate ipAddress.
	// +kubebuilder:validation:Optional
	IPAddressSelector *v1.Selector `json:"ipAddressSelector,omitempty" tf:"-"`

	// The IP protocol to which this rule applies.
	// For protocol forwarding, valid
	// options are TCP, UDP, ESP,
	// AH, SCTP, ICMP and
	// L3_DEFAULT.
	// The valid IP protocols are different for different load balancing products
	// as described in Load balancing
	// features.
	// Possible values are: TCP, UDP, ESP, AH, SCTP, ICMP.
	// +kubebuilder:validation:Optional
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// The IP Version that will be used by this global forwarding rule.
	// Possible values are: IPV4, IPV6.
	// +kubebuilder:validation:Optional
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Specifies the forwarding rule type.
	// For more information about forwarding rules, refer to
	// Forwarding rule concepts.
	// Default value is EXTERNAL.
	// Possible values are: EXTERNAL, EXTERNAL_MANAGED, INTERNAL_MANAGED, INTERNAL_SELF_MANAGED.
	// +kubebuilder:validation:Optional
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty" tf:"load_balancing_scheme,omitempty"`

	// Opaque filter criteria used by Loadbalancer to restrict routing
	// configuration to a limited set xDS compliant clients. In their xDS
	// requests to Loadbalancer, xDS clients present node metadata. If a
	// match takes place, the relevant routing configuration is made available
	// to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set
	// to MATCH_ANY, at least one of the filterLabels must match the
	// corresponding label provided in the metadata. If its filterMatchCriteria
	// is set to MATCH_ALL, then all of its filterLabels must match with
	// corresponding labels in the provided metadata.
	// metadataFilters specified here can be overridden by those specified in
	// the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their
	// loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MetadataFilters []MetadataFiltersParameters `json:"metadataFilters,omitempty" tf:"metadata_filters,omitempty"`

	// This field is not used for external load balancing.
	// For Internal TCP/UDP Load Balancing, this field identifies the network that
	// the load balanced IP should belong to for this Forwarding Rule.
	// If the subnetwork is specified, the network of the subnetwork will be used.
	// If neither subnetwork nor this field is specified, the default network will
	// be used.
	// For Private Service Connect forwarding rules that forward traffic to Google
	// APIs, a network must be provided.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// This is used in PSC consumer ForwardingRule to control whether it should try to auto-generate a DNS zone or not. Non-PSC forwarding rules do not use this field.
	// +kubebuilder:validation:Optional
	NoAutomateDNSZone *bool `json:"noAutomateDnsZone,omitempty" tf:"no_automate_dns_zone,omitempty"`

	// This field can only be used:
	// +kubebuilder:validation:Optional
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("project",false)
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Network in compute to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each sourceIpRange entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
	// +kubebuilder:validation:Optional
	SourceIPRanges []*string `json:"sourceIpRanges,omitempty" tf:"source_ip_ranges,omitempty"`

	// This field identifies the subnetwork that the load balanced IP should
	// belong to for this Forwarding Rule, used in internal load balancing and
	// network load balancing with IPv6.
	// If the network specified is in auto subnet mode, this field is optional.
	// However, a subnetwork must be specified if the network is in custom subnet
	// mode or when creating external forwarding rule with IPv6.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Subnetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// Reference to a Subnetwork in compute to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkRef *v1.Reference `json:"subnetworkRef,omitempty" tf:"-"`

	// Selector for a Subnetwork in compute to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkSelector *v1.Selector `json:"subnetworkSelector,omitempty" tf:"-"`

	// The URL of the target resource to receive the matched traffic.  For
	// regional forwarding rules, this target must be in the same region as the
	// forwarding rule. For global forwarding rules, this target must be a global
	// load balancing resource.
	// The forwarded traffic must be of a type appropriate to the target object.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.TargetSSLProxy
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Reference to a TargetSSLProxy in compute to populate target.
	// +kubebuilder:validation:Optional
	TargetRef *v1.Reference `json:"targetRef,omitempty" tf:"-"`

	// Selector for a TargetSSLProxy in compute to populate target.
	// +kubebuilder:validation:Optional
	TargetSelector *v1.Selector `json:"targetSelector,omitempty" tf:"-"`
}

type MetadataFiltersInitParameters struct {

	// The list of label value pairs that must match labels in the
	// provided metadata based on filterMatchCriteria
	// This list must not be empty and can have at the most 64 entries.
	// Structure is documented below.
	FilterLabels []FilterLabelsInitParameters `json:"filterLabels,omitempty" tf:"filter_labels,omitempty"`

	// Specifies how individual filterLabel matches within the list of
	// filterLabels contribute towards the overall metadataFilter match.
	// MATCH_ANY - At least one of the filterLabels must have a matching
	// label in the provided metadata.
	// MATCH_ALL - All filterLabels must have matching labels in the
	// provided metadata.
	// Possible values are: MATCH_ANY, MATCH_ALL.
	FilterMatchCriteria *string `json:"filterMatchCriteria,omitempty" tf:"filter_match_criteria,omitempty"`
}

type MetadataFiltersObservation struct {

	// The list of label value pairs that must match labels in the
	// provided metadata based on filterMatchCriteria
	// This list must not be empty and can have at the most 64 entries.
	// Structure is documented below.
	FilterLabels []FilterLabelsObservation `json:"filterLabels,omitempty" tf:"filter_labels,omitempty"`

	// Specifies how individual filterLabel matches within the list of
	// filterLabels contribute towards the overall metadataFilter match.
	// MATCH_ANY - At least one of the filterLabels must have a matching
	// label in the provided metadata.
	// MATCH_ALL - All filterLabels must have matching labels in the
	// provided metadata.
	// Possible values are: MATCH_ANY, MATCH_ALL.
	FilterMatchCriteria *string `json:"filterMatchCriteria,omitempty" tf:"filter_match_criteria,omitempty"`
}

type MetadataFiltersParameters struct {

	// The list of label value pairs that must match labels in the
	// provided metadata based on filterMatchCriteria
	// This list must not be empty and can have at the most 64 entries.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	FilterLabels []FilterLabelsParameters `json:"filterLabels" tf:"filter_labels,omitempty"`

	// Specifies how individual filterLabel matches within the list of
	// filterLabels contribute towards the overall metadataFilter match.
	// MATCH_ANY - At least one of the filterLabels must have a matching
	// label in the provided metadata.
	// MATCH_ALL - All filterLabels must have matching labels in the
	// provided metadata.
	// Possible values are: MATCH_ANY, MATCH_ALL.
	// +kubebuilder:validation:Optional
	FilterMatchCriteria *string `json:"filterMatchCriteria" tf:"filter_match_criteria,omitempty"`
}

// GlobalForwardingRuleSpec defines the desired state of GlobalForwardingRule
type GlobalForwardingRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalForwardingRuleParameters `json:"forProvider"`
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
	InitProvider GlobalForwardingRuleInitParameters `json:"initProvider,omitempty"`
}

// GlobalForwardingRuleStatus defines the observed state of GlobalForwardingRule.
type GlobalForwardingRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalForwardingRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalForwardingRule is the Schema for the GlobalForwardingRules API. Represents a GlobalForwardingRule resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type GlobalForwardingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalForwardingRuleSpec   `json:"spec"`
	Status            GlobalForwardingRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalForwardingRuleList contains a list of GlobalForwardingRules
type GlobalForwardingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalForwardingRule `json:"items"`
}

// Repository type metadata.
var (
	GlobalForwardingRule_Kind             = "GlobalForwardingRule"
	GlobalForwardingRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalForwardingRule_Kind}.String()
	GlobalForwardingRule_KindAPIVersion   = GlobalForwardingRule_Kind + "." + CRDGroupVersion.String()
	GlobalForwardingRule_GroupVersionKind = CRDGroupVersion.WithKind(GlobalForwardingRule_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalForwardingRule{}, &GlobalForwardingRuleList{})
}
