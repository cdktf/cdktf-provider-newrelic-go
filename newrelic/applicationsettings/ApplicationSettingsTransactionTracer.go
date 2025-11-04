// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationsettings


type ApplicationSettingsTransactionTracer struct {
	// explain_query_plans block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.75.0/docs/resources/application_settings#explain_query_plans ApplicationSettings#explain_query_plans}
	ExplainQueryPlans interface{} `field:"optional" json:"explainQueryPlans" yaml:"explainQueryPlans"`
	// sql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.75.0/docs/resources/application_settings#sql ApplicationSettings#sql}
	Sql *ApplicationSettingsTransactionTracerSql `field:"optional" json:"sql" yaml:"sql"`
	// The response time threshold value for capturing stack traces of SQL queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.75.0/docs/resources/application_settings#stack_trace_threshold_value ApplicationSettings#stack_trace_threshold_value}
	StackTraceThresholdValue *float64 `field:"optional" json:"stackTraceThresholdValue" yaml:"stackTraceThresholdValue"`
	// The type of threshold for transaction tracing, either 'APDEX_F' or 'VALUE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.75.0/docs/resources/application_settings#transaction_threshold_type ApplicationSettings#transaction_threshold_type}
	TransactionThresholdType *string `field:"optional" json:"transactionThresholdType" yaml:"transactionThresholdType"`
	// The threshold value for transaction tracing when 'transaction_threshold_type' is 'VALUE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.75.0/docs/resources/application_settings#transaction_threshold_value ApplicationSettings#transaction_threshold_value}
	TransactionThresholdValue *float64 `field:"optional" json:"transactionThresholdValue" yaml:"transactionThresholdValue"`
}

