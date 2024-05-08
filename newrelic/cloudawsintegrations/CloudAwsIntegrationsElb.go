// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudawsintegrations


type CloudAwsIntegrationsElb struct {
	// Specify each AWS region that includes the resources that you want to monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.2/docs/resources/cloud_aws_integrations#aws_regions CloudAwsIntegrations#aws_regions}
	AwsRegions *[]*string `field:"optional" json:"awsRegions" yaml:"awsRegions"`
	// Determine if extra inventory data be collected or not.
	//
	// May affect total data collection time and contribute to the Cloud provider API rate limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.2/docs/resources/cloud_aws_integrations#fetch_extended_inventory CloudAwsIntegrations#fetch_extended_inventory}
	FetchExtendedInventory interface{} `field:"optional" json:"fetchExtendedInventory" yaml:"fetchExtendedInventory"`
	// Specify if tags should be collected.
	//
	// May affect total data collection time and contribute to the Cloud provider API rate limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.2/docs/resources/cloud_aws_integrations#fetch_tags CloudAwsIntegrations#fetch_tags}
	FetchTags interface{} `field:"optional" json:"fetchTags" yaml:"fetchTags"`
	// The data polling interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.2/docs/resources/cloud_aws_integrations#metrics_polling_interval CloudAwsIntegrations#metrics_polling_interval}
	MetricsPollingInterval *float64 `field:"optional" json:"metricsPollingInterval" yaml:"metricsPollingInterval"`
}

