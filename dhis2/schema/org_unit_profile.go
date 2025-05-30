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

// checks if the OrgUnitProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgUnitProfile{}

// OrgUnitProfile struct for OrgUnitProfile
type OrgUnitProfile struct {
	Attributes []string `json:"attributes,omitempty"`
	DataItems  []string `json:"dataItems,omitempty"`
	GroupSets  []string `json:"groupSets,omitempty"`
}

// NewOrgUnitProfile instantiates a new OrgUnitProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgUnitProfile() *OrgUnitProfile {
	this := OrgUnitProfile{}
	return &this
}

// NewOrgUnitProfileWithDefaults instantiates a new OrgUnitProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgUnitProfileWithDefaults() *OrgUnitProfile {
	this := OrgUnitProfile{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *OrgUnitProfile) GetAttributes() []string {
	if o == nil || IsNil(o.Attributes) {
		var ret []string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgUnitProfile) GetAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *OrgUnitProfile) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []string and assigns it to the Attributes field.
func (o *OrgUnitProfile) SetAttributes(v []string) {
	o.Attributes = v
}

// GetDataItems returns the DataItems field value if set, zero value otherwise.
func (o *OrgUnitProfile) GetDataItems() []string {
	if o == nil || IsNil(o.DataItems) {
		var ret []string
		return ret
	}
	return o.DataItems
}

// GetDataItemsOk returns a tuple with the DataItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgUnitProfile) GetDataItemsOk() ([]string, bool) {
	if o == nil || IsNil(o.DataItems) {
		return nil, false
	}
	return o.DataItems, true
}

// HasDataItems returns a boolean if a field has been set.
func (o *OrgUnitProfile) HasDataItems() bool {
	if o != nil && !IsNil(o.DataItems) {
		return true
	}

	return false
}

// SetDataItems gets a reference to the given []string and assigns it to the DataItems field.
func (o *OrgUnitProfile) SetDataItems(v []string) {
	o.DataItems = v
}

// GetGroupSets returns the GroupSets field value if set, zero value otherwise.
func (o *OrgUnitProfile) GetGroupSets() []string {
	if o == nil || IsNil(o.GroupSets) {
		var ret []string
		return ret
	}
	return o.GroupSets
}

// GetGroupSetsOk returns a tuple with the GroupSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgUnitProfile) GetGroupSetsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupSets) {
		return nil, false
	}
	return o.GroupSets, true
}

// HasGroupSets returns a boolean if a field has been set.
func (o *OrgUnitProfile) HasGroupSets() bool {
	if o != nil && !IsNil(o.GroupSets) {
		return true
	}

	return false
}

// SetGroupSets gets a reference to the given []string and assigns it to the GroupSets field.
func (o *OrgUnitProfile) SetGroupSets(v []string) {
	o.GroupSets = v
}

func (o OrgUnitProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgUnitProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.DataItems) {
		toSerialize["dataItems"] = o.DataItems
	}
	if !IsNil(o.GroupSets) {
		toSerialize["groupSets"] = o.GroupSets
	}
	return toSerialize, nil
}

type NullableOrgUnitProfile struct {
	value *OrgUnitProfile
	isSet bool
}

func (v NullableOrgUnitProfile) Get() *OrgUnitProfile {
	return v.value
}

func (v *NullableOrgUnitProfile) Set(val *OrgUnitProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgUnitProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgUnitProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgUnitProfile(val *OrgUnitProfile) *NullableOrgUnitProfile {
	return &NullableOrgUnitProfile{value: val, isSet: true}
}

func (v NullableOrgUnitProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgUnitProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
