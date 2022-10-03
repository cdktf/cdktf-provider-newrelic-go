package cloudazureintegrations


type CloudAzureIntegrationsCostManagement struct {
	// The data polling interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_azure_integrations#metrics_polling_interval CloudAzureIntegrations#metrics_polling_interval}
	MetricsPollingInterval *float64 `field:"optional" json:"metricsPollingInterval" yaml:"metricsPollingInterval"`
	// Specify if additional cost data per tag should be collected. This field is case sensitive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_azure_integrations#tag_keys CloudAzureIntegrations#tag_keys}
	TagKeys *[]*string `field:"optional" json:"tagKeys" yaml:"tagKeys"`
}

