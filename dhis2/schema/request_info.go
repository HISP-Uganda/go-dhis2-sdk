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

// checks if the RequestInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestInfo{}

// RequestInfo struct for RequestInfo
type RequestInfo struct {
	HeaderXRequestID *string `json:"headerXRequestID,omitempty"`
}

// NewRequestInfo instantiates a new RequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestInfo() *RequestInfo {
	this := RequestInfo{}
	return &this
}

// NewRequestInfoWithDefaults instantiates a new RequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestInfoWithDefaults() *RequestInfo {
	this := RequestInfo{}
	return &this
}

// GetHeaderXRequestID returns the HeaderXRequestID field value if set, zero value otherwise.
func (o *RequestInfo) GetHeaderXRequestID() string {
	if o == nil || IsNil(o.HeaderXRequestID) {
		var ret string
		return ret
	}
	return *o.HeaderXRequestID
}

// GetHeaderXRequestIDOk returns a tuple with the HeaderXRequestID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestInfo) GetHeaderXRequestIDOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderXRequestID) {
		return nil, false
	}
	return o.HeaderXRequestID, true
}

// HasHeaderXRequestID returns a boolean if a field has been set.
func (o *RequestInfo) HasHeaderXRequestID() bool {
	if o != nil && !IsNil(o.HeaderXRequestID) {
		return true
	}

	return false
}

// SetHeaderXRequestID gets a reference to the given string and assigns it to the HeaderXRequestID field.
func (o *RequestInfo) SetHeaderXRequestID(v string) {
	o.HeaderXRequestID = &v
}

func (o RequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HeaderXRequestID) {
		toSerialize["headerXRequestID"] = o.HeaderXRequestID
	}
	return toSerialize, nil
}

type NullableRequestInfo struct {
	value *RequestInfo
	isSet bool
}

func (v NullableRequestInfo) Get() *RequestInfo {
	return v.value
}

func (v *NullableRequestInfo) Set(val *RequestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestInfo(val *RequestInfo) *NullableRequestInfo {
	return &NullableRequestInfo{value: val, isSet: true}
}

func (v NullableRequestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
