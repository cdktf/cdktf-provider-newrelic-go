// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workflow

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkflowConfig struct {
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
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/workflow#destination Workflow#destination}
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// issues_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/workflow#issues_filter Workflow#issues_filter}
	IssuesFilter *WorkflowIssuesFilter `field:"required" json:"issuesFilter" yaml:"issuesFilter"`
	// The type of the muting rule handling. One of: (NOTIFY_ALL_ISSUES, DONT_NOTIFY_FULLY_MUTED_ISSUES, DONT_NOTIFY_FULLY_OR_PARTIALLY_MUTED_ISSUES).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/workflow#muting_rules_handling Workflow#muting_rules_handling}
	MutingRulesHandling *string `field:"required" json:"mutingRulesHandling" yaml:"mutingRulesHandling"`
	// (Required) The name of the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/workflow#name Workflow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The account id of the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/workflow#account_id Workflow#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Indicates whether the destinations are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/workflow#destinations_enabled Workflow#destinations_enabled}
	DestinationsEnabled interface{} `field:"optional" json:"destinationsEnabled" yaml:"destinationsEnabled"`
	// Indicates whether the workflow is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/workflow#enabled Workflow#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// enrichments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/workflow#enrichments Workflow#enrichments}
	Enrichments *WorkflowEnrichments `field:"optional" json:"enrichments" yaml:"enrichments"`
	// Indicates whether the enrichments are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/workflow#enrichments_enabled Workflow#enrichments_enabled}
	EnrichmentsEnabled interface{} `field:"optional" json:"enrichmentsEnabled" yaml:"enrichmentsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/workflow#id Workflow#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

