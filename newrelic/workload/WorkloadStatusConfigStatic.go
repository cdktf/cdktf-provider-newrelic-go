// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workload


type WorkloadStatusConfigStatic struct {
	// Whether the static status configuration is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.0/docs/resources/workload#enabled Workload#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The status of the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.0/docs/resources/workload#status Workload#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// A description that provides additional details about the status of the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.0/docs/resources/workload#description Workload#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A short description of the status of the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.0/docs/resources/workload#summary Workload#summary}
	Summary *string `field:"optional" json:"summary" yaml:"summary"`
}

