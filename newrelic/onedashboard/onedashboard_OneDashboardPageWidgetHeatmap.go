package onedashboard


type OneDashboardPageWidgetHeatmap struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#column OneDashboard#column}.
	Column *float64 `field:"required" json:"column" yaml:"column"`
	// nrql_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#nrql_query OneDashboard#nrql_query}
	NrqlQuery interface{} `field:"required" json:"nrqlQuery" yaml:"nrqlQuery"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#row OneDashboard#row}.
	Row *float64 `field:"required" json:"row" yaml:"row"`
	// A title for the widget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#title OneDashboard#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Use this item to filter the current dashboard.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#filter_current_dashboard OneDashboard#filter_current_dashboard}
	FilterCurrentDashboard interface{} `field:"optional" json:"filterCurrentDashboard" yaml:"filterCurrentDashboard"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#height OneDashboard#height}.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#ignore_time_range OneDashboard#ignore_time_range}.
	IgnoreTimeRange interface{} `field:"optional" json:"ignoreTimeRange" yaml:"ignoreTimeRange"`
	// Related entities. Currently only supports Dashboard entities, but may allow other cases in the future.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#linked_entity_guids OneDashboard#linked_entity_guids}
	LinkedEntityGuids *[]*string `field:"optional" json:"linkedEntityGuids" yaml:"linkedEntityGuids"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#width OneDashboard#width}.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

