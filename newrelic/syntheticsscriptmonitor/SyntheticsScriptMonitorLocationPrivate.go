// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsscriptmonitor


type SyntheticsScriptMonitorLocationPrivate struct {
	// The unique identifier for the Synthetics private location in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/synthetics_script_monitor#guid SyntheticsScriptMonitor#guid}
	Guid *string `field:"required" json:"guid" yaml:"guid"`
	// The location's Verified Script Execution password (Only necessary if Verified Script Execution is enabled for the location).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/synthetics_script_monitor#vse_password SyntheticsScriptMonitor#vse_password}
	VsePassword *string `field:"optional" json:"vsePassword" yaml:"vsePassword"`
}

