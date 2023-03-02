package syntheticsbrokenlinksmonitor


type SyntheticsBrokenLinksMonitorTag struct {
	// Name of the tag key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_broken_links_monitor#key SyntheticsBrokenLinksMonitor#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Values associated with the tag key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_broken_links_monitor#values SyntheticsBrokenLinksMonitor#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

