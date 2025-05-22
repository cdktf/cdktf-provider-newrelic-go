// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type NewrelicProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs#account_id NewrelicProvider#account_id}.
	AccountId *float64 `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs#api_key NewrelicProvider#api_key}.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs#admin_api_key NewrelicProvider#admin_api_key}.
	AdminApiKey *string `field:"optional" json:"adminApiKey" yaml:"adminApiKey"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs#alias NewrelicProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs#api_url NewrelicProvider#api_url}.
	ApiUrl *string `field:"optional" json:"apiUrl" yaml:"apiUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs#cacert_file NewrelicProvider#cacert_file}.
	CacertFile *string `field:"optional" json:"cacertFile" yaml:"cacertFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs#infrastructure_api_url NewrelicProvider#infrastructure_api_url}.
	InfrastructureApiUrl *string `field:"optional" json:"infrastructureApiUrl" yaml:"infrastructureApiUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs#insecure_skip_verify NewrelicProvider#insecure_skip_verify}.
	InsecureSkipVerify interface{} `field:"optional" json:"insecureSkipVerify" yaml:"insecureSkipVerify"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs#insights_insert_key NewrelicProvider#insights_insert_key}.
	InsightsInsertKey *string `field:"optional" json:"insightsInsertKey" yaml:"insightsInsertKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs#insights_insert_url NewrelicProvider#insights_insert_url}.
	InsightsInsertUrl *string `field:"optional" json:"insightsInsertUrl" yaml:"insightsInsertUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs#insights_query_url NewrelicProvider#insights_query_url}.
	InsightsQueryUrl *string `field:"optional" json:"insightsQueryUrl" yaml:"insightsQueryUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs#nerdgraph_api_url NewrelicProvider#nerdgraph_api_url}.
	NerdgraphApiUrl *string `field:"optional" json:"nerdgraphApiUrl" yaml:"nerdgraphApiUrl"`
	// The data center for which your New Relic account is configured. Only one region per provider block is permitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs#region NewrelicProvider#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs#synthetics_api_url NewrelicProvider#synthetics_api_url}.
	SyntheticsApiUrl *string `field:"optional" json:"syntheticsApiUrl" yaml:"syntheticsApiUrl"`
}

