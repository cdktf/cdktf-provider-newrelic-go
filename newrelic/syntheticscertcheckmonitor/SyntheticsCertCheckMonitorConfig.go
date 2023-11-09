// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticscertcheckmonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsCertCheckMonitorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/resources/synthetics_cert_check_monitor#certificate_expiration SyntheticsCertCheckMonitor#certificate_expiration}.
	CertificateExpiration *float64 `field:"required" json:"certificateExpiration" yaml:"certificateExpiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/resources/synthetics_cert_check_monitor#domain SyntheticsCertCheckMonitor#domain}.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// name of the cert check monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/resources/synthetics_cert_check_monitor#name SyntheticsCertCheckMonitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The interval at which this monitor should run.
	//
	// Valid values are EVERY_MINUTE, EVERY_5_MINUTES, EVERY_10_MINUTES, EVERY_15_MINUTES, EVERY_30_MINUTES, EVERY_HOUR, EVERY_6_HOURS, EVERY_12_HOURS, or EVERY_DAY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/resources/synthetics_cert_check_monitor#period SyntheticsCertCheckMonitor#period}
	Period *string `field:"required" json:"period" yaml:"period"`
	// The monitor status (i.e. ENABLED, MUTED, DISABLED). Note: The 'MUTED' status is now deprecated, and support for this value will soon be removed from the Terraform Provider in an upcoming release. It is highly recommended for users to refrain from using this value and shift to alternatives.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/resources/synthetics_cert_check_monitor#status SyntheticsCertCheckMonitor#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// ID of the newrelic account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/resources/synthetics_cert_check_monitor#account_id SyntheticsCertCheckMonitor#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/resources/synthetics_cert_check_monitor#id SyntheticsCertCheckMonitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The locations in which this monitor should be run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/resources/synthetics_cert_check_monitor#locations_private SyntheticsCertCheckMonitor#locations_private}
	LocationsPrivate *[]*string `field:"optional" json:"locationsPrivate" yaml:"locationsPrivate"`
	// The locations in which this monitor should be run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/resources/synthetics_cert_check_monitor#locations_public SyntheticsCertCheckMonitor#locations_public}
	LocationsPublic *[]*string `field:"optional" json:"locationsPublic" yaml:"locationsPublic"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/resources/synthetics_cert_check_monitor#tag SyntheticsCertCheckMonitor#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

