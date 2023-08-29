// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package syntheticscertcheckmonitor

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SyntheticsCertCheckMonitorTagList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsCertCheckMonitorTagList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SyntheticsCertCheckMonitorTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SyntheticsCertCheckMonitorTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SyntheticsCertCheckMonitorTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SyntheticsCertCheckMonitorTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSyntheticsCertCheckMonitorTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

