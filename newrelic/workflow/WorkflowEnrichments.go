package workflow


type WorkflowEnrichments struct {
	// nrql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.21.3/docs/resources/workflow#nrql Workflow#nrql}
	Nrql interface{} `field:"required" json:"nrql" yaml:"nrql"`
}

