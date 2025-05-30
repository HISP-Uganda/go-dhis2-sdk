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

// checks if the GrantedAuthority type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantedAuthority{}

// GrantedAuthority struct for GrantedAuthority
type GrantedAuthority struct {
	Authority *string `json:"authority,omitempty"`
}

// NewGrantedAuthority instantiates a new GrantedAuthority object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantedAuthority() *GrantedAuthority {
	this := GrantedAuthority{}
	return &this
}

// NewGrantedAuthorityWithDefaults instantiates a new GrantedAuthority object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantedAuthorityWithDefaults() *GrantedAuthority {
	this := GrantedAuthority{}
	return &this
}

// GetAuthority returns the Authority field value if set, zero value otherwise.
func (o *GrantedAuthority) GetAuthority() string {
	if o == nil || IsNil(o.Authority) {
		var ret string
		return ret
	}
	return *o.Authority
}

// GetAuthorityOk returns a tuple with the Authority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantedAuthority) GetAuthorityOk() (*string, bool) {
	if o == nil || IsNil(o.Authority) {
		return nil, false
	}
	return o.Authority, true
}

// HasAuthority returns a boolean if a field has been set.
func (o *GrantedAuthority) HasAuthority() bool {
	if o != nil && !IsNil(o.Authority) {
		return true
	}

	return false
}

// SetAuthority gets a reference to the given string and assigns it to the Authority field.
func (o *GrantedAuthority) SetAuthority(v string) {
	o.Authority = &v
}

func (o GrantedAuthority) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantedAuthority) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authority) {
		toSerialize["authority"] = o.Authority
	}
	return toSerialize, nil
}

type NullableGrantedAuthority struct {
	value *GrantedAuthority
	isSet bool
}

func (v NullableGrantedAuthority) Get() *GrantedAuthority {
	return v.value
}

func (v *NullableGrantedAuthority) Set(val *GrantedAuthority) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantedAuthority) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantedAuthority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantedAuthority(val *GrantedAuthority) *NullableGrantedAuthority {
	return &NullableGrantedAuthority{value: val, isSet: true}
}

func (v NullableGrantedAuthority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantedAuthority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
