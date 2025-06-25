// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package infraalertcondition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InfraAlertConditionConfig struct {
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
	// The Infrastructure alert condition's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#name InfraAlertCondition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the alert policy where this condition should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#policy_id InfraAlertCondition#policy_id}
	PolicyId *float64 `field:"required" json:"policyId" yaml:"policyId"`
	// The type of Infrastructure alert condition. Valid values are infra_process_running, infra_metric, and infra_host_not_reporting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#type InfraAlertCondition#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The operator used to evaluate the threshold value.
	//
	// Valid values are above, below, and equal. Supported by the infra_metric and infra_process_running condition types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#comparison InfraAlertCondition#comparison}
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// critical block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#critical InfraAlertCondition#critical}
	Critical *InfraAlertConditionCritical `field:"optional" json:"critical" yaml:"critical"`
	// The description of the Infrastructure alert condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#description InfraAlertCondition#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the condition is turned on or off. Valid values are true and false. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#enabled InfraAlertCondition#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The metric event; for example, SystemSample or StorageSample. Supported by the infra_metric condition type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#event InfraAlertCondition#event}
	Event *string `field:"optional" json:"event" yaml:"event"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#id InfraAlertCondition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// For alerts on integrations, use this instead of event. Supported by the infra_metric condition type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#integration_provider InfraAlertCondition#integration_provider}
	IntegrationProvider *string `field:"optional" json:"integrationProvider" yaml:"integrationProvider"`
	// Any filters applied to processes; for example: commandName = 'java'. Supported by the infra_process_running condition type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#process_where InfraAlertCondition#process_where}
	ProcessWhere *string `field:"optional" json:"processWhere" yaml:"processWhere"`
	// Runbook URL to display in notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#runbook_url InfraAlertCondition#runbook_url}
	RunbookUrl *string `field:"optional" json:"runbookUrl" yaml:"runbookUrl"`
	// The attribute name to identify the metric being targeted;
	//
	// for example, cpuPercent, diskFreePercent, or memoryResidentSizeBytes. The underlying API will automatically populate this value for Infrastructure integrations (for example diskFreePercent), so make sure to explicitly include this value to avoid diff issues. Supported by the infra_metric condition type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#select InfraAlertCondition#select}
	Select *string `field:"optional" json:"select" yaml:"select"`
	// Determines how much time, in hours, will pass before an incident is automatically closed.
	//
	// Valid values are 1, 2, 4, 8, 12, 24, 48, or 72
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#violation_close_timer InfraAlertCondition#violation_close_timer}
	ViolationCloseTimer *float64 `field:"optional" json:"violationCloseTimer" yaml:"violationCloseTimer"`
	// warning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#warning InfraAlertCondition#warning}
	Warning *InfraAlertConditionWarning `field:"optional" json:"warning" yaml:"warning"`
	// If applicable, this identifies any Infrastructure host filters used; for example: hostname LIKE '%cassandra%'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/infra_alert_condition#where InfraAlertCondition#where}
	Where *string `field:"optional" json:"where" yaml:"where"`
}

