package onedashboard


type OneDashboardPageWidgetAreaNullValues struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#null_value OneDashboard#null_value}.
	NullValue *string `field:"optional" json:"nullValue" yaml:"nullValue"`
	// series_overrides block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#series_overrides OneDashboard#series_overrides}
	SeriesOverrides interface{} `field:"optional" json:"seriesOverrides" yaml:"seriesOverrides"`
}

