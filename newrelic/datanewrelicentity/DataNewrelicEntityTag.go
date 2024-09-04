// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datanewrelicentity


type DataNewrelicEntityTag struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.45.0/docs/data-sources/entity#key DataNewrelicEntity#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.45.0/docs/data-sources/entity#value DataNewrelicEntity#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

