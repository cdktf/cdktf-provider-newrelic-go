package servicelevel


type ServiceLevelObjectiveTimeWindow struct {
	// rolling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/service_level#rolling ServiceLevel#rolling}
	Rolling *ServiceLevelObjectiveTimeWindowRolling `field:"required" json:"rolling" yaml:"rolling"`
}

