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

// checks if the SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}

// SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct for SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
type SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct {
	Entries     []map[string]interface{} `json:"entries,omitempty"`
	Pager       *GistPager               `json:"pager,omitempty"`
	Smsoutbound []OutboundSms            `json:"smsoutbound,omitempty"`
}

// NewSmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf instantiates a new SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf() *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	this := SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}
	return &this
}

// NewSmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfWithDefaults instantiates a new SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfWithDefaults() *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	this := SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{}
	return &this
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetEntries() []map[string]interface{} {
	if o == nil || IsNil(o.Entries) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetEntriesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasEntries() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []map[string]interface{} and assigns it to the Entries field.
func (o *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetEntries(v []map[string]interface{}) {
	o.Entries = v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetPager() GistPager {
	if o == nil || IsNil(o.Pager) {
		var ret GistPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetPagerOk() (*GistPager, bool) {
	if o == nil || IsNil(o.Pager) {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasPager() bool {
	if o != nil && !IsNil(o.Pager) {
		return true
	}

	return false
}

// SetPager gets a reference to the given GistPager and assigns it to the Pager field.
func (o *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetPager(v GistPager) {
	o.Pager = &v
}

// GetSmsoutbound returns the Smsoutbound field value if set, zero value otherwise.
func (o *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetSmsoutbound() []OutboundSms {
	if o == nil || IsNil(o.Smsoutbound) {
		var ret []OutboundSms
		return ret
	}
	return o.Smsoutbound
}

// GetSmsoutboundOk returns a tuple with the Smsoutbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) GetSmsoutboundOk() ([]OutboundSms, bool) {
	if o == nil || IsNil(o.Smsoutbound) {
		return nil, false
	}
	return o.Smsoutbound, true
}

// HasSmsoutbound returns a boolean if a field has been set.
func (o *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) HasSmsoutbound() bool {
	if o != nil && !IsNil(o.Smsoutbound) {
		return true
	}

	return false
}

// SetSmsoutbound gets a reference to the given []OutboundSms and assigns it to the Smsoutbound field.
func (o *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) SetSmsoutbound(v []OutboundSms) {
	o.Smsoutbound = v
}

func (o SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entries) {
		toSerialize["entries"] = o.Entries
	}
	if !IsNil(o.Pager) {
		toSerialize["pager"] = o.Pager
	}
	if !IsNil(o.Smsoutbound) {
		toSerialize["smsoutbound"] = o.Smsoutbound
	}
	return toSerialize, nil
}

type NullableSmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf struct {
	value *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	isSet bool
}

func (v NullableSmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Get() *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	return v.value
}

func (v *NullableSmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Set(val *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf(val *SmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) *NullableSmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf {
	return &NullableSmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf{value: val, isSet: true}
}

func (v NullableSmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsOutboundGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
