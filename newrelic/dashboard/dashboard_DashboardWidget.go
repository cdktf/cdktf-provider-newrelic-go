package dashboard


type DashboardWidget struct {
	// Column position of widget from top left, starting at 1.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#column Dashboard#column}
	Column *float64 `field:"required" json:"column" yaml:"column"`
	// Row position of widget from top left, starting at 1.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#row Dashboard#row}
	Row *float64 `field:"required" json:"row" yaml:"row"`
	// A title for the widget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#title Dashboard#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// How the widget visualizes data.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#visualization Dashboard#visualization}
	Visualization *string `field:"required" json:"visualization" yaml:"visualization"`
	// The target account ID to fetch data from, if not the current account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#account_id Dashboard#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// compare_with block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#compare_with Dashboard#compare_with}
	CompareWith interface{} `field:"optional" json:"compareWith" yaml:"compareWith"`
	// The ID of a dashboard to link to from the widget's facets.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#drilldown_dashboard_id Dashboard#drilldown_dashboard_id}
	DrilldownDashboardId *float64 `field:"optional" json:"drilldownDashboardId" yaml:"drilldownDashboardId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#duration Dashboard#duration}.
	Duration *float64 `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#end_time Dashboard#end_time}.
	EndTime *float64 `field:"optional" json:"endTime" yaml:"endTime"`
	// A collection of entity ids to display data for. These are typically application IDs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#entity_ids Dashboard#entity_ids}
	EntityIds *[]*float64 `field:"optional" json:"entityIds" yaml:"entityIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#facet Dashboard#facet}.
	Facet *string `field:"optional" json:"facet" yaml:"facet"`
	// Height of the widget. Valid values are 1 to 3 inclusive. Defaults to 1.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#height Dashboard#height}
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The limit of distinct data series to display.  Requires `order_by` to be set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#limit Dashboard#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#metric Dashboard#metric}
	Metric interface{} `field:"optional" json:"metric" yaml:"metric"`
	// Description of the widget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#notes Dashboard#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// Valid NRQL query string.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#nrql Dashboard#nrql}
	Nrql *string `field:"optional" json:"nrql" yaml:"nrql"`
	// Set the order of result series.  Required when using `limit`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#order_by Dashboard#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// The markdown source to be rendered in the widget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#source Dashboard#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Threshold above which the displayed value will be styled with a red color.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#threshold_red Dashboard#threshold_red}
	ThresholdRed *float64 `field:"optional" json:"thresholdRed" yaml:"thresholdRed"`
	// Threshold above which the displayed value will be styled with a yellow color.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#threshold_yellow Dashboard#threshold_yellow}
	ThresholdYellow *float64 `field:"optional" json:"thresholdYellow" yaml:"thresholdYellow"`
	// Width of the widget. Valid values are 1 to 3 inclusive. Defaults to 1.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#width Dashboard#width}
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

