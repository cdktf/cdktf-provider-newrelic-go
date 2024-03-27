// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsstepmonitor


type SyntheticsStepMonitorTag struct {
	// Name of the tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.34.0/docs/resources/synthetics_step_monitor#key SyntheticsStepMonitor#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Values associated with the tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.34.0/docs/resources/synthetics_step_monitor#values SyntheticsStepMonitor#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

