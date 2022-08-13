// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type WorkloadEntitySearchQuery struct {
	// The query.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workload#query Workload#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

