/*
DHIS2 API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.42
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the DateFilterPeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateFilterPeriod{}

// DateFilterPeriod struct for DateFilterPeriod
type DateFilterPeriod struct {
	EndBuffer   int32              `json:"endBuffer"`
	EndDate     *time.Time         `json:"endDate,omitempty"`
	Period      RelativePeriodEnum `json:"period"`
	StartBuffer int32              `json:"startBuffer"`
	StartDate   *time.Time         `json:"startDate,omitempty"`
	Type        DatePeriodType     `json:"type"`
}

type _DateFilterPeriod DateFilterPeriod

// NewDateFilterPeriod instantiates a new DateFilterPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateFilterPeriod(endBuffer int32, period RelativePeriodEnum, startBuffer int32, type_ DatePeriodType) *DateFilterPeriod {
	this := DateFilterPeriod{}
	this.EndBuffer = endBuffer
	this.Period = period
	this.StartBuffer = startBuffer
	this.Type = type_
	return &this
}

// NewDateFilterPeriodWithDefaults instantiates a new DateFilterPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateFilterPeriodWithDefaults() *DateFilterPeriod {
	this := DateFilterPeriod{}
	return &this
}

// GetEndBuffer returns the EndBuffer field value
func (o *DateFilterPeriod) GetEndBuffer() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EndBuffer
}

// GetEndBufferOk returns a tuple with the EndBuffer field value
// and a boolean to check if the value has been set.
func (o *DateFilterPeriod) GetEndBufferOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndBuffer, true
}

// SetEndBuffer sets field value
func (o *DateFilterPeriod) SetEndBuffer(v int32) {
	o.EndBuffer = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *DateFilterPeriod) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateFilterPeriod) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *DateFilterPeriod) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *DateFilterPeriod) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetPeriod returns the Period field value
func (o *DateFilterPeriod) GetPeriod() RelativePeriodEnum {
	if o == nil {
		var ret RelativePeriodEnum
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *DateFilterPeriod) GetPeriodOk() (*RelativePeriodEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *DateFilterPeriod) SetPeriod(v RelativePeriodEnum) {
	o.Period = v
}

// GetStartBuffer returns the StartBuffer field value
func (o *DateFilterPeriod) GetStartBuffer() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartBuffer
}

// GetStartBufferOk returns a tuple with the StartBuffer field value
// and a boolean to check if the value has been set.
func (o *DateFilterPeriod) GetStartBufferOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartBuffer, true
}

// SetStartBuffer sets field value
func (o *DateFilterPeriod) SetStartBuffer(v int32) {
	o.StartBuffer = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *DateFilterPeriod) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateFilterPeriod) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *DateFilterPeriod) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *DateFilterPeriod) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetType returns the Type field value
func (o *DateFilterPeriod) GetType() DatePeriodType {
	if o == nil {
		var ret DatePeriodType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DateFilterPeriod) GetTypeOk() (*DatePeriodType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DateFilterPeriod) SetType(v DatePeriodType) {
	o.Type = v
}

func (o DateFilterPeriod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DateFilterPeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endBuffer"] = o.EndBuffer
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["period"] = o.Period
	toSerialize["startBuffer"] = o.StartBuffer
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *DateFilterPeriod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"endBuffer",
		"period",
		"startBuffer",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDateFilterPeriod := _DateFilterPeriod{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDateFilterPeriod)

	if err != nil {
		return err
	}

	*o = DateFilterPeriod(varDateFilterPeriod)

	return err
}

type NullableDateFilterPeriod struct {
	value *DateFilterPeriod
	isSet bool
}

func (v NullableDateFilterPeriod) Get() *DateFilterPeriod {
	return v.value
}

func (v *NullableDateFilterPeriod) Set(val *DateFilterPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableDateFilterPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableDateFilterPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateFilterPeriod(val *DateFilterPeriod) *NullableDateFilterPeriod {
	return &NullableDateFilterPeriod{value: val, isSet: true}
}

func (v NullableDateFilterPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateFilterPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
