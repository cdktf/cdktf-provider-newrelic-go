package onedashboard


type OneDashboardPageWidgetHeatmapUnits struct {
	// series_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.1/docs/resources/one_dashboard#series_overrides OneDashboard#series_overrides}
	SeriesOverrides interface{} `field:"optional" json:"seriesOverrides" yaml:"seriesOverrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.1/docs/resources/one_dashboard#unit OneDashboard#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

