// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudgcpintegrations


type CloudGcpIntegrationsBigQuery struct {
	// to fetch tags of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/resources/cloud_gcp_integrations#fetch_tags CloudGcpIntegrations#fetch_tags}
	FetchTags interface{} `field:"optional" json:"fetchTags" yaml:"fetchTags"`
	// the data polling interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/resources/cloud_gcp_integrations#metrics_polling_interval CloudGcpIntegrations#metrics_polling_interval}
	MetricsPollingInterval *float64 `field:"optional" json:"metricsPollingInterval" yaml:"metricsPollingInterval"`
}

