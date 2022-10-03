package syntheticsmonitorscript


type SyntheticsMonitorScriptLocation struct {
	// The monitor script location name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor_script#name SyntheticsMonitorScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The HMAC for the monitor script location. Use only one of `hmac` or `vse_password.`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor_script#hmac SyntheticsMonitorScript#hmac}
	Hmac *string `field:"optional" json:"hmac" yaml:"hmac"`
	// The password for the monitor script location used to calculate HMAC. Use only one of `vse_password` or `hmac.`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor_script#vse_password SyntheticsMonitorScript#vse_password}
	VsePassword *string `field:"optional" json:"vsePassword" yaml:"vsePassword"`
}

