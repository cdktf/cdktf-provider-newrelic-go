// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudawsgovcloudlinkaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudAwsGovcloudLinkAccountConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// access-key-id of awsGovcloud account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.46.0/docs/resources/cloud_aws_govcloud_link_account#access_key_id CloudAwsGovcloudLinkAccount#access_key_id}
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// awsGovcloud account id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.46.0/docs/resources/cloud_aws_govcloud_link_account#aws_account_id CloudAwsGovcloudLinkAccount#aws_account_id}
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// name of the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.46.0/docs/resources/cloud_aws_govcloud_link_account#name CloudAwsGovcloudLinkAccount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// secret access key of the awsGovcloud account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.46.0/docs/resources/cloud_aws_govcloud_link_account#secret_access_key CloudAwsGovcloudLinkAccount#secret_access_key}
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
	// The ID of the account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.46.0/docs/resources/cloud_aws_govcloud_link_account#account_id CloudAwsGovcloudLinkAccount#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.46.0/docs/resources/cloud_aws_govcloud_link_account#id CloudAwsGovcloudLinkAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// push or pull.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.46.0/docs/resources/cloud_aws_govcloud_link_account#metric_collection_mode CloudAwsGovcloudLinkAccount#metric_collection_mode}
	MetricCollectionMode *string `field:"optional" json:"metricCollectionMode" yaml:"metricCollectionMode"`
}

