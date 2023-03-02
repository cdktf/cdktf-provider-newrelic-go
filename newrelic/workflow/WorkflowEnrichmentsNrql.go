package workflow


type WorkflowEnrichmentsNrql struct {
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#configuration Workflow#configuration}
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// (Required) Enrichment's name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#name Workflow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

