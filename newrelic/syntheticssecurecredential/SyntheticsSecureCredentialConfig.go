// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticssecurecredential

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsSecureCredentialConfig struct {
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
	// The secure credential's key name.
	//
	// Regardless of the case used in the configuration, the provider will provide an upcased key to the underlying API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.0/docs/resources/synthetics_secure_credential#key SyntheticsSecureCredential#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The secure credential's value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.0/docs/resources/synthetics_secure_credential#value SyntheticsSecureCredential#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// The New Relic account ID where you want to create the secure credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.0/docs/resources/synthetics_secure_credential#account_id SyntheticsSecureCredential#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// The secure credential's description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.0/docs/resources/synthetics_secure_credential#description SyntheticsSecureCredential#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.0/docs/resources/synthetics_secure_credential#id SyntheticsSecureCredential#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The time the secure credential was last updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.0/docs/resources/synthetics_secure_credential#last_updated SyntheticsSecureCredential#last_updated}
	LastUpdated *string `field:"optional" json:"lastUpdated" yaml:"lastUpdated"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.0/docs/resources/synthetics_secure_credential#timeouts SyntheticsSecureCredential#timeouts}
	Timeouts *SyntheticsSecureCredentialTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

