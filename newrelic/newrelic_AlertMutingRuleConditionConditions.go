// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type AlertMutingRuleConditionConditions struct {
	// The attribute on a violation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/alert_muting_rule#attribute AlertMutingRule#attribute}
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The operator used to compare the attribute's value with the supplied value(s).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/alert_muting_rule#operator AlertMutingRule#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The value(s) to compare against the attribute's value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/alert_muting_rule#values AlertMutingRule#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

