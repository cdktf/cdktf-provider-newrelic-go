//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package notificationdestination

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NotificationDestinationPropertiesList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NotificationDestinationPropertiesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NotificationDestinationPropertiesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NotificationDestinationPropertiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NotificationDestinationPropertiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NotificationDestinationPropertiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNotificationDestinationPropertiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

