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

// checks if the OptionGroupParamsOptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionGroupParamsOptionsInner{}

// OptionGroupParamsOptionsInner struct for OptionGroupParamsOptionsInner
type OptionGroupParamsOptionsInner struct {
	// A UID for an Option object   (Java name `org.hisp.dhis.option.Option`)
	Id string `json:"id"`
}

type _OptionGroupParamsOptionsInner OptionGroupParamsOptionsInner

// NewOptionGroupParamsOptionsInner instantiates a new OptionGroupParamsOptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionGroupParamsOptionsInner(id string) *OptionGroupParamsOptionsInner {
	this := OptionGroupParamsOptionsInner{}
	this.Id = id
	return &this
}

// NewOptionGroupParamsOptionsInnerWithDefaults instantiates a new OptionGroupParamsOptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionGroupParamsOptionsInnerWithDefaults() *OptionGroupParamsOptionsInner {
	this := OptionGroupParamsOptionsInner{}
	return &this
}

// GetId returns the Id field value
func (o *OptionGroupParamsOptionsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OptionGroupParamsOptionsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OptionGroupParamsOptionsInner) SetId(v string) {
	o.Id = v
}

func (o OptionGroupParamsOptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionGroupParamsOptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *OptionGroupParamsOptionsInner) UnmarshalJSON(data []byte) (err error) {
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

	varOptionGroupParamsOptionsInner := _OptionGroupParamsOptionsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOptionGroupParamsOptionsInner)

	if err != nil {
		return err
	}

	*o = OptionGroupParamsOptionsInner(varOptionGroupParamsOptionsInner)

	return err
}

type NullableOptionGroupParamsOptionsInner struct {
	value *OptionGroupParamsOptionsInner
	isSet bool
}

func (v NullableOptionGroupParamsOptionsInner) Get() *OptionGroupParamsOptionsInner {
	return v.value
}

func (v *NullableOptionGroupParamsOptionsInner) Set(val *OptionGroupParamsOptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionGroupParamsOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionGroupParamsOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionGroupParamsOptionsInner(val *OptionGroupParamsOptionsInner) *NullableOptionGroupParamsOptionsInner {
	return &NullableOptionGroupParamsOptionsInner{value: val, isSet: true}
}

func (v NullableOptionGroupParamsOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionGroupParamsOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
