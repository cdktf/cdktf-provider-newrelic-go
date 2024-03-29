// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package insightsevent


type InsightsEventEventAttribute struct {
	// The name of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.34.1/docs/resources/insights_event#key InsightsEvent#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.34.1/docs/resources/insights_event#value InsightsEvent#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Specify the type for the attribute value.
	//
	// This is useful when passing integer or float values to Insights. Allowed values are string, int, or float. Defaults to string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.34.1/docs/resources/insights_event#type InsightsEvent#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

