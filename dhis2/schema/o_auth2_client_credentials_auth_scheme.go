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

// checks if the OAuth2ClientCredentialsAuthScheme type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ClientCredentialsAuthScheme{}

// OAuth2ClientCredentialsAuthScheme struct for OAuth2ClientCredentialsAuthScheme
type OAuth2ClientCredentialsAuthScheme struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	TokenUri     string `json:"tokenUri"`
}

type _OAuth2ClientCredentialsAuthScheme OAuth2ClientCredentialsAuthScheme

// NewOAuth2ClientCredentialsAuthScheme instantiates a new OAuth2ClientCredentialsAuthScheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientCredentialsAuthScheme(clientId string, clientSecret string, tokenUri string) *OAuth2ClientCredentialsAuthScheme {
	this := OAuth2ClientCredentialsAuthScheme{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.TokenUri = tokenUri
	return &this
}

// NewOAuth2ClientCredentialsAuthSchemeWithDefaults instantiates a new OAuth2ClientCredentialsAuthScheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientCredentialsAuthSchemeWithDefaults() *OAuth2ClientCredentialsAuthScheme {
	this := OAuth2ClientCredentialsAuthScheme{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *OAuth2ClientCredentialsAuthScheme) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientCredentialsAuthScheme) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *OAuth2ClientCredentialsAuthScheme) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *OAuth2ClientCredentialsAuthScheme) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientCredentialsAuthScheme) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *OAuth2ClientCredentialsAuthScheme) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetTokenUri returns the TokenUri field value
func (o *OAuth2ClientCredentialsAuthScheme) GetTokenUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenUri
}

// GetTokenUriOk returns a tuple with the TokenUri field value
// and a boolean to check if the value has been set.
func (o *OAuth2ClientCredentialsAuthScheme) GetTokenUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenUri, true
}

// SetTokenUri sets field value
func (o *OAuth2ClientCredentialsAuthScheme) SetTokenUri(v string) {
	o.TokenUri = v
}

func (o OAuth2ClientCredentialsAuthScheme) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ClientCredentialsAuthScheme) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientId"] = o.ClientId
	toSerialize["clientSecret"] = o.ClientSecret
	toSerialize["tokenUri"] = o.TokenUri
	return toSerialize, nil
}

func (o *OAuth2ClientCredentialsAuthScheme) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientId",
		"clientSecret",
		"tokenUri",
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

	varOAuth2ClientCredentialsAuthScheme := _OAuth2ClientCredentialsAuthScheme{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOAuth2ClientCredentialsAuthScheme)

	if err != nil {
		return err
	}

	*o = OAuth2ClientCredentialsAuthScheme(varOAuth2ClientCredentialsAuthScheme)

	return err
}

type NullableOAuth2ClientCredentialsAuthScheme struct {
	value *OAuth2ClientCredentialsAuthScheme
	isSet bool
}

func (v NullableOAuth2ClientCredentialsAuthScheme) Get() *OAuth2ClientCredentialsAuthScheme {
	return v.value
}

func (v *NullableOAuth2ClientCredentialsAuthScheme) Set(val *OAuth2ClientCredentialsAuthScheme) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientCredentialsAuthScheme) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientCredentialsAuthScheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientCredentialsAuthScheme(val *OAuth2ClientCredentialsAuthScheme) *NullableOAuth2ClientCredentialsAuthScheme {
	return &NullableOAuth2ClientCredentialsAuthScheme{value: val, isSet: true}
}

func (v NullableOAuth2ClientCredentialsAuthScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientCredentialsAuthScheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
