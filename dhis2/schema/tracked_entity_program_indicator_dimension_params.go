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

// checks if the TrackedEntityProgramIndicatorDimensionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackedEntityProgramIndicatorDimensionParams{}

// TrackedEntityProgramIndicatorDimensionParams struct for TrackedEntityProgramIndicatorDimensionParams
type TrackedEntityProgramIndicatorDimensionParams struct {
	Filter           *string                 `json:"filter,omitempty"`
	LegendSet        *LegendSetParams        `json:"legendSet,omitempty"`
	ProgramIndicator *ProgramIndicatorParams `json:"programIndicator,omitempty"`
}

// NewTrackedEntityProgramIndicatorDimensionParams instantiates a new TrackedEntityProgramIndicatorDimensionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackedEntityProgramIndicatorDimensionParams() *TrackedEntityProgramIndicatorDimensionParams {
	this := TrackedEntityProgramIndicatorDimensionParams{}
	return &this
}

// NewTrackedEntityProgramIndicatorDimensionParamsWithDefaults instantiates a new TrackedEntityProgramIndicatorDimensionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackedEntityProgramIndicatorDimensionParamsWithDefaults() *TrackedEntityProgramIndicatorDimensionParams {
	this := TrackedEntityProgramIndicatorDimensionParams{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TrackedEntityProgramIndicatorDimensionParams) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackedEntityProgramIndicatorDimensionParams) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TrackedEntityProgramIndicatorDimensionParams) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *TrackedEntityProgramIndicatorDimensionParams) SetFilter(v string) {
	o.Filter = &v
}

// GetLegendSet returns the LegendSet field value if set, zero value otherwise.
func (o *TrackedEntityProgramIndicatorDimensionParams) GetLegendSet() LegendSetParams {
	if o == nil || IsNil(o.LegendSet) {
		var ret LegendSetParams
		return ret
	}
	return *o.LegendSet
}

// GetLegendSetOk returns a tuple with the LegendSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackedEntityProgramIndicatorDimensionParams) GetLegendSetOk() (*LegendSetParams, bool) {
	if o == nil || IsNil(o.LegendSet) {
		return nil, false
	}
	return o.LegendSet, true
}

// HasLegendSet returns a boolean if a field has been set.
func (o *TrackedEntityProgramIndicatorDimensionParams) HasLegendSet() bool {
	if o != nil && !IsNil(o.LegendSet) {
		return true
	}

	return false
}

// SetLegendSet gets a reference to the given LegendSetParams and assigns it to the LegendSet field.
func (o *TrackedEntityProgramIndicatorDimensionParams) SetLegendSet(v LegendSetParams) {
	o.LegendSet = &v
}

// GetProgramIndicator returns the ProgramIndicator field value if set, zero value otherwise.
func (o *TrackedEntityProgramIndicatorDimensionParams) GetProgramIndicator() ProgramIndicatorParams {
	if o == nil || IsNil(o.ProgramIndicator) {
		var ret ProgramIndicatorParams
		return ret
	}
	return *o.ProgramIndicator
}

// GetProgramIndicatorOk returns a tuple with the ProgramIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackedEntityProgramIndicatorDimensionParams) GetProgramIndicatorOk() (*ProgramIndicatorParams, bool) {
	if o == nil || IsNil(o.ProgramIndicator) {
		return nil, false
	}
	return o.ProgramIndicator, true
}

// HasProgramIndicator returns a boolean if a field has been set.
func (o *TrackedEntityProgramIndicatorDimensionParams) HasProgramIndicator() bool {
	if o != nil && !IsNil(o.ProgramIndicator) {
		return true
	}

	return false
}

// SetProgramIndicator gets a reference to the given ProgramIndicatorParams and assigns it to the ProgramIndicator field.
func (o *TrackedEntityProgramIndicatorDimensionParams) SetProgramIndicator(v ProgramIndicatorParams) {
	o.ProgramIndicator = &v
}

func (o TrackedEntityProgramIndicatorDimensionParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrackedEntityProgramIndicatorDimensionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.LegendSet) {
		toSerialize["legendSet"] = o.LegendSet
	}
	if !IsNil(o.ProgramIndicator) {
		toSerialize["programIndicator"] = o.ProgramIndicator
	}
	return toSerialize, nil
}

type NullableTrackedEntityProgramIndicatorDimensionParams struct {
	value *TrackedEntityProgramIndicatorDimensionParams
	isSet bool
}

func (v NullableTrackedEntityProgramIndicatorDimensionParams) Get() *TrackedEntityProgramIndicatorDimensionParams {
	return v.value
}

func (v *NullableTrackedEntityProgramIndicatorDimensionParams) Set(val *TrackedEntityProgramIndicatorDimensionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackedEntityProgramIndicatorDimensionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackedEntityProgramIndicatorDimensionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackedEntityProgramIndicatorDimensionParams(val *TrackedEntityProgramIndicatorDimensionParams) *NullableTrackedEntityProgramIndicatorDimensionParams {
	return &NullableTrackedEntityProgramIndicatorDimensionParams{value: val, isSet: true}
}

func (v NullableTrackedEntityProgramIndicatorDimensionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackedEntityProgramIndicatorDimensionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
