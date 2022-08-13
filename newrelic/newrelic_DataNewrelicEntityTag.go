// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type DataNewrelicEntityTag struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/entity#key DataNewrelicEntity#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/entity#value DataNewrelicEntity#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

