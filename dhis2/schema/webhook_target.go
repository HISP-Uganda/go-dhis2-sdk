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

// checks if the WebhookTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookTarget{}

// WebhookTarget struct for WebhookTarget
type WebhookTarget struct {
	Auth        *RouteAuth         `json:"auth,omitempty"`
	ClientId    string             `json:"clientId"`
	ContentType string             `json:"contentType"`
	Headers     *map[string]string `json:"headers,omitempty"`
	Type        *string            `json:"type,omitempty"`
	Url         string             `json:"url"`
}

type _WebhookTarget WebhookTarget

// NewWebhookTarget instantiates a new WebhookTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookTarget(clientId string, contentType string, url string) *WebhookTarget {
	this := WebhookTarget{}
	this.ClientId = clientId
	this.ContentType = contentType
	this.Url = url
	return &this
}

// NewWebhookTargetWithDefaults instantiates a new WebhookTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookTargetWithDefaults() *WebhookTarget {
	this := WebhookTarget{}
	return &this
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *WebhookTarget) GetAuth() RouteAuth {
	if o == nil || IsNil(o.Auth) {
		var ret RouteAuth
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookTarget) GetAuthOk() (*RouteAuth, bool) {
	if o == nil || IsNil(o.Auth) {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *WebhookTarget) HasAuth() bool {
	if o != nil && !IsNil(o.Auth) {
		return true
	}

	return false
}

// SetAuth gets a reference to the given RouteAuth and assigns it to the Auth field.
func (o *WebhookTarget) SetAuth(v RouteAuth) {
	o.Auth = &v
}

// GetClientId returns the ClientId field value
func (o *WebhookTarget) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *WebhookTarget) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *WebhookTarget) SetClientId(v string) {
	o.ClientId = v
}

// GetContentType returns the ContentType field value
func (o *WebhookTarget) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *WebhookTarget) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *WebhookTarget) SetContentType(v string) {
	o.ContentType = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *WebhookTarget) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookTarget) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *WebhookTarget) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *WebhookTarget) SetHeaders(v map[string]string) {
	o.Headers = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WebhookTarget) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookTarget) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WebhookTarget) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WebhookTarget) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value
func (o *WebhookTarget) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebhookTarget) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WebhookTarget) SetUrl(v string) {
	o.Url = v
}

func (o WebhookTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Auth) {
		toSerialize["auth"] = o.Auth
	}
	toSerialize["clientId"] = o.ClientId
	toSerialize["contentType"] = o.ContentType
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *WebhookTarget) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientId",
		"contentType",
		"url",
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

	varWebhookTarget := _WebhookTarget{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookTarget)

	if err != nil {
		return err
	}

	*o = WebhookTarget(varWebhookTarget)

	return err
}

type NullableWebhookTarget struct {
	value *WebhookTarget
	isSet bool
}

func (v NullableWebhookTarget) Get() *WebhookTarget {
	return v.value
}

func (v *NullableWebhookTarget) Set(val *WebhookTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookTarget(val *WebhookTarget) *NullableWebhookTarget {
	return &NullableWebhookTarget{value: val, isSet: true}
}

func (v NullableWebhookTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
