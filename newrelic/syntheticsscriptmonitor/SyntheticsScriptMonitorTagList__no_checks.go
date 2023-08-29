// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package syntheticsscriptmonitor

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SyntheticsScriptMonitorTagList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsScriptMonitorTagList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SyntheticsScriptMonitorTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SyntheticsScriptMonitorTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SyntheticsScriptMonitorTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SyntheticsScriptMonitorTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSyntheticsScriptMonitorTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

