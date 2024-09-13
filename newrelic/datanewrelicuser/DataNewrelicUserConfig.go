// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datanewrelicuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataNewrelicUserConfig struct {
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
	// The ID of the Authentication Domain the user being queried would belong to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.46.0/docs/data-sources/user#authentication_domain_id DataNewrelicUser#authentication_domain_id}
	AuthenticationDomainId *string `field:"required" json:"authenticationDomainId" yaml:"authenticationDomainId"`
	// The email ID of the user to be queried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.46.0/docs/data-sources/user#email_id DataNewrelicUser#email_id}
	EmailId *string `field:"optional" json:"emailId" yaml:"emailId"`
	// The name of the user to be queried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.46.0/docs/data-sources/user#name DataNewrelicUser#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

