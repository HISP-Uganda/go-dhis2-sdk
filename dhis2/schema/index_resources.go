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

// checks if the IndexResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexResources{}

// IndexResources struct for IndexResources
type IndexResources struct {
	Resources []IndexResource `json:"resources,omitempty"`
}

// NewIndexResources instantiates a new IndexResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexResources() *IndexResources {
	this := IndexResources{}
	return &this
}

// NewIndexResourcesWithDefaults instantiates a new IndexResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexResourcesWithDefaults() *IndexResources {
	this := IndexResources{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *IndexResources) GetResources() []IndexResource {
	if o == nil || IsNil(o.Resources) {
		var ret []IndexResource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexResources) GetResourcesOk() ([]IndexResource, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *IndexResources) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []IndexResource and assigns it to the Resources field.
func (o *IndexResources) SetResources(v []IndexResource) {
	o.Resources = v
}

func (o IndexResources) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return toSerialize, nil
}

type NullableIndexResources struct {
	value *IndexResources
	isSet bool
}

func (v NullableIndexResources) Get() *IndexResources {
	return v.value
}

func (v *NullableIndexResources) Set(val *IndexResources) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexResources) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexResources(val *IndexResources) *NullableIndexResources {
	return &NullableIndexResources{value: val, isSet: true}
}

func (v NullableIndexResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
