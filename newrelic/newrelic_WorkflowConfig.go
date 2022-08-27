// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkflowConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// destination_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#destination_configuration Workflow#destination_configuration}
	DestinationConfiguration interface{} `field:"required" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// issues_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#issues_filter Workflow#issues_filter}
	IssuesFilter *WorkflowIssuesFilter `field:"required" json:"issuesFilter" yaml:"issuesFilter"`
	// The type of the muting rule handling. One of: (NOTIFY_ALL_ISSUES, DONT_NOTIFY_FULLY_MUTED_ISSUES, DONT_NOTIFY_FULLY_OR_PARTIALLY_MUTED_ISSUES).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#muting_rules_handling Workflow#muting_rules_handling}
	MutingRulesHandling *string `field:"required" json:"mutingRulesHandling" yaml:"mutingRulesHandling"`
	// (Required) The name of the workflow.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#name Workflow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The account id of the workflow.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#account_id Workflow#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Indicates whether the destinations are enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#destinations_enabled Workflow#destinations_enabled}
	DestinationsEnabled interface{} `field:"optional" json:"destinationsEnabled" yaml:"destinationsEnabled"`
	// enrichments block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#enrichments Workflow#enrichments}
	Enrichments *WorkflowEnrichments `field:"optional" json:"enrichments" yaml:"enrichments"`
	// Indicates whether the enrichments are enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#enrichments_enabled Workflow#enrichments_enabled}
	EnrichmentsEnabled interface{} `field:"optional" json:"enrichmentsEnabled" yaml:"enrichmentsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#id Workflow#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Indicates whether the workflow is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#workflow_enabled Workflow#workflow_enabled}
	WorkflowEnabled interface{} `field:"optional" json:"workflowEnabled" yaml:"workflowEnabled"`
}

