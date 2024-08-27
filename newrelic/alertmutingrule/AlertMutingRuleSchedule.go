// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertmutingrule


type AlertMutingRuleSchedule struct {
	// The time zone that applies to the MutingRule schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/alert_muting_rule#time_zone AlertMutingRule#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// The datetime stamp when the MutingRule schedule should stop repeating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/alert_muting_rule#end_repeat AlertMutingRule#end_repeat}
	EndRepeat *string `field:"optional" json:"endRepeat" yaml:"endRepeat"`
	// The datetime stamp representing when the MutingRule should end.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/alert_muting_rule#end_time AlertMutingRule#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The frequency the MutingRule schedule repeats. One of [DAILY, WEEKLY, MONTHLY].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/alert_muting_rule#repeat AlertMutingRule#repeat}
	Repeat *string `field:"optional" json:"repeat" yaml:"repeat"`
	// The number of times the MutingRule schedule should repeat.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/alert_muting_rule#repeat_count AlertMutingRule#repeat_count}
	RepeatCount *float64 `field:"optional" json:"repeatCount" yaml:"repeatCount"`
	// The datetime stamp representing when the MutingRule should start.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/alert_muting_rule#start_time AlertMutingRule#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// The day(s) of the week that a MutingRule should repeat when the repeat field is set to WEEKLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/alert_muting_rule#weekly_repeat_days AlertMutingRule#weekly_repeat_days}
	WeeklyRepeatDays *[]*string `field:"optional" json:"weeklyRepeatDays" yaml:"weeklyRepeatDays"`
}

