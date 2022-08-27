// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type SyntheticsStepMonitorTag struct {
	// Name of the tag key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_step_monitor#key SyntheticsStepMonitor#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Values associated with the tag key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_step_monitor#values SyntheticsStepMonitor#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

