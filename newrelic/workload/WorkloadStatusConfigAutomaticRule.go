package workload


type WorkloadStatusConfigAutomaticRule struct {
	// rollup block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workload#rollup Workload#rollup}
	Rollup *WorkloadStatusConfigAutomaticRuleRollup `field:"required" json:"rollup" yaml:"rollup"`
	// A list of entity GUIDs composing the rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workload#entity_guids Workload#entity_guids}
	EntityGuids *[]*string `field:"optional" json:"entityGuids" yaml:"entityGuids"`
	// nrql_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workload#nrql_query Workload#nrql_query}
	NrqlQuery interface{} `field:"optional" json:"nrqlQuery" yaml:"nrqlQuery"`
}

