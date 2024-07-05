// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticscertcheckmonitor


type SyntheticsCertCheckMonitorTag struct {
	// Name of the tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/synthetics_cert_check_monitor#key SyntheticsCertCheckMonitor#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Values associated with the tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/synthetics_cert_check_monitor#values SyntheticsCertCheckMonitor#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

