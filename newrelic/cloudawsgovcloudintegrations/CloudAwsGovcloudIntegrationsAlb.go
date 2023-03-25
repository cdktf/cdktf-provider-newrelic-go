package cloudawsgovcloudintegrations


type CloudAwsGovcloudIntegrationsAlb struct {
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
	// Specify if tags should be collected.
	//
	// May affect total data collection time and contribute to the Cloud provider API rate limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_govcloud_integrations#fetch_tags CloudAwsGovcloudIntegrations#fetch_tags}
	FetchTags interface{} `field:"optional" json:"fetchTags" yaml:"fetchTags"`
	// Specify each name or prefix for the LBs that you want to monitor. Filter values are case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_govcloud_integrations#load_balancer_prefixes CloudAwsGovcloudIntegrations#load_balancer_prefixes}
	LoadBalancerPrefixes *[]*string `field:"optional" json:"loadBalancerPrefixes" yaml:"loadBalancerPrefixes"`
	// The data polling interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_govcloud_integrations#metrics_polling_interval CloudAwsGovcloudIntegrations#metrics_polling_interval}
	MetricsPollingInterval *float64 `field:"optional" json:"metricsPollingInterval" yaml:"metricsPollingInterval"`
	// Specify a Tag key associated with the resources that you want to monitor. Filter values are case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_govcloud_integrations#tag_key CloudAwsGovcloudIntegrations#tag_key}
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// Specify a Tag value associated with the resources that you want to monitor. Filter values are case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_govcloud_integrations#tag_value CloudAwsGovcloudIntegrations#tag_value}
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}
