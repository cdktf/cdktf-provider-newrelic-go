// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetFunnelInitialSorting struct {
	// Defines the sort order. Either ascending or descending.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.57.0/docs/resources/one_dashboard#direction OneDashboard#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The column name to be sorted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.57.0/docs/resources/one_dashboard#name OneDashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

