package pluginsalertcondition


type PluginsAlertConditionTerm struct {
	// In minutes, must be in the range of 5 to 120, inclusive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/plugins_alert_condition#duration PluginsAlertCondition#duration}
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// Must be 0 or greater.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/plugins_alert_condition#threshold PluginsAlertCondition#threshold}
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// One of `all` or `any`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/plugins_alert_condition#time_function PluginsAlertCondition#time_function}
	TimeFunction *string `field:"required" json:"timeFunction" yaml:"timeFunction"`
	// One of `above`, `below`, or `equal`. Defaults to equal.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/plugins_alert_condition#operator PluginsAlertCondition#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// One of `critical` or `warning`. Defaults to critical.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/plugins_alert_condition#priority PluginsAlertCondition#priority}
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
}

