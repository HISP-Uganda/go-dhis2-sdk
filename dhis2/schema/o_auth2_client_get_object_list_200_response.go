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

// checks if the OAuth2ClientGetObjectList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ClientGetObjectList200Response{}

// OAuth2ClientGetObjectList200Response struct for OAuth2ClientGetObjectList200Response
type OAuth2ClientGetObjectList200Response struct {
	Pager         *Pager              `json:"pager,omitempty"`
	OAuth2Clients []Dhis2OAuth2Client `json:"oAuth2Clients,omitempty"`
}

// NewOAuth2ClientGetObjectList200Response instantiates a new OAuth2ClientGetObjectList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientGetObjectList200Response() *OAuth2ClientGetObjectList200Response {
	this := OAuth2ClientGetObjectList200Response{}
	return &this
}

// NewOAuth2ClientGetObjectList200ResponseWithDefaults instantiates a new OAuth2ClientGetObjectList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientGetObjectList200ResponseWithDefaults() *OAuth2ClientGetObjectList200Response {
	this := OAuth2ClientGetObjectList200Response{}
	return &this
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *OAuth2ClientGetObjectList200Response) GetPager() Pager {
	if o == nil || IsNil(o.Pager) {
		var ret Pager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientGetObjectList200Response) GetPagerOk() (*Pager, bool) {
	if o == nil || IsNil(o.Pager) {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *OAuth2ClientGetObjectList200Response) HasPager() bool {
	if o != nil && !IsNil(o.Pager) {
		return true
	}

	return false
}

// SetPager gets a reference to the given Pager and assigns it to the Pager field.
func (o *OAuth2ClientGetObjectList200Response) SetPager(v Pager) {
	o.Pager = &v
}

// GetOAuth2Clients returns the OAuth2Clients field value if set, zero value otherwise.
func (o *OAuth2ClientGetObjectList200Response) GetOAuth2Clients() []Dhis2OAuth2Client {
	if o == nil || IsNil(o.OAuth2Clients) {
		var ret []Dhis2OAuth2Client
		return ret
	}
	return o.OAuth2Clients
}

// GetOAuth2ClientsOk returns a tuple with the OAuth2Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientGetObjectList200Response) GetOAuth2ClientsOk() ([]Dhis2OAuth2Client, bool) {
	if o == nil || IsNil(o.OAuth2Clients) {
		return nil, false
	}
	return o.OAuth2Clients, true
}

// HasOAuth2Clients returns a boolean if a field has been set.
func (o *OAuth2ClientGetObjectList200Response) HasOAuth2Clients() bool {
	if o != nil && !IsNil(o.OAuth2Clients) {
		return true
	}

	return false
}

// SetOAuth2Clients gets a reference to the given []Dhis2OAuth2Client and assigns it to the OAuth2Clients field.
func (o *OAuth2ClientGetObjectList200Response) SetOAuth2Clients(v []Dhis2OAuth2Client) {
	o.OAuth2Clients = v
}

func (o OAuth2ClientGetObjectList200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ClientGetObjectList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pager) {
		toSerialize["pager"] = o.Pager
	}
	if !IsNil(o.OAuth2Clients) {
		toSerialize["oAuth2Clients"] = o.OAuth2Clients
	}
	return toSerialize, nil
}

type NullableOAuth2ClientGetObjectList200Response struct {
	value *OAuth2ClientGetObjectList200Response
	isSet bool
}

func (v NullableOAuth2ClientGetObjectList200Response) Get() *OAuth2ClientGetObjectList200Response {
	return v.value
}

func (v *NullableOAuth2ClientGetObjectList200Response) Set(val *OAuth2ClientGetObjectList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientGetObjectList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientGetObjectList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientGetObjectList200Response(val *OAuth2ClientGetObjectList200Response) *NullableOAuth2ClientGetObjectList200Response {
	return &NullableOAuth2ClientGetObjectList200Response{value: val, isSet: true}
}

func (v NullableOAuth2ClientGetObjectList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientGetObjectList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
