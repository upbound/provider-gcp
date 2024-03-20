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

type ForwardingRuleInitParameters struct {

	// The ports, portRange, and allPorts fields are mutually exclusive.
	// Only packets addressed to ports in the specified range will be forwarded
	// to the backends configured with this forwarding rule.
	// The allPorts field has the following limitations:
	AllPorts *bool `json:"allPorts,omitempty" tf:"all_ports,omitempty"`

	// This field is used along with the backend_service field for
	// internal load balancing or with the target field for internal
	// TargetInstance.
	// If the field is set to TRUE, clients can access ILB from all
	// regions.
	// Otherwise only allows access from clients in the same region as the
	// internal load balancer.
	AllowGlobalAccess *bool `json:"allowGlobalAccess,omitempty" tf:"allow_global_access,omitempty"`

	// This is used in PSC consumer ForwardingRule to control whether the PSC endpoint can be accessed from another region.
	AllowPscGlobalAccess *bool `json:"allowPscGlobalAccess,omitempty" tf:"allow_psc_global_access,omitempty"`

	// Identifies the backend service to which the forwarding rule sends traffic.
	// Required for Internal TCP/UDP Load Balancing and Network Load Balancing;
	// must be omitted for all other load balancer types.
	// +crossplane:generate:reference:type=RegionBackendService
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	BackendService *string `json:"backendService,omitempty" tf:"backend_service,omitempty"`

	// Reference to a RegionBackendService to populate backendService.
	// +kubebuilder:validation:Optional
	BackendServiceRef *v1.Reference `json:"backendServiceRef,omitempty" tf:"-"`

	// Selector for a RegionBackendService to populate backendService.
	// +kubebuilder:validation:Optional
	BackendServiceSelector *v1.Selector `json:"backendServiceSelector,omitempty" tf:"-"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// IP address for which this forwarding rule accepts traffic. When a client
	// sends traffic to this IP address, the forwarding rule directs the traffic
	// to the referenced target or backendService.
	// While creating a forwarding rule, specifying an IPAddress is
	// required under the following circumstances:
	// +crossplane:generate:reference:type=Address
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Reference to a Address to populate ipAddress.
	// +kubebuilder:validation:Optional
	IPAddressRef *v1.Reference `json:"ipAddressRef,omitempty" tf:"-"`

	// Selector for a Address to populate ipAddress.
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
	// A Forwarding Rule with protocol L3_DEFAULT can attach with target instance or
	// backend service with UNSPECIFIED protocol.
	// A forwarding rule with "L3_DEFAULT" IPProtocal cannot be attached to a backend service with TCP or UDP.
	// Possible values are: TCP, UDP, ESP, AH, SCTP, ICMP, L3_DEFAULT.
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// The IP address version that will be used by this forwarding rule.
	// Valid options are IPV4 and IPV6.
	// If not set, the IPv4 address will be used by default.
	// Possible values are: IPV4, IPV6.
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Indicates whether or not this load balancer can be used as a collector for
	// packet mirroring. To prevent mirroring loops, instances behind this
	// load balancer will not have their traffic mirrored even if a
	// PacketMirroring rule applies to them.
	// This can only be set to true for load balancers that have their
	// loadBalancingScheme set to INTERNAL.
	IsMirroringCollector *bool `json:"isMirroringCollector,omitempty" tf:"is_mirroring_collector,omitempty"`

	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Specifies the forwarding rule type.
	// For more information about forwarding rules, refer to
	// Forwarding rule concepts.
	// Default value is EXTERNAL.
	// Possible values are: EXTERNAL, EXTERNAL_MANAGED, INTERNAL, INTERNAL_MANAGED.
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty" tf:"load_balancing_scheme,omitempty"`

	// This field is not used for external load balancing.
	// For Internal TCP/UDP Load Balancing, this field identifies the network that
	// the load balanced IP should belong to for this Forwarding Rule.
	// If the subnetwork is specified, the network of the subnetwork will be used.
	// If neither subnetwork nor this field is specified, the default network will
	// be used.
	// For Private Service Connect forwarding rules that forward traffic to Google
	// APIs, a network must be provided.
	// +crossplane:generate:reference:type=Network
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// This signifies the networking tier used for configuring
	// this load balancer and can only take the following values:
	// PREMIUM, STANDARD.
	// For regional ForwardingRule, the valid values are PREMIUM and
	// STANDARD. For GlobalForwardingRule, the valid value is
	// PREMIUM.
	// If this field is not specified, it is assumed to be PREMIUM.
	// If IPAddress is specified, this value must be equal to the
	// networkTier of the Address.
	// Possible values are: PREMIUM, STANDARD.
	NetworkTier *string `json:"networkTier,omitempty" tf:"network_tier,omitempty"`

	// This is used in PSC consumer ForwardingRule to control whether it should try to auto-generate a DNS zone or not. Non-PSC forwarding rules do not use this field.
	NoAutomateDNSZone *bool `json:"noAutomateDnsZone,omitempty" tf:"no_automate_dns_zone,omitempty"`

	// The ports, portRange, and allPorts fields are mutually exclusive.
	// Only packets addressed to ports in the specified range will be forwarded
	// to the backends configured with this forwarding rule.
	// The portRange field has the following limitations:
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// The ports, portRange, and allPorts fields are mutually exclusive.
	// Only packets addressed to ports in the specified range will be forwarded
	// to the backends configured with this forwarding rule.
	// The ports field has the following limitations:
	// +listType=set
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// this is used in psc consumer forwardingrule to make provider recreate the forwardingrule when the status is closed
	RecreateClosedPsc *bool `json:"recreateClosedPsc,omitempty" tf:"recreate_closed_psc,omitempty"`

	// Service Directory resources to register this forwarding rule with.
	// Currently, only supports a single Service Directory resource.
	// Structure is documented below.
	ServiceDirectoryRegistrations []ServiceDirectoryRegistrationsInitParameters `json:"serviceDirectoryRegistrations,omitempty" tf:"service_directory_registrations,omitempty"`

	// An optional prefix to the service name for this Forwarding Rule.
	// If specified, will be the first label of the fully qualified service
	// name.
	// The label must be 1-63 characters long, and comply with RFC1035.
	// Specifically, the label must be 1-63 characters long and match the
	// regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first
	// character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	// This field is only used for INTERNAL load balancing.
	ServiceLabel *string `json:"serviceLabel,omitempty" tf:"service_label,omitempty"`

	// If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each sourceIpRange entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
	SourceIPRanges []*string `json:"sourceIpRanges,omitempty" tf:"source_ip_ranges,omitempty"`

	// This field identifies the subnetwork that the load balanced IP should
	// belong to for this Forwarding Rule, used in internal load balancing and
	// network load balancing with IPv6.
	// If the network specified is in auto subnet mode, this field is optional.
	// However, a subnetwork must be specified if the network is in custom subnet
	// mode or when creating external forwarding rule with IPv6.
	// +crossplane:generate:reference:type=Subnetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// Reference to a Subnetwork to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkRef *v1.Reference `json:"subnetworkRef,omitempty" tf:"-"`

	// Selector for a Subnetwork to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkSelector *v1.Selector `json:"subnetworkSelector,omitempty" tf:"-"`

	// is set to targetGrpcProxy and
	// validateForProxyless is set to true, the
	// IPAddress should be set to 0.0.0.0.
	// +crossplane:generate:reference:type=RegionTargetHTTPProxy
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Reference to a RegionTargetHTTPProxy to populate target.
	// +kubebuilder:validation:Optional
	TargetRef *v1.Reference `json:"targetRef,omitempty" tf:"-"`

	// Selector for a RegionTargetHTTPProxy to populate target.
	// +kubebuilder:validation:Optional
	TargetSelector *v1.Selector `json:"targetSelector,omitempty" tf:"-"`
}

type ForwardingRuleObservation struct {

	// The ports, portRange, and allPorts fields are mutually exclusive.
	// Only packets addressed to ports in the specified range will be forwarded
	// to the backends configured with this forwarding rule.
	// The allPorts field has the following limitations:
	AllPorts *bool `json:"allPorts,omitempty" tf:"all_ports,omitempty"`

	// This field is used along with the backend_service field for
	// internal load balancing or with the target field for internal
	// TargetInstance.
	// If the field is set to TRUE, clients can access ILB from all
	// regions.
	// Otherwise only allows access from clients in the same region as the
	// internal load balancer.
	AllowGlobalAccess *bool `json:"allowGlobalAccess,omitempty" tf:"allow_global_access,omitempty"`

	// This is used in PSC consumer ForwardingRule to control whether the PSC endpoint can be accessed from another region.
	AllowPscGlobalAccess *bool `json:"allowPscGlobalAccess,omitempty" tf:"allow_psc_global_access,omitempty"`

	// Identifies the backend service to which the forwarding rule sends traffic.
	// Required for Internal TCP/UDP Load Balancing and Network Load Balancing;
	// must be omitted for all other load balancer types.
	BackendService *string `json:"backendService,omitempty" tf:"backend_service,omitempty"`

	// [Output Only] The URL for the corresponding base Forwarding Rule. By base Forwarding Rule, we mean the Forwarding Rule that has the same IP address, protocol, and port settings with the current Forwarding Rule, but without sourceIPRanges specified. Always empty if the current Forwarding Rule does not have sourceIPRanges specified.
	BaseForwardingRule *string `json:"baseForwardingRule,omitempty" tf:"base_forwarding_rule,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// for all of the labels present on the resource.
	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP address for which this forwarding rule accepts traffic. When a client
	// sends traffic to this IP address, the forwarding rule directs the traffic
	// to the referenced target or backendService.
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
	// A Forwarding Rule with protocol L3_DEFAULT can attach with target instance or
	// backend service with UNSPECIFIED protocol.
	// A forwarding rule with "L3_DEFAULT" IPProtocal cannot be attached to a backend service with TCP or UDP.
	// Possible values are: TCP, UDP, ESP, AH, SCTP, ICMP, L3_DEFAULT.
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// The IP address version that will be used by this forwarding rule.
	// Valid options are IPV4 and IPV6.
	// If not set, the IPv4 address will be used by default.
	// Possible values are: IPV4, IPV6.
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Indicates whether or not this load balancer can be used as a collector for
	// packet mirroring. To prevent mirroring loops, instances behind this
	// load balancer will not have their traffic mirrored even if a
	// PacketMirroring rule applies to them.
	// This can only be set to true for load balancers that have their
	// loadBalancingScheme set to INTERNAL.
	IsMirroringCollector *bool `json:"isMirroringCollector,omitempty" tf:"is_mirroring_collector,omitempty"`

	// The fingerprint used for optimistic locking of this resource.  Used
	// internally during updates.
	LabelFingerprint *string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`

	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Specifies the forwarding rule type.
	// For more information about forwarding rules, refer to
	// Forwarding rule concepts.
	// Default value is EXTERNAL.
	// Possible values are: EXTERNAL, EXTERNAL_MANAGED, INTERNAL, INTERNAL_MANAGED.
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty" tf:"load_balancing_scheme,omitempty"`

	// This field is not used for external load balancing.
	// For Internal TCP/UDP Load Balancing, this field identifies the network that
	// the load balanced IP should belong to for this Forwarding Rule.
	// If the subnetwork is specified, the network of the subnetwork will be used.
	// If neither subnetwork nor this field is specified, the default network will
	// be used.
	// For Private Service Connect forwarding rules that forward traffic to Google
	// APIs, a network must be provided.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// This signifies the networking tier used for configuring
	// this load balancer and can only take the following values:
	// PREMIUM, STANDARD.
	// For regional ForwardingRule, the valid values are PREMIUM and
	// STANDARD. For GlobalForwardingRule, the valid value is
	// PREMIUM.
	// If this field is not specified, it is assumed to be PREMIUM.
	// If IPAddress is specified, this value must be equal to the
	// networkTier of the Address.
	// Possible values are: PREMIUM, STANDARD.
	NetworkTier *string `json:"networkTier,omitempty" tf:"network_tier,omitempty"`

	// This is used in PSC consumer ForwardingRule to control whether it should try to auto-generate a DNS zone or not. Non-PSC forwarding rules do not use this field.
	NoAutomateDNSZone *bool `json:"noAutomateDnsZone,omitempty" tf:"no_automate_dns_zone,omitempty"`

	// The ports, portRange, and allPorts fields are mutually exclusive.
	// Only packets addressed to ports in the specified range will be forwarded
	// to the backends configured with this forwarding rule.
	// The portRange field has the following limitations:
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// The ports, portRange, and allPorts fields are mutually exclusive.
	// Only packets addressed to ports in the specified range will be forwarded
	// to the backends configured with this forwarding rule.
	// The ports field has the following limitations:
	// +listType=set
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The PSC connection id of the PSC Forwarding Rule.
	PscConnectionID *string `json:"pscConnectionId,omitempty" tf:"psc_connection_id,omitempty"`

	// The PSC connection status of the PSC Forwarding Rule. Possible values: STATUS_UNSPECIFIED, PENDING, ACCEPTED, REJECTED, CLOSED
	PscConnectionStatus *string `json:"pscConnectionStatus,omitempty" tf:"psc_connection_status,omitempty"`

	// this is used in psc consumer forwardingrule to make provider recreate the forwardingrule when the status is closed
	RecreateClosedPsc *bool `json:"recreateClosedPsc,omitempty" tf:"recreate_closed_psc,omitempty"`

	// A reference to the region where the regional forwarding rule resides.
	// This field is not applicable to global forwarding rules.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// Service Directory resources to register this forwarding rule with.
	// Currently, only supports a single Service Directory resource.
	// Structure is documented below.
	ServiceDirectoryRegistrations []ServiceDirectoryRegistrationsObservation `json:"serviceDirectoryRegistrations,omitempty" tf:"service_directory_registrations,omitempty"`

	// An optional prefix to the service name for this Forwarding Rule.
	// If specified, will be the first label of the fully qualified service
	// name.
	// The label must be 1-63 characters long, and comply with RFC1035.
	// Specifically, the label must be 1-63 characters long and match the
	// regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first
	// character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	// This field is only used for INTERNAL load balancing.
	ServiceLabel *string `json:"serviceLabel,omitempty" tf:"service_label,omitempty"`

	// The internal fully qualified service name for this Forwarding Rule.
	// This field is only used for INTERNAL load balancing.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each sourceIpRange entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
	SourceIPRanges []*string `json:"sourceIpRanges,omitempty" tf:"source_ip_ranges,omitempty"`

	// This field identifies the subnetwork that the load balanced IP should
	// belong to for this Forwarding Rule, used in internal load balancing and
	// network load balancing with IPv6.
	// If the network specified is in auto subnet mode, this field is optional.
	// However, a subnetwork must be specified if the network is in custom subnet
	// mode or when creating external forwarding rule with IPv6.
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// is set to targetGrpcProxy and
	// validateForProxyless is set to true, the
	// IPAddress should be set to 0.0.0.0.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`
}

type ForwardingRuleParameters struct {

	// The ports, portRange, and allPorts fields are mutually exclusive.
	// Only packets addressed to ports in the specified range will be forwarded
	// to the backends configured with this forwarding rule.
	// The allPorts field has the following limitations:
	// +kubebuilder:validation:Optional
	AllPorts *bool `json:"allPorts,omitempty" tf:"all_ports,omitempty"`

	// This field is used along with the backend_service field for
	// internal load balancing or with the target field for internal
	// TargetInstance.
	// If the field is set to TRUE, clients can access ILB from all
	// regions.
	// Otherwise only allows access from clients in the same region as the
	// internal load balancer.
	// +kubebuilder:validation:Optional
	AllowGlobalAccess *bool `json:"allowGlobalAccess,omitempty" tf:"allow_global_access,omitempty"`

	// This is used in PSC consumer ForwardingRule to control whether the PSC endpoint can be accessed from another region.
	// +kubebuilder:validation:Optional
	AllowPscGlobalAccess *bool `json:"allowPscGlobalAccess,omitempty" tf:"allow_psc_global_access,omitempty"`

	// Identifies the backend service to which the forwarding rule sends traffic.
	// Required for Internal TCP/UDP Load Balancing and Network Load Balancing;
	// must be omitted for all other load balancer types.
	// +crossplane:generate:reference:type=RegionBackendService
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	// +kubebuilder:validation:Optional
	BackendService *string `json:"backendService,omitempty" tf:"backend_service,omitempty"`

	// Reference to a RegionBackendService to populate backendService.
	// +kubebuilder:validation:Optional
	BackendServiceRef *v1.Reference `json:"backendServiceRef,omitempty" tf:"-"`

	// Selector for a RegionBackendService to populate backendService.
	// +kubebuilder:validation:Optional
	BackendServiceSelector *v1.Selector `json:"backendServiceSelector,omitempty" tf:"-"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// IP address for which this forwarding rule accepts traffic. When a client
	// sends traffic to this IP address, the forwarding rule directs the traffic
	// to the referenced target or backendService.
	// While creating a forwarding rule, specifying an IPAddress is
	// required under the following circumstances:
	// +crossplane:generate:reference:type=Address
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Reference to a Address to populate ipAddress.
	// +kubebuilder:validation:Optional
	IPAddressRef *v1.Reference `json:"ipAddressRef,omitempty" tf:"-"`

	// Selector for a Address to populate ipAddress.
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
	// A Forwarding Rule with protocol L3_DEFAULT can attach with target instance or
	// backend service with UNSPECIFIED protocol.
	// A forwarding rule with "L3_DEFAULT" IPProtocal cannot be attached to a backend service with TCP or UDP.
	// Possible values are: TCP, UDP, ESP, AH, SCTP, ICMP, L3_DEFAULT.
	// +kubebuilder:validation:Optional
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// The IP address version that will be used by this forwarding rule.
	// Valid options are IPV4 and IPV6.
	// If not set, the IPv4 address will be used by default.
	// Possible values are: IPV4, IPV6.
	// +kubebuilder:validation:Optional
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Indicates whether or not this load balancer can be used as a collector for
	// packet mirroring. To prevent mirroring loops, instances behind this
	// load balancer will not have their traffic mirrored even if a
	// PacketMirroring rule applies to them.
	// This can only be set to true for load balancers that have their
	// loadBalancingScheme set to INTERNAL.
	// +kubebuilder:validation:Optional
	IsMirroringCollector *bool `json:"isMirroringCollector,omitempty" tf:"is_mirroring_collector,omitempty"`

	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Specifies the forwarding rule type.
	// For more information about forwarding rules, refer to
	// Forwarding rule concepts.
	// Default value is EXTERNAL.
	// Possible values are: EXTERNAL, EXTERNAL_MANAGED, INTERNAL, INTERNAL_MANAGED.
	// +kubebuilder:validation:Optional
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty" tf:"load_balancing_scheme,omitempty"`

	// This field is not used for external load balancing.
	// For Internal TCP/UDP Load Balancing, this field identifies the network that
	// the load balanced IP should belong to for this Forwarding Rule.
	// If the subnetwork is specified, the network of the subnetwork will be used.
	// If neither subnetwork nor this field is specified, the default network will
	// be used.
	// For Private Service Connect forwarding rules that forward traffic to Google
	// APIs, a network must be provided.
	// +crossplane:generate:reference:type=Network
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// This signifies the networking tier used for configuring
	// this load balancer and can only take the following values:
	// PREMIUM, STANDARD.
	// For regional ForwardingRule, the valid values are PREMIUM and
	// STANDARD. For GlobalForwardingRule, the valid value is
	// PREMIUM.
	// If this field is not specified, it is assumed to be PREMIUM.
	// If IPAddress is specified, this value must be equal to the
	// networkTier of the Address.
	// Possible values are: PREMIUM, STANDARD.
	// +kubebuilder:validation:Optional
	NetworkTier *string `json:"networkTier,omitempty" tf:"network_tier,omitempty"`

	// This is used in PSC consumer ForwardingRule to control whether it should try to auto-generate a DNS zone or not. Non-PSC forwarding rules do not use this field.
	// +kubebuilder:validation:Optional
	NoAutomateDNSZone *bool `json:"noAutomateDnsZone,omitempty" tf:"no_automate_dns_zone,omitempty"`

	// The ports, portRange, and allPorts fields are mutually exclusive.
	// Only packets addressed to ports in the specified range will be forwarded
	// to the backends configured with this forwarding rule.
	// The portRange field has the following limitations:
	// +kubebuilder:validation:Optional
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// The ports, portRange, and allPorts fields are mutually exclusive.
	// Only packets addressed to ports in the specified range will be forwarded
	// to the backends configured with this forwarding rule.
	// The ports field has the following limitations:
	// +kubebuilder:validation:Optional
	// +listType=set
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// this is used in psc consumer forwardingrule to make provider recreate the forwardingrule when the status is closed
	// +kubebuilder:validation:Optional
	RecreateClosedPsc *bool `json:"recreateClosedPsc,omitempty" tf:"recreate_closed_psc,omitempty"`

	// A reference to the region where the regional forwarding rule resides.
	// This field is not applicable to global forwarding rules.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// Service Directory resources to register this forwarding rule with.
	// Currently, only supports a single Service Directory resource.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ServiceDirectoryRegistrations []ServiceDirectoryRegistrationsParameters `json:"serviceDirectoryRegistrations,omitempty" tf:"service_directory_registrations,omitempty"`

	// An optional prefix to the service name for this Forwarding Rule.
	// If specified, will be the first label of the fully qualified service
	// name.
	// The label must be 1-63 characters long, and comply with RFC1035.
	// Specifically, the label must be 1-63 characters long and match the
	// regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first
	// character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	// This field is only used for INTERNAL load balancing.
	// +kubebuilder:validation:Optional
	ServiceLabel *string `json:"serviceLabel,omitempty" tf:"service_label,omitempty"`

	// If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each sourceIpRange entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
	// +kubebuilder:validation:Optional
	SourceIPRanges []*string `json:"sourceIpRanges,omitempty" tf:"source_ip_ranges,omitempty"`

	// This field identifies the subnetwork that the load balanced IP should
	// belong to for this Forwarding Rule, used in internal load balancing and
	// network load balancing with IPv6.
	// If the network specified is in auto subnet mode, this field is optional.
	// However, a subnetwork must be specified if the network is in custom subnet
	// mode or when creating external forwarding rule with IPv6.
	// +crossplane:generate:reference:type=Subnetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// Reference to a Subnetwork to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkRef *v1.Reference `json:"subnetworkRef,omitempty" tf:"-"`

	// Selector for a Subnetwork to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkSelector *v1.Selector `json:"subnetworkSelector,omitempty" tf:"-"`

	// is set to targetGrpcProxy and
	// validateForProxyless is set to true, the
	// IPAddress should be set to 0.0.0.0.
	// +crossplane:generate:reference:type=RegionTargetHTTPProxy
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Reference to a RegionTargetHTTPProxy to populate target.
	// +kubebuilder:validation:Optional
	TargetRef *v1.Reference `json:"targetRef,omitempty" tf:"-"`

	// Selector for a RegionTargetHTTPProxy to populate target.
	// +kubebuilder:validation:Optional
	TargetSelector *v1.Selector `json:"targetSelector,omitempty" tf:"-"`
}

type ServiceDirectoryRegistrationsInitParameters struct {

	// Service Directory namespace to register the forwarding rule under.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Service Directory service to register the forwarding rule under.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type ServiceDirectoryRegistrationsObservation struct {

	// Service Directory namespace to register the forwarding rule under.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Service Directory service to register the forwarding rule under.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type ServiceDirectoryRegistrationsParameters struct {

	// Service Directory namespace to register the forwarding rule under.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Service Directory service to register the forwarding rule under.
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

// ForwardingRuleSpec defines the desired state of ForwardingRule
type ForwardingRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ForwardingRuleParameters `json:"forProvider"`
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
	InitProvider ForwardingRuleInitParameters `json:"initProvider,omitempty"`
}

// ForwardingRuleStatus defines the observed state of ForwardingRule.
type ForwardingRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ForwardingRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ForwardingRule is the Schema for the ForwardingRules API. A ForwardingRule resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ForwardingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ForwardingRuleSpec   `json:"spec"`
	Status            ForwardingRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ForwardingRuleList contains a list of ForwardingRules
type ForwardingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ForwardingRule `json:"items"`
}

// Repository type metadata.
var (
	ForwardingRule_Kind             = "ForwardingRule"
	ForwardingRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ForwardingRule_Kind}.String()
	ForwardingRule_KindAPIVersion   = ForwardingRule_Kind + "." + CRDGroupVersion.String()
	ForwardingRule_GroupVersionKind = CRDGroupVersion.WithKind(ForwardingRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ForwardingRule{}, &ForwardingRuleList{})
}
