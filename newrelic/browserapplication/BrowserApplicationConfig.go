// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package browserapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BrowserApplicationConfig struct {
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
	// The name of the application to monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.34.1/docs/resources/browser_application#name BrowserApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.34.1/docs/resources/browser_application#account_id BrowserApplication#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Configure cookies. The default is enabled: true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.34.1/docs/resources/browser_application#cookies_enabled BrowserApplication#cookies_enabled}
	CookiesEnabled interface{} `field:"optional" json:"cookiesEnabled" yaml:"cookiesEnabled"`
	// Configure distributed tracing in browser apps. The default is enabled: true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.34.1/docs/resources/browser_application#distributed_tracing_enabled BrowserApplication#distributed_tracing_enabled}
	DistributedTracingEnabled interface{} `field:"optional" json:"distributedTracingEnabled" yaml:"distributedTracingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.34.1/docs/resources/browser_application#id BrowserApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Determines which browser loader is configured. The default is "SPA".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.34.1/docs/resources/browser_application#loader_type BrowserApplication#loader_type}
	LoaderType *string `field:"optional" json:"loaderType" yaml:"loaderType"`
}

