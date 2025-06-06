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

// checks if the DashboardGetObjectList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardGetObjectList200Response{}

// DashboardGetObjectList200Response struct for DashboardGetObjectList200Response
type DashboardGetObjectList200Response struct {
	Pager      *Pager      `json:"pager,omitempty"`
	Dashboards []Dashboard `json:"dashboards,omitempty"`
}

// NewDashboardGetObjectList200Response instantiates a new DashboardGetObjectList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardGetObjectList200Response() *DashboardGetObjectList200Response {
	this := DashboardGetObjectList200Response{}
	return &this
}

// NewDashboardGetObjectList200ResponseWithDefaults instantiates a new DashboardGetObjectList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardGetObjectList200ResponseWithDefaults() *DashboardGetObjectList200Response {
	this := DashboardGetObjectList200Response{}
	return &this
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *DashboardGetObjectList200Response) GetPager() Pager {
	if o == nil || IsNil(o.Pager) {
		var ret Pager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardGetObjectList200Response) GetPagerOk() (*Pager, bool) {
	if o == nil || IsNil(o.Pager) {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *DashboardGetObjectList200Response) HasPager() bool {
	if o != nil && !IsNil(o.Pager) {
		return true
	}

	return false
}

// SetPager gets a reference to the given Pager and assigns it to the Pager field.
func (o *DashboardGetObjectList200Response) SetPager(v Pager) {
	o.Pager = &v
}

// GetDashboards returns the Dashboards field value if set, zero value otherwise.
func (o *DashboardGetObjectList200Response) GetDashboards() []Dashboard {
	if o == nil || IsNil(o.Dashboards) {
		var ret []Dashboard
		return ret
	}
	return o.Dashboards
}

// GetDashboardsOk returns a tuple with the Dashboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardGetObjectList200Response) GetDashboardsOk() ([]Dashboard, bool) {
	if o == nil || IsNil(o.Dashboards) {
		return nil, false
	}
	return o.Dashboards, true
}

// HasDashboards returns a boolean if a field has been set.
func (o *DashboardGetObjectList200Response) HasDashboards() bool {
	if o != nil && !IsNil(o.Dashboards) {
		return true
	}

	return false
}

// SetDashboards gets a reference to the given []Dashboard and assigns it to the Dashboards field.
func (o *DashboardGetObjectList200Response) SetDashboards(v []Dashboard) {
	o.Dashboards = v
}

func (o DashboardGetObjectList200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardGetObjectList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pager) {
		toSerialize["pager"] = o.Pager
	}
	if !IsNil(o.Dashboards) {
		toSerialize["dashboards"] = o.Dashboards
	}
	return toSerialize, nil
}

type NullableDashboardGetObjectList200Response struct {
	value *DashboardGetObjectList200Response
	isSet bool
}

func (v NullableDashboardGetObjectList200Response) Get() *DashboardGetObjectList200Response {
	return v.value
}

func (v *NullableDashboardGetObjectList200Response) Set(val *DashboardGetObjectList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardGetObjectList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardGetObjectList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardGetObjectList200Response(val *DashboardGetObjectList200Response) *NullableDashboardGetObjectList200Response {
	return &NullableDashboardGetObjectList200Response{value: val, isSet: true}
}

func (v NullableDashboardGetObjectList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardGetObjectList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
