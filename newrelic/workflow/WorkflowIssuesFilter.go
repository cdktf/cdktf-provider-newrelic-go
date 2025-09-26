// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workflow


type WorkflowIssuesFilter struct {
	// (Required) Filter's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/workflow#name Workflow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// (Required) The type of the filter. One of: (FILTER, VIEW).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/workflow#type Workflow#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// predicate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.2/docs/resources/workflow#predicate Workflow#predicate}
	Predicate interface{} `field:"optional" json:"predicate" yaml:"predicate"`
}

