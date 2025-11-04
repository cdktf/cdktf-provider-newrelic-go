// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package entitytags

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EntityTagsTagList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EntityTagsTagList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EntityTagsTagList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EntityTagsTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EntityTagsTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EntityTagsTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EntityTagsTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEntityTagsTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

