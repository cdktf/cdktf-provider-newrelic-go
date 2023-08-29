// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsmonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsMonitorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#name SyntheticsMonitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The monitor status (i.e. ENABLED, MUTED, DISABLED).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#status SyntheticsMonitor#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// The monitor type. Valid values are SIMPLE AND BROWSER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#type SyntheticsMonitor#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// ID of the newrelic account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#account_id SyntheticsMonitor#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Bypass HEAD request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#bypass_head_request SyntheticsMonitor#bypass_head_request}
	BypassHeadRequest interface{} `field:"optional" json:"bypassHeadRequest" yaml:"bypassHeadRequest"`
	// custom_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#custom_header SyntheticsMonitor#custom_header}
	CustomHeader interface{} `field:"optional" json:"customHeader" yaml:"customHeader"`
	// Capture a screenshot during job execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#enable_screenshot_on_failure_and_script SyntheticsMonitor#enable_screenshot_on_failure_and_script}
	EnableScreenshotOnFailureAndScript interface{} `field:"optional" json:"enableScreenshotOnFailureAndScript" yaml:"enableScreenshotOnFailureAndScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#id SyntheticsMonitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The locations in which this monitor should be run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#locations_private SyntheticsMonitor#locations_private}
	LocationsPrivate *[]*string `field:"optional" json:"locationsPrivate" yaml:"locationsPrivate"`
	// The locations in which this monitor should be run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#locations_public SyntheticsMonitor#locations_public}
	LocationsPublic *[]*string `field:"optional" json:"locationsPublic" yaml:"locationsPublic"`
	// The interval at which this monitor should run.
	//
	// Valid values are EVERY_MINUTE, EVERY_5_MINUTES, EVERY_10_MINUTES, EVERY_15_MINUTES, EVERY_30_MINUTES, EVERY_HOUR, EVERY_6_HOURS, EVERY_12_HOURS, or EVERY_DAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#period SyntheticsMonitor#period}
	Period *string `field:"optional" json:"period" yaml:"period"`
	// The runtime type that the monitor will run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#runtime_type SyntheticsMonitor#runtime_type}
	RuntimeType *string `field:"optional" json:"runtimeType" yaml:"runtimeType"`
	// The specific version of the runtime type selected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#runtime_type_version SyntheticsMonitor#runtime_type_version}
	RuntimeTypeVersion *string `field:"optional" json:"runtimeTypeVersion" yaml:"runtimeTypeVersion"`
	// The programing language that should execute the script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#script_language SyntheticsMonitor#script_language}
	ScriptLanguage *string `field:"optional" json:"scriptLanguage" yaml:"scriptLanguage"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#tag SyntheticsMonitor#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
	// Fail the monitor check if redirected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#treat_redirect_as_failure SyntheticsMonitor#treat_redirect_as_failure}
	TreatRedirectAsFailure interface{} `field:"optional" json:"treatRedirectAsFailure" yaml:"treatRedirectAsFailure"`
	// The URI for the monitor to hit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#uri SyntheticsMonitor#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
	// The string to validate against in the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#validation_string SyntheticsMonitor#validation_string}
	ValidationString *string `field:"optional" json:"validationString" yaml:"validationString"`
	// Verify SSL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/synthetics_monitor#verify_ssl SyntheticsMonitor#verify_ssl}
	VerifySsl interface{} `field:"optional" json:"verifySsl" yaml:"verifySsl"`
}

