// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudazureintegrations


type CloudAzureIntegrationsMonitor struct {
	// A flag that specifies if the integration is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.0/docs/resources/cloud_azure_integrations#enabled CloudAzureIntegrations#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specify resource tags in 'key:value' form to be excluded from monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.0/docs/resources/cloud_azure_integrations#exclude_tags CloudAzureIntegrations#exclude_tags}
	ExcludeTags *[]*string `field:"optional" json:"excludeTags" yaml:"excludeTags"`
	// Specify resource tags in 'key:value' form to be monitored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.0/docs/resources/cloud_azure_integrations#include_tags CloudAzureIntegrations#include_tags}
	IncludeTags *[]*string `field:"optional" json:"includeTags" yaml:"includeTags"`
	// The data polling interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.0/docs/resources/cloud_azure_integrations#metrics_polling_interval CloudAzureIntegrations#metrics_polling_interval}
	MetricsPollingInterval *float64 `field:"optional" json:"metricsPollingInterval" yaml:"metricsPollingInterval"`
	// Specify each Resource group associated with the resources that you want to monitor. Filter values are case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.0/docs/resources/cloud_azure_integrations#resource_groups CloudAzureIntegrations#resource_groups}
	ResourceGroups *[]*string `field:"optional" json:"resourceGroups" yaml:"resourceGroups"`
	// Specify each Azure resource type that needs to be monitored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.0/docs/resources/cloud_azure_integrations#resource_types CloudAzureIntegrations#resource_types}
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

