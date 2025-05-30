/*
DHIS2 API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.42
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
	"time"
)

// checks if the Dhis2Info type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dhis2Info{}

// Dhis2Info struct for Dhis2Info
type Dhis2Info struct {
	BuildTime  *time.Time `json:"buildTime,omitempty"`
	Revision   *string    `json:"revision,omitempty"`
	ServerDate *time.Time `json:"serverDate,omitempty"`
	SystemId   *string    `json:"systemId,omitempty"`
	Version    *string    `json:"version,omitempty"`
}

// NewDhis2Info instantiates a new Dhis2Info object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhis2Info() *Dhis2Info {
	this := Dhis2Info{}
	return &this
}

// NewDhis2InfoWithDefaults instantiates a new Dhis2Info object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhis2InfoWithDefaults() *Dhis2Info {
	this := Dhis2Info{}
	return &this
}

// GetBuildTime returns the BuildTime field value if set, zero value otherwise.
func (o *Dhis2Info) GetBuildTime() time.Time {
	if o == nil || IsNil(o.BuildTime) {
		var ret time.Time
		return ret
	}
	return *o.BuildTime
}

// GetBuildTimeOk returns a tuple with the BuildTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dhis2Info) GetBuildTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BuildTime) {
		return nil, false
	}
	return o.BuildTime, true
}

// HasBuildTime returns a boolean if a field has been set.
func (o *Dhis2Info) HasBuildTime() bool {
	if o != nil && !IsNil(o.BuildTime) {
		return true
	}

	return false
}

// SetBuildTime gets a reference to the given time.Time and assigns it to the BuildTime field.
func (o *Dhis2Info) SetBuildTime(v time.Time) {
	o.BuildTime = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *Dhis2Info) GetRevision() string {
	if o == nil || IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dhis2Info) GetRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *Dhis2Info) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *Dhis2Info) SetRevision(v string) {
	o.Revision = &v
}

// GetServerDate returns the ServerDate field value if set, zero value otherwise.
func (o *Dhis2Info) GetServerDate() time.Time {
	if o == nil || IsNil(o.ServerDate) {
		var ret time.Time
		return ret
	}
	return *o.ServerDate
}

// GetServerDateOk returns a tuple with the ServerDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dhis2Info) GetServerDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ServerDate) {
		return nil, false
	}
	return o.ServerDate, true
}

// HasServerDate returns a boolean if a field has been set.
func (o *Dhis2Info) HasServerDate() bool {
	if o != nil && !IsNil(o.ServerDate) {
		return true
	}

	return false
}

// SetServerDate gets a reference to the given time.Time and assigns it to the ServerDate field.
func (o *Dhis2Info) SetServerDate(v time.Time) {
	o.ServerDate = &v
}

// GetSystemId returns the SystemId field value if set, zero value otherwise.
func (o *Dhis2Info) GetSystemId() string {
	if o == nil || IsNil(o.SystemId) {
		var ret string
		return ret
	}
	return *o.SystemId
}

// GetSystemIdOk returns a tuple with the SystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dhis2Info) GetSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemId) {
		return nil, false
	}
	return o.SystemId, true
}

// HasSystemId returns a boolean if a field has been set.
func (o *Dhis2Info) HasSystemId() bool {
	if o != nil && !IsNil(o.SystemId) {
		return true
	}

	return false
}

// SetSystemId gets a reference to the given string and assigns it to the SystemId field.
func (o *Dhis2Info) SetSystemId(v string) {
	o.SystemId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Dhis2Info) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dhis2Info) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Dhis2Info) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Dhis2Info) SetVersion(v string) {
	o.Version = &v
}

func (o Dhis2Info) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dhis2Info) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuildTime) {
		toSerialize["buildTime"] = o.BuildTime
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.ServerDate) {
		toSerialize["serverDate"] = o.ServerDate
	}
	if !IsNil(o.SystemId) {
		toSerialize["systemId"] = o.SystemId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableDhis2Info struct {
	value *Dhis2Info
	isSet bool
}

func (v NullableDhis2Info) Get() *Dhis2Info {
	return v.value
}

func (v *NullableDhis2Info) Set(val *Dhis2Info) {
	v.value = val
	v.isSet = true
}

func (v NullableDhis2Info) IsSet() bool {
	return v.isSet
}

func (v *NullableDhis2Info) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhis2Info(val *Dhis2Info) *NullableDhis2Info {
	return &NullableDhis2Info{value: val, isSet: true}
}

func (v NullableDhis2Info) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhis2Info) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
