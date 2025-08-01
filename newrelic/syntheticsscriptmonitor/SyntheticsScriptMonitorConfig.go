// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsscriptmonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsScriptMonitorConfig struct {
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
	// The title of this monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#name SyntheticsScriptMonitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The interval at which this monitor should run.
	//
	// Valid values are EVERY_MINUTE, EVERY_5_MINUTES, EVERY_10_MINUTES, EVERY_15_MINUTES, EVERY_30_MINUTES, EVERY_HOUR, EVERY_6_HOURS, EVERY_12_HOURS, or EVERY_DAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#period SyntheticsScriptMonitor#period}
	Period *string `field:"required" json:"period" yaml:"period"`
	// The monitor status (ENABLED or DISABLED).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#status SyntheticsScriptMonitor#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// The monitor type. Valid values are SCRIPT_BROWSER, and SCRIPT_API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#type SyntheticsScriptMonitor#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// ID of the newrelic account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#account_id SyntheticsScriptMonitor#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// The multiple browsers list on which synthetic monitors will run. Valid values are array of CHROME,and FIREFOX.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#browsers SyntheticsScriptMonitor#browsers}
	Browsers *[]*string `field:"optional" json:"browsers" yaml:"browsers"`
	// The device orientation the user would like to represent. Valid values are LANDSCAPE, PORTRAIT, or NONE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#device_orientation SyntheticsScriptMonitor#device_orientation}
	DeviceOrientation *string `field:"optional" json:"deviceOrientation" yaml:"deviceOrientation"`
	// The multiple devices list on which synthetic monitors will run.
	//
	// Valid values are array of DESKTOP, MOBILE_LANDSCAPE, MOBILE_PORTRAIT, TABLET_LANDSCAPE and TABLET_PORTRAIT
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#devices SyntheticsScriptMonitor#devices}
	Devices *[]*string `field:"optional" json:"devices" yaml:"devices"`
	// The device type that a user can select. Valid values are MOBILE, TABLET, or NONE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#device_type SyntheticsScriptMonitor#device_type}
	DeviceType *string `field:"optional" json:"deviceType" yaml:"deviceType"`
	// Capture a screenshot during job execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#enable_screenshot_on_failure_and_script SyntheticsScriptMonitor#enable_screenshot_on_failure_and_script}
	EnableScreenshotOnFailureAndScript interface{} `field:"optional" json:"enableScreenshotOnFailureAndScript" yaml:"enableScreenshotOnFailureAndScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#id SyntheticsScriptMonitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// location_private block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#location_private SyntheticsScriptMonitor#location_private}
	LocationPrivate interface{} `field:"optional" json:"locationPrivate" yaml:"locationPrivate"`
	// The public location(s) that the monitor will run jobs from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#locations_public SyntheticsScriptMonitor#locations_public}
	LocationsPublic *[]*string `field:"optional" json:"locationsPublic" yaml:"locationsPublic"`
	// The runtime type that the monitor will run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#runtime_type SyntheticsScriptMonitor#runtime_type}
	RuntimeType *string `field:"optional" json:"runtimeType" yaml:"runtimeType"`
	// The specific semver version of the runtime type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#runtime_type_version SyntheticsScriptMonitor#runtime_type_version}
	RuntimeTypeVersion *string `field:"optional" json:"runtimeTypeVersion" yaml:"runtimeTypeVersion"`
	// The script that the monitor runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#script SyntheticsScriptMonitor#script}
	Script *string `field:"optional" json:"script" yaml:"script"`
	// The programing language that should execute the script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#script_language SyntheticsScriptMonitor#script_language}
	ScriptLanguage *string `field:"optional" json:"scriptLanguage" yaml:"scriptLanguage"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#tag SyntheticsScriptMonitor#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
	// A boolean attribute to be set true by the customer, if they would like to use the unsupported legacy runtime of Synthetic Monitors by means of an exemption given until the October 22, 2024 Legacy Runtime EOL.
	//
	// Setting this attribute to true would allow skipping validation performed by the the New Relic Terraform Provider starting v3.43.0 to disallow using the legacy runtime with new monitors. This would, hence, allow creation of monitors in the legacy runtime until the October 22, 2024 Legacy Runtime EOL, if exempt by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/synthetics_script_monitor#use_unsupported_legacy_runtime SyntheticsScriptMonitor#use_unsupported_legacy_runtime}
	UseUnsupportedLegacyRuntime interface{} `field:"optional" json:"useUnsupportedLegacyRuntime" yaml:"useUnsupportedLegacyRuntime"`
}

