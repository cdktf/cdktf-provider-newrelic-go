// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package nrqlalertcondition

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NrqlAlertConditionTermList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NrqlAlertConditionTermList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NrqlAlertConditionTermList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NrqlAlertConditionTermList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NrqlAlertConditionTermList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NrqlAlertConditionTermList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NrqlAlertConditionTermList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNrqlAlertConditionTermListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

