// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboardraw


type OneDashboardRawPage struct {
	// The dashboard page's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.30.0/docs/resources/one_dashboard_raw#name OneDashboardRaw#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The dashboard page's description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.30.0/docs/resources/one_dashboard_raw#description OneDashboardRaw#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// widget block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.30.0/docs/resources/one_dashboard_raw#widget OneDashboardRaw#widget}
	Widget interface{} `field:"optional" json:"widget" yaml:"widget"`
}

