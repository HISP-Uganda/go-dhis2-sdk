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

// checks if the CategoryDimensionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryDimensionParams{}

// CategoryDimensionParams struct for CategoryDimensionParams
type CategoryDimensionParams struct {
	Category        *CategoryParams                               `json:"category,omitempty"`
	CategoryOptions []CategoryDimensionParamsCategoryOptionsInner `json:"categoryOptions,omitempty"`
}

// NewCategoryDimensionParams instantiates a new CategoryDimensionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryDimensionParams() *CategoryDimensionParams {
	this := CategoryDimensionParams{}
	return &this
}

// NewCategoryDimensionParamsWithDefaults instantiates a new CategoryDimensionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryDimensionParamsWithDefaults() *CategoryDimensionParams {
	this := CategoryDimensionParams{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CategoryDimensionParams) GetCategory() CategoryParams {
	if o == nil || IsNil(o.Category) {
		var ret CategoryParams
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryDimensionParams) GetCategoryOk() (*CategoryParams, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CategoryDimensionParams) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given CategoryParams and assigns it to the Category field.
func (o *CategoryDimensionParams) SetCategory(v CategoryParams) {
	o.Category = &v
}

// GetCategoryOptions returns the CategoryOptions field value if set, zero value otherwise.
func (o *CategoryDimensionParams) GetCategoryOptions() []CategoryDimensionParamsCategoryOptionsInner {
	if o == nil || IsNil(o.CategoryOptions) {
		var ret []CategoryDimensionParamsCategoryOptionsInner
		return ret
	}
	return o.CategoryOptions
}

// GetCategoryOptionsOk returns a tuple with the CategoryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryDimensionParams) GetCategoryOptionsOk() ([]CategoryDimensionParamsCategoryOptionsInner, bool) {
	if o == nil || IsNil(o.CategoryOptions) {
		return nil, false
	}
	return o.CategoryOptions, true
}

// HasCategoryOptions returns a boolean if a field has been set.
func (o *CategoryDimensionParams) HasCategoryOptions() bool {
	if o != nil && !IsNil(o.CategoryOptions) {
		return true
	}

	return false
}

// SetCategoryOptions gets a reference to the given []CategoryDimensionParamsCategoryOptionsInner and assigns it to the CategoryOptions field.
func (o *CategoryDimensionParams) SetCategoryOptions(v []CategoryDimensionParamsCategoryOptionsInner) {
	o.CategoryOptions = v
}

func (o CategoryDimensionParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryDimensionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.CategoryOptions) {
		toSerialize["categoryOptions"] = o.CategoryOptions
	}
	return toSerialize, nil
}

type NullableCategoryDimensionParams struct {
	value *CategoryDimensionParams
	isSet bool
}

func (v NullableCategoryDimensionParams) Get() *CategoryDimensionParams {
	return v.value
}

func (v *NullableCategoryDimensionParams) Set(val *CategoryDimensionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryDimensionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryDimensionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryDimensionParams(val *CategoryDimensionParams) *NullableCategoryDimensionParams {
	return &NullableCategoryDimensionParams{value: val, isSet: true}
}

func (v NullableCategoryDimensionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryDimensionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
