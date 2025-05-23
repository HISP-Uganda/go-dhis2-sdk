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

// checks if the TrackerTrigramIndexJobParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackerTrigramIndexJobParameters{}

// TrackerTrigramIndexJobParameters struct for TrackerTrigramIndexJobParameters
type TrackerTrigramIndexJobParameters struct {
	Attributes        []string `json:"attributes,omitempty"`
	SkipIndexDeletion *bool    `json:"skipIndexDeletion,omitempty"`
}

// NewTrackerTrigramIndexJobParameters instantiates a new TrackerTrigramIndexJobParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackerTrigramIndexJobParameters() *TrackerTrigramIndexJobParameters {
	this := TrackerTrigramIndexJobParameters{}
	return &this
}

// NewTrackerTrigramIndexJobParametersWithDefaults instantiates a new TrackerTrigramIndexJobParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackerTrigramIndexJobParametersWithDefaults() *TrackerTrigramIndexJobParameters {
	this := TrackerTrigramIndexJobParameters{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TrackerTrigramIndexJobParameters) GetAttributes() []string {
	if o == nil || IsNil(o.Attributes) {
		var ret []string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerTrigramIndexJobParameters) GetAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TrackerTrigramIndexJobParameters) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []string and assigns it to the Attributes field.
func (o *TrackerTrigramIndexJobParameters) SetAttributes(v []string) {
	o.Attributes = v
}

// GetSkipIndexDeletion returns the SkipIndexDeletion field value if set, zero value otherwise.
func (o *TrackerTrigramIndexJobParameters) GetSkipIndexDeletion() bool {
	if o == nil || IsNil(o.SkipIndexDeletion) {
		var ret bool
		return ret
	}
	return *o.SkipIndexDeletion
}

// GetSkipIndexDeletionOk returns a tuple with the SkipIndexDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerTrigramIndexJobParameters) GetSkipIndexDeletionOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipIndexDeletion) {
		return nil, false
	}
	return o.SkipIndexDeletion, true
}

// HasSkipIndexDeletion returns a boolean if a field has been set.
func (o *TrackerTrigramIndexJobParameters) HasSkipIndexDeletion() bool {
	if o != nil && !IsNil(o.SkipIndexDeletion) {
		return true
	}

	return false
}

// SetSkipIndexDeletion gets a reference to the given bool and assigns it to the SkipIndexDeletion field.
func (o *TrackerTrigramIndexJobParameters) SetSkipIndexDeletion(v bool) {
	o.SkipIndexDeletion = &v
}

func (o TrackerTrigramIndexJobParameters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrackerTrigramIndexJobParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.SkipIndexDeletion) {
		toSerialize["skipIndexDeletion"] = o.SkipIndexDeletion
	}
	return toSerialize, nil
}

type NullableTrackerTrigramIndexJobParameters struct {
	value *TrackerTrigramIndexJobParameters
	isSet bool
}

func (v NullableTrackerTrigramIndexJobParameters) Get() *TrackerTrigramIndexJobParameters {
	return v.value
}

func (v *NullableTrackerTrigramIndexJobParameters) Set(val *TrackerTrigramIndexJobParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackerTrigramIndexJobParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackerTrigramIndexJobParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackerTrigramIndexJobParameters(val *TrackerTrigramIndexJobParameters) *NullableTrackerTrigramIndexJobParameters {
	return &NullableTrackerTrigramIndexJobParameters{value: val, isSet: true}
}

func (v NullableTrackerTrigramIndexJobParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackerTrigramIndexJobParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
