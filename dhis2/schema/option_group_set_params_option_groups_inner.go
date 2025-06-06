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
)

// checks if the OptionGroupSetParamsOptionGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionGroupSetParamsOptionGroupsInner{}

// OptionGroupSetParamsOptionGroupsInner struct for OptionGroupSetParamsOptionGroupsInner
type OptionGroupSetParamsOptionGroupsInner struct {
	// A UID for an OptionGroup object   (Java name `org.hisp.dhis.option.OptionGroup`)
	Id string `json:"id"`
}

type _OptionGroupSetParamsOptionGroupsInner OptionGroupSetParamsOptionGroupsInner

// NewOptionGroupSetParamsOptionGroupsInner instantiates a new OptionGroupSetParamsOptionGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionGroupSetParamsOptionGroupsInner(id string) *OptionGroupSetParamsOptionGroupsInner {
	this := OptionGroupSetParamsOptionGroupsInner{}
	this.Id = id
	return &this
}

// NewOptionGroupSetParamsOptionGroupsInnerWithDefaults instantiates a new OptionGroupSetParamsOptionGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionGroupSetParamsOptionGroupsInnerWithDefaults() *OptionGroupSetParamsOptionGroupsInner {
	this := OptionGroupSetParamsOptionGroupsInner{}
	return &this
}

// GetId returns the Id field value
func (o *OptionGroupSetParamsOptionGroupsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OptionGroupSetParamsOptionGroupsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OptionGroupSetParamsOptionGroupsInner) SetId(v string) {
	o.Id = v
}

func (o OptionGroupSetParamsOptionGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionGroupSetParamsOptionGroupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *OptionGroupSetParamsOptionGroupsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varOptionGroupSetParamsOptionGroupsInner := _OptionGroupSetParamsOptionGroupsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOptionGroupSetParamsOptionGroupsInner)

	if err != nil {
		return err
	}

	*o = OptionGroupSetParamsOptionGroupsInner(varOptionGroupSetParamsOptionGroupsInner)

	return err
}

type NullableOptionGroupSetParamsOptionGroupsInner struct {
	value *OptionGroupSetParamsOptionGroupsInner
	isSet bool
}

func (v NullableOptionGroupSetParamsOptionGroupsInner) Get() *OptionGroupSetParamsOptionGroupsInner {
	return v.value
}

func (v *NullableOptionGroupSetParamsOptionGroupsInner) Set(val *OptionGroupSetParamsOptionGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionGroupSetParamsOptionGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionGroupSetParamsOptionGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionGroupSetParamsOptionGroupsInner(val *OptionGroupSetParamsOptionGroupsInner) *NullableOptionGroupSetParamsOptionGroupsInner {
	return &NullableOptionGroupSetParamsOptionGroupsInner{value: val, isSet: true}
}

func (v NullableOptionGroupSetParamsOptionGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionGroupSetParamsOptionGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
