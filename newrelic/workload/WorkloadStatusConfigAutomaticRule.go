// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workload


type WorkloadStatusConfigAutomaticRule struct {
	// rollup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.74.0/docs/resources/workload#rollup Workload#rollup}
	Rollup *WorkloadStatusConfigAutomaticRuleRollup `field:"required" json:"rollup" yaml:"rollup"`
	// A list of entity GUIDs composing the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.74.0/docs/resources/workload#entity_guids Workload#entity_guids}
	EntityGuids *[]*string `field:"optional" json:"entityGuids" yaml:"entityGuids"`
	// nrql_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.74.0/docs/resources/workload#nrql_query Workload#nrql_query}
	NrqlQuery interface{} `field:"optional" json:"nrqlQuery" yaml:"nrqlQuery"`
}

