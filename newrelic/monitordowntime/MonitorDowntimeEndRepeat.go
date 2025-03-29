// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordowntime


type MonitorDowntimeEndRepeat struct {
	// A date, on which the Monitor Downtime's repeat cycle is expected to end.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.60.0/docs/resources/monitor_downtime#on_date MonitorDowntime#on_date}
	OnDate *string `field:"optional" json:"onDate" yaml:"onDate"`
	// Number of repetitions after which the Monitor Downtime's repeat cycle is expected to end.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.60.0/docs/resources/monitor_downtime#on_repeat MonitorDowntime#on_repeat}
	OnRepeat *float64 `field:"optional" json:"onRepeat" yaml:"onRepeat"`
}

