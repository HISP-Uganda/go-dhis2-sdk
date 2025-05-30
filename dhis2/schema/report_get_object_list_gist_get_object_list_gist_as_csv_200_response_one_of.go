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

// checks if the ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}

// ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct for ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
type ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct {
	Entries []map[string]interface{} `json:"entries,omitempty"`
	Pager   *GistPager               `json:"pager,omitempty"`
	Reports []Report                 `json:"reports,omitempty"`
}

// NewReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf instantiates a new ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf() *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	this := ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}
	return &this
}

// NewReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfWithDefaults instantiates a new ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfWithDefaults() *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	this := ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}
	return &this
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetEntries() []map[string]interface{} {
	if o == nil || IsNil(o.Entries) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetEntriesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasEntries() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []map[string]interface{} and assigns it to the Entries field.
func (o *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetEntries(v []map[string]interface{}) {
	o.Entries = v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetPager() GistPager {
	if o == nil || IsNil(o.Pager) {
		var ret GistPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetPagerOk() (*GistPager, bool) {
	if o == nil || IsNil(o.Pager) {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasPager() bool {
	if o != nil && !IsNil(o.Pager) {
		return true
	}

	return false
}

// SetPager gets a reference to the given GistPager and assigns it to the Pager field.
func (o *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetPager(v GistPager) {
	o.Pager = &v
}

// GetReports returns the Reports field value if set, zero value otherwise.
func (o *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetReports() []Report {
	if o == nil || IsNil(o.Reports) {
		var ret []Report
		return ret
	}
	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetReportsOk() ([]Report, bool) {
	if o == nil || IsNil(o.Reports) {
		return nil, false
	}
	return o.Reports, true
}

// HasReports returns a boolean if a field has been set.
func (o *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasReports() bool {
	if o != nil && !IsNil(o.Reports) {
		return true
	}

	return false
}

// SetReports gets a reference to the given []Report and assigns it to the Reports field.
func (o *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetReports(v []Report) {
	o.Reports = v
}

func (o ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entries) {
		toSerialize["entries"] = o.Entries
	}
	if !IsNil(o.Pager) {
		toSerialize["pager"] = o.Pager
	}
	if !IsNil(o.Reports) {
		toSerialize["reports"] = o.Reports
	}
	return toSerialize, nil
}

type NullableReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct {
	value *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	isSet bool
}

func (v NullableReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Get() *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	return v.value
}

func (v *NullableReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Set(val *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf(val *ReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) *NullableReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	return &NullableReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{value: val, isSet: true}
}

func (v NullableReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
