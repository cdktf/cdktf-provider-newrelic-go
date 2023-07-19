package onedashboard


type OneDashboardPageWidgetLineNullValues struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.2/docs/resources/one_dashboard#null_value OneDashboard#null_value}.
	NullValue *string `field:"optional" json:"nullValue" yaml:"nullValue"`
	// series_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.2/docs/resources/one_dashboard#series_overrides OneDashboard#series_overrides}
	SeriesOverrides interface{} `field:"optional" json:"seriesOverrides" yaml:"seriesOverrides"`
}

