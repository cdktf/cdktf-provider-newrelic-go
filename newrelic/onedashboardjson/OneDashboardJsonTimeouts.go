// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboardjson


type OneDashboardJsonTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.4/docs/resources/one_dashboard_json#create OneDashboardJson#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.4/docs/resources/one_dashboard_json#update OneDashboardJson#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

