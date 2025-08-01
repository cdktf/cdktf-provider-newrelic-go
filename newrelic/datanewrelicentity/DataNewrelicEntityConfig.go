// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datanewrelicentity

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataNewrelicEntityConfig struct {
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
	// The name of the entity in New Relic One.
	//
	// The first entity matching this name for the given search parameters will be returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/data-sources/entity#name DataNewrelicEntity#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The New Relic account ID;
	//
	// if specified, constrains the data source to return an entity belonging to the account with this ID, of all matching entities retrieved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/data-sources/entity#account_id DataNewrelicEntity#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// The entity's domain.
	//
	// Valid values are APM, BROWSER, INFRA, MOBILE, SYNTH, and EXT. If not specified, all domains are searched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/data-sources/entity#domain DataNewrelicEntity#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/data-sources/entity#entity_tags DataNewrelicEntity#entity_tags}.
	EntityTags *string `field:"optional" json:"entityTags" yaml:"entityTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/data-sources/entity#id DataNewrelicEntity#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Ignore case when searching the entity name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/data-sources/entity#ignore_case DataNewrelicEntity#ignore_case}
	IgnoreCase interface{} `field:"optional" json:"ignoreCase" yaml:"ignoreCase"`
	// A boolean attribute which when set to true, does not throw an error if the queried entity is not found.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/data-sources/entity#ignore_not_found DataNewrelicEntity#ignore_not_found}
	IgnoreNotFound interface{} `field:"optional" json:"ignoreNotFound" yaml:"ignoreNotFound"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/data-sources/entity#tag DataNewrelicEntity#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
	// The entity's type. Valid values are APPLICATION, DASHBOARD, HOST, MONITOR, SERVICE and WORKLOAD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/data-sources/entity#type DataNewrelicEntity#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

