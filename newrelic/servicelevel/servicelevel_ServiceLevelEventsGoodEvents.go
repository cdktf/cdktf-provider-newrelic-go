package servicelevel


type ServiceLevelEventsGoodEvents struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/service_level#from ServiceLevel#from}.
	From *string `field:"required" json:"from" yaml:"from"`
	// select block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/service_level#select ServiceLevel#select}
	Select *ServiceLevelEventsGoodEventsSelect `field:"optional" json:"select" yaml:"select"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/service_level#where ServiceLevel#where}.
	Where *string `field:"optional" json:"where" yaml:"where"`
}

