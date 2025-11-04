// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workflow

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkflowEnrichmentsNrqlList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkflowEnrichmentsNrqlList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkflowEnrichmentsNrqlList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkflowEnrichmentsNrqlList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkflowEnrichmentsNrqlList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkflowEnrichmentsNrqlList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkflowEnrichmentsNrqlList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkflowEnrichmentsNrqlListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

