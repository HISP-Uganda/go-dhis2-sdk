/*
DHIS2 API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.42
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the MetadataItemParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataItemParams{}

// MetadataItemParams struct for MetadataItemParams
type MetadataItemParams struct {
	AggregationType      AggregationType                  `json:"aggregationType"`
	Code                 *string                          `json:"code,omitempty"`
	Description          *string                          `json:"description,omitempty"`
	DimensionItemType    DimensionItemType                `json:"dimensionItemType"`
	DimensionType        DimensionType                    `json:"dimensionType"`
	EndDate              *time.Time                       `json:"endDate,omitempty"`
	Expression           *string                          `json:"expression,omitempty"`
	IndicatorType        *MetadataItemParamsIndicatorType `json:"indicatorType,omitempty"`
	LegendSet            *string                          `json:"legendSet,omitempty"`
	Name                 *string                          `json:"name,omitempty"`
	Options              []map[string]string              `json:"options,omitempty"`
	StartDate            *time.Time                       `json:"startDate,omitempty"`
	Style                *ObjectStyle                     `json:"style,omitempty"`
	TotalAggregationType TotalAggregationType             `json:"totalAggregationType"`
	Uid                  *string                          `json:"uid,omitempty"`
	ValueType            ValueType                        `json:"valueType"`
}

type _MetadataItemParams MetadataItemParams

// NewMetadataItemParams instantiates a new MetadataItemParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataItemParams(aggregationType AggregationType, dimensionItemType DimensionItemType, dimensionType DimensionType, totalAggregationType TotalAggregationType, valueType ValueType) *MetadataItemParams {
	this := MetadataItemParams{}
	this.AggregationType = aggregationType
	this.DimensionItemType = dimensionItemType
	this.DimensionType = dimensionType
	this.TotalAggregationType = totalAggregationType
	this.ValueType = valueType
	return &this
}

// NewMetadataItemParamsWithDefaults instantiates a new MetadataItemParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataItemParamsWithDefaults() *MetadataItemParams {
	this := MetadataItemParams{}
	return &this
}

// GetAggregationType returns the AggregationType field value
func (o *MetadataItemParams) GetAggregationType() AggregationType {
	if o == nil {
		var ret AggregationType
		return ret
	}

	return o.AggregationType
}

// GetAggregationTypeOk returns a tuple with the AggregationType field value
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetAggregationTypeOk() (*AggregationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AggregationType, true
}

// SetAggregationType sets field value
func (o *MetadataItemParams) SetAggregationType(v AggregationType) {
	o.AggregationType = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *MetadataItemParams) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *MetadataItemParams) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *MetadataItemParams) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MetadataItemParams) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MetadataItemParams) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MetadataItemParams) SetDescription(v string) {
	o.Description = &v
}

// GetDimensionItemType returns the DimensionItemType field value
func (o *MetadataItemParams) GetDimensionItemType() DimensionItemType {
	if o == nil {
		var ret DimensionItemType
		return ret
	}

	return o.DimensionItemType
}

// GetDimensionItemTypeOk returns a tuple with the DimensionItemType field value
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetDimensionItemTypeOk() (*DimensionItemType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DimensionItemType, true
}

// SetDimensionItemType sets field value
func (o *MetadataItemParams) SetDimensionItemType(v DimensionItemType) {
	o.DimensionItemType = v
}

// GetDimensionType returns the DimensionType field value
func (o *MetadataItemParams) GetDimensionType() DimensionType {
	if o == nil {
		var ret DimensionType
		return ret
	}

	return o.DimensionType
}

// GetDimensionTypeOk returns a tuple with the DimensionType field value
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetDimensionTypeOk() (*DimensionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DimensionType, true
}

// SetDimensionType sets field value
func (o *MetadataItemParams) SetDimensionType(v DimensionType) {
	o.DimensionType = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *MetadataItemParams) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *MetadataItemParams) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *MetadataItemParams) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *MetadataItemParams) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *MetadataItemParams) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *MetadataItemParams) SetExpression(v string) {
	o.Expression = &v
}

// GetIndicatorType returns the IndicatorType field value if set, zero value otherwise.
func (o *MetadataItemParams) GetIndicatorType() MetadataItemParamsIndicatorType {
	if o == nil || IsNil(o.IndicatorType) {
		var ret MetadataItemParamsIndicatorType
		return ret
	}
	return *o.IndicatorType
}

// GetIndicatorTypeOk returns a tuple with the IndicatorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetIndicatorTypeOk() (*MetadataItemParamsIndicatorType, bool) {
	if o == nil || IsNil(o.IndicatorType) {
		return nil, false
	}
	return o.IndicatorType, true
}

// HasIndicatorType returns a boolean if a field has been set.
func (o *MetadataItemParams) HasIndicatorType() bool {
	if o != nil && !IsNil(o.IndicatorType) {
		return true
	}

	return false
}

// SetIndicatorType gets a reference to the given MetadataItemParamsIndicatorType and assigns it to the IndicatorType field.
func (o *MetadataItemParams) SetIndicatorType(v MetadataItemParamsIndicatorType) {
	o.IndicatorType = &v
}

// GetLegendSet returns the LegendSet field value if set, zero value otherwise.
func (o *MetadataItemParams) GetLegendSet() string {
	if o == nil || IsNil(o.LegendSet) {
		var ret string
		return ret
	}
	return *o.LegendSet
}

// GetLegendSetOk returns a tuple with the LegendSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetLegendSetOk() (*string, bool) {
	if o == nil || IsNil(o.LegendSet) {
		return nil, false
	}
	return o.LegendSet, true
}

// HasLegendSet returns a boolean if a field has been set.
func (o *MetadataItemParams) HasLegendSet() bool {
	if o != nil && !IsNil(o.LegendSet) {
		return true
	}

	return false
}

// SetLegendSet gets a reference to the given string and assigns it to the LegendSet field.
func (o *MetadataItemParams) SetLegendSet(v string) {
	o.LegendSet = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetadataItemParams) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetadataItemParams) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MetadataItemParams) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *MetadataItemParams) GetOptions() []map[string]string {
	if o == nil || IsNil(o.Options) {
		var ret []map[string]string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetOptionsOk() ([]map[string]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *MetadataItemParams) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []map[string]string and assigns it to the Options field.
func (o *MetadataItemParams) SetOptions(v []map[string]string) {
	o.Options = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *MetadataItemParams) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *MetadataItemParams) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *MetadataItemParams) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *MetadataItemParams) GetStyle() ObjectStyle {
	if o == nil || IsNil(o.Style) {
		var ret ObjectStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetStyleOk() (*ObjectStyle, bool) {
	if o == nil || IsNil(o.Style) {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *MetadataItemParams) HasStyle() bool {
	if o != nil && !IsNil(o.Style) {
		return true
	}

	return false
}

// SetStyle gets a reference to the given ObjectStyle and assigns it to the Style field.
func (o *MetadataItemParams) SetStyle(v ObjectStyle) {
	o.Style = &v
}

// GetTotalAggregationType returns the TotalAggregationType field value
func (o *MetadataItemParams) GetTotalAggregationType() TotalAggregationType {
	if o == nil {
		var ret TotalAggregationType
		return ret
	}

	return o.TotalAggregationType
}

// GetTotalAggregationTypeOk returns a tuple with the TotalAggregationType field value
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetTotalAggregationTypeOk() (*TotalAggregationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAggregationType, true
}

// SetTotalAggregationType sets field value
func (o *MetadataItemParams) SetTotalAggregationType(v TotalAggregationType) {
	o.TotalAggregationType = v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *MetadataItemParams) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *MetadataItemParams) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *MetadataItemParams) SetUid(v string) {
	o.Uid = &v
}

// GetValueType returns the ValueType field value
func (o *MetadataItemParams) GetValueType() ValueType {
	if o == nil {
		var ret ValueType
		return ret
	}

	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *MetadataItemParams) GetValueTypeOk() (*ValueType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value
func (o *MetadataItemParams) SetValueType(v ValueType) {
	o.ValueType = v
}

func (o MetadataItemParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataItemParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aggregationType"] = o.AggregationType
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["dimensionItemType"] = o.DimensionItemType
	toSerialize["dimensionType"] = o.DimensionType
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.IndicatorType) {
		toSerialize["indicatorType"] = o.IndicatorType
	}
	if !IsNil(o.LegendSet) {
		toSerialize["legendSet"] = o.LegendSet
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.Style) {
		toSerialize["style"] = o.Style
	}
	toSerialize["totalAggregationType"] = o.TotalAggregationType
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	toSerialize["valueType"] = o.ValueType
	return toSerialize, nil
}

func (o *MetadataItemParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"aggregationType",
		"dimensionItemType",
		"dimensionType",
		"totalAggregationType",
		"valueType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMetadataItemParams := _MetadataItemParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMetadataItemParams)

	if err != nil {
		return err
	}

	*o = MetadataItemParams(varMetadataItemParams)

	return err
}

type NullableMetadataItemParams struct {
	value *MetadataItemParams
	isSet bool
}

func (v NullableMetadataItemParams) Get() *MetadataItemParams {
	return v.value
}

func (v *NullableMetadataItemParams) Set(val *MetadataItemParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataItemParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataItemParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataItemParams(val *MetadataItemParams) *NullableMetadataItemParams {
	return &NullableMetadataItemParams{value: val, isSet: true}
}

func (v NullableMetadataItemParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataItemParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
