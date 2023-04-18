package workload


type WorkloadStatusConfigAutomatic struct {
	// Whether the automatic status configuration is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/workload#enabled Workload#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// remaining_entities_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/workload#remaining_entities_rule Workload#remaining_entities_rule}
	RemainingEntitiesRule *WorkloadStatusConfigAutomaticRemainingEntitiesRule `field:"optional" json:"remainingEntitiesRule" yaml:"remainingEntitiesRule"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/workload#rule Workload#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
}

