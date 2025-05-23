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

// checks if the ValidationNotificationTemplateParamsValidationRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidationNotificationTemplateParamsValidationRulesInner{}

// ValidationNotificationTemplateParamsValidationRulesInner struct for ValidationNotificationTemplateParamsValidationRulesInner
type ValidationNotificationTemplateParamsValidationRulesInner struct {
	// A UID for an ValidationRule object   (Java name `org.hisp.dhis.validation.ValidationRule`)
	Id string `json:"id"`
}

type _ValidationNotificationTemplateParamsValidationRulesInner ValidationNotificationTemplateParamsValidationRulesInner

// NewValidationNotificationTemplateParamsValidationRulesInner instantiates a new ValidationNotificationTemplateParamsValidationRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationNotificationTemplateParamsValidationRulesInner(id string) *ValidationNotificationTemplateParamsValidationRulesInner {
	this := ValidationNotificationTemplateParamsValidationRulesInner{}
	this.Id = id
	return &this
}

// NewValidationNotificationTemplateParamsValidationRulesInnerWithDefaults instantiates a new ValidationNotificationTemplateParamsValidationRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationNotificationTemplateParamsValidationRulesInnerWithDefaults() *ValidationNotificationTemplateParamsValidationRulesInner {
	this := ValidationNotificationTemplateParamsValidationRulesInner{}
	return &this
}

// GetId returns the Id field value
func (o *ValidationNotificationTemplateParamsValidationRulesInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ValidationNotificationTemplateParamsValidationRulesInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ValidationNotificationTemplateParamsValidationRulesInner) SetId(v string) {
	o.Id = v
}

func (o ValidationNotificationTemplateParamsValidationRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidationNotificationTemplateParamsValidationRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ValidationNotificationTemplateParamsValidationRulesInner) UnmarshalJSON(data []byte) (err error) {
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

	varValidationNotificationTemplateParamsValidationRulesInner := _ValidationNotificationTemplateParamsValidationRulesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varValidationNotificationTemplateParamsValidationRulesInner)

	if err != nil {
		return err
	}

	*o = ValidationNotificationTemplateParamsValidationRulesInner(varValidationNotificationTemplateParamsValidationRulesInner)

	return err
}

type NullableValidationNotificationTemplateParamsValidationRulesInner struct {
	value *ValidationNotificationTemplateParamsValidationRulesInner
	isSet bool
}

func (v NullableValidationNotificationTemplateParamsValidationRulesInner) Get() *ValidationNotificationTemplateParamsValidationRulesInner {
	return v.value
}

func (v *NullableValidationNotificationTemplateParamsValidationRulesInner) Set(val *ValidationNotificationTemplateParamsValidationRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationNotificationTemplateParamsValidationRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationNotificationTemplateParamsValidationRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationNotificationTemplateParamsValidationRulesInner(val *ValidationNotificationTemplateParamsValidationRulesInner) *NullableValidationNotificationTemplateParamsValidationRulesInner {
	return &NullableValidationNotificationTemplateParamsValidationRulesInner{value: val, isSet: true}
}

func (v NullableValidationNotificationTemplateParamsValidationRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationNotificationTemplateParamsValidationRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
