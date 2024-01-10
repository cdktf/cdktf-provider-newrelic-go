// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package notificationdestination

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NotificationDestinationPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NotificationDestinationPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NotificationDestinationPropertyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NotificationDestinationPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NotificationDestinationPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NotificationDestinationPropertyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NotificationDestinationPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNotificationDestinationPropertyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

