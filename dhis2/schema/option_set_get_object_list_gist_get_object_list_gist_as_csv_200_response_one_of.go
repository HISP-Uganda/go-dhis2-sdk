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

// checks if the OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}

// OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct for OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
type OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct {
	Entries    []map[string]interface{} `json:"entries,omitempty"`
	Pager      *GistPager               `json:"pager,omitempty"`
	OptionSets []OptionSet              `json:"optionSets,omitempty"`
}

// NewOptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf instantiates a new OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf() *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	this := OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}
	return &this
}

// NewOptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfWithDefaults instantiates a new OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfWithDefaults() *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	this := OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}
	return &this
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetEntries() []map[string]interface{} {
	if o == nil || IsNil(o.Entries) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetEntriesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasEntries() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []map[string]interface{} and assigns it to the Entries field.
func (o *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetEntries(v []map[string]interface{}) {
	o.Entries = v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetPager() GistPager {
	if o == nil || IsNil(o.Pager) {
		var ret GistPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetPagerOk() (*GistPager, bool) {
	if o == nil || IsNil(o.Pager) {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasPager() bool {
	if o != nil && !IsNil(o.Pager) {
		return true
	}

	return false
}

// SetPager gets a reference to the given GistPager and assigns it to the Pager field.
func (o *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetPager(v GistPager) {
	o.Pager = &v
}

// GetOptionSets returns the OptionSets field value if set, zero value otherwise.
func (o *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetOptionSets() []OptionSet {
	if o == nil || IsNil(o.OptionSets) {
		var ret []OptionSet
		return ret
	}
	return o.OptionSets
}

// GetOptionSetsOk returns a tuple with the OptionSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetOptionSetsOk() ([]OptionSet, bool) {
	if o == nil || IsNil(o.OptionSets) {
		return nil, false
	}
	return o.OptionSets, true
}

// HasOptionSets returns a boolean if a field has been set.
func (o *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasOptionSets() bool {
	if o != nil && !IsNil(o.OptionSets) {
		return true
	}

	return false
}

// SetOptionSets gets a reference to the given []OptionSet and assigns it to the OptionSets field.
func (o *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetOptionSets(v []OptionSet) {
	o.OptionSets = v
}

func (o OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entries) {
		toSerialize["entries"] = o.Entries
	}
	if !IsNil(o.Pager) {
		toSerialize["pager"] = o.Pager
	}
	if !IsNil(o.OptionSets) {
		toSerialize["optionSets"] = o.OptionSets
	}
	return toSerialize, nil
}

type NullableOptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct {
	value *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	isSet bool
}

func (v NullableOptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Get() *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	return v.value
}

func (v *NullableOptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Set(val *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf(val *OptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) *NullableOptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	return &NullableOptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{value: val, isSet: true}
}

func (v NullableOptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
