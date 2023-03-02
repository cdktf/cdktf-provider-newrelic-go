package notificationdestination


type NotificationDestinationAuthToken struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#token NotificationDestination#token}.
	Token *string `field:"required" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#prefix NotificationDestination#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

