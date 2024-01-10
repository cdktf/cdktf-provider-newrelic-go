// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package onedashboard

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OneDashboardPageList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OneDashboardPageList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OneDashboardPageList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OneDashboardPageList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OneDashboardPageList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OneDashboardPageList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OneDashboardPageList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOneDashboardPageListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

