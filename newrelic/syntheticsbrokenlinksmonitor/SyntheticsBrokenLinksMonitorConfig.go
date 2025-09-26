// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsbrokenlinksmonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsBrokenLinksMonitorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/synthetics_broken_links_monitor#name SyntheticsBrokenLinksMonitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The interval at which this monitor should run.
	//
	// Valid values are EVERY_MINUTE, EVERY_5_MINUTES, EVERY_10_MINUTES, EVERY_15_MINUTES, EVERY_30_MINUTES, EVERY_HOUR, EVERY_6_HOURS, EVERY_12_HOURS, or EVERY_DAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/synthetics_broken_links_monitor#period SyntheticsBrokenLinksMonitor#period}
	Period *string `field:"required" json:"period" yaml:"period"`
	// The monitor status (ENABLED or DISABLED).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/synthetics_broken_links_monitor#status SyntheticsBrokenLinksMonitor#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// The URI the monitor runs against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/synthetics_broken_links_monitor#uri SyntheticsBrokenLinksMonitor#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// ID of the newrelic account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/synthetics_broken_links_monitor#account_id SyntheticsBrokenLinksMonitor#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/synthetics_broken_links_monitor#id SyntheticsBrokenLinksMonitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List private location GUIDs for which the monitor will run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/synthetics_broken_links_monitor#locations_private SyntheticsBrokenLinksMonitor#locations_private}
	LocationsPrivate *[]*string `field:"optional" json:"locationsPrivate" yaml:"locationsPrivate"`
	// Publicly available location names in which the monitor will run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/synthetics_broken_links_monitor#locations_public SyntheticsBrokenLinksMonitor#locations_public}
	LocationsPublic *[]*string `field:"optional" json:"locationsPublic" yaml:"locationsPublic"`
	// The runtime type that the monitor will run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/synthetics_broken_links_monitor#runtime_type SyntheticsBrokenLinksMonitor#runtime_type}
	RuntimeType *string `field:"optional" json:"runtimeType" yaml:"runtimeType"`
	// The specific semver version of the runtime type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/synthetics_broken_links_monitor#runtime_type_version SyntheticsBrokenLinksMonitor#runtime_type_version}
	RuntimeTypeVersion *string `field:"optional" json:"runtimeTypeVersion" yaml:"runtimeTypeVersion"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/synthetics_broken_links_monitor#tag SyntheticsBrokenLinksMonitor#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
	// A boolean attribute to be set true by the customer, if they would like to use the unsupported legacy runtime of Synthetic Monitors by means of an exemption given until the October 22, 2024 Legacy Runtime EOL.
	//
	// Setting this attribute to true would allow skipping validation performed by the the New Relic Terraform Provider starting v3.43.0 to disallow using the legacy runtime with new monitors. This would, hence, allow creation of monitors in the legacy runtime until the October 22, 2024 Legacy Runtime EOL, if exempt by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/synthetics_broken_links_monitor#use_unsupported_legacy_runtime SyntheticsBrokenLinksMonitor#use_unsupported_legacy_runtime}
	UseUnsupportedLegacyRuntime interface{} `field:"optional" json:"useUnsupportedLegacyRuntime" yaml:"useUnsupportedLegacyRuntime"`
}

