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

// checks if the IndexResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexResource{}

// IndexResource struct for IndexResource
type IndexResource struct {
	DisplayName *string `json:"displayName,omitempty"`
	Href        *string `json:"href,omitempty"`
	Plural      *string `json:"plural,omitempty"`
	Singular    *string `json:"singular,omitempty"`
}

// NewIndexResource instantiates a new IndexResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexResource() *IndexResource {
	this := IndexResource{}
	return &this
}

// NewIndexResourceWithDefaults instantiates a new IndexResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexResourceWithDefaults() *IndexResource {
	this := IndexResource{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IndexResource) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexResource) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IndexResource) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IndexResource) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *IndexResource) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexResource) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *IndexResource) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *IndexResource) SetHref(v string) {
	o.Href = &v
}

// GetPlural returns the Plural field value if set, zero value otherwise.
func (o *IndexResource) GetPlural() string {
	if o == nil || IsNil(o.Plural) {
		var ret string
		return ret
	}
	return *o.Plural
}

// GetPluralOk returns a tuple with the Plural field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexResource) GetPluralOk() (*string, bool) {
	if o == nil || IsNil(o.Plural) {
		return nil, false
	}
	return o.Plural, true
}

// HasPlural returns a boolean if a field has been set.
func (o *IndexResource) HasPlural() bool {
	if o != nil && !IsNil(o.Plural) {
		return true
	}

	return false
}

// SetPlural gets a reference to the given string and assigns it to the Plural field.
func (o *IndexResource) SetPlural(v string) {
	o.Plural = &v
}

// GetSingular returns the Singular field value if set, zero value otherwise.
func (o *IndexResource) GetSingular() string {
	if o == nil || IsNil(o.Singular) {
		var ret string
		return ret
	}
	return *o.Singular
}

// GetSingularOk returns a tuple with the Singular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexResource) GetSingularOk() (*string, bool) {
	if o == nil || IsNil(o.Singular) {
		return nil, false
	}
	return o.Singular, true
}

// HasSingular returns a boolean if a field has been set.
func (o *IndexResource) HasSingular() bool {
	if o != nil && !IsNil(o.Singular) {
		return true
	}

	return false
}

// SetSingular gets a reference to the given string and assigns it to the Singular field.
func (o *IndexResource) SetSingular(v string) {
	o.Singular = &v
}

func (o IndexResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Plural) {
		toSerialize["plural"] = o.Plural
	}
	if !IsNil(o.Singular) {
		toSerialize["singular"] = o.Singular
	}
	return toSerialize, nil
}

type NullableIndexResource struct {
	value *IndexResource
	isSet bool
}

func (v NullableIndexResource) Get() *IndexResource {
	return v.value
}

func (v *NullableIndexResource) Set(val *IndexResource) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexResource) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexResource(val *IndexResource) *NullableIndexResource {
	return &NullableIndexResource{value: val, isSet: true}
}

func (v NullableIndexResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
