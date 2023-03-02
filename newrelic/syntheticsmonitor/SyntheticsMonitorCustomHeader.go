package syntheticsmonitor


type SyntheticsMonitorCustomHeader struct {
	// Header name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#name SyntheticsMonitor#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Header value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#value SyntheticsMonitor#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

