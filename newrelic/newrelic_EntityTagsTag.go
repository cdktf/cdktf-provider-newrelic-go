// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type EntityTagsTag struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/entity_tags#key EntityTags#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag values.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/entity_tags#values EntityTags#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

