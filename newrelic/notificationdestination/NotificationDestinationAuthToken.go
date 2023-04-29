package notificationdestination


type NotificationDestinationAuthToken struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.21.3/docs/resources/notification_destination#token NotificationDestination#token}.
	Token *string `field:"required" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.21.3/docs/resources/notification_destination#prefix NotificationDestination#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

