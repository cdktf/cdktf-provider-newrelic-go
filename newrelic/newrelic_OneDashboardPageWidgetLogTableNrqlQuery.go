// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type OneDashboardPageWidgetLogTableNrqlQuery struct {
	// The NRQL query.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#query OneDashboard#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The account id used for the NRQL query.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/one_dashboard#account_id OneDashboard#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
}

