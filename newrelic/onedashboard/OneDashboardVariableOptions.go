// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardVariableOptions struct {
	// Only applies to variables of type NRQL.
	//
	// With this turned on, query condition defined with the variable will not be included in the query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.57.1/docs/resources/one_dashboard#excluded OneDashboard#excluded}
	Excluded interface{} `field:"optional" json:"excluded" yaml:"excluded"`
	// Only applies to variables of type NRQL.
	//
	// With this turned on, the time range for the NRQL query will override the time picker on dashboards and other pages. Turn this off to use the time picker as normal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.57.1/docs/resources/one_dashboard#ignore_time_range OneDashboard#ignore_time_range}
	IgnoreTimeRange interface{} `field:"optional" json:"ignoreTimeRange" yaml:"ignoreTimeRange"`
}

