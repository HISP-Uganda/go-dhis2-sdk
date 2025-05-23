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

// checks if the CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}

// CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct for CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
type CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct {
	Entries                 []map[string]interface{} `json:"entries,omitempty"`
	Pager                   *GistPager               `json:"pager,omitempty"`
	CategoryOptionGroupSets []CategoryOptionGroupSet `json:"categoryOptionGroupSets,omitempty"`
}

// NewCategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf instantiates a new CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf() *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	this := CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}
	return &this
}

// NewCategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfWithDefaults instantiates a new CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfWithDefaults() *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	this := CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}
	return &this
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetEntries() []map[string]interface{} {
	if o == nil || IsNil(o.Entries) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetEntriesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasEntries() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []map[string]interface{} and assigns it to the Entries field.
func (o *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetEntries(v []map[string]interface{}) {
	o.Entries = v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetPager() GistPager {
	if o == nil || IsNil(o.Pager) {
		var ret GistPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetPagerOk() (*GistPager, bool) {
	if o == nil || IsNil(o.Pager) {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasPager() bool {
	if o != nil && !IsNil(o.Pager) {
		return true
	}

	return false
}

// SetPager gets a reference to the given GistPager and assigns it to the Pager field.
func (o *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetPager(v GistPager) {
	o.Pager = &v
}

// GetCategoryOptionGroupSets returns the CategoryOptionGroupSets field value if set, zero value otherwise.
func (o *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetCategoryOptionGroupSets() []CategoryOptionGroupSet {
	if o == nil || IsNil(o.CategoryOptionGroupSets) {
		var ret []CategoryOptionGroupSet
		return ret
	}
	return o.CategoryOptionGroupSets
}

// GetCategoryOptionGroupSetsOk returns a tuple with the CategoryOptionGroupSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetCategoryOptionGroupSetsOk() ([]CategoryOptionGroupSet, bool) {
	if o == nil || IsNil(o.CategoryOptionGroupSets) {
		return nil, false
	}
	return o.CategoryOptionGroupSets, true
}

// HasCategoryOptionGroupSets returns a boolean if a field has been set.
func (o *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasCategoryOptionGroupSets() bool {
	if o != nil && !IsNil(o.CategoryOptionGroupSets) {
		return true
	}

	return false
}

// SetCategoryOptionGroupSets gets a reference to the given []CategoryOptionGroupSet and assigns it to the CategoryOptionGroupSets field.
func (o *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetCategoryOptionGroupSets(v []CategoryOptionGroupSet) {
	o.CategoryOptionGroupSets = v
}

func (o CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entries) {
		toSerialize["entries"] = o.Entries
	}
	if !IsNil(o.Pager) {
		toSerialize["pager"] = o.Pager
	}
	if !IsNil(o.CategoryOptionGroupSets) {
		toSerialize["categoryOptionGroupSets"] = o.CategoryOptionGroupSets
	}
	return toSerialize, nil
}

type NullableCategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct {
	value *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	isSet bool
}

func (v NullableCategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Get() *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	return v.value
}

func (v *NullableCategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Set(val *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf(val *CategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) *NullableCategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	return &NullableCategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{value: val, isSet: true}
}

func (v NullableCategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryOptionGroupSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
