// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type WorkflowEnrichmentsNrqlConfigurations struct {
	// enrichment's NRQL query.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#query Workflow#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

