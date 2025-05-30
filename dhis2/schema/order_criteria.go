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

// checks if the OrderCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCriteria{}

// OrderCriteria struct for OrderCriteria
type OrderCriteria struct {
	Direction SortDirection `json:"direction"`
	Field     *string       `json:"field,omitempty"`
}

type _OrderCriteria OrderCriteria

// NewOrderCriteria instantiates a new OrderCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCriteria(direction SortDirection) *OrderCriteria {
	this := OrderCriteria{}
	this.Direction = direction
	return &this
}

// NewOrderCriteriaWithDefaults instantiates a new OrderCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCriteriaWithDefaults() *OrderCriteria {
	this := OrderCriteria{}
	return &this
}

// GetDirection returns the Direction field value
func (o *OrderCriteria) GetDirection() SortDirection {
	if o == nil {
		var ret SortDirection
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *OrderCriteria) GetDirectionOk() (*SortDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *OrderCriteria) SetDirection(v SortDirection) {
	o.Direction = v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *OrderCriteria) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCriteria) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *OrderCriteria) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *OrderCriteria) SetField(v string) {
	o.Field = &v
}

func (o OrderCriteria) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["direction"] = o.Direction
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	return toSerialize, nil
}

func (o *OrderCriteria) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"direction",
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

	varOrderCriteria := _OrderCriteria{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderCriteria)

	if err != nil {
		return err
	}

	*o = OrderCriteria(varOrderCriteria)

	return err
}

type NullableOrderCriteria struct {
	value *OrderCriteria
	isSet bool
}

func (v NullableOrderCriteria) Get() *OrderCriteria {
	return v.value
}

func (v *NullableOrderCriteria) Set(val *OrderCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCriteria(val *OrderCriteria) *NullableOrderCriteria {
	return &NullableOrderCriteria{value: val, isSet: true}
}

func (v NullableOrderCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
