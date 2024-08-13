// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsprivatelocation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsPrivateLocationConfig struct {
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
	// The private location description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.42.0/docs/resources/synthetics_private_location#description SyntheticsPrivateLocation#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the private location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.42.0/docs/resources/synthetics_private_location#name SyntheticsPrivateLocation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.42.0/docs/resources/synthetics_private_location#account_id SyntheticsPrivateLocation#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.42.0/docs/resources/synthetics_private_location#id SyntheticsPrivateLocation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The private location requires a password to edit if value is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.42.0/docs/resources/synthetics_private_location#verified_script_execution SyntheticsPrivateLocation#verified_script_execution}
	VerifiedScriptExecution interface{} `field:"optional" json:"verifiedScriptExecution" yaml:"verifiedScriptExecution"`
}

