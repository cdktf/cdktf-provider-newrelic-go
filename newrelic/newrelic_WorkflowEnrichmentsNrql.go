// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type WorkflowEnrichmentsNrql struct {
	// configurations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#configurations Workflow#configurations}
	Configurations interface{} `field:"required" json:"configurations" yaml:"configurations"`
	// (Required) Enrichment's name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#name Workflow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

