// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package insightsevent

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InsightsEventEventAttributeList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_InsightsEventEventAttributeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_InsightsEventEventAttributeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InsightsEventEventAttributeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InsightsEventEventAttributeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_InsightsEventEventAttributeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewInsightsEventEventAttributeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

