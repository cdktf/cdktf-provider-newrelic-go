//go:build no_runtime_type_checking

package notificationchannel

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NotificationChannelPropertiesList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NotificationChannelPropertiesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NotificationChannelPropertiesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NotificationChannelPropertiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NotificationChannelPropertiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NotificationChannelPropertiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNotificationChannelPropertiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

