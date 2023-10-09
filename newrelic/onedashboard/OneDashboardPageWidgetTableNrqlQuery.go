// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetTableNrqlQuery struct {
	// The NRQL query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.2/docs/resources/one_dashboard#query OneDashboard#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The account id used for the NRQL query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.2/docs/resources/one_dashboard#account_id OneDashboard#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
}

