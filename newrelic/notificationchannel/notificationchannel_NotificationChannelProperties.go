package notificationchannel


type NotificationChannelProperties struct {
	// Notification channel property key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_channel#key NotificationChannel#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Notification channel property value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_channel#value NotificationChannel#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Notification channel property display key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_channel#display_value NotificationChannel#display_value}
	DisplayValue *string `field:"optional" json:"displayValue" yaml:"displayValue"`
	// Notification channel property label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_channel#label NotificationChannel#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
}

