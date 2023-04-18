package workflow


type WorkflowDestination struct {
	// (Required) Destination's channel id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/workflow#channel_id Workflow#channel_id}
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// List of triggers to notify about in this destination configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/workflow#notification_triggers Workflow#notification_triggers}
	NotificationTriggers *[]*string `field:"optional" json:"notificationTriggers" yaml:"notificationTriggers"`
}

