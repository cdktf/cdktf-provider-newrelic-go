package onedashboard


type OneDashboardPageWidgetMarkdownColors struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#color OneDashboard#color}.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// series_overrides block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#series_overrides OneDashboard#series_overrides}
	SeriesOverrides interface{} `field:"optional" json:"seriesOverrides" yaml:"seriesOverrides"`
}

