// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardVariable struct {
	// The variable identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.1/docs/resources/one_dashboard#name OneDashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicates the strategy to apply when replacing a variable in a NRQL query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.1/docs/resources/one_dashboard#replacement_strategy OneDashboard#replacement_strategy}
	ReplacementStrategy *string `field:"required" json:"replacementStrategy" yaml:"replacementStrategy"`
	// Human-friendly display string for this variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.1/docs/resources/one_dashboard#title OneDashboard#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Specifies the data type of the variable and where its possible values may come from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.1/docs/resources/one_dashboard#type OneDashboard#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Default values for this variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.1/docs/resources/one_dashboard#default_values OneDashboard#default_values}
	DefaultValues *[]*string `field:"optional" json:"defaultValues" yaml:"defaultValues"`
	// Indicates whether this variable supports multiple selection or not. Only applies to variables of type NRQL or ENUM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.1/docs/resources/one_dashboard#is_multi_selection OneDashboard#is_multi_selection}
	IsMultiSelection interface{} `field:"optional" json:"isMultiSelection" yaml:"isMultiSelection"`
	// item block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.1/docs/resources/one_dashboard#item OneDashboard#item}
	Item interface{} `field:"optional" json:"item" yaml:"item"`
	// nrql_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.1/docs/resources/one_dashboard#nrql_query OneDashboard#nrql_query}
	NrqlQuery *OneDashboardVariableNrqlQuery `field:"optional" json:"nrqlQuery" yaml:"nrqlQuery"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.1/docs/resources/one_dashboard#options OneDashboard#options}
	Options interface{} `field:"optional" json:"options" yaml:"options"`
}

