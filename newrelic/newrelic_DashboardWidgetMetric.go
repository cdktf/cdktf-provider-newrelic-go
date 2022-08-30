// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type DashboardWidgetMetric struct {
	// The metric name to display.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The metric scope.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#scope Dashboard#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The metric units.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#units Dashboard#units}
	Units *string `field:"optional" json:"units" yaml:"units"`
	// The metric values to display.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#values Dashboard#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

