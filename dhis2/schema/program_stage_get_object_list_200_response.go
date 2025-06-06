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

// checks if the ProgramStageGetObjectList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgramStageGetObjectList200Response{}

// ProgramStageGetObjectList200Response struct for ProgramStageGetObjectList200Response
type ProgramStageGetObjectList200Response struct {
	Pager         *Pager         `json:"pager,omitempty"`
	ProgramStages []ProgramStage `json:"programStages,omitempty"`
}

// NewProgramStageGetObjectList200Response instantiates a new ProgramStageGetObjectList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramStageGetObjectList200Response() *ProgramStageGetObjectList200Response {
	this := ProgramStageGetObjectList200Response{}
	return &this
}

// NewProgramStageGetObjectList200ResponseWithDefaults instantiates a new ProgramStageGetObjectList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramStageGetObjectList200ResponseWithDefaults() *ProgramStageGetObjectList200Response {
	this := ProgramStageGetObjectList200Response{}
	return &this
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *ProgramStageGetObjectList200Response) GetPager() Pager {
	if o == nil || IsNil(o.Pager) {
		var ret Pager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramStageGetObjectList200Response) GetPagerOk() (*Pager, bool) {
	if o == nil || IsNil(o.Pager) {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *ProgramStageGetObjectList200Response) HasPager() bool {
	if o != nil && !IsNil(o.Pager) {
		return true
	}

	return false
}

// SetPager gets a reference to the given Pager and assigns it to the Pager field.
func (o *ProgramStageGetObjectList200Response) SetPager(v Pager) {
	o.Pager = &v
}

// GetProgramStages returns the ProgramStages field value if set, zero value otherwise.
func (o *ProgramStageGetObjectList200Response) GetProgramStages() []ProgramStage {
	if o == nil || IsNil(o.ProgramStages) {
		var ret []ProgramStage
		return ret
	}
	return o.ProgramStages
}

// GetProgramStagesOk returns a tuple with the ProgramStages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramStageGetObjectList200Response) GetProgramStagesOk() ([]ProgramStage, bool) {
	if o == nil || IsNil(o.ProgramStages) {
		return nil, false
	}
	return o.ProgramStages, true
}

// HasProgramStages returns a boolean if a field has been set.
func (o *ProgramStageGetObjectList200Response) HasProgramStages() bool {
	if o != nil && !IsNil(o.ProgramStages) {
		return true
	}

	return false
}

// SetProgramStages gets a reference to the given []ProgramStage and assigns it to the ProgramStages field.
func (o *ProgramStageGetObjectList200Response) SetProgramStages(v []ProgramStage) {
	o.ProgramStages = v
}

func (o ProgramStageGetObjectList200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProgramStageGetObjectList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pager) {
		toSerialize["pager"] = o.Pager
	}
	if !IsNil(o.ProgramStages) {
		toSerialize["programStages"] = o.ProgramStages
	}
	return toSerialize, nil
}

type NullableProgramStageGetObjectList200Response struct {
	value *ProgramStageGetObjectList200Response
	isSet bool
}

func (v NullableProgramStageGetObjectList200Response) Get() *ProgramStageGetObjectList200Response {
	return v.value
}

func (v *NullableProgramStageGetObjectList200Response) Set(val *ProgramStageGetObjectList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramStageGetObjectList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramStageGetObjectList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramStageGetObjectList200Response(val *ProgramStageGetObjectList200Response) *NullableProgramStageGetObjectList200Response {
	return &NullableProgramStageGetObjectList200Response{value: val, isSet: true}
}

func (v NullableProgramStageGetObjectList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramStageGetObjectList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
