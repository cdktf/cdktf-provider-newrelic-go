// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsMonitorConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The interval (in minutes) at which this monitor should run.
	//
	// Valid values are 1, 5, 10, 15, 30, 60, 360, 720, or 1440.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#frequency SyntheticsMonitor#frequency}
	Frequency *float64 `field:"required" json:"frequency" yaml:"frequency"`
	// The locations in which this monitor should be run.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#locations SyntheticsMonitor#locations}
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
	// The title of this monitor.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#name SyntheticsMonitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The monitor status (i.e. ENABLED, MUTED, DISABLED).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#status SyntheticsMonitor#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// The monitor type. Valid values are SIMPLE, BROWSER, SCRIPT_BROWSER, and SCRIPT_API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#type SyntheticsMonitor#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Bypass HEAD request.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#bypass_head_request SyntheticsMonitor#bypass_head_request}
	BypassHeadRequest interface{} `field:"optional" json:"bypassHeadRequest" yaml:"bypassHeadRequest"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#id SyntheticsMonitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The base threshold (in seconds) to calculate the apdex score for use in the SLA report. (Default 7 seconds).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#sla_threshold SyntheticsMonitor#sla_threshold}
	SlaThreshold *float64 `field:"optional" json:"slaThreshold" yaml:"slaThreshold"`
	// Fail the monitor check if redirected.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#treat_redirect_as_failure SyntheticsMonitor#treat_redirect_as_failure}
	TreatRedirectAsFailure interface{} `field:"optional" json:"treatRedirectAsFailure" yaml:"treatRedirectAsFailure"`
	// The URI for the monitor to hit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#uri SyntheticsMonitor#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
	// The string to validate against in the response.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#validation_string SyntheticsMonitor#validation_string}
	ValidationString *string `field:"optional" json:"validationString" yaml:"validationString"`
	// Verify SSL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_monitor#verify_ssl SyntheticsMonitor#verify_ssl}
	VerifySsl interface{} `field:"optional" json:"verifySsl" yaml:"verifySsl"`
}

