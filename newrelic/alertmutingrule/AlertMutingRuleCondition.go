package alertmutingrule


type AlertMutingRuleCondition struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/alert_muting_rule#conditions AlertMutingRule#conditions}
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// The operator used to combine all the MutingRuleConditions within the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/alert_muting_rule#operator AlertMutingRule#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
}

