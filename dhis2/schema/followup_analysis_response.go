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

// checks if the FollowupAnalysisResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowupAnalysisResponse{}

// FollowupAnalysisResponse struct for FollowupAnalysisResponse
type FollowupAnalysisResponse struct {
	FollowupValues []FollowupValue           `json:"followupValues,omitempty"`
	Metadata       *FollowupAnalysisMetadata `json:"metadata,omitempty"`
}

// NewFollowupAnalysisResponse instantiates a new FollowupAnalysisResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowupAnalysisResponse() *FollowupAnalysisResponse {
	this := FollowupAnalysisResponse{}
	return &this
}

// NewFollowupAnalysisResponseWithDefaults instantiates a new FollowupAnalysisResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowupAnalysisResponseWithDefaults() *FollowupAnalysisResponse {
	this := FollowupAnalysisResponse{}
	return &this
}

// GetFollowupValues returns the FollowupValues field value if set, zero value otherwise.
func (o *FollowupAnalysisResponse) GetFollowupValues() []FollowupValue {
	if o == nil || IsNil(o.FollowupValues) {
		var ret []FollowupValue
		return ret
	}
	return o.FollowupValues
}

// GetFollowupValuesOk returns a tuple with the FollowupValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FollowupAnalysisResponse) GetFollowupValuesOk() ([]FollowupValue, bool) {
	if o == nil || IsNil(o.FollowupValues) {
		return nil, false
	}
	return o.FollowupValues, true
}

// HasFollowupValues returns a boolean if a field has been set.
func (o *FollowupAnalysisResponse) HasFollowupValues() bool {
	if o != nil && !IsNil(o.FollowupValues) {
		return true
	}

	return false
}

// SetFollowupValues gets a reference to the given []FollowupValue and assigns it to the FollowupValues field.
func (o *FollowupAnalysisResponse) SetFollowupValues(v []FollowupValue) {
	o.FollowupValues = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *FollowupAnalysisResponse) GetMetadata() FollowupAnalysisMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret FollowupAnalysisMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FollowupAnalysisResponse) GetMetadataOk() (*FollowupAnalysisMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *FollowupAnalysisResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given FollowupAnalysisMetadata and assigns it to the Metadata field.
func (o *FollowupAnalysisResponse) SetMetadata(v FollowupAnalysisMetadata) {
	o.Metadata = &v
}

func (o FollowupAnalysisResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowupAnalysisResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FollowupValues) {
		toSerialize["followupValues"] = o.FollowupValues
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableFollowupAnalysisResponse struct {
	value *FollowupAnalysisResponse
	isSet bool
}

func (v NullableFollowupAnalysisResponse) Get() *FollowupAnalysisResponse {
	return v.value
}

func (v *NullableFollowupAnalysisResponse) Set(val *FollowupAnalysisResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowupAnalysisResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowupAnalysisResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowupAnalysisResponse(val *FollowupAnalysisResponse) *NullableFollowupAnalysisResponse {
	return &NullableFollowupAnalysisResponse{value: val, isSet: true}
}

func (v NullableFollowupAnalysisResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowupAnalysisResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
