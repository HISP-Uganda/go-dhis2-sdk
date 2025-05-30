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

// checks if the OptionGetObjectList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionGetObjectList200Response{}

// OptionGetObjectList200Response struct for OptionGetObjectList200Response
type OptionGetObjectList200Response struct {
	Pager   *Pager   `json:"pager,omitempty"`
	Options []Option `json:"options,omitempty"`
}

// NewOptionGetObjectList200Response instantiates a new OptionGetObjectList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionGetObjectList200Response() *OptionGetObjectList200Response {
	this := OptionGetObjectList200Response{}
	return &this
}

// NewOptionGetObjectList200ResponseWithDefaults instantiates a new OptionGetObjectList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionGetObjectList200ResponseWithDefaults() *OptionGetObjectList200Response {
	this := OptionGetObjectList200Response{}
	return &this
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *OptionGetObjectList200Response) GetPager() Pager {
	if o == nil || IsNil(o.Pager) {
		var ret Pager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGetObjectList200Response) GetPagerOk() (*Pager, bool) {
	if o == nil || IsNil(o.Pager) {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *OptionGetObjectList200Response) HasPager() bool {
	if o != nil && !IsNil(o.Pager) {
		return true
	}

	return false
}

// SetPager gets a reference to the given Pager and assigns it to the Pager field.
func (o *OptionGetObjectList200Response) SetPager(v Pager) {
	o.Pager = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *OptionGetObjectList200Response) GetOptions() []Option {
	if o == nil || IsNil(o.Options) {
		var ret []Option
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGetObjectList200Response) GetOptionsOk() ([]Option, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *OptionGetObjectList200Response) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []Option and assigns it to the Options field.
func (o *OptionGetObjectList200Response) SetOptions(v []Option) {
	o.Options = v
}

func (o OptionGetObjectList200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionGetObjectList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pager) {
		toSerialize["pager"] = o.Pager
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableOptionGetObjectList200Response struct {
	value *OptionGetObjectList200Response
	isSet bool
}

func (v NullableOptionGetObjectList200Response) Get() *OptionGetObjectList200Response {
	return v.value
}

func (v *NullableOptionGetObjectList200Response) Set(val *OptionGetObjectList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionGetObjectList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionGetObjectList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionGetObjectList200Response(val *OptionGetObjectList200Response) *NullableOptionGetObjectList200Response {
	return &NullableOptionGetObjectList200Response{value: val, isSet: true}
}

func (v NullableOptionGetObjectList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionGetObjectList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
