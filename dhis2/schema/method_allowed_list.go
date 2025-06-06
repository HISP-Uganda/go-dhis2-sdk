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

// checks if the MethodAllowedList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MethodAllowedList{}

// MethodAllowedList struct for MethodAllowedList
type MethodAllowedList struct {
	AllowedMethods []string `json:"allowedMethods,omitempty"`
	Type           *string  `json:"type,omitempty"`
}

// NewMethodAllowedList instantiates a new MethodAllowedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMethodAllowedList() *MethodAllowedList {
	this := MethodAllowedList{}
	return &this
}

// NewMethodAllowedListWithDefaults instantiates a new MethodAllowedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMethodAllowedListWithDefaults() *MethodAllowedList {
	this := MethodAllowedList{}
	return &this
}

// GetAllowedMethods returns the AllowedMethods field value if set, zero value otherwise.
func (o *MethodAllowedList) GetAllowedMethods() []string {
	if o == nil || IsNil(o.AllowedMethods) {
		var ret []string
		return ret
	}
	return o.AllowedMethods
}

// GetAllowedMethodsOk returns a tuple with the AllowedMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MethodAllowedList) GetAllowedMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedMethods) {
		return nil, false
	}
	return o.AllowedMethods, true
}

// HasAllowedMethods returns a boolean if a field has been set.
func (o *MethodAllowedList) HasAllowedMethods() bool {
	if o != nil && !IsNil(o.AllowedMethods) {
		return true
	}

	return false
}

// SetAllowedMethods gets a reference to the given []string and assigns it to the AllowedMethods field.
func (o *MethodAllowedList) SetAllowedMethods(v []string) {
	o.AllowedMethods = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MethodAllowedList) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MethodAllowedList) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MethodAllowedList) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MethodAllowedList) SetType(v string) {
	o.Type = &v
}

func (o MethodAllowedList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MethodAllowedList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedMethods) {
		toSerialize["allowedMethods"] = o.AllowedMethods
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMethodAllowedList struct {
	value *MethodAllowedList
	isSet bool
}

func (v NullableMethodAllowedList) Get() *MethodAllowedList {
	return v.value
}

func (v *NullableMethodAllowedList) Set(val *MethodAllowedList) {
	v.value = val
	v.isSet = true
}

func (v NullableMethodAllowedList) IsSet() bool {
	return v.isSet
}

func (v *NullableMethodAllowedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMethodAllowedList(val *MethodAllowedList) *NullableMethodAllowedList {
	return &NullableMethodAllowedList{value: val, isSet: true}
}

func (v NullableMethodAllowedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMethodAllowedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
