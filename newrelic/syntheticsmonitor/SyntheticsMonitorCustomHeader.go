// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsmonitor


type SyntheticsMonitorCustomHeader struct {
	// Header name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.56.0/docs/resources/synthetics_monitor#name SyntheticsMonitor#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.56.0/docs/resources/synthetics_monitor#value SyntheticsMonitor#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

