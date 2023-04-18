package servicelevel


type ServiceLevelObjectiveTimeWindowRolling struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/service_level#count ServiceLevel#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/service_level#unit ServiceLevel#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
}

