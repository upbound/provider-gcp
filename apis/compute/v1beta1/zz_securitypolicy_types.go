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

type AdaptiveProtectionConfigInitParameters struct {

	// Configuration for Google Cloud Armor Adaptive Protection Layer 7 DDoS Defense. Structure is documented below.
	Layer7DdosDefenseConfig []Layer7DdosDefenseConfigInitParameters `json:"layer7DdosDefenseConfig,omitempty" tf:"layer_7_ddos_defense_config,omitempty"`
}

type AdaptiveProtectionConfigObservation struct {

	// Configuration for Google Cloud Armor Adaptive Protection Layer 7 DDoS Defense. Structure is documented below.
	Layer7DdosDefenseConfig []Layer7DdosDefenseConfigObservation `json:"layer7DdosDefenseConfig,omitempty" tf:"layer_7_ddos_defense_config,omitempty"`
}

type AdaptiveProtectionConfigParameters struct {

	// Configuration for Google Cloud Armor Adaptive Protection Layer 7 DDoS Defense. Structure is documented below.
	// +kubebuilder:validation:Optional
	Layer7DdosDefenseConfig []Layer7DdosDefenseConfigParameters `json:"layer7DdosDefenseConfig,omitempty" tf:"layer_7_ddos_defense_config,omitempty"`
}

type AdvancedOptionsConfigInitParameters struct {

	// Custom configuration to apply the JSON parsing. Only applicable when
	// json_parsing is set to STANDARD. Structure is documented below.
	JSONCustomConfig []JSONCustomConfigInitParameters `json:"jsonCustomConfig,omitempty" tf:"json_custom_config,omitempty"`

	// Whether or not to JSON parse the payload body. Defaults to DISABLED.
	JSONParsing *string `json:"jsonParsing,omitempty" tf:"json_parsing,omitempty"`

	// Log level to use. Defaults to NORMAL.
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type AdvancedOptionsConfigObservation struct {

	// Custom configuration to apply the JSON parsing. Only applicable when
	// json_parsing is set to STANDARD. Structure is documented below.
	JSONCustomConfig []JSONCustomConfigObservation `json:"jsonCustomConfig,omitempty" tf:"json_custom_config,omitempty"`

	// Whether or not to JSON parse the payload body. Defaults to DISABLED.
	JSONParsing *string `json:"jsonParsing,omitempty" tf:"json_parsing,omitempty"`

	// Log level to use. Defaults to NORMAL.
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type AdvancedOptionsConfigParameters struct {

	// Custom configuration to apply the JSON parsing. Only applicable when
	// json_parsing is set to STANDARD. Structure is documented below.
	// +kubebuilder:validation:Optional
	JSONCustomConfig []JSONCustomConfigParameters `json:"jsonCustomConfig,omitempty" tf:"json_custom_config,omitempty"`

	// Whether or not to JSON parse the payload body. Defaults to DISABLED.
	// +kubebuilder:validation:Optional
	JSONParsing *string `json:"jsonParsing,omitempty" tf:"json_parsing,omitempty"`

	// Log level to use. Defaults to NORMAL.
	// +kubebuilder:validation:Optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`
}

type BanThresholdInitParameters struct {

	// Number of HTTP(S) requests for calculating the threshold.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// Interval over which the threshold is computed.
	IntervalSec *float64 `json:"intervalSec,omitempty" tf:"interval_sec,omitempty"`
}

type BanThresholdObservation struct {

	// Number of HTTP(S) requests for calculating the threshold.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// Interval over which the threshold is computed.
	IntervalSec *float64 `json:"intervalSec,omitempty" tf:"interval_sec,omitempty"`
}

type BanThresholdParameters struct {

	// Number of HTTP(S) requests for calculating the threshold.
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count" tf:"count,omitempty"`

	// Interval over which the threshold is computed.
	// +kubebuilder:validation:Optional
	IntervalSec *float64 `json:"intervalSec" tf:"interval_sec,omitempty"`
}

type ConfigInitParameters struct {

	// field in config.
	// +listType=set
	SrcIPRanges []*string `json:"srcIpRanges,omitempty" tf:"src_ip_ranges,omitempty"`
}

type ConfigObservation struct {

	// field in config.
	// +listType=set
	SrcIPRanges []*string `json:"srcIpRanges,omitempty" tf:"src_ip_ranges,omitempty"`
}

type ConfigParameters struct {

	// field in config.
	// +kubebuilder:validation:Optional
	// +listType=set
	SrcIPRanges []*string `json:"srcIpRanges" tf:"src_ip_ranges,omitempty"`
}

type ExceedRedirectOptionsInitParameters struct {

	// Target for the redirect action. This is required if the type is EXTERNAL_302 and cannot be specified for GOOGLE_RECAPTCHA.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Type of the redirect action.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ExceedRedirectOptionsObservation struct {

	// Target for the redirect action. This is required if the type is EXTERNAL_302 and cannot be specified for GOOGLE_RECAPTCHA.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Type of the redirect action.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ExceedRedirectOptionsParameters struct {

	// Target for the redirect action. This is required if the type is EXTERNAL_302 and cannot be specified for GOOGLE_RECAPTCHA.
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Type of the redirect action.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type ExprInitParameters struct {

	// Textual representation of an expression in Common Expression Language syntax.
	// The application context of the containing message determines which well-known feature set of CEL is supported.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`
}

type ExprObservation struct {

	// Textual representation of an expression in Common Expression Language syntax.
	// The application context of the containing message determines which well-known feature set of CEL is supported.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`
}

type ExprParameters struct {

	// Textual representation of an expression in Common Expression Language syntax.
	// The application context of the containing message determines which well-known feature set of CEL is supported.
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`
}

type JSONCustomConfigInitParameters struct {

	// A list of custom Content-Type header values to apply the JSON parsing. The
	// format of the Content-Type header values is defined in
	// RFC 1341. When configuring a custom Content-Type header
	// value, only the type/subtype needs to be specified, and the parameters should be excluded.
	// +listType=set
	ContentTypes []*string `json:"contentTypes,omitempty" tf:"content_types,omitempty"`
}

type JSONCustomConfigObservation struct {

	// A list of custom Content-Type header values to apply the JSON parsing. The
	// format of the Content-Type header values is defined in
	// RFC 1341. When configuring a custom Content-Type header
	// value, only the type/subtype needs to be specified, and the parameters should be excluded.
	// +listType=set
	ContentTypes []*string `json:"contentTypes,omitempty" tf:"content_types,omitempty"`
}

type JSONCustomConfigParameters struct {

	// A list of custom Content-Type header values to apply the JSON parsing. The
	// format of the Content-Type header values is defined in
	// RFC 1341. When configuring a custom Content-Type header
	// value, only the type/subtype needs to be specified, and the parameters should be excluded.
	// +kubebuilder:validation:Optional
	// +listType=set
	ContentTypes []*string `json:"contentTypes" tf:"content_types,omitempty"`
}

type Layer7DdosDefenseConfigInitParameters struct {

	// If set to true, enables CAAP for L7 DDoS detection.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Rule visibility can be one of the following: STANDARD - opaque rules. (default) PREMIUM - transparent rules.
	RuleVisibility *string `json:"ruleVisibility,omitempty" tf:"rule_visibility,omitempty"`
}

type Layer7DdosDefenseConfigObservation struct {

	// If set to true, enables CAAP for L7 DDoS detection.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Rule visibility can be one of the following: STANDARD - opaque rules. (default) PREMIUM - transparent rules.
	RuleVisibility *string `json:"ruleVisibility,omitempty" tf:"rule_visibility,omitempty"`
}

type Layer7DdosDefenseConfigParameters struct {

	// If set to true, enables CAAP for L7 DDoS detection.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Rule visibility can be one of the following: STANDARD - opaque rules. (default) PREMIUM - transparent rules.
	// +kubebuilder:validation:Optional
	RuleVisibility *string `json:"ruleVisibility,omitempty" tf:"rule_visibility,omitempty"`
}

type RateLimitOptionsInitParameters struct {

	// Can only be specified if the action for the rule is "rate_based_ban".
	// If specified, determines the time (in seconds) the traffic will continue to be banned by the rate limit after the rate falls below the threshold.
	BanDurationSec *float64 `json:"banDurationSec,omitempty" tf:"ban_duration_sec,omitempty"`

	// Can only be specified if the action for the rule is "rate_based_ban".
	// If specified, the key will be banned for the configured 'ban_duration_sec' when the number of requests that exceed the 'rate_limit_threshold' also
	// exceed this 'ban_threshold'. Structure is documented below.
	BanThreshold []BanThresholdInitParameters `json:"banThreshold,omitempty" tf:"ban_threshold,omitempty"`

	// Action to take for requests that are under the configured rate limit threshold. Valid option is "allow" only.
	ConformAction *string `json:"conformAction,omitempty" tf:"conform_action,omitempty"`

	// Determines the key to enforce the rate_limit_threshold on. If not specified, defaults to "ALL".
	EnforceOnKey *string `json:"enforceOnKey,omitempty" tf:"enforce_on_key,omitempty"`

	// Rate limit key name applicable only for the following key types: HTTP_HEADER -- Name of the HTTP header whose value is taken as the key value. HTTP_COOKIE -- Name of the HTTP cookie whose value is taken as the key value.
	EnforceOnKeyName *string `json:"enforceOnKeyName,omitempty" tf:"enforce_on_key_name,omitempty"`

	// When a request is denied, returns the HTTP response code specified.
	// Valid options are "deny()" where valid values for status are 403, 404, 429, and 502.
	ExceedAction *string `json:"exceedAction,omitempty" tf:"exceed_action,omitempty"`

	// block supports:
	ExceedRedirectOptions []ExceedRedirectOptionsInitParameters `json:"exceedRedirectOptions,omitempty" tf:"exceed_redirect_options,omitempty"`

	// Threshold at which to begin ratelimiting. Structure is documented below.
	RateLimitThreshold []RateLimitThresholdInitParameters `json:"rateLimitThreshold,omitempty" tf:"rate_limit_threshold,omitempty"`
}

type RateLimitOptionsObservation struct {

	// Can only be specified if the action for the rule is "rate_based_ban".
	// If specified, determines the time (in seconds) the traffic will continue to be banned by the rate limit after the rate falls below the threshold.
	BanDurationSec *float64 `json:"banDurationSec,omitempty" tf:"ban_duration_sec,omitempty"`

	// Can only be specified if the action for the rule is "rate_based_ban".
	// If specified, the key will be banned for the configured 'ban_duration_sec' when the number of requests that exceed the 'rate_limit_threshold' also
	// exceed this 'ban_threshold'. Structure is documented below.
	BanThreshold []BanThresholdObservation `json:"banThreshold,omitempty" tf:"ban_threshold,omitempty"`

	// Action to take for requests that are under the configured rate limit threshold. Valid option is "allow" only.
	ConformAction *string `json:"conformAction,omitempty" tf:"conform_action,omitempty"`

	// Determines the key to enforce the rate_limit_threshold on. If not specified, defaults to "ALL".
	EnforceOnKey *string `json:"enforceOnKey,omitempty" tf:"enforce_on_key,omitempty"`

	// Rate limit key name applicable only for the following key types: HTTP_HEADER -- Name of the HTTP header whose value is taken as the key value. HTTP_COOKIE -- Name of the HTTP cookie whose value is taken as the key value.
	EnforceOnKeyName *string `json:"enforceOnKeyName,omitempty" tf:"enforce_on_key_name,omitempty"`

	// When a request is denied, returns the HTTP response code specified.
	// Valid options are "deny()" where valid values for status are 403, 404, 429, and 502.
	ExceedAction *string `json:"exceedAction,omitempty" tf:"exceed_action,omitempty"`

	// block supports:
	ExceedRedirectOptions []ExceedRedirectOptionsObservation `json:"exceedRedirectOptions,omitempty" tf:"exceed_redirect_options,omitempty"`

	// Threshold at which to begin ratelimiting. Structure is documented below.
	RateLimitThreshold []RateLimitThresholdObservation `json:"rateLimitThreshold,omitempty" tf:"rate_limit_threshold,omitempty"`
}

type RateLimitOptionsParameters struct {

	// Can only be specified if the action for the rule is "rate_based_ban".
	// If specified, determines the time (in seconds) the traffic will continue to be banned by the rate limit after the rate falls below the threshold.
	// +kubebuilder:validation:Optional
	BanDurationSec *float64 `json:"banDurationSec,omitempty" tf:"ban_duration_sec,omitempty"`

	// Can only be specified if the action for the rule is "rate_based_ban".
	// If specified, the key will be banned for the configured 'ban_duration_sec' when the number of requests that exceed the 'rate_limit_threshold' also
	// exceed this 'ban_threshold'. Structure is documented below.
	// +kubebuilder:validation:Optional
	BanThreshold []BanThresholdParameters `json:"banThreshold,omitempty" tf:"ban_threshold,omitempty"`

	// Action to take for requests that are under the configured rate limit threshold. Valid option is "allow" only.
	// +kubebuilder:validation:Optional
	ConformAction *string `json:"conformAction" tf:"conform_action,omitempty"`

	// Determines the key to enforce the rate_limit_threshold on. If not specified, defaults to "ALL".
	// +kubebuilder:validation:Optional
	EnforceOnKey *string `json:"enforceOnKey,omitempty" tf:"enforce_on_key,omitempty"`

	// Rate limit key name applicable only for the following key types: HTTP_HEADER -- Name of the HTTP header whose value is taken as the key value. HTTP_COOKIE -- Name of the HTTP cookie whose value is taken as the key value.
	// +kubebuilder:validation:Optional
	EnforceOnKeyName *string `json:"enforceOnKeyName,omitempty" tf:"enforce_on_key_name,omitempty"`

	// When a request is denied, returns the HTTP response code specified.
	// Valid options are "deny()" where valid values for status are 403, 404, 429, and 502.
	// +kubebuilder:validation:Optional
	ExceedAction *string `json:"exceedAction" tf:"exceed_action,omitempty"`

	// block supports:
	// +kubebuilder:validation:Optional
	ExceedRedirectOptions []ExceedRedirectOptionsParameters `json:"exceedRedirectOptions,omitempty" tf:"exceed_redirect_options,omitempty"`

	// Threshold at which to begin ratelimiting. Structure is documented below.
	// +kubebuilder:validation:Optional
	RateLimitThreshold []RateLimitThresholdParameters `json:"rateLimitThreshold" tf:"rate_limit_threshold,omitempty"`
}

type RateLimitThresholdInitParameters struct {

	// Number of HTTP(S) requests for calculating the threshold.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// Interval over which the threshold is computed.
	IntervalSec *float64 `json:"intervalSec,omitempty" tf:"interval_sec,omitempty"`
}

type RateLimitThresholdObservation struct {

	// Number of HTTP(S) requests for calculating the threshold.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// Interval over which the threshold is computed.
	IntervalSec *float64 `json:"intervalSec,omitempty" tf:"interval_sec,omitempty"`
}

type RateLimitThresholdParameters struct {

	// Number of HTTP(S) requests for calculating the threshold.
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count" tf:"count,omitempty"`

	// Interval over which the threshold is computed.
	// +kubebuilder:validation:Optional
	IntervalSec *float64 `json:"intervalSec" tf:"interval_sec,omitempty"`
}

type RecaptchaOptionsConfigInitParameters struct {

	// A field to supply a reCAPTCHA site key to be used for all the rules using the redirect action with the type of GOOGLE_RECAPTCHA under the security policy. The specified site key needs to be created from the reCAPTCHA API. The user is responsible for the validity of the specified site key. If not specified, a Google-managed site key is used.
	RedirectSiteKey *string `json:"redirectSiteKey,omitempty" tf:"redirect_site_key,omitempty"`
}

type RecaptchaOptionsConfigObservation struct {

	// A field to supply a reCAPTCHA site key to be used for all the rules using the redirect action with the type of GOOGLE_RECAPTCHA under the security policy. The specified site key needs to be created from the reCAPTCHA API. The user is responsible for the validity of the specified site key. If not specified, a Google-managed site key is used.
	RedirectSiteKey *string `json:"redirectSiteKey,omitempty" tf:"redirect_site_key,omitempty"`
}

type RecaptchaOptionsConfigParameters struct {

	// A field to supply a reCAPTCHA site key to be used for all the rules using the redirect action with the type of GOOGLE_RECAPTCHA under the security policy. The specified site key needs to be created from the reCAPTCHA API. The user is responsible for the validity of the specified site key. If not specified, a Google-managed site key is used.
	// +kubebuilder:validation:Optional
	RedirectSiteKey *string `json:"redirectSiteKey" tf:"redirect_site_key,omitempty"`
}

type RedirectOptionsInitParameters struct {

	// Target for the redirect action. This is required if the type is EXTERNAL_302 and cannot be specified for GOOGLE_RECAPTCHA.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Type of the redirect action.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RedirectOptionsObservation struct {

	// Target for the redirect action. This is required if the type is EXTERNAL_302 and cannot be specified for GOOGLE_RECAPTCHA.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Type of the redirect action.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RedirectOptionsParameters struct {

	// Target for the redirect action. This is required if the type is EXTERNAL_302 and cannot be specified for GOOGLE_RECAPTCHA.
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Type of the redirect action.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type RequestHeadersToAddsInitParameters struct {

	// The name of the header to set.
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`

	// The value to set the named header to.
	HeaderValue *string `json:"headerValue,omitempty" tf:"header_value,omitempty"`
}

type RequestHeadersToAddsObservation struct {

	// The name of the header to set.
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`

	// The value to set the named header to.
	HeaderValue *string `json:"headerValue,omitempty" tf:"header_value,omitempty"`
}

type RequestHeadersToAddsParameters struct {

	// The name of the header to set.
	// +kubebuilder:validation:Optional
	HeaderName *string `json:"headerName" tf:"header_name,omitempty"`

	// The value to set the named header to.
	// +kubebuilder:validation:Optional
	HeaderValue *string `json:"headerValue,omitempty" tf:"header_value,omitempty"`
}

type RuleHeaderActionInitParameters struct {

	// The list of request headers to add or overwrite if they're already present. Structure is documented below.
	RequestHeadersToAdds []RequestHeadersToAddsInitParameters `json:"requestHeadersToAdds,omitempty" tf:"request_headers_to_adds,omitempty"`
}

type RuleHeaderActionObservation struct {

	// The list of request headers to add or overwrite if they're already present. Structure is documented below.
	RequestHeadersToAdds []RequestHeadersToAddsObservation `json:"requestHeadersToAdds,omitempty" tf:"request_headers_to_adds,omitempty"`
}

type RuleHeaderActionParameters struct {

	// The list of request headers to add or overwrite if they're already present. Structure is documented below.
	// +kubebuilder:validation:Optional
	RequestHeadersToAdds []RequestHeadersToAddsParameters `json:"requestHeadersToAdds" tf:"request_headers_to_adds,omitempty"`
}

type RuleInitParameters struct {

	// Action to take when match matches the request. Valid values:
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// An optional description of this rule. Max size is 64.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Additional actions that are performed on headers. Structure is documented below.
	HeaderAction []RuleHeaderActionInitParameters `json:"headerAction,omitempty" tf:"header_action,omitempty"`

	// A match condition that incoming traffic is evaluated against.
	// If it evaluates to true, the corresponding action is enforced. Structure is documented below.
	Match []RuleMatchInitParameters `json:"match,omitempty" tf:"match,omitempty"`

	// When set to true, the action specified above is not enforced.
	// Stackdriver logs for requests that trigger a preview action are annotated as such.
	Preview *bool `json:"preview,omitempty" tf:"preview,omitempty"`

	// An unique positive integer indicating the priority of evaluation for a rule.
	// Rules are evaluated from highest priority (lowest numerically) to lowest priority (highest numerically) in order.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Must be specified if the action is "rate_based_ban" or "throttle". Cannot be specified for other actions. Structure is documented below.
	RateLimitOptions []RateLimitOptionsInitParameters `json:"rateLimitOptions,omitempty" tf:"rate_limit_options,omitempty"`

	// Can be specified if the action is "redirect". Cannot be specified for other actions. Structure is documented below.
	RedirectOptions []RedirectOptionsInitParameters `json:"redirectOptions,omitempty" tf:"redirect_options,omitempty"`
}

type RuleMatchInitParameters struct {

	// The configuration options available when specifying versioned_expr.
	// This field must be specified if versioned_expr is specified and cannot be specified if versioned_expr is not specified.
	// Structure is documented below.
	Config []ConfigInitParameters `json:"config,omitempty" tf:"config,omitempty"`

	// User defined CEVAL expression. A CEVAL expression is used to specify match criteria
	// such as origin.ip, source.region_code and contents in the request header.
	// Structure is documented below.
	Expr []ExprInitParameters `json:"expr,omitempty" tf:"expr,omitempty"`

	// Predefined rule expression. If this field is specified, config must also be specified.
	// Available options:
	VersionedExpr *string `json:"versionedExpr,omitempty" tf:"versioned_expr,omitempty"`
}

type RuleMatchObservation struct {

	// The configuration options available when specifying versioned_expr.
	// This field must be specified if versioned_expr is specified and cannot be specified if versioned_expr is not specified.
	// Structure is documented below.
	Config []ConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	// User defined CEVAL expression. A CEVAL expression is used to specify match criteria
	// such as origin.ip, source.region_code and contents in the request header.
	// Structure is documented below.
	Expr []ExprObservation `json:"expr,omitempty" tf:"expr,omitempty"`

	// Predefined rule expression. If this field is specified, config must also be specified.
	// Available options:
	VersionedExpr *string `json:"versionedExpr,omitempty" tf:"versioned_expr,omitempty"`
}

type RuleMatchParameters struct {

	// The configuration options available when specifying versioned_expr.
	// This field must be specified if versioned_expr is specified and cannot be specified if versioned_expr is not specified.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Config []ConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// User defined CEVAL expression. A CEVAL expression is used to specify match criteria
	// such as origin.ip, source.region_code and contents in the request header.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Expr []ExprParameters `json:"expr,omitempty" tf:"expr,omitempty"`

	// Predefined rule expression. If this field is specified, config must also be specified.
	// Available options:
	// +kubebuilder:validation:Optional
	VersionedExpr *string `json:"versionedExpr,omitempty" tf:"versioned_expr,omitempty"`
}

type RuleObservation struct {

	// Action to take when match matches the request. Valid values:
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// An optional description of this rule. Max size is 64.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Additional actions that are performed on headers. Structure is documented below.
	HeaderAction []RuleHeaderActionObservation `json:"headerAction,omitempty" tf:"header_action,omitempty"`

	// A match condition that incoming traffic is evaluated against.
	// If it evaluates to true, the corresponding action is enforced. Structure is documented below.
	Match []RuleMatchObservation `json:"match,omitempty" tf:"match,omitempty"`

	// When set to true, the action specified above is not enforced.
	// Stackdriver logs for requests that trigger a preview action are annotated as such.
	Preview *bool `json:"preview,omitempty" tf:"preview,omitempty"`

	// An unique positive integer indicating the priority of evaluation for a rule.
	// Rules are evaluated from highest priority (lowest numerically) to lowest priority (highest numerically) in order.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Must be specified if the action is "rate_based_ban" or "throttle". Cannot be specified for other actions. Structure is documented below.
	RateLimitOptions []RateLimitOptionsObservation `json:"rateLimitOptions,omitempty" tf:"rate_limit_options,omitempty"`

	// Can be specified if the action is "redirect". Cannot be specified for other actions. Structure is documented below.
	RedirectOptions []RedirectOptionsObservation `json:"redirectOptions,omitempty" tf:"redirect_options,omitempty"`
}

type RuleParameters struct {

	// Action to take when match matches the request. Valid values:
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// An optional description of this rule. Max size is 64.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Additional actions that are performed on headers. Structure is documented below.
	// +kubebuilder:validation:Optional
	HeaderAction []RuleHeaderActionParameters `json:"headerAction,omitempty" tf:"header_action,omitempty"`

	// A match condition that incoming traffic is evaluated against.
	// If it evaluates to true, the corresponding action is enforced. Structure is documented below.
	// +kubebuilder:validation:Optional
	Match []RuleMatchParameters `json:"match" tf:"match,omitempty"`

	// When set to true, the action specified above is not enforced.
	// Stackdriver logs for requests that trigger a preview action are annotated as such.
	// +kubebuilder:validation:Optional
	Preview *bool `json:"preview,omitempty" tf:"preview,omitempty"`

	// An unique positive integer indicating the priority of evaluation for a rule.
	// Rules are evaluated from highest priority (lowest numerically) to lowest priority (highest numerically) in order.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// Must be specified if the action is "rate_based_ban" or "throttle". Cannot be specified for other actions. Structure is documented below.
	// +kubebuilder:validation:Optional
	RateLimitOptions []RateLimitOptionsParameters `json:"rateLimitOptions,omitempty" tf:"rate_limit_options,omitempty"`

	// Can be specified if the action is "redirect". Cannot be specified for other actions. Structure is documented below.
	// +kubebuilder:validation:Optional
	RedirectOptions []RedirectOptionsParameters `json:"redirectOptions,omitempty" tf:"redirect_options,omitempty"`
}

type SecurityPolicyInitParameters struct {

	// Configuration for Google Cloud Armor Adaptive Protection. Structure is documented below.
	AdaptiveProtectionConfig []AdaptiveProtectionConfigInitParameters `json:"adaptiveProtectionConfig,omitempty" tf:"adaptive_protection_config,omitempty"`

	// Advanced Configuration Options.
	// Structure is documented below.
	AdvancedOptionsConfig []AdvancedOptionsConfigInitParameters `json:"advancedOptionsConfig,omitempty" tf:"advanced_options_config,omitempty"`

	// An optional description of this security policy. Max size is 2048.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// reCAPTCHA Configuration Options. Structure is documented below.
	RecaptchaOptionsConfig []RecaptchaOptionsConfigInitParameters `json:"recaptchaOptionsConfig,omitempty" tf:"recaptcha_options_config,omitempty"`

	// The set of rules that belong to this policy. There must always be a default
	// rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a
	// security policy, a default rule with action "allow" will be added. Structure is documented below.
	Rule []RuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// The type indicates the intended use of the security policy. This field can be set only at resource creation time.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SecurityPolicyObservation struct {

	// Configuration for Google Cloud Armor Adaptive Protection. Structure is documented below.
	AdaptiveProtectionConfig []AdaptiveProtectionConfigObservation `json:"adaptiveProtectionConfig,omitempty" tf:"adaptive_protection_config,omitempty"`

	// Advanced Configuration Options.
	// Structure is documented below.
	AdvancedOptionsConfig []AdvancedOptionsConfigObservation `json:"advancedOptionsConfig,omitempty" tf:"advanced_options_config,omitempty"`

	// An optional description of this security policy. Max size is 2048.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Fingerprint of this resource.
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/securityPolicies/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// reCAPTCHA Configuration Options. Structure is documented below.
	RecaptchaOptionsConfig []RecaptchaOptionsConfigObservation `json:"recaptchaOptionsConfig,omitempty" tf:"recaptcha_options_config,omitempty"`

	// The set of rules that belong to this policy. There must always be a default
	// rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a
	// security policy, a default rule with action "allow" will be added. Structure is documented below.
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The type indicates the intended use of the security policy. This field can be set only at resource creation time.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SecurityPolicyParameters struct {

	// Configuration for Google Cloud Armor Adaptive Protection. Structure is documented below.
	// +kubebuilder:validation:Optional
	AdaptiveProtectionConfig []AdaptiveProtectionConfigParameters `json:"adaptiveProtectionConfig,omitempty" tf:"adaptive_protection_config,omitempty"`

	// Advanced Configuration Options.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AdvancedOptionsConfig []AdvancedOptionsConfigParameters `json:"advancedOptionsConfig,omitempty" tf:"advanced_options_config,omitempty"`

	// An optional description of this security policy. Max size is 2048.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// reCAPTCHA Configuration Options. Structure is documented below.
	// +kubebuilder:validation:Optional
	RecaptchaOptionsConfig []RecaptchaOptionsConfigParameters `json:"recaptchaOptionsConfig,omitempty" tf:"recaptcha_options_config,omitempty"`

	// The set of rules that belong to this policy. There must always be a default
	// rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a
	// security policy, a default rule with action "allow" will be added. Structure is documented below.
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// The type indicates the intended use of the security policy. This field can be set only at resource creation time.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// SecurityPolicySpec defines the desired state of SecurityPolicy
type SecurityPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityPolicyParameters `json:"forProvider"`
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
	InitProvider SecurityPolicyInitParameters `json:"initProvider,omitempty"`
}

// SecurityPolicyStatus defines the observed state of SecurityPolicy.
type SecurityPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecurityPolicy is the Schema for the SecurityPolicys API. Creates a Security Policy resource for Google Compute Engine.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type SecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityPolicySpec   `json:"spec"`
	Status            SecurityPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityPolicyList contains a list of SecurityPolicys
type SecurityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityPolicy `json:"items"`
}

// Repository type metadata.
var (
	SecurityPolicy_Kind             = "SecurityPolicy"
	SecurityPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityPolicy_Kind}.String()
	SecurityPolicy_KindAPIVersion   = SecurityPolicy_Kind + "." + CRDGroupVersion.String()
	SecurityPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SecurityPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityPolicy{}, &SecurityPolicyList{})
}
