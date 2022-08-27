// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type SyntheticsCertCheckMonitorTag struct {
	// Name of the tag key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_cert_check_monitor#key SyntheticsCertCheckMonitor#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Values associated with the tag key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_cert_check_monitor#values SyntheticsCertCheckMonitor#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

