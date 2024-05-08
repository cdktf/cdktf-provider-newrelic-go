// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nrqldroprule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NrqlDropRuleConfig struct {
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
	// The drop rule action (drop_data, drop_attributes, or drop_attributes_from_metric_aggregates).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.2/docs/resources/nrql_drop_rule#action NrqlDropRule#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Explains which data to apply the drop rule to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.2/docs/resources/nrql_drop_rule#nrql NrqlDropRule#nrql}
	Nrql *string `field:"required" json:"nrql" yaml:"nrql"`
	// Account with the NRQL drop rule will be put.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.2/docs/resources/nrql_drop_rule#account_id NrqlDropRule#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Provides additional information about the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.2/docs/resources/nrql_drop_rule#description NrqlDropRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.2/docs/resources/nrql_drop_rule#id NrqlDropRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

