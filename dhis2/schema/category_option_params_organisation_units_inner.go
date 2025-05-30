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

// checks if the CategoryOptionParamsOrganisationUnitsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryOptionParamsOrganisationUnitsInner{}

// CategoryOptionParamsOrganisationUnitsInner struct for CategoryOptionParamsOrganisationUnitsInner
type CategoryOptionParamsOrganisationUnitsInner struct {
	// A UID for an OrganisationUnit object   (Java name `org.hisp.dhis.organisationunit.OrganisationUnit`)
	Id string `json:"id"`
}

type _CategoryOptionParamsOrganisationUnitsInner CategoryOptionParamsOrganisationUnitsInner

// NewCategoryOptionParamsOrganisationUnitsInner instantiates a new CategoryOptionParamsOrganisationUnitsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryOptionParamsOrganisationUnitsInner(id string) *CategoryOptionParamsOrganisationUnitsInner {
	this := CategoryOptionParamsOrganisationUnitsInner{}
	this.Id = id
	return &this
}

// NewCategoryOptionParamsOrganisationUnitsInnerWithDefaults instantiates a new CategoryOptionParamsOrganisationUnitsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryOptionParamsOrganisationUnitsInnerWithDefaults() *CategoryOptionParamsOrganisationUnitsInner {
	this := CategoryOptionParamsOrganisationUnitsInner{}
	return &this
}

// GetId returns the Id field value
func (o *CategoryOptionParamsOrganisationUnitsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CategoryOptionParamsOrganisationUnitsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CategoryOptionParamsOrganisationUnitsInner) SetId(v string) {
	o.Id = v
}

func (o CategoryOptionParamsOrganisationUnitsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryOptionParamsOrganisationUnitsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *CategoryOptionParamsOrganisationUnitsInner) UnmarshalJSON(data []byte) (err error) {
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

	varCategoryOptionParamsOrganisationUnitsInner := _CategoryOptionParamsOrganisationUnitsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCategoryOptionParamsOrganisationUnitsInner)

	if err != nil {
		return err
	}

	*o = CategoryOptionParamsOrganisationUnitsInner(varCategoryOptionParamsOrganisationUnitsInner)

	return err
}

type NullableCategoryOptionParamsOrganisationUnitsInner struct {
	value *CategoryOptionParamsOrganisationUnitsInner
	isSet bool
}

func (v NullableCategoryOptionParamsOrganisationUnitsInner) Get() *CategoryOptionParamsOrganisationUnitsInner {
	return v.value
}

func (v *NullableCategoryOptionParamsOrganisationUnitsInner) Set(val *CategoryOptionParamsOrganisationUnitsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryOptionParamsOrganisationUnitsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryOptionParamsOrganisationUnitsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryOptionParamsOrganisationUnitsInner(val *CategoryOptionParamsOrganisationUnitsInner) *NullableCategoryOptionParamsOrganisationUnitsInner {
	return &NullableCategoryOptionParamsOrganisationUnitsInner{value: val, isSet: true}
}

func (v NullableCategoryOptionParamsOrganisationUnitsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryOptionParamsOrganisationUnitsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
