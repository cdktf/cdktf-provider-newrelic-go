package servicelevel


type ServiceLevelEventsBadEventsSelect struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/service_level#function ServiceLevel#function}.
	Function *string `field:"required" json:"function" yaml:"function"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/service_level#attribute ServiceLevel#attribute}.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// The event threshold to use in the SELECT clause.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/service_level#threshold ServiceLevel#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
}

