// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type DashboardWidgetCompareWithPresentation struct {
	// The color for the rendered data.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#color Dashboard#color}
	Color *string `field:"required" json:"color" yaml:"color"`
	// The name for the rendered data.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

