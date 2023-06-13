package onedashboard


type OneDashboardPageWidgetLineColorsSeriesOverrides struct {
	// Color code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/one_dashboard#color OneDashboard#color}
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Series name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/one_dashboard#series_name OneDashboard#series_name}
	SeriesName *string `field:"optional" json:"seriesName" yaml:"seriesName"`
}

