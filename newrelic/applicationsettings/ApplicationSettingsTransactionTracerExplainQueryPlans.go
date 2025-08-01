// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationsettings


type ApplicationSettingsTransactionTracerExplainQueryPlans struct {
	// The type of threshold for explain plans, either 'APDEX_F' or 'VALUE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/application_settings#query_plan_threshold_type ApplicationSettings#query_plan_threshold_type}
	QueryPlanThresholdType *string `field:"optional" json:"queryPlanThresholdType" yaml:"queryPlanThresholdType"`
	// The threshold value for explain plans when 'query_plan_threshold_type' is 'VALUE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/application_settings#query_plan_threshold_value ApplicationSettings#query_plan_threshold_value}
	QueryPlanThresholdValue *float64 `field:"optional" json:"queryPlanThresholdValue" yaml:"queryPlanThresholdValue"`
}

