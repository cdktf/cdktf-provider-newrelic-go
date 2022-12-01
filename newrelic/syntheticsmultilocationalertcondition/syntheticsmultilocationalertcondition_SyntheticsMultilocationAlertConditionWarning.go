package syntheticsmultilocationalertcondition


type SyntheticsMultilocationAlertConditionWarning struct {
	// The minimum number of monitor locations that must be concurrently failing before an incident is opened.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/synthetics_multilocation_alert_condition#threshold SyntheticsMultilocationAlertCondition#threshold}
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
}

