// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetLineNrqlQuery struct {
	// The NRQL query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.68.0/docs/resources/one_dashboard#query OneDashboard#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The account ID(s) used for the NRQL query.
	//
	// Can be a single account ID or multiple account IDs in a JSON-encoded array.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.68.0/docs/resources/one_dashboard#account_id OneDashboard#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
}

