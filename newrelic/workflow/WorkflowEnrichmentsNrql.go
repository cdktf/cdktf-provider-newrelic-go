// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workflow


type WorkflowEnrichmentsNrql struct {
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.1/docs/resources/workflow#configuration Workflow#configuration}
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// (Required) Enrichment's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.1/docs/resources/workflow#name Workflow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

