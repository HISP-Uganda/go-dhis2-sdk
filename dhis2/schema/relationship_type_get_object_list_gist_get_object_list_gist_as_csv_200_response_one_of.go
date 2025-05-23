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

// checks if the RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}

// RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct for RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
type RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct {
	Entries           []map[string]interface{} `json:"entries,omitempty"`
	Pager             *GistPager               `json:"pager,omitempty"`
	RelationshipTypes []RelationshipType       `json:"relationshipTypes,omitempty"`
}

// NewRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf instantiates a new RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf() *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	this := RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}
	return &this
}

// NewRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfWithDefaults instantiates a new RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfWithDefaults() *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	this := RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}
	return &this
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetEntries() []map[string]interface{} {
	if o == nil || IsNil(o.Entries) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetEntriesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasEntries() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []map[string]interface{} and assigns it to the Entries field.
func (o *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetEntries(v []map[string]interface{}) {
	o.Entries = v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetPager() GistPager {
	if o == nil || IsNil(o.Pager) {
		var ret GistPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetPagerOk() (*GistPager, bool) {
	if o == nil || IsNil(o.Pager) {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasPager() bool {
	if o != nil && !IsNil(o.Pager) {
		return true
	}

	return false
}

// SetPager gets a reference to the given GistPager and assigns it to the Pager field.
func (o *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetPager(v GistPager) {
	o.Pager = &v
}

// GetRelationshipTypes returns the RelationshipTypes field value if set, zero value otherwise.
func (o *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetRelationshipTypes() []RelationshipType {
	if o == nil || IsNil(o.RelationshipTypes) {
		var ret []RelationshipType
		return ret
	}
	return o.RelationshipTypes
}

// GetRelationshipTypesOk returns a tuple with the RelationshipTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetRelationshipTypesOk() ([]RelationshipType, bool) {
	if o == nil || IsNil(o.RelationshipTypes) {
		return nil, false
	}
	return o.RelationshipTypes, true
}

// HasRelationshipTypes returns a boolean if a field has been set.
func (o *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasRelationshipTypes() bool {
	if o != nil && !IsNil(o.RelationshipTypes) {
		return true
	}

	return false
}

// SetRelationshipTypes gets a reference to the given []RelationshipType and assigns it to the RelationshipTypes field.
func (o *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetRelationshipTypes(v []RelationshipType) {
	o.RelationshipTypes = v
}

func (o RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entries) {
		toSerialize["entries"] = o.Entries
	}
	if !IsNil(o.Pager) {
		toSerialize["pager"] = o.Pager
	}
	if !IsNil(o.RelationshipTypes) {
		toSerialize["relationshipTypes"] = o.RelationshipTypes
	}
	return toSerialize, nil
}

type NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct {
	value *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	isSet bool
}

func (v NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Get() *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	return v.value
}

func (v *NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Set(val *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf(val *RelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) *NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	return &NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{value: val, isSet: true}
}

func (v NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipTypeGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
