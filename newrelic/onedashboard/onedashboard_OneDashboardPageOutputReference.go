package onedashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v3/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v3/onedashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OneDashboardPageOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	Guid() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WidgetArea() OneDashboardPageWidgetAreaList
	WidgetAreaInput() interface{}
	WidgetBar() OneDashboardPageWidgetBarList
	WidgetBarInput() interface{}
	WidgetBillboard() OneDashboardPageWidgetBillboardList
	WidgetBillboardInput() interface{}
	WidgetBullet() OneDashboardPageWidgetBulletList
	WidgetBulletInput() interface{}
	WidgetFunnel() OneDashboardPageWidgetFunnelList
	WidgetFunnelInput() interface{}
	WidgetHeatmap() OneDashboardPageWidgetHeatmapList
	WidgetHeatmapInput() interface{}
	WidgetHistogram() OneDashboardPageWidgetHistogramList
	WidgetHistogramInput() interface{}
	WidgetJson() OneDashboardPageWidgetJsonList
	WidgetJsonInput() interface{}
	WidgetLine() OneDashboardPageWidgetLineList
	WidgetLineInput() interface{}
	WidgetLogTable() OneDashboardPageWidgetLogTableList
	WidgetLogTableInput() interface{}
	WidgetMarkdown() OneDashboardPageWidgetMarkdownList
	WidgetMarkdownInput() interface{}
	WidgetPie() OneDashboardPageWidgetPieList
	WidgetPieInput() interface{}
	WidgetStackedBar() OneDashboardPageWidgetStackedBarList
	WidgetStackedBarInput() interface{}
	WidgetTable() OneDashboardPageWidgetTableList
	WidgetTableInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutWidgetArea(value interface{})
	PutWidgetBar(value interface{})
	PutWidgetBillboard(value interface{})
	PutWidgetBullet(value interface{})
	PutWidgetFunnel(value interface{})
	PutWidgetHeatmap(value interface{})
	PutWidgetHistogram(value interface{})
	PutWidgetJson(value interface{})
	PutWidgetLine(value interface{})
	PutWidgetLogTable(value interface{})
	PutWidgetMarkdown(value interface{})
	PutWidgetPie(value interface{})
	PutWidgetStackedBar(value interface{})
	PutWidgetTable(value interface{})
	ResetDescription()
	ResetWidgetArea()
	ResetWidgetBar()
	ResetWidgetBillboard()
	ResetWidgetBullet()
	ResetWidgetFunnel()
	ResetWidgetHeatmap()
	ResetWidgetHistogram()
	ResetWidgetJson()
	ResetWidgetLine()
	ResetWidgetLogTable()
	ResetWidgetMarkdown()
	ResetWidgetPie()
	ResetWidgetStackedBar()
	ResetWidgetTable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OneDashboardPageOutputReference
type jsiiProxy_OneDashboardPageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OneDashboardPageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) Guid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetArea() OneDashboardPageWidgetAreaList {
	var returns OneDashboardPageWidgetAreaList
	_jsii_.Get(
		j,
		"widgetArea",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetAreaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetAreaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetBar() OneDashboardPageWidgetBarList {
	var returns OneDashboardPageWidgetBarList
	_jsii_.Get(
		j,
		"widgetBar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetBarInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetBarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetBillboard() OneDashboardPageWidgetBillboardList {
	var returns OneDashboardPageWidgetBillboardList
	_jsii_.Get(
		j,
		"widgetBillboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetBillboardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetBillboardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetBullet() OneDashboardPageWidgetBulletList {
	var returns OneDashboardPageWidgetBulletList
	_jsii_.Get(
		j,
		"widgetBullet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetBulletInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetBulletInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetFunnel() OneDashboardPageWidgetFunnelList {
	var returns OneDashboardPageWidgetFunnelList
	_jsii_.Get(
		j,
		"widgetFunnel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetFunnelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetFunnelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetHeatmap() OneDashboardPageWidgetHeatmapList {
	var returns OneDashboardPageWidgetHeatmapList
	_jsii_.Get(
		j,
		"widgetHeatmap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetHeatmapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetHeatmapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetHistogram() OneDashboardPageWidgetHistogramList {
	var returns OneDashboardPageWidgetHistogramList
	_jsii_.Get(
		j,
		"widgetHistogram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetHistogramInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetHistogramInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetJson() OneDashboardPageWidgetJsonList {
	var returns OneDashboardPageWidgetJsonList
	_jsii_.Get(
		j,
		"widgetJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetJsonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetLine() OneDashboardPageWidgetLineList {
	var returns OneDashboardPageWidgetLineList
	_jsii_.Get(
		j,
		"widgetLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetLineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetLogTable() OneDashboardPageWidgetLogTableList {
	var returns OneDashboardPageWidgetLogTableList
	_jsii_.Get(
		j,
		"widgetLogTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetLogTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetLogTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetMarkdown() OneDashboardPageWidgetMarkdownList {
	var returns OneDashboardPageWidgetMarkdownList
	_jsii_.Get(
		j,
		"widgetMarkdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetMarkdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetMarkdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetPie() OneDashboardPageWidgetPieList {
	var returns OneDashboardPageWidgetPieList
	_jsii_.Get(
		j,
		"widgetPie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetPieInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetPieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetStackedBar() OneDashboardPageWidgetStackedBarList {
	var returns OneDashboardPageWidgetStackedBarList
	_jsii_.Get(
		j,
		"widgetStackedBar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetStackedBarInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetStackedBarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetTable() OneDashboardPageWidgetTableList {
	var returns OneDashboardPageWidgetTableList
	_jsii_.Get(
		j,
		"widgetTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageOutputReference) WidgetTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetTableInput",
		&returns,
	)
	return returns
}


func NewOneDashboardPageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OneDashboardPageOutputReference {
	_init_.Initialize()

	if err := validateNewOneDashboardPageOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OneDashboardPageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.oneDashboard.OneDashboardPageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOneDashboardPageOutputReference_Override(o OneDashboardPageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.oneDashboard.OneDashboardPageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OneDashboardPageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageOutputReference) PutWidgetArea(value interface{}) {
	if err := o.validatePutWidgetAreaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWidgetArea",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) PutWidgetBar(value interface{}) {
	if err := o.validatePutWidgetBarParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWidgetBar",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) PutWidgetBillboard(value interface{}) {
	if err := o.validatePutWidgetBillboardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWidgetBillboard",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) PutWidgetBullet(value interface{}) {
	if err := o.validatePutWidgetBulletParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWidgetBullet",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) PutWidgetFunnel(value interface{}) {
	if err := o.validatePutWidgetFunnelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWidgetFunnel",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) PutWidgetHeatmap(value interface{}) {
	if err := o.validatePutWidgetHeatmapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWidgetHeatmap",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) PutWidgetHistogram(value interface{}) {
	if err := o.validatePutWidgetHistogramParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWidgetHistogram",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) PutWidgetJson(value interface{}) {
	if err := o.validatePutWidgetJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWidgetJson",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) PutWidgetLine(value interface{}) {
	if err := o.validatePutWidgetLineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWidgetLine",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) PutWidgetLogTable(value interface{}) {
	if err := o.validatePutWidgetLogTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWidgetLogTable",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) PutWidgetMarkdown(value interface{}) {
	if err := o.validatePutWidgetMarkdownParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWidgetMarkdown",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) PutWidgetPie(value interface{}) {
	if err := o.validatePutWidgetPieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWidgetPie",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) PutWidgetStackedBar(value interface{}) {
	if err := o.validatePutWidgetStackedBarParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWidgetStackedBar",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) PutWidgetTable(value interface{}) {
	if err := o.validatePutWidgetTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWidgetTable",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		o,
		"resetDescription",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ResetWidgetArea() {
	_jsii_.InvokeVoid(
		o,
		"resetWidgetArea",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ResetWidgetBar() {
	_jsii_.InvokeVoid(
		o,
		"resetWidgetBar",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ResetWidgetBillboard() {
	_jsii_.InvokeVoid(
		o,
		"resetWidgetBillboard",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ResetWidgetBullet() {
	_jsii_.InvokeVoid(
		o,
		"resetWidgetBullet",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ResetWidgetFunnel() {
	_jsii_.InvokeVoid(
		o,
		"resetWidgetFunnel",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ResetWidgetHeatmap() {
	_jsii_.InvokeVoid(
		o,
		"resetWidgetHeatmap",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ResetWidgetHistogram() {
	_jsii_.InvokeVoid(
		o,
		"resetWidgetHistogram",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ResetWidgetJson() {
	_jsii_.InvokeVoid(
		o,
		"resetWidgetJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ResetWidgetLine() {
	_jsii_.InvokeVoid(
		o,
		"resetWidgetLine",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ResetWidgetLogTable() {
	_jsii_.InvokeVoid(
		o,
		"resetWidgetLogTable",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ResetWidgetMarkdown() {
	_jsii_.InvokeVoid(
		o,
		"resetWidgetMarkdown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ResetWidgetPie() {
	_jsii_.InvokeVoid(
		o,
		"resetWidgetPie",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ResetWidgetStackedBar() {
	_jsii_.InvokeVoid(
		o,
		"resetWidgetStackedBar",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ResetWidgetTable() {
	_jsii_.InvokeVoid(
		o,
		"resetWidgetTable",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

