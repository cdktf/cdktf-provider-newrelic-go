// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NrqlAlertConditionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The title of the condition.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#name NrqlAlertCondition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// nrql block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#nrql NrqlAlertCondition#nrql}
	Nrql *NrqlAlertConditionNrql `field:"required" json:"nrql" yaml:"nrql"`
	// The ID of the policy where this condition should be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#policy_id NrqlAlertCondition#policy_id}
	PolicyId *float64 `field:"required" json:"policyId" yaml:"policyId"`
	// The New Relic account ID for managing your NRQL alert conditions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#account_id NrqlAlertCondition#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// How long we wait for data that belongs in each aggregation window.
	//
	// Depending on your data, a longer delay may increase accuracy but delay notifications. Use aggregationDelay with the EVENT_FLOW and CADENCE aggregation methods.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#aggregation_delay NrqlAlertCondition#aggregation_delay}
	AggregationDelay *string `field:"optional" json:"aggregationDelay" yaml:"aggregationDelay"`
	// The method that determines when we consider an aggregation window to be complete so that we can evaluate the signal for violations.
	//
	// Default is CADENCE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#aggregation_method NrqlAlertCondition#aggregation_method}
	AggregationMethod *string `field:"optional" json:"aggregationMethod" yaml:"aggregationMethod"`
	// How long we wait after each data point arrives to make sure we've processed the whole batch.
	//
	// Use aggregationTimer with the EVENT_TIMER aggregation method.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#aggregation_timer NrqlAlertCondition#aggregation_timer}
	AggregationTimer *string `field:"optional" json:"aggregationTimer" yaml:"aggregationTimer"`
	// The duration of the time window used to evaluate the NRQL query, in seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#aggregation_window NrqlAlertCondition#aggregation_window}
	AggregationWindow *float64 `field:"optional" json:"aggregationWindow" yaml:"aggregationWindow"`
	// The baseline direction of a baseline NRQL alert condition. Valid values are: 'LOWER_ONLY', 'UPPER_AND_LOWER', 'UPPER_ONLY' (case insensitive).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#baseline_direction NrqlAlertCondition#baseline_direction}
	BaselineDirection *string `field:"optional" json:"baselineDirection" yaml:"baselineDirection"`
	// Whether to close all open violations when the signal expires.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#close_violations_on_expiration NrqlAlertCondition#close_violations_on_expiration}
	CloseViolationsOnExpiration interface{} `field:"optional" json:"closeViolationsOnExpiration" yaml:"closeViolationsOnExpiration"`
	// critical block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#critical NrqlAlertCondition#critical}
	Critical *NrqlAlertConditionCritical `field:"optional" json:"critical" yaml:"critical"`
	// The description of the NRQL alert condition.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#description NrqlAlertCondition#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether or not to enable the alert condition.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#enabled NrqlAlertCondition#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The amount of time (in seconds) to wait before considering the signal expired.
	//
	// Must be in the range of 30 to 172800 (inclusive)
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#expiration_duration NrqlAlertCondition#expiration_duration}
	ExpirationDuration *float64 `field:"optional" json:"expirationDuration" yaml:"expirationDuration"`
	// Which strategy to use when filling gaps in the signal.
	//
	// If static, the 'fill value' will be used for filling gaps in the signal. Valid values are: 'NONE', 'LAST_VALUE', or 'STATIC' (case insensitive).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#fill_option NrqlAlertCondition#fill_option}
	FillOption *string `field:"optional" json:"fillOption" yaml:"fillOption"`
	// If using the 'static' fill option, this value will be used for filling gaps in the signal.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#fill_value NrqlAlertCondition#fill_value}
	FillValue *float64 `field:"optional" json:"fillValue" yaml:"fillValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#id NrqlAlertCondition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether to create a new violation to capture that the signal expired.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#open_violation_on_expiration NrqlAlertCondition#open_violation_on_expiration}
	OpenViolationOnExpiration interface{} `field:"optional" json:"openViolationOnExpiration" yaml:"openViolationOnExpiration"`
	// Runbook URL to display in notifications.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#runbook_url NrqlAlertCondition#runbook_url}
	RunbookUrl *string `field:"optional" json:"runbookUrl" yaml:"runbookUrl"`
	// The duration of overlapping timewindows used to smooth the chart line, in seconds.
	//
	// Must be a factor of `aggregation_window` and less than the aggregation window. It should be greater or equal to 30 seconds if `aggregation_window` is less than or equal to 3600 seconds, or greater or equal to `aggregation_window / 120` if `aggregation_window` is greater than 3600 seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#slide_by NrqlAlertCondition#slide_by}
	SlideBy *float64 `field:"optional" json:"slideBy" yaml:"slideBy"`
	// term block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#term NrqlAlertCondition#term}
	Term interface{} `field:"optional" json:"term" yaml:"term"`
	// The type of NRQL alert condition to create. Valid values are: 'static', 'baseline'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#type NrqlAlertCondition#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Values are: 'single_value' (deprecated) or 'sum' (deprecated).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#value_function NrqlAlertCondition#value_function}
	ValueFunction *string `field:"optional" json:"valueFunction" yaml:"valueFunction"`
	// Sets a time limit, in hours, that will automatically force-close a long-lasting violation after the time limit you select.
	//
	// Possible values are 'ONE_HOUR', 'TWO_HOURS', 'FOUR_HOURS', 'EIGHT_HOURS', 'TWELVE_HOURS', 'TWENTY_FOUR_HOURS', 'THIRTY_DAYS' (case insensitive).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#violation_time_limit NrqlAlertCondition#violation_time_limit}
	ViolationTimeLimit *string `field:"optional" json:"violationTimeLimit" yaml:"violationTimeLimit"`
	// Sets a time limit, in seconds, that will automatically force-close a long-lasting violation after the time limit you select.
	//
	// Must be in the range of 300 to 2592000 (inclusive)
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#violation_time_limit_seconds NrqlAlertCondition#violation_time_limit_seconds}
	ViolationTimeLimitSeconds *float64 `field:"optional" json:"violationTimeLimitSeconds" yaml:"violationTimeLimitSeconds"`
	// warning block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition#warning NrqlAlertCondition#warning}
	Warning *NrqlAlertConditionWarning `field:"optional" json:"warning" yaml:"warning"`
}

