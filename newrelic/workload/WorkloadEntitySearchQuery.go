package workload


type WorkloadEntitySearchQuery struct {
	// The query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/workload#query Workload#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

