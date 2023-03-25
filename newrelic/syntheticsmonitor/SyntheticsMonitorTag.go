package syntheticsmonitor


type SyntheticsMonitorTag struct {
	// Name of the tag key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#key SyntheticsMonitor#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Values associated with the tag key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#values SyntheticsMonitor#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}
