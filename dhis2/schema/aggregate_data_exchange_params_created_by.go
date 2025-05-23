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

// checks if the AggregateDataExchangeParamsCreatedBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregateDataExchangeParamsCreatedBy{}

// AggregateDataExchangeParamsCreatedBy struct for AggregateDataExchangeParamsCreatedBy
type AggregateDataExchangeParamsCreatedBy struct {
	// A UID for an User object   (Java name `org.hisp.dhis.user.User`)
	Id string `json:"id"`
}

type _AggregateDataExchangeParamsCreatedBy AggregateDataExchangeParamsCreatedBy

// NewAggregateDataExchangeParamsCreatedBy instantiates a new AggregateDataExchangeParamsCreatedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregateDataExchangeParamsCreatedBy(id string) *AggregateDataExchangeParamsCreatedBy {
	this := AggregateDataExchangeParamsCreatedBy{}
	this.Id = id
	return &this
}

// NewAggregateDataExchangeParamsCreatedByWithDefaults instantiates a new AggregateDataExchangeParamsCreatedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregateDataExchangeParamsCreatedByWithDefaults() *AggregateDataExchangeParamsCreatedBy {
	this := AggregateDataExchangeParamsCreatedBy{}
	return &this
}

// GetId returns the Id field value
func (o *AggregateDataExchangeParamsCreatedBy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AggregateDataExchangeParamsCreatedBy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AggregateDataExchangeParamsCreatedBy) SetId(v string) {
	o.Id = v
}

func (o AggregateDataExchangeParamsCreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregateDataExchangeParamsCreatedBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AggregateDataExchangeParamsCreatedBy) UnmarshalJSON(data []byte) (err error) {
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

	varAggregateDataExchangeParamsCreatedBy := _AggregateDataExchangeParamsCreatedBy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAggregateDataExchangeParamsCreatedBy)

	if err != nil {
		return err
	}

	*o = AggregateDataExchangeParamsCreatedBy(varAggregateDataExchangeParamsCreatedBy)

	return err
}

type NullableAggregateDataExchangeParamsCreatedBy struct {
	value *AggregateDataExchangeParamsCreatedBy
	isSet bool
}

func (v NullableAggregateDataExchangeParamsCreatedBy) Get() *AggregateDataExchangeParamsCreatedBy {
	return v.value
}

func (v *NullableAggregateDataExchangeParamsCreatedBy) Set(val *AggregateDataExchangeParamsCreatedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregateDataExchangeParamsCreatedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregateDataExchangeParamsCreatedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregateDataExchangeParamsCreatedBy(val *AggregateDataExchangeParamsCreatedBy) *NullableAggregateDataExchangeParamsCreatedBy {
	return &NullableAggregateDataExchangeParamsCreatedBy{value: val, isSet: true}
}

func (v NullableAggregateDataExchangeParamsCreatedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregateDataExchangeParamsCreatedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
