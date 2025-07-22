// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datanewreliccloudaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataNewrelicCloudAccountConfig struct {
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
	// The cloud provider of the account, e.g. aws, gcp, azure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.64.0/docs/data-sources/cloud_account#cloud_provider DataNewrelicCloudAccount#cloud_provider}
	CloudProvider *string `field:"required" json:"cloudProvider" yaml:"cloudProvider"`
	// The name of the cloud account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.64.0/docs/data-sources/cloud_account#name DataNewrelicCloudAccount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the New Relic account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.64.0/docs/data-sources/cloud_account#account_id DataNewrelicCloudAccount#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.64.0/docs/data-sources/cloud_account#id DataNewrelicCloudAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

