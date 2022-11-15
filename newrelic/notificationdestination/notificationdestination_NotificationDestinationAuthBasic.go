package notificationdestination


type NotificationDestinationAuthBasic struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#password NotificationDestination#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#user NotificationDestination#user}.
	User *string `field:"required" json:"user" yaml:"user"`
}

