package syntheticsscriptmonitor


type SyntheticsScriptMonitorLocationPrivate struct {
	// The unique identifier for the Synthetics private location in New Relic.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_script_monitor#guid SyntheticsScriptMonitor#guid}
	Guid *string `field:"required" json:"guid" yaml:"guid"`
	// The location's Verified Script Execution password (Only necessary if Verified Script Execution is enabled for the location).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_script_monitor#vse_password SyntheticsScriptMonitor#vse_password}
	VsePassword *string `field:"optional" json:"vsePassword" yaml:"vsePassword"`
}

