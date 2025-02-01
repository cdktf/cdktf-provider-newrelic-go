// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsstepmonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsStepMonitorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/synthetics_step_monitor#name SyntheticsStepMonitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The interval at which this monitor should run.
	//
	// Valid values are EVERY_MINUTE, EVERY_5_MINUTES, EVERY_10_MINUTES, EVERY_15_MINUTES, EVERY_30_MINUTES, EVERY_HOUR, EVERY_6_HOURS, EVERY_12_HOURS, or EVERY_DAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/synthetics_step_monitor#period SyntheticsStepMonitor#period}
	Period *string `field:"required" json:"period" yaml:"period"`
	// The monitor status (ENABLED or DISABLED).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/synthetics_step_monitor#status SyntheticsStepMonitor#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// steps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/synthetics_step_monitor#steps SyntheticsStepMonitor#steps}
	Steps interface{} `field:"required" json:"steps" yaml:"steps"`
	// ID of the newrelic account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/synthetics_step_monitor#account_id SyntheticsStepMonitor#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// The multiple browsers list on which synthetic monitors will run. Valid values are array of CHROME,and FIREFOX.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/synthetics_step_monitor#browsers SyntheticsStepMonitor#browsers}
	Browsers *[]*string `field:"optional" json:"browsers" yaml:"browsers"`
	// The multiple devices list on which synthetic monitors will run.
	//
	// Valid values are array of DESKTOP, MOBILE_LANDSCAPE, MOBILE_PORTRAIT, TABLET_LANDSCAPE and TABLET_PORTRAIT
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/synthetics_step_monitor#devices SyntheticsStepMonitor#devices}
	Devices *[]*string `field:"optional" json:"devices" yaml:"devices"`
	// Capture a screenshot during job execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/synthetics_step_monitor#enable_screenshot_on_failure_and_script SyntheticsStepMonitor#enable_screenshot_on_failure_and_script}
	EnableScreenshotOnFailureAndScript interface{} `field:"optional" json:"enableScreenshotOnFailureAndScript" yaml:"enableScreenshotOnFailureAndScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/synthetics_step_monitor#id SyntheticsStepMonitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// location_private block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/synthetics_step_monitor#location_private SyntheticsStepMonitor#location_private}
	LocationPrivate interface{} `field:"optional" json:"locationPrivate" yaml:"locationPrivate"`
	// The public location(s) that the monitor will run jobs from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/synthetics_step_monitor#locations_public SyntheticsStepMonitor#locations_public}
	LocationsPublic *[]*string `field:"optional" json:"locationsPublic" yaml:"locationsPublic"`
	// The runtime type that the monitor will run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/synthetics_step_monitor#runtime_type SyntheticsStepMonitor#runtime_type}
	RuntimeType *string `field:"optional" json:"runtimeType" yaml:"runtimeType"`
	// The specific semver version of the runtime type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/synthetics_step_monitor#runtime_type_version SyntheticsStepMonitor#runtime_type_version}
	RuntimeTypeVersion *string `field:"optional" json:"runtimeTypeVersion" yaml:"runtimeTypeVersion"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/synthetics_step_monitor#tag SyntheticsStepMonitor#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
	// A boolean attribute to be set true by the customer, if they would like to use the unsupported legacy runtime of Synthetic Monitors by means of an exemption given until the October 22, 2024 Legacy Runtime EOL.
	//
	// Setting this attribute to true would allow skipping validation performed by the the New Relic Terraform Provider starting v3.43.0 to disallow using the legacy runtime with new monitors. This would, hence, allow creation of monitors in the legacy runtime until the October 22, 2024 Legacy Runtime EOL, if exempt by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/synthetics_step_monitor#use_unsupported_legacy_runtime SyntheticsStepMonitor#use_unsupported_legacy_runtime}
	UseUnsupportedLegacyRuntime interface{} `field:"optional" json:"useUnsupportedLegacyRuntime" yaml:"useUnsupportedLegacyRuntime"`
}

