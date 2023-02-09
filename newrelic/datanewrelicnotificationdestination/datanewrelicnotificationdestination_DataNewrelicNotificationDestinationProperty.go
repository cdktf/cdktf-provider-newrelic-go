package datanewrelicnotificationdestination


type DataNewrelicNotificationDestinationProperty struct {
	// Notification property key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/notification_destination#key DataNewrelicNotificationDestination#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Notification property value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/notification_destination#value DataNewrelicNotificationDestination#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Notification property display key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/notification_destination#display_value DataNewrelicNotificationDestination#display_value}
	DisplayValue *string `field:"optional" json:"displayValue" yaml:"displayValue"`
	// Notification property label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/notification_destination#label DataNewrelicNotificationDestination#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
}

