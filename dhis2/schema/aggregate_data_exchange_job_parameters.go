/*
DHIS2 API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.42
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// checks if the AggregateDataExchangeJobParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregateDataExchangeJobParameters{}

// AggregateDataExchangeJobParameters struct for AggregateDataExchangeJobParameters
type AggregateDataExchangeJobParameters struct {
	DataExchangeIds []string `json:"dataExchangeIds,omitempty"`
}

// NewAggregateDataExchangeJobParameters instantiates a new AggregateDataExchangeJobParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregateDataExchangeJobParameters() *AggregateDataExchangeJobParameters {
	this := AggregateDataExchangeJobParameters{}
	return &this
}

// NewAggregateDataExchangeJobParametersWithDefaults instantiates a new AggregateDataExchangeJobParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregateDataExchangeJobParametersWithDefaults() *AggregateDataExchangeJobParameters {
	this := AggregateDataExchangeJobParameters{}
	return &this
}

// GetDataExchangeIds returns the DataExchangeIds field value if set, zero value otherwise.
func (o *AggregateDataExchangeJobParameters) GetDataExchangeIds() []string {
	if o == nil || IsNil(o.DataExchangeIds) {
		var ret []string
		return ret
	}
	return o.DataExchangeIds
}

// GetDataExchangeIdsOk returns a tuple with the DataExchangeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateDataExchangeJobParameters) GetDataExchangeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.DataExchangeIds) {
		return nil, false
	}
	return o.DataExchangeIds, true
}

// HasDataExchangeIds returns a boolean if a field has been set.
func (o *AggregateDataExchangeJobParameters) HasDataExchangeIds() bool {
	if o != nil && !IsNil(o.DataExchangeIds) {
		return true
	}

	return false
}

// SetDataExchangeIds gets a reference to the given []string and assigns it to the DataExchangeIds field.
func (o *AggregateDataExchangeJobParameters) SetDataExchangeIds(v []string) {
	o.DataExchangeIds = v
}

func (o AggregateDataExchangeJobParameters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregateDataExchangeJobParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataExchangeIds) {
		toSerialize["dataExchangeIds"] = o.DataExchangeIds
	}
	return toSerialize, nil
}

type NullableAggregateDataExchangeJobParameters struct {
	value *AggregateDataExchangeJobParameters
	isSet bool
}

func (v NullableAggregateDataExchangeJobParameters) Get() *AggregateDataExchangeJobParameters {
	return v.value
}

func (v *NullableAggregateDataExchangeJobParameters) Set(val *AggregateDataExchangeJobParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregateDataExchangeJobParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregateDataExchangeJobParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregateDataExchangeJobParameters(val *AggregateDataExchangeJobParameters) *NullableAggregateDataExchangeJobParameters {
	return &NullableAggregateDataExchangeJobParameters{value: val, isSet: true}
}

func (v NullableAggregateDataExchangeJobParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregateDataExchangeJobParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
