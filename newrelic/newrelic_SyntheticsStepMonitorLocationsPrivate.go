// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type SyntheticsStepMonitorLocationsPrivate struct {
	// The unique identifier for the Synthetics private location in New Relic.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_step_monitor#guid SyntheticsStepMonitor#guid}
	Guid *string `field:"required" json:"guid" yaml:"guid"`
	// The location's Verified Script Execution password (Only necessary if Verified Script Execution is enabled for the location).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_step_monitor#vse_password SyntheticsStepMonitor#vse_password}
	VsePassword *string `field:"optional" json:"vsePassword" yaml:"vsePassword"`
}

