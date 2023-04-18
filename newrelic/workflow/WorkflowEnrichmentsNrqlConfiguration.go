package workflow


type WorkflowEnrichmentsNrqlConfiguration struct {
	// enrichment's NRQL query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/workflow#query Workflow#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

