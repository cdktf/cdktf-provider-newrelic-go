// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nrqlalertcondition


type NrqlAlertConditionNrql struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.1/docs/resources/nrql_alert_condition#query NrqlAlertCondition#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
	// The New Relic account ID to use as the basis for the NRQL alert condition's `query`;
	//
	// will default to `account_id` if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.1/docs/resources/nrql_alert_condition#data_account_id NrqlAlertCondition#data_account_id}
	DataAccountId *float64 `field:"optional" json:"dataAccountId" yaml:"dataAccountId"`
	// NRQL queries are evaluated in one-minute time windows.
	//
	// The start time depends on the value you provide in the NRQL condition's `evaluation_offset`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.1/docs/resources/nrql_alert_condition#evaluation_offset NrqlAlertCondition#evaluation_offset}
	EvaluationOffset *float64 `field:"optional" json:"evaluationOffset" yaml:"evaluationOffset"`
	// NRQL queries are evaluated in one-minute time windows.
	//
	// The start time depends on the value you provide in the NRQL condition's `since_value`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.1/docs/resources/nrql_alert_condition#since_value NrqlAlertCondition#since_value}
	SinceValue *string `field:"optional" json:"sinceValue" yaml:"sinceValue"`
}

