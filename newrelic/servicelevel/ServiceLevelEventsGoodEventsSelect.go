package servicelevel


type ServiceLevelEventsGoodEventsSelect struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/service_level#function ServiceLevel#function}.
	Function *string `field:"required" json:"function" yaml:"function"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/service_level#attribute ServiceLevel#attribute}.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
}

