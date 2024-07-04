// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordowntime


type MonitorDowntimeFrequency struct {
	// A numerical list of days of a month on which the Monitor Downtime is scheduled to run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.0/docs/resources/monitor_downtime#days_of_month MonitorDowntime#days_of_month}
	DaysOfMonth *[]*float64 `field:"optional" json:"daysOfMonth" yaml:"daysOfMonth"`
	// days_of_week block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.0/docs/resources/monitor_downtime#days_of_week MonitorDowntime#days_of_week}
	DaysOfWeek *MonitorDowntimeFrequencyDaysOfWeek `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
}

