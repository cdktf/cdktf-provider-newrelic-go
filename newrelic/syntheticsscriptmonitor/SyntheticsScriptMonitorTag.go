// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsscriptmonitor


type SyntheticsScriptMonitorTag struct {
	// Name of the tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.69.1/docs/resources/synthetics_script_monitor#key SyntheticsScriptMonitor#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Values associated with the tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.69.1/docs/resources/synthetics_script_monitor#values SyntheticsScriptMonitor#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

