// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type WorkflowEnrichments struct {
	// nrql block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#nrql Workflow#nrql}
	Nrql interface{} `field:"required" json:"nrql" yaml:"nrql"`
}

