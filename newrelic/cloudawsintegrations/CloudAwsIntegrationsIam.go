// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudawsintegrations


type CloudAwsIntegrationsIam struct {
	// The data polling interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.0/docs/resources/cloud_aws_integrations#metrics_polling_interval CloudAwsIntegrations#metrics_polling_interval}
	MetricsPollingInterval *float64 `field:"optional" json:"metricsPollingInterval" yaml:"metricsPollingInterval"`
	// Specify a Tag key associated with the resources that you want to monitor. Filter values are case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.0/docs/resources/cloud_aws_integrations#tag_key CloudAwsIntegrations#tag_key}
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// Specify a Tag value associated with the resources that you want to monitor. Filter values are case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.0/docs/resources/cloud_aws_integrations#tag_value CloudAwsIntegrations#tag_value}
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}

