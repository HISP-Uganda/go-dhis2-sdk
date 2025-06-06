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

// checks if the FileResourceWebMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileResourceWebMessageResponse{}

// FileResourceWebMessageResponse struct for FileResourceWebMessageResponse
type FileResourceWebMessageResponse struct {
	FileResource *FileResource `json:"fileResource,omitempty"`
	ResponseType *string       `json:"responseType,omitempty"`
}

// NewFileResourceWebMessageResponse instantiates a new FileResourceWebMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileResourceWebMessageResponse() *FileResourceWebMessageResponse {
	this := FileResourceWebMessageResponse{}
	return &this
}

// NewFileResourceWebMessageResponseWithDefaults instantiates a new FileResourceWebMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileResourceWebMessageResponseWithDefaults() *FileResourceWebMessageResponse {
	this := FileResourceWebMessageResponse{}
	return &this
}

// GetFileResource returns the FileResource field value if set, zero value otherwise.
func (o *FileResourceWebMessageResponse) GetFileResource() FileResource {
	if o == nil || IsNil(o.FileResource) {
		var ret FileResource
		return ret
	}
	return *o.FileResource
}

// GetFileResourceOk returns a tuple with the FileResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileResourceWebMessageResponse) GetFileResourceOk() (*FileResource, bool) {
	if o == nil || IsNil(o.FileResource) {
		return nil, false
	}
	return o.FileResource, true
}

// HasFileResource returns a boolean if a field has been set.
func (o *FileResourceWebMessageResponse) HasFileResource() bool {
	if o != nil && !IsNil(o.FileResource) {
		return true
	}

	return false
}

// SetFileResource gets a reference to the given FileResource and assigns it to the FileResource field.
func (o *FileResourceWebMessageResponse) SetFileResource(v FileResource) {
	o.FileResource = &v
}

// GetResponseType returns the ResponseType field value if set, zero value otherwise.
func (o *FileResourceWebMessageResponse) GetResponseType() string {
	if o == nil || IsNil(o.ResponseType) {
		var ret string
		return ret
	}
	return *o.ResponseType
}

// GetResponseTypeOk returns a tuple with the ResponseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileResourceWebMessageResponse) GetResponseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseType) {
		return nil, false
	}
	return o.ResponseType, true
}

// HasResponseType returns a boolean if a field has been set.
func (o *FileResourceWebMessageResponse) HasResponseType() bool {
	if o != nil && !IsNil(o.ResponseType) {
		return true
	}

	return false
}

// SetResponseType gets a reference to the given string and assigns it to the ResponseType field.
func (o *FileResourceWebMessageResponse) SetResponseType(v string) {
	o.ResponseType = &v
}

func (o FileResourceWebMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileResourceWebMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileResource) {
		toSerialize["fileResource"] = o.FileResource
	}
	if !IsNil(o.ResponseType) {
		toSerialize["responseType"] = o.ResponseType
	}
	return toSerialize, nil
}

type NullableFileResourceWebMessageResponse struct {
	value *FileResourceWebMessageResponse
	isSet bool
}

func (v NullableFileResourceWebMessageResponse) Get() *FileResourceWebMessageResponse {
	return v.value
}

func (v *NullableFileResourceWebMessageResponse) Set(val *FileResourceWebMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileResourceWebMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileResourceWebMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileResourceWebMessageResponse(val *FileResourceWebMessageResponse) *NullableFileResourceWebMessageResponse {
	return &NullableFileResourceWebMessageResponse{value: val, isSet: true}
}

func (v NullableFileResourceWebMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileResourceWebMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
