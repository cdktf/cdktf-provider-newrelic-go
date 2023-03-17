package onedashboard


type OneDashboardPageWidgetHistogramUnits struct {
	// series_overrides block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#series_overrides OneDashboard#series_overrides}
	SeriesOverrides interface{} `field:"optional" json:"seriesOverrides" yaml:"seriesOverrides"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#unit OneDashboard#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

