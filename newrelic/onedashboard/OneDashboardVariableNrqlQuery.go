package onedashboard


type OneDashboardVariableNrqlQuery struct {
	// NRQL formatted query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/one_dashboard#query OneDashboard#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// New Relic account ID(s) to issue the query against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/one_dashboard#account_ids OneDashboard#account_ids}
	AccountIds *[]*float64 `field:"optional" json:"accountIds" yaml:"accountIds"`
}

