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

// checks if the AppSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppSettings{}

// AppSettings struct for AppSettings
type AppSettings struct {
	DashboardWidget *DashboardWidgetAppSettings `json:"dashboardWidget,omitempty"`
}

// NewAppSettings instantiates a new AppSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppSettings() *AppSettings {
	this := AppSettings{}
	return &this
}

// NewAppSettingsWithDefaults instantiates a new AppSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppSettingsWithDefaults() *AppSettings {
	this := AppSettings{}
	return &this
}

// GetDashboardWidget returns the DashboardWidget field value if set, zero value otherwise.
func (o *AppSettings) GetDashboardWidget() DashboardWidgetAppSettings {
	if o == nil || IsNil(o.DashboardWidget) {
		var ret DashboardWidgetAppSettings
		return ret
	}
	return *o.DashboardWidget
}

// GetDashboardWidgetOk returns a tuple with the DashboardWidget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSettings) GetDashboardWidgetOk() (*DashboardWidgetAppSettings, bool) {
	if o == nil || IsNil(o.DashboardWidget) {
		return nil, false
	}
	return o.DashboardWidget, true
}

// HasDashboardWidget returns a boolean if a field has been set.
func (o *AppSettings) HasDashboardWidget() bool {
	if o != nil && !IsNil(o.DashboardWidget) {
		return true
	}

	return false
}

// SetDashboardWidget gets a reference to the given DashboardWidgetAppSettings and assigns it to the DashboardWidget field.
func (o *AppSettings) SetDashboardWidget(v DashboardWidgetAppSettings) {
	o.DashboardWidget = &v
}

func (o AppSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DashboardWidget) {
		toSerialize["dashboardWidget"] = o.DashboardWidget
	}
	return toSerialize, nil
}

type NullableAppSettings struct {
	value *AppSettings
	isSet bool
}

func (v NullableAppSettings) Get() *AppSettings {
	return v.value
}

func (v *NullableAppSettings) Set(val *AppSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAppSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAppSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppSettings(val *AppSettings) *NullableAppSettings {
	return &NullableAppSettings{value: val, isSet: true}
}

func (v NullableAppSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
