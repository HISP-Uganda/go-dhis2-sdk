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

// checks if the PredictorGroupGetObjectList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PredictorGroupGetObjectList200Response{}

// PredictorGroupGetObjectList200Response struct for PredictorGroupGetObjectList200Response
type PredictorGroupGetObjectList200Response struct {
	Pager           *Pager           `json:"pager,omitempty"`
	PredictorGroups []PredictorGroup `json:"predictorGroups,omitempty"`
}

// NewPredictorGroupGetObjectList200Response instantiates a new PredictorGroupGetObjectList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredictorGroupGetObjectList200Response() *PredictorGroupGetObjectList200Response {
	this := PredictorGroupGetObjectList200Response{}
	return &this
}

// NewPredictorGroupGetObjectList200ResponseWithDefaults instantiates a new PredictorGroupGetObjectList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredictorGroupGetObjectList200ResponseWithDefaults() *PredictorGroupGetObjectList200Response {
	this := PredictorGroupGetObjectList200Response{}
	return &this
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *PredictorGroupGetObjectList200Response) GetPager() Pager {
	if o == nil || IsNil(o.Pager) {
		var ret Pager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictorGroupGetObjectList200Response) GetPagerOk() (*Pager, bool) {
	if o == nil || IsNil(o.Pager) {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *PredictorGroupGetObjectList200Response) HasPager() bool {
	if o != nil && !IsNil(o.Pager) {
		return true
	}

	return false
}

// SetPager gets a reference to the given Pager and assigns it to the Pager field.
func (o *PredictorGroupGetObjectList200Response) SetPager(v Pager) {
	o.Pager = &v
}

// GetPredictorGroups returns the PredictorGroups field value if set, zero value otherwise.
func (o *PredictorGroupGetObjectList200Response) GetPredictorGroups() []PredictorGroup {
	if o == nil || IsNil(o.PredictorGroups) {
		var ret []PredictorGroup
		return ret
	}
	return o.PredictorGroups
}

// GetPredictorGroupsOk returns a tuple with the PredictorGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictorGroupGetObjectList200Response) GetPredictorGroupsOk() ([]PredictorGroup, bool) {
	if o == nil || IsNil(o.PredictorGroups) {
		return nil, false
	}
	return o.PredictorGroups, true
}

// HasPredictorGroups returns a boolean if a field has been set.
func (o *PredictorGroupGetObjectList200Response) HasPredictorGroups() bool {
	if o != nil && !IsNil(o.PredictorGroups) {
		return true
	}

	return false
}

// SetPredictorGroups gets a reference to the given []PredictorGroup and assigns it to the PredictorGroups field.
func (o *PredictorGroupGetObjectList200Response) SetPredictorGroups(v []PredictorGroup) {
	o.PredictorGroups = v
}

func (o PredictorGroupGetObjectList200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredictorGroupGetObjectList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pager) {
		toSerialize["pager"] = o.Pager
	}
	if !IsNil(o.PredictorGroups) {
		toSerialize["predictorGroups"] = o.PredictorGroups
	}
	return toSerialize, nil
}

type NullablePredictorGroupGetObjectList200Response struct {
	value *PredictorGroupGetObjectList200Response
	isSet bool
}

func (v NullablePredictorGroupGetObjectList200Response) Get() *PredictorGroupGetObjectList200Response {
	return v.value
}

func (v *NullablePredictorGroupGetObjectList200Response) Set(val *PredictorGroupGetObjectList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePredictorGroupGetObjectList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePredictorGroupGetObjectList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredictorGroupGetObjectList200Response(val *PredictorGroupGetObjectList200Response) *NullablePredictorGroupGetObjectList200Response {
	return &NullablePredictorGroupGetObjectList200Response{value: val, isSet: true}
}

func (v NullablePredictorGroupGetObjectList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredictorGroupGetObjectList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
