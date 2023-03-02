package workflow


type WorkflowDestination struct {
	// (Required) Destination's channel id.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#channel_id Workflow#channel_id}
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// List of triggers to notify about in this destination configuration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#notification_triggers Workflow#notification_triggers}
	NotificationTriggers *[]*string `field:"optional" json:"notificationTriggers" yaml:"notificationTriggers"`
}

