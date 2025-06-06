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

// checks if the LockException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LockException{}

// LockException struct for LockException
type LockException struct {
	Created          *time.Time              `json:"created,omitempty"`
	DataSet          *BaseIdentifiableObject `json:"dataSet,omitempty"`
	Name             *string                 `json:"name,omitempty"`
	OrganisationUnit *BaseIdentifiableObject `json:"organisationUnit,omitempty"`
	Period           *string                 `json:"period,omitempty"`
}

// NewLockException instantiates a new LockException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLockException() *LockException {
	this := LockException{}
	return &this
}

// NewLockExceptionWithDefaults instantiates a new LockException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLockExceptionWithDefaults() *LockException {
	this := LockException{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *LockException) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockException) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *LockException) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *LockException) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDataSet returns the DataSet field value if set, zero value otherwise.
func (o *LockException) GetDataSet() BaseIdentifiableObject {
	if o == nil || IsNil(o.DataSet) {
		var ret BaseIdentifiableObject
		return ret
	}
	return *o.DataSet
}

// GetDataSetOk returns a tuple with the DataSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockException) GetDataSetOk() (*BaseIdentifiableObject, bool) {
	if o == nil || IsNil(o.DataSet) {
		return nil, false
	}
	return o.DataSet, true
}

// HasDataSet returns a boolean if a field has been set.
func (o *LockException) HasDataSet() bool {
	if o != nil && !IsNil(o.DataSet) {
		return true
	}

	return false
}

// SetDataSet gets a reference to the given BaseIdentifiableObject and assigns it to the DataSet field.
func (o *LockException) SetDataSet(v BaseIdentifiableObject) {
	o.DataSet = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LockException) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockException) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LockException) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LockException) SetName(v string) {
	o.Name = &v
}

// GetOrganisationUnit returns the OrganisationUnit field value if set, zero value otherwise.
func (o *LockException) GetOrganisationUnit() BaseIdentifiableObject {
	if o == nil || IsNil(o.OrganisationUnit) {
		var ret BaseIdentifiableObject
		return ret
	}
	return *o.OrganisationUnit
}

// GetOrganisationUnitOk returns a tuple with the OrganisationUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockException) GetOrganisationUnitOk() (*BaseIdentifiableObject, bool) {
	if o == nil || IsNil(o.OrganisationUnit) {
		return nil, false
	}
	return o.OrganisationUnit, true
}

// HasOrganisationUnit returns a boolean if a field has been set.
func (o *LockException) HasOrganisationUnit() bool {
	if o != nil && !IsNil(o.OrganisationUnit) {
		return true
	}

	return false
}

// SetOrganisationUnit gets a reference to the given BaseIdentifiableObject and assigns it to the OrganisationUnit field.
func (o *LockException) SetOrganisationUnit(v BaseIdentifiableObject) {
	o.OrganisationUnit = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *LockException) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockException) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *LockException) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *LockException) SetPeriod(v string) {
	o.Period = &v
}

func (o LockException) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LockException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.DataSet) {
		toSerialize["dataSet"] = o.DataSet
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrganisationUnit) {
		toSerialize["organisationUnit"] = o.OrganisationUnit
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	return toSerialize, nil
}

type NullableLockException struct {
	value *LockException
	isSet bool
}

func (v NullableLockException) Get() *LockException {
	return v.value
}

func (v *NullableLockException) Set(val *LockException) {
	v.value = val
	v.isSet = true
}

func (v NullableLockException) IsSet() bool {
	return v.isSet
}

func (v *NullableLockException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLockException(val *LockException) *NullableLockException {
	return &NullableLockException{value: val, isSet: true}
}

func (v NullableLockException) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLockException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
