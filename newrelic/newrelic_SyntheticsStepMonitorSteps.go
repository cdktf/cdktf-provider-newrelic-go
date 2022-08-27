// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type SyntheticsStepMonitorSteps struct {
	// The position of the step within the script ranging from 0-100.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_step_monitor#ordinal SyntheticsStepMonitor#ordinal}
	Ordinal *float64 `field:"required" json:"ordinal" yaml:"ordinal"`
	// The type of step to be added to the script.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_step_monitor#type SyntheticsStepMonitor#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The metadata values related to the check the step performs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_step_monitor#values SyntheticsStepMonitor#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

