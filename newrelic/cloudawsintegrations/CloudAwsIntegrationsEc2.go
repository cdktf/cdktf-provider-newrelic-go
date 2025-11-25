// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudawsintegrations


type CloudAwsIntegrationsEc2 struct {
	// Specify each AWS region that includes the resources that you want to monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.1/docs/resources/cloud_aws_integrations#aws_regions CloudAwsIntegrations#aws_regions}
	AwsRegions *[]*string `field:"optional" json:"awsRegions" yaml:"awsRegions"`
	// Specify if the old legacy metadata and tag names have to be kept, it will consume more ingest data size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.1/docs/resources/cloud_aws_integrations#duplicate_ec2_tags CloudAwsIntegrations#duplicate_ec2_tags}
	DuplicateEc2Tags interface{} `field:"optional" json:"duplicateEc2Tags" yaml:"duplicateEc2Tags"`
	// Specify if IP addresses of ec2 instance should be collected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.1/docs/resources/cloud_aws_integrations#fetch_ip_addresses CloudAwsIntegrations#fetch_ip_addresses}
	FetchIpAddresses interface{} `field:"optional" json:"fetchIpAddresses" yaml:"fetchIpAddresses"`
	// The data polling interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.1/docs/resources/cloud_aws_integrations#metrics_polling_interval CloudAwsIntegrations#metrics_polling_interval}
	MetricsPollingInterval *float64 `field:"optional" json:"metricsPollingInterval" yaml:"metricsPollingInterval"`
	// Specify a Tag key associated with the resources that you want to monitor. Filter values are case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.1/docs/resources/cloud_aws_integrations#tag_key CloudAwsIntegrations#tag_key}
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// Specify a Tag value associated with the resources that you want to monitor. Filter values are case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.1/docs/resources/cloud_aws_integrations#tag_value CloudAwsIntegrations#tag_value}
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}

