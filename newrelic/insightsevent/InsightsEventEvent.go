package insightsevent


type InsightsEventEvent struct {
	// attribute block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/insights_event#attribute InsightsEvent#attribute}
	Attribute interface{} `field:"required" json:"attribute" yaml:"attribute"`
	// The event's name. Can be a combination of alphanumeric characters, underscores, and colons.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/insights_event#type InsightsEvent#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Must be a Unix epoch timestamp. You can define timestamps either in seconds or in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/insights_event#timestamp InsightsEvent#timestamp}
	Timestamp *float64 `field:"optional" json:"timestamp" yaml:"timestamp"`
}
