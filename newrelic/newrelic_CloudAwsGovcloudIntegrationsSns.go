// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type CloudAwsGovcloudIntegrationsSns struct {
	// Specify each AWS region that includes the resources that you want to monitor.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_govcloud_integrations#aws_regions CloudAwsGovcloudIntegrations#aws_regions}
	AwsRegions *[]*string `field:"optional" json:"awsRegions" yaml:"awsRegions"`
	// Determine if extra inventory data be collected or not.
	//
	// May affect total data collection time and contribute to the Cloud provider API rate limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_govcloud_integrations#fetch_extended_inventory CloudAwsGovcloudIntegrations#fetch_extended_inventory}
	FetchExtendedInventory interface{} `field:"optional" json:"fetchExtendedInventory" yaml:"fetchExtendedInventory"`
	// The data polling interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_govcloud_integrations#metrics_polling_interval CloudAwsGovcloudIntegrations#metrics_polling_interval}
	MetricsPollingInterval *float64 `field:"optional" json:"metricsPollingInterval" yaml:"metricsPollingInterval"`
}
