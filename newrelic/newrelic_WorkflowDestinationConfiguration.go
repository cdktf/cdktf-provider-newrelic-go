// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic


type WorkflowDestinationConfiguration struct {
	// (Required) Destination's channel id.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#channel_id Workflow#channel_id}
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
}

