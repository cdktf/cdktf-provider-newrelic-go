package workflow


type WorkflowEnrichmentsNrqlConfiguration struct {
	// enrichment's NRQL query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.22.0/docs/resources/workflow#query Workflow#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

