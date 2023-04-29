package workload


type WorkloadStatusConfigAutomaticRemainingEntitiesRule struct {
	// remaining_entities_rule_rollup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.21.3/docs/resources/workload#remaining_entities_rule_rollup Workload#remaining_entities_rule_rollup}
	RemainingEntitiesRuleRollup *WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollup `field:"required" json:"remainingEntitiesRuleRollup" yaml:"remainingEntitiesRuleRollup"`
}

