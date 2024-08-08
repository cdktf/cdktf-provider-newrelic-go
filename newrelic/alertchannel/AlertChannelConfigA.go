// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertchannel


type AlertChannelConfigA struct {
	// The API key for integrating with OpsGenie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#api_key AlertChannel#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Specifies an authentication password for use with a channel. Supported by the webhook channel type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#auth_password AlertChannel#auth_password}
	AuthPassword *string `field:"optional" json:"authPassword" yaml:"authPassword"`
	// Specifies an authentication method for use with a channel.
	//
	// Supported by the webhook channel type. Only HTTP basic authentication is currently supported via the value BASIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#auth_type AlertChannel#auth_type}
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// Specifies an authentication username for use with a channel. Supported by the webhook channel type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#auth_username AlertChannel#auth_username}
	AuthUsername *string `field:"optional" json:"authUsername" yaml:"authUsername"`
	// The base URL of the webhook destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#base_url AlertChannel#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// The Slack channel to send notifications to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#channel AlertChannel#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// A map of key/value pairs that represents extra HTTP headers to be sent along with the webhook payload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#headers AlertChannel#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// Use instead of headers if the desired payload is more complex than a list of key/value pairs (e.g. a set of headers that makes use of nested objects). The value provided should be a valid JSON string with escaped double quotes. Conflicts with headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#headers_string AlertChannel#headers_string}
	HeadersString *string `field:"optional" json:"headersString" yaml:"headersString"`
	// true or false.
	//
	// Flag for whether or not to attach a JSON document containing information about the associated alert to the email that is sent to recipients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#include_json_attachment AlertChannel#include_json_attachment}
	IncludeJsonAttachment *string `field:"optional" json:"includeJsonAttachment" yaml:"includeJsonAttachment"`
	// The key for integrating with VictorOps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#key AlertChannel#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A map of key/value pairs that represents the webhook payload. Must provide payload_type if setting this argument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#payload AlertChannel#payload}
	Payload *map[string]*string `field:"optional" json:"payload" yaml:"payload"`
	// Use instead of payload if the desired payload is more complex than a list of key/value pairs (e.g. a payload that makes use of nested objects). The value provided should be a valid JSON string with escaped double quotes. Conflicts with payload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#payload_string AlertChannel#payload_string}
	PayloadString *string `field:"optional" json:"payloadString" yaml:"payloadString"`
	// Can either be application/json or application/x-www-form-urlencoded. The payload_type argument is required if payload is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#payload_type AlertChannel#payload_type}
	PayloadType *string `field:"optional" json:"payloadType" yaml:"payloadType"`
	// A set of recipients for targeting notifications. Multiple values are comma separated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#recipients AlertChannel#recipients}
	Recipients *string `field:"optional" json:"recipients" yaml:"recipients"`
	// The data center region to store your data. Valid values are US and EU. Default is US.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#region AlertChannel#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The route key for integrating with VictorOps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#route_key AlertChannel#route_key}
	RouteKey *string `field:"optional" json:"routeKey" yaml:"routeKey"`
	// Specifies the service key for integrating with Pagerduty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#service_key AlertChannel#service_key}
	ServiceKey *string `field:"optional" json:"serviceKey" yaml:"serviceKey"`
	// A set of tags for targeting notifications. Multiple values are comma separated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#tags AlertChannel#tags}
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
	// A set of teams for targeting notifications. Multiple values are comma separated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#teams AlertChannel#teams}
	Teams *string `field:"optional" json:"teams" yaml:"teams"`
	// Your organization's Slack URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#url AlertChannel#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The user ID for use with the user channel type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.41.1/docs/resources/alert_channel#user_id AlertChannel#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

