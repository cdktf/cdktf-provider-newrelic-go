package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v3/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v3/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetOutputReference interface {
	cdktf.ComplexObject
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
	Column() *float64
	SetColumn(val *float64)
	ColumnInput() *float64
	CompareWith() DashboardWidgetCompareWithList
	CompareWithInput() interface{}
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
	DrilldownDashboardId() *float64
	SetDrilldownDashboardId(val *float64)
	DrilldownDashboardIdInput() *float64
	Duration() *float64
	SetDuration(val *float64)
	DurationInput() *float64
	EndTime() *float64
	SetEndTime(val *float64)
	EndTimeInput() *float64
	EntityIds() *[]*float64
	SetEntityIds(val *[]*float64)
	EntityIdsInput() *[]*float64
	Facet() *string
	SetFacet(val *string)
	FacetInput() *string
	// Experimental.
	Fqn() *string
	Height() *float64
	SetHeight(val *float64)
	HeightInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Limit() *float64
	SetLimit(val *float64)
	LimitInput() *float64
	Metric() DashboardWidgetMetricList
	MetricInput() interface{}
	Notes() *string
	SetNotes(val *string)
	NotesInput() *string
	Nrql() *string
	SetNrql(val *string)
	NrqlInput() *string
	OrderBy() *string
	SetOrderBy(val *string)
	OrderByInput() *string
	RawMetricName() *string
	Row() *float64
	SetRow(val *float64)
	RowInput() *float64
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThresholdRed() *float64
	SetThresholdRed(val *float64)
	ThresholdRedInput() *float64
	ThresholdYellow() *float64
	SetThresholdYellow(val *float64)
	ThresholdYellowInput() *float64
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	Visualization() *string
	SetVisualization(val *string)
	VisualizationInput() *string
	WidgetId() *float64
	Width() *float64
	SetWidth(val *float64)
	WidthInput() *float64
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
	PutCompareWith(value interface{})
	PutMetric(value interface{})
	ResetAccountId()
	ResetCompareWith()
	ResetDrilldownDashboardId()
	ResetDuration()
	ResetEndTime()
	ResetEntityIds()
	ResetFacet()
	ResetHeight()
	ResetLimit()
	ResetMetric()
	ResetNotes()
	ResetNrql()
	ResetOrderBy()
	ResetSource()
	ResetThresholdRed()
	ResetThresholdYellow()
	ResetWidth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetOutputReference
type jsiiProxy_DashboardWidgetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetOutputReference) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Column() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ColumnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) CompareWith() DashboardWidgetCompareWithList {
	var returns DashboardWidgetCompareWithList
	_jsii_.Get(
		j,
		"compareWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) CompareWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compareWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) DrilldownDashboardId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drilldownDashboardId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) DrilldownDashboardIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drilldownDashboardIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Duration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) DurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) EndTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) EndTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) EntityIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"entityIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) EntityIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"entityIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Facet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"facet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) FacetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"facetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) HeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) LimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Metric() DashboardWidgetMetricList {
	var returns DashboardWidgetMetricList
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) MetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Notes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) NotesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Nrql() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nrql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) NrqlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nrqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) OrderBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) OrderByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) RawMetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rawMetricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Row() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"row",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) RowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ThresholdRed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdRed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ThresholdRedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdRedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ThresholdYellow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdYellow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ThresholdYellowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdYellowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Visualization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) VisualizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visualizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) WidgetId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widgetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) WidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widthInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardWidgetOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.dashboard.DashboardWidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardWidgetOutputReference_Override(d DashboardWidgetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.dashboard.DashboardWidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetAccountId(val *float64) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetColumn(val *float64) {
	if err := j.validateSetColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"column",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetDrilldownDashboardId(val *float64) {
	if err := j.validateSetDrilldownDashboardIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drilldownDashboardId",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetDuration(val *float64) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetEndTime(val *float64) {
	if err := j.validateSetEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetEntityIds(val *[]*float64) {
	if err := j.validateSetEntityIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityIds",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetFacet(val *string) {
	if err := j.validateSetFacetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"facet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetHeight(val *float64) {
	if err := j.validateSetHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"height",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetLimit(val *float64) {
	if err := j.validateSetLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetNotes(val *string) {
	if err := j.validateSetNotesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notes",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetNrql(val *string) {
	if err := j.validateSetNrqlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nrql",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetOrderBy(val *string) {
	if err := j.validateSetOrderByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orderBy",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetRow(val *float64) {
	if err := j.validateSetRowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"row",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetThresholdRed(val *float64) {
	if err := j.validateSetThresholdRedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdRed",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetThresholdYellow(val *float64) {
	if err := j.validateSetThresholdYellowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdYellow",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetVisualization(val *string) {
	if err := j.validateSetVisualizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visualization",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetWidth(val *float64) {
	if err := j.validateSetWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"width",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutCompareWith(value interface{}) {
	if err := d.validatePutCompareWithParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCompareWith",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutMetric(value interface{}) {
	if err := d.validatePutMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetric",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetAccountId() {
	_jsii_.InvokeVoid(
		d,
		"resetAccountId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetCompareWith() {
	_jsii_.InvokeVoid(
		d,
		"resetCompareWith",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetDrilldownDashboardId() {
	_jsii_.InvokeVoid(
		d,
		"resetDrilldownDashboardId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetDuration() {
	_jsii_.InvokeVoid(
		d,
		"resetDuration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetEndTime() {
	_jsii_.InvokeVoid(
		d,
		"resetEndTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetEntityIds() {
	_jsii_.InvokeVoid(
		d,
		"resetEntityIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetFacet() {
	_jsii_.InvokeVoid(
		d,
		"resetFacet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetHeight() {
	_jsii_.InvokeVoid(
		d,
		"resetHeight",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetMetric() {
	_jsii_.InvokeVoid(
		d,
		"resetMetric",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetNotes() {
	_jsii_.InvokeVoid(
		d,
		"resetNotes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetNrql() {
	_jsii_.InvokeVoid(
		d,
		"resetNrql",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetOrderBy() {
	_jsii_.InvokeVoid(
		d,
		"resetOrderBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		d,
		"resetSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetThresholdRed() {
	_jsii_.InvokeVoid(
		d,
		"resetThresholdRed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetThresholdYellow() {
	_jsii_.InvokeVoid(
		d,
		"resetThresholdYellow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetWidth() {
	_jsii_.InvokeVoid(
		d,
		"resetWidth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

