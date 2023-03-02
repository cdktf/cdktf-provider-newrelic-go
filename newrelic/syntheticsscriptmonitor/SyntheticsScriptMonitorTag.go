package syntheticsscriptmonitor


type SyntheticsScriptMonitorTag struct {
	// Name of the tag key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_script_monitor#key SyntheticsScriptMonitor#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Values associated with the tag key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_script_monitor#values SyntheticsScriptMonitor#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

