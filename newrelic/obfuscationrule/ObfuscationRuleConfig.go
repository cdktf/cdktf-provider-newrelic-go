// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obfuscationrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObfuscationRuleConfig struct {
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
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.73.0/docs/resources/obfuscation_rule#action ObfuscationRule#action}
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// Whether the rule should be applied or not to incoming data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.73.0/docs/resources/obfuscation_rule#enabled ObfuscationRule#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// NRQL for determining whether a given log record should have obfuscation actions applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.73.0/docs/resources/obfuscation_rule#filter ObfuscationRule#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Name of rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.73.0/docs/resources/obfuscation_rule#name ObfuscationRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The account id associated with the obfuscation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.73.0/docs/resources/obfuscation_rule#account_id ObfuscationRule#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Description of rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.73.0/docs/resources/obfuscation_rule#description ObfuscationRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.73.0/docs/resources/obfuscation_rule#id ObfuscationRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

