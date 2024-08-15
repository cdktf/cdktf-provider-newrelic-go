// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsstepmonitor


type SyntheticsStepMonitorSteps struct {
	// The position of the step within the script ranging from 0-100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.42.1/docs/resources/synthetics_step_monitor#ordinal SyntheticsStepMonitor#ordinal}
	Ordinal *float64 `field:"required" json:"ordinal" yaml:"ordinal"`
	// The type of step to be added to the script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.42.1/docs/resources/synthetics_step_monitor#type SyntheticsStepMonitor#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The metadata values related to the check the step performs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.42.1/docs/resources/synthetics_step_monitor#values SyntheticsStepMonitor#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

