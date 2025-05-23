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

// checks if the MergeParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MergeParams{}

// MergeParams struct for MergeParams
type MergeParams struct {
	DataMergeStrategy DataMergeStrategy `json:"dataMergeStrategy"`
	DeleteSources     *bool             `json:"deleteSources,omitempty"`
	Sources           []string          `json:"sources,omitempty"`
	Target            *string           `json:"target,omitempty"`
}

type _MergeParams MergeParams

// NewMergeParams instantiates a new MergeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeParams(dataMergeStrategy DataMergeStrategy) *MergeParams {
	this := MergeParams{}
	this.DataMergeStrategy = dataMergeStrategy
	return &this
}

// NewMergeParamsWithDefaults instantiates a new MergeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeParamsWithDefaults() *MergeParams {
	this := MergeParams{}
	return &this
}

// GetDataMergeStrategy returns the DataMergeStrategy field value
func (o *MergeParams) GetDataMergeStrategy() DataMergeStrategy {
	if o == nil {
		var ret DataMergeStrategy
		return ret
	}

	return o.DataMergeStrategy
}

// GetDataMergeStrategyOk returns a tuple with the DataMergeStrategy field value
// and a boolean to check if the value has been set.
func (o *MergeParams) GetDataMergeStrategyOk() (*DataMergeStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataMergeStrategy, true
}

// SetDataMergeStrategy sets field value
func (o *MergeParams) SetDataMergeStrategy(v DataMergeStrategy) {
	o.DataMergeStrategy = v
}

// GetDeleteSources returns the DeleteSources field value if set, zero value otherwise.
func (o *MergeParams) GetDeleteSources() bool {
	if o == nil || IsNil(o.DeleteSources) {
		var ret bool
		return ret
	}
	return *o.DeleteSources
}

// GetDeleteSourcesOk returns a tuple with the DeleteSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeParams) GetDeleteSourcesOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteSources) {
		return nil, false
	}
	return o.DeleteSources, true
}

// HasDeleteSources returns a boolean if a field has been set.
func (o *MergeParams) HasDeleteSources() bool {
	if o != nil && !IsNil(o.DeleteSources) {
		return true
	}

	return false
}

// SetDeleteSources gets a reference to the given bool and assigns it to the DeleteSources field.
func (o *MergeParams) SetDeleteSources(v bool) {
	o.DeleteSources = &v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *MergeParams) GetSources() []string {
	if o == nil || IsNil(o.Sources) {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeParams) GetSourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *MergeParams) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []string and assigns it to the Sources field.
func (o *MergeParams) SetSources(v []string) {
	o.Sources = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *MergeParams) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeParams) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *MergeParams) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *MergeParams) SetTarget(v string) {
	o.Target = &v
}

func (o MergeParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dataMergeStrategy"] = o.DataMergeStrategy
	if !IsNil(o.DeleteSources) {
		toSerialize["deleteSources"] = o.DeleteSources
	}
	if !IsNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

func (o *MergeParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dataMergeStrategy",
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

	varMergeParams := _MergeParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMergeParams)

	if err != nil {
		return err
	}

	*o = MergeParams(varMergeParams)

	return err
}

type NullableMergeParams struct {
	value *MergeParams
	isSet bool
}

func (v NullableMergeParams) Get() *MergeParams {
	return v.value
}

func (v *NullableMergeParams) Set(val *MergeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeParams(val *MergeParams) *NullableMergeParams {
	return &NullableMergeParams{value: val, isSet: true}
}

func (v NullableMergeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
