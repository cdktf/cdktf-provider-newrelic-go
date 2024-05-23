// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudawsgovcloudintegrations


type CloudAwsGovcloudIntegrationsRedShift struct {
	// Specify each AWS region that includes the resources that you want to monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.1/docs/resources/cloud_aws_govcloud_integrations#aws_regions CloudAwsGovcloudIntegrations#aws_regions}
	AwsRegions *[]*string `field:"optional" json:"awsRegions" yaml:"awsRegions"`
	// The data polling interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.1/docs/resources/cloud_aws_govcloud_integrations#metrics_polling_interval CloudAwsGovcloudIntegrations#metrics_polling_interval}
	MetricsPollingInterval *float64 `field:"optional" json:"metricsPollingInterval" yaml:"metricsPollingInterval"`
	// Specify a Tag key associated with the resources that you want to monitor. Filter values are case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.1/docs/resources/cloud_aws_govcloud_integrations#tag_key CloudAwsGovcloudIntegrations#tag_key}
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// Specify a Tag value associated with the resources that you want to monitor. Filter values are case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.1/docs/resources/cloud_aws_govcloud_integrations#tag_value CloudAwsGovcloudIntegrations#tag_value}
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}

