// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type ServiceLevelObjectiveTimeWindowRolling struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/service_level#count ServiceLevel#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/service_level#unit ServiceLevel#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
}

