package workload


type WorkloadStatusConfigAutomaticRuleNrqlQuery struct {
	// The entity search query that is used to perform the search of a group of entities.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workload#query Workload#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

