// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertcondition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlertConditionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// The instance IDs associated with this condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/alert_condition#entities AlertCondition#entities}
	Entities *[]*float64 `field:"required" json:"entities" yaml:"entities"`
	// The metric field accepts parameters based on the type set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/alert_condition#metric AlertCondition#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// The title of the condition. Must be between 1 and 128 characters, inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/alert_condition#name AlertCondition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the policy where this condition should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/alert_condition#policy_id AlertCondition#policy_id}
	PolicyId *float64 `field:"required" json:"policyId" yaml:"policyId"`
	// term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/alert_condition#term AlertCondition#term}
	Term interface{} `field:"required" json:"term" yaml:"term"`
	// The type of condition. One of: (browser_metric, mobile_metric, servers_metric, apm_app_metric, apm_jvm_metric, apm_kt_metric).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/alert_condition#type AlertCondition#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// One of (application, instance).
	//
	// Choose application for most scenarios. If you are using the JVM plugin in New Relic, the instance setting allows your condition to trigger for specific app instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/alert_condition#condition_scope AlertCondition#condition_scope}
	ConditionScope *string `field:"optional" json:"conditionScope" yaml:"conditionScope"`
	// Whether the condition is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/alert_condition#enabled AlertCondition#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A valid Garbage Collection metric e.g. GC/G1 Young Generation. This is required if you are using apm_jvm_metric with gc_cpu_time condition type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/alert_condition#gc_metric AlertCondition#gc_metric}
	GcMetric *string `field:"optional" json:"gcMetric" yaml:"gcMetric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/alert_condition#id AlertCondition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Runbook URL to display in notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/alert_condition#runbook_url AlertCondition#runbook_url}
	RunbookUrl *string `field:"optional" json:"runbookUrl" yaml:"runbookUrl"`
	// A custom metric to be evaluated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/alert_condition#user_defined_metric AlertCondition#user_defined_metric}
	UserDefinedMetric *string `field:"optional" json:"userDefinedMetric" yaml:"userDefinedMetric"`
	// One of: (average, min, max, total, sample_size, percent, rate).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/alert_condition#user_defined_value_function AlertCondition#user_defined_value_function}
	UserDefinedValueFunction *string `field:"optional" json:"userDefinedValueFunction" yaml:"userDefinedValueFunction"`
	// Automatically close instance-based incidents, including JVM health metric incidents, after the number of hours specified.
	//
	// Must be between 1 and 720 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/alert_condition#violation_close_timer AlertCondition#violation_close_timer}
	ViolationCloseTimer *float64 `field:"optional" json:"violationCloseTimer" yaml:"violationCloseTimer"`
}

