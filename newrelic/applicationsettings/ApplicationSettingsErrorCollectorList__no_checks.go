// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package applicationsettings

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationSettingsErrorCollectorList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettingsErrorCollectorList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettingsErrorCollectorList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationSettingsErrorCollectorListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

