package syntheticsmultilocationalertcondition


type SyntheticsMultilocationAlertConditionCritical struct {
	// The minimum number of monitor locations that must be concurrently failing before an incident is opened.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.22.0/docs/resources/synthetics_multilocation_alert_condition#threshold SyntheticsMultilocationAlertCondition#threshold}
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
}

