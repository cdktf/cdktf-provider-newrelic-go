// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordowntime


type MonitorDowntimeFrequencyDaysOfWeek struct {
	// An occurrence of the day selected within the month.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.75.2/docs/resources/monitor_downtime#ordinal_day_of_month MonitorDowntime#ordinal_day_of_month}
	OrdinalDayOfMonth *string `field:"required" json:"ordinalDayOfMonth" yaml:"ordinalDayOfMonth"`
	// The day of the week on which the Monitor Downtime would run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.75.2/docs/resources/monitor_downtime#week_day MonitorDowntime#week_day}
	WeekDay *string `field:"required" json:"weekDay" yaml:"weekDay"`
}

