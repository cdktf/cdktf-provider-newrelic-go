// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workload


type WorkloadEntitySearchQuery struct {
	// A valid entity search query; empty, and null values are considered invalid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.73.0/docs/resources/workload#query Workload#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

