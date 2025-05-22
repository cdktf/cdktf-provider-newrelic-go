// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package insightsevent


type InsightsEventEvent struct {
	// attribute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs/resources/insights_event#attribute InsightsEvent#attribute}
	Attribute interface{} `field:"required" json:"attribute" yaml:"attribute"`
	// The event's name. Can be a combination of alphanumeric characters, underscores, and colons.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs/resources/insights_event#type InsightsEvent#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Must be a Unix epoch timestamp. You can define timestamps either in seconds or in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs/resources/insights_event#timestamp InsightsEvent#timestamp}
	Timestamp *float64 `field:"optional" json:"timestamp" yaml:"timestamp"`
}

