// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudgcpintegrations


type CloudGcpIntegrationsRouter struct {
	// the data polling interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.52.1/docs/resources/cloud_gcp_integrations#metrics_polling_interval CloudGcpIntegrations#metrics_polling_interval}
	MetricsPollingInterval *float64 `field:"optional" json:"metricsPollingInterval" yaml:"metricsPollingInterval"`
}

