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

// checks if the CategoryOptionGroupSetDimensionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryOptionGroupSetDimensionParams{}

// CategoryOptionGroupSetDimensionParams struct for CategoryOptionGroupSetDimensionParams
type CategoryOptionGroupSetDimensionParams struct {
	CategoryOptionGroupSet *CategoryOptionGroupSetParams                                    `json:"categoryOptionGroupSet,omitempty"`
	CategoryOptionGroups   []CategoryOptionGroupSetDimensionParamsCategoryOptionGroupsInner `json:"categoryOptionGroups,omitempty"`
}

// NewCategoryOptionGroupSetDimensionParams instantiates a new CategoryOptionGroupSetDimensionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryOptionGroupSetDimensionParams() *CategoryOptionGroupSetDimensionParams {
	this := CategoryOptionGroupSetDimensionParams{}
	return &this
}

// NewCategoryOptionGroupSetDimensionParamsWithDefaults instantiates a new CategoryOptionGroupSetDimensionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryOptionGroupSetDimensionParamsWithDefaults() *CategoryOptionGroupSetDimensionParams {
	this := CategoryOptionGroupSetDimensionParams{}
	return &this
}

// GetCategoryOptionGroupSet returns the CategoryOptionGroupSet field value if set, zero value otherwise.
func (o *CategoryOptionGroupSetDimensionParams) GetCategoryOptionGroupSet() CategoryOptionGroupSetParams {
	if o == nil || IsNil(o.CategoryOptionGroupSet) {
		var ret CategoryOptionGroupSetParams
		return ret
	}
	return *o.CategoryOptionGroupSet
}

// GetCategoryOptionGroupSetOk returns a tuple with the CategoryOptionGroupSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroupSetDimensionParams) GetCategoryOptionGroupSetOk() (*CategoryOptionGroupSetParams, bool) {
	if o == nil || IsNil(o.CategoryOptionGroupSet) {
		return nil, false
	}
	return o.CategoryOptionGroupSet, true
}

// HasCategoryOptionGroupSet returns a boolean if a field has been set.
func (o *CategoryOptionGroupSetDimensionParams) HasCategoryOptionGroupSet() bool {
	if o != nil && !IsNil(o.CategoryOptionGroupSet) {
		return true
	}

	return false
}

// SetCategoryOptionGroupSet gets a reference to the given CategoryOptionGroupSetParams and assigns it to the CategoryOptionGroupSet field.
func (o *CategoryOptionGroupSetDimensionParams) SetCategoryOptionGroupSet(v CategoryOptionGroupSetParams) {
	o.CategoryOptionGroupSet = &v
}

// GetCategoryOptionGroups returns the CategoryOptionGroups field value if set, zero value otherwise.
func (o *CategoryOptionGroupSetDimensionParams) GetCategoryOptionGroups() []CategoryOptionGroupSetDimensionParamsCategoryOptionGroupsInner {
	if o == nil || IsNil(o.CategoryOptionGroups) {
		var ret []CategoryOptionGroupSetDimensionParamsCategoryOptionGroupsInner
		return ret
	}
	return o.CategoryOptionGroups
}

// GetCategoryOptionGroupsOk returns a tuple with the CategoryOptionGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroupSetDimensionParams) GetCategoryOptionGroupsOk() ([]CategoryOptionGroupSetDimensionParamsCategoryOptionGroupsInner, bool) {
	if o == nil || IsNil(o.CategoryOptionGroups) {
		return nil, false
	}
	return o.CategoryOptionGroups, true
}

// HasCategoryOptionGroups returns a boolean if a field has been set.
func (o *CategoryOptionGroupSetDimensionParams) HasCategoryOptionGroups() bool {
	if o != nil && !IsNil(o.CategoryOptionGroups) {
		return true
	}

	return false
}

// SetCategoryOptionGroups gets a reference to the given []CategoryOptionGroupSetDimensionParamsCategoryOptionGroupsInner and assigns it to the CategoryOptionGroups field.
func (o *CategoryOptionGroupSetDimensionParams) SetCategoryOptionGroups(v []CategoryOptionGroupSetDimensionParamsCategoryOptionGroupsInner) {
	o.CategoryOptionGroups = v
}

func (o CategoryOptionGroupSetDimensionParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryOptionGroupSetDimensionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoryOptionGroupSet) {
		toSerialize["categoryOptionGroupSet"] = o.CategoryOptionGroupSet
	}
	if !IsNil(o.CategoryOptionGroups) {
		toSerialize["categoryOptionGroups"] = o.CategoryOptionGroups
	}
	return toSerialize, nil
}

type NullableCategoryOptionGroupSetDimensionParams struct {
	value *CategoryOptionGroupSetDimensionParams
	isSet bool
}

func (v NullableCategoryOptionGroupSetDimensionParams) Get() *CategoryOptionGroupSetDimensionParams {
	return v.value
}

func (v *NullableCategoryOptionGroupSetDimensionParams) Set(val *CategoryOptionGroupSetDimensionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryOptionGroupSetDimensionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryOptionGroupSetDimensionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryOptionGroupSetDimensionParams(val *CategoryOptionGroupSetDimensionParams) *NullableCategoryOptionGroupSetDimensionParams {
	return &NullableCategoryOptionGroupSetDimensionParams{value: val, isSet: true}
}

func (v NullableCategoryOptionGroupSetDimensionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryOptionGroupSetDimensionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
