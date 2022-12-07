package onedashboard


type OneDashboardVariableItem struct {
	// A possible variable value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#value OneDashboard#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// A human-friendly display string for this value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#title OneDashboard#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

