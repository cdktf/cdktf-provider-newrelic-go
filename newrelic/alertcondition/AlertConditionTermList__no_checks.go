// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package alertcondition

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertConditionTermList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AlertConditionTermList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlertConditionTermList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlertConditionTermList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlertConditionTermList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlertConditionTermList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlertConditionTermList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlertConditionTermListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

