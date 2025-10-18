// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertcondition


type AlertConditionTerm struct {
	// In minutes, must be in the range of 5 to 120, inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/alert_condition#duration AlertCondition#duration}
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// Must be 0 or greater.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/alert_condition#threshold AlertCondition#threshold}
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// One of (all, any).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/alert_condition#time_function AlertCondition#time_function}
	TimeFunction *string `field:"required" json:"timeFunction" yaml:"timeFunction"`
	// One of (above, below, equal). Defaults to equal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/alert_condition#operator AlertCondition#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// One of (critical, warning). Defaults to critical.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/alert_condition#priority AlertCondition#priority}
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
}

