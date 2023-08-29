// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workload

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRuleNrqlQueryList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRuleNrqlQueryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRuleNrqlQueryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRuleNrqlQueryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRuleNrqlQueryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRuleNrqlQueryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkloadStatusConfigAutomaticRuleNrqlQueryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

