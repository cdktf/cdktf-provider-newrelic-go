// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetAreaTooltip struct {
	// Tooltip display mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.67.0/docs/resources/one_dashboard#mode OneDashboard#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

