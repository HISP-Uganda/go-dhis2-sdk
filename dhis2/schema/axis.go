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

// checks if the Axis type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Axis{}

// Axis struct for Axis
type Axis struct {
	Axis            *int32  `json:"axis,omitempty"`
	DimensionalItem *string `json:"dimensionalItem,omitempty"`
}

// NewAxis instantiates a new Axis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAxis() *Axis {
	this := Axis{}
	return &this
}

// NewAxisWithDefaults instantiates a new Axis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAxisWithDefaults() *Axis {
	this := Axis{}
	return &this
}

// GetAxis returns the Axis field value if set, zero value otherwise.
func (o *Axis) GetAxis() int32 {
	if o == nil || IsNil(o.Axis) {
		var ret int32
		return ret
	}
	return *o.Axis
}

// GetAxisOk returns a tuple with the Axis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Axis) GetAxisOk() (*int32, bool) {
	if o == nil || IsNil(o.Axis) {
		return nil, false
	}
	return o.Axis, true
}

// HasAxis returns a boolean if a field has been set.
func (o *Axis) HasAxis() bool {
	if o != nil && !IsNil(o.Axis) {
		return true
	}

	return false
}

// SetAxis gets a reference to the given int32 and assigns it to the Axis field.
func (o *Axis) SetAxis(v int32) {
	o.Axis = &v
}

// GetDimensionalItem returns the DimensionalItem field value if set, zero value otherwise.
func (o *Axis) GetDimensionalItem() string {
	if o == nil || IsNil(o.DimensionalItem) {
		var ret string
		return ret
	}
	return *o.DimensionalItem
}

// GetDimensionalItemOk returns a tuple with the DimensionalItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Axis) GetDimensionalItemOk() (*string, bool) {
	if o == nil || IsNil(o.DimensionalItem) {
		return nil, false
	}
	return o.DimensionalItem, true
}

// HasDimensionalItem returns a boolean if a field has been set.
func (o *Axis) HasDimensionalItem() bool {
	if o != nil && !IsNil(o.DimensionalItem) {
		return true
	}

	return false
}

// SetDimensionalItem gets a reference to the given string and assigns it to the DimensionalItem field.
func (o *Axis) SetDimensionalItem(v string) {
	o.DimensionalItem = &v
}

func (o Axis) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Axis) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Axis) {
		toSerialize["axis"] = o.Axis
	}
	if !IsNil(o.DimensionalItem) {
		toSerialize["dimensionalItem"] = o.DimensionalItem
	}
	return toSerialize, nil
}

type NullableAxis struct {
	value *Axis
	isSet bool
}

func (v NullableAxis) Get() *Axis {
	return v.value
}

func (v *NullableAxis) Set(val *Axis) {
	v.value = val
	v.isSet = true
}

func (v NullableAxis) IsSet() bool {
	return v.isSet
}

func (v *NullableAxis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAxis(val *Axis) *NullableAxis {
	return &NullableAxis{value: val, isSet: true}
}

func (v NullableAxis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAxis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
