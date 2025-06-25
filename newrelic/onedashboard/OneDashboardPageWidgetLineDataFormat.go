// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetLineDataFormat struct {
	// The column name to be sorted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/one_dashboard#name OneDashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Defines the type of the mentioned column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/one_dashboard#type OneDashboard#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Defines the format of the mentioned type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/one_dashboard#format OneDashboard#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The precision of the type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.63.0/docs/resources/one_dashboard#precision OneDashboard#precision}
	Precision *float64 `field:"optional" json:"precision" yaml:"precision"`
}

