// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obfuscationexpression

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObfuscationExpressionConfig struct {
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
	// Name of expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.74.0/docs/resources/obfuscation_expression#name ObfuscationExpression#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Regex of expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.74.0/docs/resources/obfuscation_expression#regex ObfuscationExpression#regex}
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// The account id associated with the obfuscation expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.74.0/docs/resources/obfuscation_expression#account_id ObfuscationExpression#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Description of expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.74.0/docs/resources/obfuscation_expression#description ObfuscationExpression#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.74.0/docs/resources/obfuscation_expression#id ObfuscationExpression#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

