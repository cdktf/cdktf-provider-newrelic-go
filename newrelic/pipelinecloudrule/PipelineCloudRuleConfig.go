// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipelinecloudrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipelineCloudRuleConfig struct {
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
	// The name of the rule. This must be unique within an account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/pipeline_cloud_rule#name PipelineCloudRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The NRQL query that defines which data will be processed by this pipeline cloud rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/pipeline_cloud_rule#nrql PipelineCloudRule#nrql}
	Nrql *string `field:"required" json:"nrql" yaml:"nrql"`
	// The account ID where the Pipeline Cloud rule will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/pipeline_cloud_rule#account_id PipelineCloudRule#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Provides additional information about the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/pipeline_cloud_rule#description PipelineCloudRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/pipeline_cloud_rule#id PipelineCloudRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

