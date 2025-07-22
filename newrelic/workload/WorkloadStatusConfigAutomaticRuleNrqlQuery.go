// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workload


type WorkloadStatusConfigAutomaticRuleNrqlQuery struct {
	// The entity search query that is used to perform the search of a group of entities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.64.0/docs/resources/workload#query Workload#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

