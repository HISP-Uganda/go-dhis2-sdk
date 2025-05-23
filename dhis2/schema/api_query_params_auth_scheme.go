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

// checks if the ApiQueryParamsAuthScheme type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiQueryParamsAuthScheme{}

// ApiQueryParamsAuthScheme struct for ApiQueryParamsAuthScheme
type ApiQueryParamsAuthScheme struct {
	QueryParams map[string]string `json:"queryParams"`
}

type _ApiQueryParamsAuthScheme ApiQueryParamsAuthScheme

// NewApiQueryParamsAuthScheme instantiates a new ApiQueryParamsAuthScheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiQueryParamsAuthScheme(queryParams map[string]string) *ApiQueryParamsAuthScheme {
	this := ApiQueryParamsAuthScheme{}
	this.QueryParams = queryParams
	return &this
}

// NewApiQueryParamsAuthSchemeWithDefaults instantiates a new ApiQueryParamsAuthScheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiQueryParamsAuthSchemeWithDefaults() *ApiQueryParamsAuthScheme {
	this := ApiQueryParamsAuthScheme{}
	return &this
}

// GetQueryParams returns the QueryParams field value
func (o *ApiQueryParamsAuthScheme) GetQueryParams() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.QueryParams
}

// GetQueryParamsOk returns a tuple with the QueryParams field value
// and a boolean to check if the value has been set.
func (o *ApiQueryParamsAuthScheme) GetQueryParamsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryParams, true
}

// SetQueryParams sets field value
func (o *ApiQueryParamsAuthScheme) SetQueryParams(v map[string]string) {
	o.QueryParams = v
}

func (o ApiQueryParamsAuthScheme) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiQueryParamsAuthScheme) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryParams"] = o.QueryParams
	return toSerialize, nil
}

func (o *ApiQueryParamsAuthScheme) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"queryParams",
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

	varApiQueryParamsAuthScheme := _ApiQueryParamsAuthScheme{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiQueryParamsAuthScheme)

	if err != nil {
		return err
	}

	*o = ApiQueryParamsAuthScheme(varApiQueryParamsAuthScheme)

	return err
}

type NullableApiQueryParamsAuthScheme struct {
	value *ApiQueryParamsAuthScheme
	isSet bool
}

func (v NullableApiQueryParamsAuthScheme) Get() *ApiQueryParamsAuthScheme {
	return v.value
}

func (v *NullableApiQueryParamsAuthScheme) Set(val *ApiQueryParamsAuthScheme) {
	v.value = val
	v.isSet = true
}

func (v NullableApiQueryParamsAuthScheme) IsSet() bool {
	return v.isSet
}

func (v *NullableApiQueryParamsAuthScheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiQueryParamsAuthScheme(val *ApiQueryParamsAuthScheme) *NullableApiQueryParamsAuthScheme {
	return &NullableApiQueryParamsAuthScheme{value: val, isSet: true}
}

func (v NullableApiQueryParamsAuthScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiQueryParamsAuthScheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
