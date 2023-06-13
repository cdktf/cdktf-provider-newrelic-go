package cloudawsintegrations


type CloudAwsIntegrationsBilling struct {
	// The data polling interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/cloud_aws_integrations#metrics_polling_interval CloudAwsIntegrations#metrics_polling_interval}
	MetricsPollingInterval *float64 `field:"optional" json:"metricsPollingInterval" yaml:"metricsPollingInterval"`
}

