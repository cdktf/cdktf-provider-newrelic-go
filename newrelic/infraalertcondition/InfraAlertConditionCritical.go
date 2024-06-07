// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package infraalertcondition


type InfraAlertConditionCritical struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.38.0/docs/resources/infra_alert_condition#duration InfraAlertCondition#duration}.
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.38.0/docs/resources/infra_alert_condition#time_function InfraAlertCondition#time_function}.
	TimeFunction *string `field:"optional" json:"timeFunction" yaml:"timeFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.38.0/docs/resources/infra_alert_condition#value InfraAlertCondition#value}.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

