// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nrqlalertcondition


type NrqlAlertConditionCritical struct {
	// For baseline conditions must be in range [1, 1000].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.0/docs/resources/nrql_alert_condition#threshold NrqlAlertCondition#threshold}
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Violations will not change system health status for this term.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.0/docs/resources/nrql_alert_condition#disable_health_status_reporting NrqlAlertCondition#disable_health_status_reporting}
	DisableHealthStatusReporting interface{} `field:"optional" json:"disableHealthStatusReporting" yaml:"disableHealthStatusReporting"`
	// In minutes, must be in the range of 1 to 120 (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.0/docs/resources/nrql_alert_condition#duration NrqlAlertCondition#duration}
	Duration *float64 `field:"optional" json:"duration" yaml:"duration"`
	// One of (above, above_or_equals, below, below_or_equals, equals, not_equals). Defaults to 'equals'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.0/docs/resources/nrql_alert_condition#operator NrqlAlertCondition#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// prediction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.0/docs/resources/nrql_alert_condition#prediction NrqlAlertCondition#prediction}
	Prediction *NrqlAlertConditionCriticalPrediction `field:"optional" json:"prediction" yaml:"prediction"`
	// The duration, in seconds, that the threshold must violate in order to create an incident.
	//
	// Value must be a multiple of the 'aggregation_window' (which has a default of 60 seconds). Value must be within 120-86400 seconds for baseline conditions, and within 60-86400 seconds for static conditions
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.0/docs/resources/nrql_alert_condition#threshold_duration NrqlAlertCondition#threshold_duration}
	ThresholdDuration *float64 `field:"optional" json:"thresholdDuration" yaml:"thresholdDuration"`
	// The criteria for how many data points must be in violation for the specified threshold duration.
	//
	// Valid values are: 'ALL' or 'AT_LEAST_ONCE' (case insensitive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.0/docs/resources/nrql_alert_condition#threshold_occurrences NrqlAlertCondition#threshold_occurrences}
	ThresholdOccurrences *string `field:"optional" json:"thresholdOccurrences" yaml:"thresholdOccurrences"`
	// Valid values are: 'all' or 'any'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.0/docs/resources/nrql_alert_condition#time_function NrqlAlertCondition#time_function}
	TimeFunction *string `field:"optional" json:"timeFunction" yaml:"timeFunction"`
}

