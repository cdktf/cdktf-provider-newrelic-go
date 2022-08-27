// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type NotificationDestinationProperty struct {
	// Notification property key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#key NotificationDestination#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Notification property value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#value NotificationDestination#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Notification property display key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#display_value NotificationDestination#display_value}
	DisplayValue *string `field:"optional" json:"displayValue" yaml:"displayValue"`
	// Notification property label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#label NotificationDestination#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
}

