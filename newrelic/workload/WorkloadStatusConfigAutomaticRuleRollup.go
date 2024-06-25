// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workload


type WorkloadStatusConfigAutomaticRuleRollup struct {
	// The rollup strategy that is applied to a group of entities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.38.1/docs/resources/workload#strategy Workload#strategy}
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// Type of threshold defined for the rule.
	//
	// This is an optional field that only applies when strategy is WORST_STATUS_WINS. Use a threshold to roll up the worst status only after a certain amount of entities are not operational.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.38.1/docs/resources/workload#threshold_type Workload#threshold_type}
	ThresholdType *string `field:"optional" json:"thresholdType" yaml:"thresholdType"`
	// Threshold value defined for the rule.
	//
	// This optional field is used in combination with thresholdType. If the threshold type is null, the threshold value will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.38.1/docs/resources/workload#threshold_value Workload#threshold_value}
	ThresholdValue *float64 `field:"optional" json:"thresholdValue" yaml:"thresholdValue"`
}

