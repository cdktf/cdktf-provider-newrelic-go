package cloudazureintegrations


type CloudAzureIntegrationsRedisCache struct {
	// The data polling interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_azure_integrations#metrics_polling_interval CloudAzureIntegrations#metrics_polling_interval}
	MetricsPollingInterval *float64 `field:"optional" json:"metricsPollingInterval" yaml:"metricsPollingInterval"`
	// Specify each Resource group associated with the resources that you want to monitor. Filter values are case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_azure_integrations#resource_groups CloudAzureIntegrations#resource_groups}
	ResourceGroups *[]*string `field:"optional" json:"resourceGroups" yaml:"resourceGroups"`
}

