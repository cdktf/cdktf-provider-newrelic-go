// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package entitytags


type EntityTagsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.69.0/docs/resources/entity_tags#create EntityTags#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

