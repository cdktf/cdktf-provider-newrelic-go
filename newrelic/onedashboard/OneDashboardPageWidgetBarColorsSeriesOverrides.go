package onedashboard


type OneDashboardPageWidgetBarColorsSeriesOverrides struct {
	// Color code.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#color OneDashboard#color}
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Series name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#series_name OneDashboard#series_name}
	SeriesName *string `field:"optional" json:"seriesName" yaml:"seriesName"`
}

