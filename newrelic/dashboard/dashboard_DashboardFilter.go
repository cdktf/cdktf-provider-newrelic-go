package dashboard


type DashboardFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#event_types Dashboard#event_types}.
	EventTypes *[]*string `field:"required" json:"eventTypes" yaml:"eventTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#attributes Dashboard#attributes}.
	Attributes *[]*string `field:"optional" json:"attributes" yaml:"attributes"`
}

