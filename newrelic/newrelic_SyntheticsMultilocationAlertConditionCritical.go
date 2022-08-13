// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type SyntheticsMultilocationAlertConditionCritical struct {
	// The minimum number of monitor locations that must be concurrently failing before a violation is opened.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_multilocation_alert_condition#threshold SyntheticsMultilocationAlertCondition#threshold}
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
}

