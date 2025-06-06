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

// checks if the EventDataFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventDataFilter{}

// EventDataFilter struct for EventDataFilter
type EventDataFilter struct {
	DataItem   *string           `json:"dataItem,omitempty"`
	DateFilter *DateFilterPeriod `json:"dateFilter,omitempty"`
	Eq         *string           `json:"eq,omitempty"`
	Ge         *string           `json:"ge,omitempty"`
	Gt         *string           `json:"gt,omitempty"`
	In         []string          `json:"in,omitempty"`
	Le         *string           `json:"le,omitempty"`
	Like       *string           `json:"like,omitempty"`
	Lt         *string           `json:"lt,omitempty"`
	Null       *bool             `json:"null,omitempty"`
}

// NewEventDataFilter instantiates a new EventDataFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventDataFilter() *EventDataFilter {
	this := EventDataFilter{}
	return &this
}

// NewEventDataFilterWithDefaults instantiates a new EventDataFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventDataFilterWithDefaults() *EventDataFilter {
	this := EventDataFilter{}
	return &this
}

// GetDataItem returns the DataItem field value if set, zero value otherwise.
func (o *EventDataFilter) GetDataItem() string {
	if o == nil || IsNil(o.DataItem) {
		var ret string
		return ret
	}
	return *o.DataItem
}

// GetDataItemOk returns a tuple with the DataItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDataFilter) GetDataItemOk() (*string, bool) {
	if o == nil || IsNil(o.DataItem) {
		return nil, false
	}
	return o.DataItem, true
}

// HasDataItem returns a boolean if a field has been set.
func (o *EventDataFilter) HasDataItem() bool {
	if o != nil && !IsNil(o.DataItem) {
		return true
	}

	return false
}

// SetDataItem gets a reference to the given string and assigns it to the DataItem field.
func (o *EventDataFilter) SetDataItem(v string) {
	o.DataItem = &v
}

// GetDateFilter returns the DateFilter field value if set, zero value otherwise.
func (o *EventDataFilter) GetDateFilter() DateFilterPeriod {
	if o == nil || IsNil(o.DateFilter) {
		var ret DateFilterPeriod
		return ret
	}
	return *o.DateFilter
}

// GetDateFilterOk returns a tuple with the DateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDataFilter) GetDateFilterOk() (*DateFilterPeriod, bool) {
	if o == nil || IsNil(o.DateFilter) {
		return nil, false
	}
	return o.DateFilter, true
}

// HasDateFilter returns a boolean if a field has been set.
func (o *EventDataFilter) HasDateFilter() bool {
	if o != nil && !IsNil(o.DateFilter) {
		return true
	}

	return false
}

// SetDateFilter gets a reference to the given DateFilterPeriod and assigns it to the DateFilter field.
func (o *EventDataFilter) SetDateFilter(v DateFilterPeriod) {
	o.DateFilter = &v
}

// GetEq returns the Eq field value if set, zero value otherwise.
func (o *EventDataFilter) GetEq() string {
	if o == nil || IsNil(o.Eq) {
		var ret string
		return ret
	}
	return *o.Eq
}

// GetEqOk returns a tuple with the Eq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDataFilter) GetEqOk() (*string, bool) {
	if o == nil || IsNil(o.Eq) {
		return nil, false
	}
	return o.Eq, true
}

// HasEq returns a boolean if a field has been set.
func (o *EventDataFilter) HasEq() bool {
	if o != nil && !IsNil(o.Eq) {
		return true
	}

	return false
}

// SetEq gets a reference to the given string and assigns it to the Eq field.
func (o *EventDataFilter) SetEq(v string) {
	o.Eq = &v
}

// GetGe returns the Ge field value if set, zero value otherwise.
func (o *EventDataFilter) GetGe() string {
	if o == nil || IsNil(o.Ge) {
		var ret string
		return ret
	}
	return *o.Ge
}

// GetGeOk returns a tuple with the Ge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDataFilter) GetGeOk() (*string, bool) {
	if o == nil || IsNil(o.Ge) {
		return nil, false
	}
	return o.Ge, true
}

// HasGe returns a boolean if a field has been set.
func (o *EventDataFilter) HasGe() bool {
	if o != nil && !IsNil(o.Ge) {
		return true
	}

	return false
}

// SetGe gets a reference to the given string and assigns it to the Ge field.
func (o *EventDataFilter) SetGe(v string) {
	o.Ge = &v
}

// GetGt returns the Gt field value if set, zero value otherwise.
func (o *EventDataFilter) GetGt() string {
	if o == nil || IsNil(o.Gt) {
		var ret string
		return ret
	}
	return *o.Gt
}

// GetGtOk returns a tuple with the Gt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDataFilter) GetGtOk() (*string, bool) {
	if o == nil || IsNil(o.Gt) {
		return nil, false
	}
	return o.Gt, true
}

// HasGt returns a boolean if a field has been set.
func (o *EventDataFilter) HasGt() bool {
	if o != nil && !IsNil(o.Gt) {
		return true
	}

	return false
}

// SetGt gets a reference to the given string and assigns it to the Gt field.
func (o *EventDataFilter) SetGt(v string) {
	o.Gt = &v
}

// GetIn returns the In field value if set, zero value otherwise.
func (o *EventDataFilter) GetIn() []string {
	if o == nil || IsNil(o.In) {
		var ret []string
		return ret
	}
	return o.In
}

// GetInOk returns a tuple with the In field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDataFilter) GetInOk() ([]string, bool) {
	if o == nil || IsNil(o.In) {
		return nil, false
	}
	return o.In, true
}

// HasIn returns a boolean if a field has been set.
func (o *EventDataFilter) HasIn() bool {
	if o != nil && !IsNil(o.In) {
		return true
	}

	return false
}

// SetIn gets a reference to the given []string and assigns it to the In field.
func (o *EventDataFilter) SetIn(v []string) {
	o.In = v
}

// GetLe returns the Le field value if set, zero value otherwise.
func (o *EventDataFilter) GetLe() string {
	if o == nil || IsNil(o.Le) {
		var ret string
		return ret
	}
	return *o.Le
}

// GetLeOk returns a tuple with the Le field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDataFilter) GetLeOk() (*string, bool) {
	if o == nil || IsNil(o.Le) {
		return nil, false
	}
	return o.Le, true
}

// HasLe returns a boolean if a field has been set.
func (o *EventDataFilter) HasLe() bool {
	if o != nil && !IsNil(o.Le) {
		return true
	}

	return false
}

// SetLe gets a reference to the given string and assigns it to the Le field.
func (o *EventDataFilter) SetLe(v string) {
	o.Le = &v
}

// GetLike returns the Like field value if set, zero value otherwise.
func (o *EventDataFilter) GetLike() string {
	if o == nil || IsNil(o.Like) {
		var ret string
		return ret
	}
	return *o.Like
}

// GetLikeOk returns a tuple with the Like field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDataFilter) GetLikeOk() (*string, bool) {
	if o == nil || IsNil(o.Like) {
		return nil, false
	}
	return o.Like, true
}

// HasLike returns a boolean if a field has been set.
func (o *EventDataFilter) HasLike() bool {
	if o != nil && !IsNil(o.Like) {
		return true
	}

	return false
}

// SetLike gets a reference to the given string and assigns it to the Like field.
func (o *EventDataFilter) SetLike(v string) {
	o.Like = &v
}

// GetLt returns the Lt field value if set, zero value otherwise.
func (o *EventDataFilter) GetLt() string {
	if o == nil || IsNil(o.Lt) {
		var ret string
		return ret
	}
	return *o.Lt
}

// GetLtOk returns a tuple with the Lt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDataFilter) GetLtOk() (*string, bool) {
	if o == nil || IsNil(o.Lt) {
		return nil, false
	}
	return o.Lt, true
}

// HasLt returns a boolean if a field has been set.
func (o *EventDataFilter) HasLt() bool {
	if o != nil && !IsNil(o.Lt) {
		return true
	}

	return false
}

// SetLt gets a reference to the given string and assigns it to the Lt field.
func (o *EventDataFilter) SetLt(v string) {
	o.Lt = &v
}

// GetNull returns the Null field value if set, zero value otherwise.
func (o *EventDataFilter) GetNull() bool {
	if o == nil || IsNil(o.Null) {
		var ret bool
		return ret
	}
	return *o.Null
}

// GetNullOk returns a tuple with the Null field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventDataFilter) GetNullOk() (*bool, bool) {
	if o == nil || IsNil(o.Null) {
		return nil, false
	}
	return o.Null, true
}

// HasNull returns a boolean if a field has been set.
func (o *EventDataFilter) HasNull() bool {
	if o != nil && !IsNil(o.Null) {
		return true
	}

	return false
}

// SetNull gets a reference to the given bool and assigns it to the Null field.
func (o *EventDataFilter) SetNull(v bool) {
	o.Null = &v
}

func (o EventDataFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventDataFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataItem) {
		toSerialize["dataItem"] = o.DataItem
	}
	if !IsNil(o.DateFilter) {
		toSerialize["dateFilter"] = o.DateFilter
	}
	if !IsNil(o.Eq) {
		toSerialize["eq"] = o.Eq
	}
	if !IsNil(o.Ge) {
		toSerialize["ge"] = o.Ge
	}
	if !IsNil(o.Gt) {
		toSerialize["gt"] = o.Gt
	}
	if !IsNil(o.In) {
		toSerialize["in"] = o.In
	}
	if !IsNil(o.Le) {
		toSerialize["le"] = o.Le
	}
	if !IsNil(o.Like) {
		toSerialize["like"] = o.Like
	}
	if !IsNil(o.Lt) {
		toSerialize["lt"] = o.Lt
	}
	if !IsNil(o.Null) {
		toSerialize["null"] = o.Null
	}
	return toSerialize, nil
}

type NullableEventDataFilter struct {
	value *EventDataFilter
	isSet bool
}

func (v NullableEventDataFilter) Get() *EventDataFilter {
	return v.value
}

func (v *NullableEventDataFilter) Set(val *EventDataFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableEventDataFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableEventDataFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventDataFilter(val *EventDataFilter) *NullableEventDataFilter {
	return &NullableEventDataFilter{value: val, isSet: true}
}

func (v NullableEventDataFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventDataFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
