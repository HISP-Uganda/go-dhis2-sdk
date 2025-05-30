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

// checks if the ExchangeTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeTarget{}

// ExchangeTarget struct for ExchangeTarget
type ExchangeTarget struct {
	Api     *Api           `json:"api,omitempty"`
	Request *TargetRequest `json:"request,omitempty"`
	Type    TargetType     `json:"type"`
}

type _ExchangeTarget ExchangeTarget

// NewExchangeTarget instantiates a new ExchangeTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeTarget(type_ TargetType) *ExchangeTarget {
	this := ExchangeTarget{}
	this.Type = type_
	return &this
}

// NewExchangeTargetWithDefaults instantiates a new ExchangeTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeTargetWithDefaults() *ExchangeTarget {
	this := ExchangeTarget{}
	return &this
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *ExchangeTarget) GetApi() Api {
	if o == nil || IsNil(o.Api) {
		var ret Api
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeTarget) GetApiOk() (*Api, bool) {
	if o == nil || IsNil(o.Api) {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *ExchangeTarget) HasApi() bool {
	if o != nil && !IsNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given Api and assigns it to the Api field.
func (o *ExchangeTarget) SetApi(v Api) {
	o.Api = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *ExchangeTarget) GetRequest() TargetRequest {
	if o == nil || IsNil(o.Request) {
		var ret TargetRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeTarget) GetRequestOk() (*TargetRequest, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *ExchangeTarget) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given TargetRequest and assigns it to the Request field.
func (o *ExchangeTarget) SetRequest(v TargetRequest) {
	o.Request = &v
}

// GetType returns the Type field value
func (o *ExchangeTarget) GetType() TargetType {
	if o == nil {
		var ret TargetType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExchangeTarget) GetTypeOk() (*TargetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExchangeTarget) SetType(v TargetType) {
	o.Type = v
}

func (o ExchangeTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ExchangeTarget) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varExchangeTarget := _ExchangeTarget{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExchangeTarget)

	if err != nil {
		return err
	}

	*o = ExchangeTarget(varExchangeTarget)

	return err
}

type NullableExchangeTarget struct {
	value *ExchangeTarget
	isSet bool
}

func (v NullableExchangeTarget) Get() *ExchangeTarget {
	return v.value
}

func (v *NullableExchangeTarget) Set(val *ExchangeTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeTarget(val *ExchangeTarget) *NullableExchangeTarget {
	return &NullableExchangeTarget{value: val, isSet: true}
}

func (v NullableExchangeTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
