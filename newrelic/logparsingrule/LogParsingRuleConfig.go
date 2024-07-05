// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logparsingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogParsingRuleConfig struct {
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
	// Whether or not this rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/log_parsing_rule#enabled LogParsingRule#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The Grok of what to parse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/log_parsing_rule#grok LogParsingRule#grok}
	Grok *string `field:"required" json:"grok" yaml:"grok"`
	// The Lucene to match events to the parsing rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/log_parsing_rule#lucene LogParsingRule#lucene}
	Lucene *string `field:"required" json:"lucene" yaml:"lucene"`
	// A description of what this parsing rule represents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/log_parsing_rule#name LogParsingRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The NRQL to match events to the parsing rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/log_parsing_rule#nrql LogParsingRule#nrql}
	Nrql *string `field:"required" json:"nrql" yaml:"nrql"`
	// The account id associated with the obfuscation expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/log_parsing_rule#account_id LogParsingRule#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// The parsing rule will apply to value of this attribute.
	//
	// If field is not provided, value will default to message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/log_parsing_rule#attribute LogParsingRule#attribute}
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/log_parsing_rule#id LogParsingRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the Grok pattern matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/log_parsing_rule#matched LogParsingRule#matched}
	Matched interface{} `field:"optional" json:"matched" yaml:"matched"`
}

