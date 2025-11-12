// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workflow


type WorkflowIssuesFilterPredicate struct {
	// (Required) predicate's attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.75.4/docs/resources/workflow#attribute Workflow#attribute}
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The type of the operator.
	//
	// One of: (CONTAINS, DOES_NOT_CONTAIN, DOES_NOT_EQUAL, DOES_NOT_EXACTLY_MATCH, ENDS_WITH, EQUAL, EXACTLY_MATCHES, GREATER_OR_EQUAL, GREATER_THAN, IS, IS_NOT, LESS_OR_EQUAL, LESS_THAN, STARTS_WITH).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.75.4/docs/resources/workflow#operator Workflow#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// List of predicate values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.75.4/docs/resources/workflow#values Workflow#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

