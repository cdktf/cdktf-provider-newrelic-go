// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keytransaction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyTransactionConfig struct {
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
	// The acceptable amount of the time spent in the backend before customers get frustrated (Apdex target).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.52.0/docs/resources/key_transaction#apdex_index KeyTransaction#apdex_index}
	ApdexIndex *float64 `field:"required" json:"apdexIndex" yaml:"apdexIndex"`
	// The GUID of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.52.0/docs/resources/key_transaction#application_guid KeyTransaction#application_guid}
	ApplicationGuid *string `field:"required" json:"applicationGuid" yaml:"applicationGuid"`
	// The acceptable amount of time for rendering a page in a browser before customers get frustrated (browser Apdex target).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.52.0/docs/resources/key_transaction#browser_apdex_target KeyTransaction#browser_apdex_target}
	BrowserApdexTarget *float64 `field:"required" json:"browserApdexTarget" yaml:"browserApdexTarget"`
	// The name of the metric underlying this key transaction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.52.0/docs/resources/key_transaction#metric_name KeyTransaction#metric_name}
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The name of the key transaction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.52.0/docs/resources/key_transaction#name KeyTransaction#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.52.0/docs/resources/key_transaction#id KeyTransaction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

