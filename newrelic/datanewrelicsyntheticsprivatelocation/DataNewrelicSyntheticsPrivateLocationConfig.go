// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datanewrelicsyntheticsprivatelocation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataNewrelicSyntheticsPrivateLocationConfig struct {
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
	// The name of the Synthetics monitor private location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/data-sources/synthetics_private_location#name DataNewrelicSyntheticsPrivateLocation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/data-sources/synthetics_private_location#account_id DataNewrelicSyntheticsPrivateLocation#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/data-sources/synthetics_private_location#id DataNewrelicSyntheticsPrivateLocation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The key of the queried private location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/data-sources/synthetics_private_location#key DataNewrelicSyntheticsPrivateLocation#key}
	Key *[]*string `field:"optional" json:"key" yaml:"key"`
}

