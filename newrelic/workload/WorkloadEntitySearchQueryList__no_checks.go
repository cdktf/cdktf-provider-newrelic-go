// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workload

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkloadEntitySearchQueryList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkloadEntitySearchQueryList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkloadEntitySearchQueryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkloadEntitySearchQueryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkloadEntitySearchQueryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkloadEntitySearchQueryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkloadEntitySearchQueryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkloadEntitySearchQueryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

