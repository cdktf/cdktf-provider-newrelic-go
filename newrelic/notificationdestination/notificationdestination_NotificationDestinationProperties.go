package notificationdestination


type NotificationDestinationProperties struct {
	// Notification destination property key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#key NotificationDestination#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Notification destination property value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#value NotificationDestination#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Notification destination property display key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#display_value NotificationDestination#display_value}
	DisplayValue *string `field:"optional" json:"displayValue" yaml:"displayValue"`
	// Notification destination property label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#label NotificationDestination#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
}

