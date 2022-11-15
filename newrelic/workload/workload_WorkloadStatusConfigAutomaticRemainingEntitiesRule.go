package workload


type WorkloadStatusConfigAutomaticRemainingEntitiesRule struct {
	// remaining_entities_rule_rollup block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workload#remaining_entities_rule_rollup Workload#remaining_entities_rule_rollup}
	RemainingEntitiesRuleRollup *WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollup `field:"required" json:"remainingEntitiesRuleRollup" yaml:"remainingEntitiesRuleRollup"`
}

