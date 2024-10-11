// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventstometricsrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventsToMetricsRuleConfig struct {
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
	// The name of the rule. This must be unique within an account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.50.0/docs/resources/events_to_metrics_rule#name EventsToMetricsRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Explains how to create metrics from events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.50.0/docs/resources/events_to_metrics_rule#nrql EventsToMetricsRule#nrql}
	Nrql *string `field:"required" json:"nrql" yaml:"nrql"`
	// Account with the event and where the metrics will be put.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.50.0/docs/resources/events_to_metrics_rule#account_id EventsToMetricsRule#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Provides additional information about the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.50.0/docs/resources/events_to_metrics_rule#description EventsToMetricsRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// True means this rule is enabled. False means the rule is currently not creating metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.50.0/docs/resources/events_to_metrics_rule#enabled EventsToMetricsRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.50.0/docs/resources/events_to_metrics_rule#id EventsToMetricsRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

