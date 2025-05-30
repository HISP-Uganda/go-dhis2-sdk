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

// checks if the DimensionGetObjectList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DimensionGetObjectList200Response{}

// DimensionGetObjectList200Response struct for DimensionGetObjectList200Response
type DimensionGetObjectList200Response struct {
	Pager      *Pager                  `json:"pager,omitempty"`
	Dimensions []BaseDimensionalObject `json:"dimensions,omitempty"`
}

// NewDimensionGetObjectList200Response instantiates a new DimensionGetObjectList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionGetObjectList200Response() *DimensionGetObjectList200Response {
	this := DimensionGetObjectList200Response{}
	return &this
}

// NewDimensionGetObjectList200ResponseWithDefaults instantiates a new DimensionGetObjectList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionGetObjectList200ResponseWithDefaults() *DimensionGetObjectList200Response {
	this := DimensionGetObjectList200Response{}
	return &this
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *DimensionGetObjectList200Response) GetPager() Pager {
	if o == nil || IsNil(o.Pager) {
		var ret Pager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionGetObjectList200Response) GetPagerOk() (*Pager, bool) {
	if o == nil || IsNil(o.Pager) {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *DimensionGetObjectList200Response) HasPager() bool {
	if o != nil && !IsNil(o.Pager) {
		return true
	}

	return false
}

// SetPager gets a reference to the given Pager and assigns it to the Pager field.
func (o *DimensionGetObjectList200Response) SetPager(v Pager) {
	o.Pager = &v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *DimensionGetObjectList200Response) GetDimensions() []BaseDimensionalObject {
	if o == nil || IsNil(o.Dimensions) {
		var ret []BaseDimensionalObject
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionGetObjectList200Response) GetDimensionsOk() ([]BaseDimensionalObject, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *DimensionGetObjectList200Response) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given []BaseDimensionalObject and assigns it to the Dimensions field.
func (o *DimensionGetObjectList200Response) SetDimensions(v []BaseDimensionalObject) {
	o.Dimensions = v
}

func (o DimensionGetObjectList200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DimensionGetObjectList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pager) {
		toSerialize["pager"] = o.Pager
	}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	return toSerialize, nil
}

type NullableDimensionGetObjectList200Response struct {
	value *DimensionGetObjectList200Response
	isSet bool
}

func (v NullableDimensionGetObjectList200Response) Get() *DimensionGetObjectList200Response {
	return v.value
}

func (v *NullableDimensionGetObjectList200Response) Set(val *DimensionGetObjectList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDimensionGetObjectList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDimensionGetObjectList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDimensionGetObjectList200Response(val *DimensionGetObjectList200Response) *NullableDimensionGetObjectList200Response {
	return &NullableDimensionGetObjectList200Response{value: val, isSet: true}
}

func (v NullableDimensionGetObjectList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDimensionGetObjectList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
