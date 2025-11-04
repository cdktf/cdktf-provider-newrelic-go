// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package applicationsettings

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationSettingsTransactionTracerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationSettingsTransactionTracerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

