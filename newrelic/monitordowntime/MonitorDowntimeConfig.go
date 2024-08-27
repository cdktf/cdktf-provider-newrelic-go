// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordowntime

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorDowntimeConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A datetime stamp signifying the end of the Monitor Downtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/monitor_downtime#end_time MonitorDowntime#end_time}
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// An identifier of the type of Monitor Downtime to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/monitor_downtime#mode MonitorDowntime#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// A name to identify the Monitor Downtime to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/monitor_downtime#name MonitorDowntime#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A datetime stamp signifying the start of the Monitor Downtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/monitor_downtime#start_time MonitorDowntime#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// The timezone that applies to the Monitor Downtime schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/monitor_downtime#time_zone MonitorDowntime#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// The ID of the New Relic account in which the Monitor Downtime shall be created.
	//
	// Defaults to the `account_id` in the provider{} configuration if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/monitor_downtime#account_id MonitorDowntime#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// end_repeat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/monitor_downtime#end_repeat MonitorDowntime#end_repeat}
	EndRepeat *MonitorDowntimeEndRepeat `field:"optional" json:"endRepeat" yaml:"endRepeat"`
	// frequency block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/monitor_downtime#frequency MonitorDowntime#frequency}
	Frequency *MonitorDowntimeFrequency `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/monitor_downtime#id MonitorDowntime#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A list of maintenance days to be included with the created weekly Monitor Downtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/monitor_downtime#maintenance_days MonitorDowntime#maintenance_days}
	MaintenanceDays *[]*string `field:"optional" json:"maintenanceDays" yaml:"maintenanceDays"`
	// A list of GUIDs of monitors, to which the created Monitor Downtime shall be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.43.0/docs/resources/monitor_downtime#monitor_guids MonitorDowntime#monitor_guids}
	MonitorGuids *[]*string `field:"optional" json:"monitorGuids" yaml:"monitorGuids"`
}

