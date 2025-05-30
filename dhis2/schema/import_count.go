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

// checks if the ImportCount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportCount{}

// ImportCount struct for ImportCount
type ImportCount struct {
	Deleted  int32 `json:"deleted"`
	Ignored  int32 `json:"ignored"`
	Imported int32 `json:"imported"`
	Updated  int32 `json:"updated"`
}

type _ImportCount ImportCount

// NewImportCount instantiates a new ImportCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportCount(deleted int32, ignored int32, imported int32, updated int32) *ImportCount {
	this := ImportCount{}
	this.Deleted = deleted
	this.Ignored = ignored
	this.Imported = imported
	this.Updated = updated
	return &this
}

// NewImportCountWithDefaults instantiates a new ImportCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportCountWithDefaults() *ImportCount {
	this := ImportCount{}
	return &this
}

// GetDeleted returns the Deleted field value
func (o *ImportCount) GetDeleted() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *ImportCount) GetDeletedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *ImportCount) SetDeleted(v int32) {
	o.Deleted = v
}

// GetIgnored returns the Ignored field value
func (o *ImportCount) GetIgnored() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ignored
}

// GetIgnoredOk returns a tuple with the Ignored field value
// and a boolean to check if the value has been set.
func (o *ImportCount) GetIgnoredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ignored, true
}

// SetIgnored sets field value
func (o *ImportCount) SetIgnored(v int32) {
	o.Ignored = v
}

// GetImported returns the Imported field value
func (o *ImportCount) GetImported() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Imported
}

// GetImportedOk returns a tuple with the Imported field value
// and a boolean to check if the value has been set.
func (o *ImportCount) GetImportedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Imported, true
}

// SetImported sets field value
func (o *ImportCount) SetImported(v int32) {
	o.Imported = v
}

// GetUpdated returns the Updated field value
func (o *ImportCount) GetUpdated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *ImportCount) GetUpdatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *ImportCount) SetUpdated(v int32) {
	o.Updated = v
}

func (o ImportCount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportCount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deleted"] = o.Deleted
	toSerialize["ignored"] = o.Ignored
	toSerialize["imported"] = o.Imported
	toSerialize["updated"] = o.Updated
	return toSerialize, nil
}

func (o *ImportCount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deleted",
		"ignored",
		"imported",
		"updated",
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

	varImportCount := _ImportCount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImportCount)

	if err != nil {
		return err
	}

	*o = ImportCount(varImportCount)

	return err
}

type NullableImportCount struct {
	value *ImportCount
	isSet bool
}

func (v NullableImportCount) Get() *ImportCount {
	return v.value
}

func (v *NullableImportCount) Set(val *ImportCount) {
	v.value = val
	v.isSet = true
}

func (v NullableImportCount) IsSet() bool {
	return v.isSet
}

func (v *NullableImportCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportCount(val *ImportCount) *NullableImportCount {
	return &NullableImportCount{value: val, isSet: true}
}

func (v NullableImportCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
