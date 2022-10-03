package dashboard


type DashboardWidgetCompareWith struct {
	// The offset duration for the COMPARE WITH clause.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#offset_duration Dashboard#offset_duration}
	OffsetDuration *string `field:"required" json:"offsetDuration" yaml:"offsetDuration"`
	// presentation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#presentation Dashboard#presentation}
	Presentation *DashboardWidgetCompareWithPresentation `field:"required" json:"presentation" yaml:"presentation"`
}

