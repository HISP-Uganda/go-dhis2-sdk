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

// checks if the JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}

// JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct for JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
type JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct {
	Entries           []map[string]interface{} `json:"entries,omitempty"`
	Pager             *GistPager               `json:"pager,omitempty"`
	JobConfigurations []JobConfiguration       `json:"jobConfigurations,omitempty"`
}

// NewJobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf instantiates a new JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf() *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	this := JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}
	return &this
}

// NewJobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfWithDefaults instantiates a new JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfWithDefaults() *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	this := JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}
	return &this
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetEntries() []map[string]interface{} {
	if o == nil || IsNil(o.Entries) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetEntriesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasEntries() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []map[string]interface{} and assigns it to the Entries field.
func (o *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetEntries(v []map[string]interface{}) {
	o.Entries = v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetPager() GistPager {
	if o == nil || IsNil(o.Pager) {
		var ret GistPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetPagerOk() (*GistPager, bool) {
	if o == nil || IsNil(o.Pager) {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasPager() bool {
	if o != nil && !IsNil(o.Pager) {
		return true
	}

	return false
}

// SetPager gets a reference to the given GistPager and assigns it to the Pager field.
func (o *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetPager(v GistPager) {
	o.Pager = &v
}

// GetJobConfigurations returns the JobConfigurations field value if set, zero value otherwise.
func (o *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetJobConfigurations() []JobConfiguration {
	if o == nil || IsNil(o.JobConfigurations) {
		var ret []JobConfiguration
		return ret
	}
	return o.JobConfigurations
}

// GetJobConfigurationsOk returns a tuple with the JobConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetJobConfigurationsOk() ([]JobConfiguration, bool) {
	if o == nil || IsNil(o.JobConfigurations) {
		return nil, false
	}
	return o.JobConfigurations, true
}

// HasJobConfigurations returns a boolean if a field has been set.
func (o *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasJobConfigurations() bool {
	if o != nil && !IsNil(o.JobConfigurations) {
		return true
	}

	return false
}

// SetJobConfigurations gets a reference to the given []JobConfiguration and assigns it to the JobConfigurations field.
func (o *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetJobConfigurations(v []JobConfiguration) {
	o.JobConfigurations = v
}

func (o JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entries) {
		toSerialize["entries"] = o.Entries
	}
	if !IsNil(o.Pager) {
		toSerialize["pager"] = o.Pager
	}
	if !IsNil(o.JobConfigurations) {
		toSerialize["jobConfigurations"] = o.JobConfigurations
	}
	return toSerialize, nil
}

type NullableJobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct {
	value *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	isSet bool
}

func (v NullableJobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Get() *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	return v.value
}

func (v *NullableJobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Set(val *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableJobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableJobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf(val *JobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) *NullableJobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	return &NullableJobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{value: val, isSet: true}
}

func (v NullableJobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobConfigurationGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
