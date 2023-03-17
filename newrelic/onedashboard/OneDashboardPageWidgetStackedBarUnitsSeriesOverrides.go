package onedashboard


type OneDashboardPageWidgetStackedBarUnitsSeriesOverrides struct {
	// Series name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#series_name OneDashboard#series_name}
	SeriesName *string `field:"optional" json:"seriesName" yaml:"seriesName"`
	// Unit name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#unit OneDashboard#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

