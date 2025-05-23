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

// checks if the SchedulerEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulerEntry{}

// SchedulerEntry struct for SchedulerEntry
type SchedulerEntry struct {
	Configurable                     *bool               `json:"configurable,omitempty"`
	CronExpression                   *string             `json:"cronExpression,omitempty"`
	Delay                            *int32              `json:"delay,omitempty"`
	Enabled                          *bool               `json:"enabled,omitempty"`
	MaxDelayedExecutionTime          *time.Time          `json:"maxDelayedExecutionTime,omitempty"`
	Name                             *string             `json:"name,omitempty"`
	NextExecutionTime                *time.Time          `json:"nextExecutionTime,omitempty"`
	SecondsToMaxDelayedExecutionTime *int64              `json:"secondsToMaxDelayedExecutionTime,omitempty"`
	SecondsToNextExecutionTime       *int64              `json:"secondsToNextExecutionTime,omitempty"`
	Sequence                         []SchedulerEntryJob `json:"sequence,omitempty"`
	Status                           JobStatus           `json:"status"`
	Type                             *string             `json:"type,omitempty"`
}

type _SchedulerEntry SchedulerEntry

// NewSchedulerEntry instantiates a new SchedulerEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulerEntry(status JobStatus) *SchedulerEntry {
	this := SchedulerEntry{}
	this.Status = status
	return &this
}

// NewSchedulerEntryWithDefaults instantiates a new SchedulerEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulerEntryWithDefaults() *SchedulerEntry {
	this := SchedulerEntry{}
	return &this
}

// GetConfigurable returns the Configurable field value if set, zero value otherwise.
func (o *SchedulerEntry) GetConfigurable() bool {
	if o == nil || IsNil(o.Configurable) {
		var ret bool
		return ret
	}
	return *o.Configurable
}

// GetConfigurableOk returns a tuple with the Configurable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerEntry) GetConfigurableOk() (*bool, bool) {
	if o == nil || IsNil(o.Configurable) {
		return nil, false
	}
	return o.Configurable, true
}

// HasConfigurable returns a boolean if a field has been set.
func (o *SchedulerEntry) HasConfigurable() bool {
	if o != nil && !IsNil(o.Configurable) {
		return true
	}

	return false
}

// SetConfigurable gets a reference to the given bool and assigns it to the Configurable field.
func (o *SchedulerEntry) SetConfigurable(v bool) {
	o.Configurable = &v
}

// GetCronExpression returns the CronExpression field value if set, zero value otherwise.
func (o *SchedulerEntry) GetCronExpression() string {
	if o == nil || IsNil(o.CronExpression) {
		var ret string
		return ret
	}
	return *o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerEntry) GetCronExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.CronExpression) {
		return nil, false
	}
	return o.CronExpression, true
}

// HasCronExpression returns a boolean if a field has been set.
func (o *SchedulerEntry) HasCronExpression() bool {
	if o != nil && !IsNil(o.CronExpression) {
		return true
	}

	return false
}

// SetCronExpression gets a reference to the given string and assigns it to the CronExpression field.
func (o *SchedulerEntry) SetCronExpression(v string) {
	o.CronExpression = &v
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *SchedulerEntry) GetDelay() int32 {
	if o == nil || IsNil(o.Delay) {
		var ret int32
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerEntry) GetDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.Delay) {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *SchedulerEntry) HasDelay() bool {
	if o != nil && !IsNil(o.Delay) {
		return true
	}

	return false
}

// SetDelay gets a reference to the given int32 and assigns it to the Delay field.
func (o *SchedulerEntry) SetDelay(v int32) {
	o.Delay = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SchedulerEntry) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerEntry) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SchedulerEntry) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SchedulerEntry) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMaxDelayedExecutionTime returns the MaxDelayedExecutionTime field value if set, zero value otherwise.
func (o *SchedulerEntry) GetMaxDelayedExecutionTime() time.Time {
	if o == nil || IsNil(o.MaxDelayedExecutionTime) {
		var ret time.Time
		return ret
	}
	return *o.MaxDelayedExecutionTime
}

// GetMaxDelayedExecutionTimeOk returns a tuple with the MaxDelayedExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerEntry) GetMaxDelayedExecutionTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MaxDelayedExecutionTime) {
		return nil, false
	}
	return o.MaxDelayedExecutionTime, true
}

// HasMaxDelayedExecutionTime returns a boolean if a field has been set.
func (o *SchedulerEntry) HasMaxDelayedExecutionTime() bool {
	if o != nil && !IsNil(o.MaxDelayedExecutionTime) {
		return true
	}

	return false
}

// SetMaxDelayedExecutionTime gets a reference to the given time.Time and assigns it to the MaxDelayedExecutionTime field.
func (o *SchedulerEntry) SetMaxDelayedExecutionTime(v time.Time) {
	o.MaxDelayedExecutionTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SchedulerEntry) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerEntry) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SchedulerEntry) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SchedulerEntry) SetName(v string) {
	o.Name = &v
}

// GetNextExecutionTime returns the NextExecutionTime field value if set, zero value otherwise.
func (o *SchedulerEntry) GetNextExecutionTime() time.Time {
	if o == nil || IsNil(o.NextExecutionTime) {
		var ret time.Time
		return ret
	}
	return *o.NextExecutionTime
}

// GetNextExecutionTimeOk returns a tuple with the NextExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerEntry) GetNextExecutionTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NextExecutionTime) {
		return nil, false
	}
	return o.NextExecutionTime, true
}

// HasNextExecutionTime returns a boolean if a field has been set.
func (o *SchedulerEntry) HasNextExecutionTime() bool {
	if o != nil && !IsNil(o.NextExecutionTime) {
		return true
	}

	return false
}

// SetNextExecutionTime gets a reference to the given time.Time and assigns it to the NextExecutionTime field.
func (o *SchedulerEntry) SetNextExecutionTime(v time.Time) {
	o.NextExecutionTime = &v
}

// GetSecondsToMaxDelayedExecutionTime returns the SecondsToMaxDelayedExecutionTime field value if set, zero value otherwise.
func (o *SchedulerEntry) GetSecondsToMaxDelayedExecutionTime() int64 {
	if o == nil || IsNil(o.SecondsToMaxDelayedExecutionTime) {
		var ret int64
		return ret
	}
	return *o.SecondsToMaxDelayedExecutionTime
}

// GetSecondsToMaxDelayedExecutionTimeOk returns a tuple with the SecondsToMaxDelayedExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerEntry) GetSecondsToMaxDelayedExecutionTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.SecondsToMaxDelayedExecutionTime) {
		return nil, false
	}
	return o.SecondsToMaxDelayedExecutionTime, true
}

// HasSecondsToMaxDelayedExecutionTime returns a boolean if a field has been set.
func (o *SchedulerEntry) HasSecondsToMaxDelayedExecutionTime() bool {
	if o != nil && !IsNil(o.SecondsToMaxDelayedExecutionTime) {
		return true
	}

	return false
}

// SetSecondsToMaxDelayedExecutionTime gets a reference to the given int64 and assigns it to the SecondsToMaxDelayedExecutionTime field.
func (o *SchedulerEntry) SetSecondsToMaxDelayedExecutionTime(v int64) {
	o.SecondsToMaxDelayedExecutionTime = &v
}

// GetSecondsToNextExecutionTime returns the SecondsToNextExecutionTime field value if set, zero value otherwise.
func (o *SchedulerEntry) GetSecondsToNextExecutionTime() int64 {
	if o == nil || IsNil(o.SecondsToNextExecutionTime) {
		var ret int64
		return ret
	}
	return *o.SecondsToNextExecutionTime
}

// GetSecondsToNextExecutionTimeOk returns a tuple with the SecondsToNextExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerEntry) GetSecondsToNextExecutionTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.SecondsToNextExecutionTime) {
		return nil, false
	}
	return o.SecondsToNextExecutionTime, true
}

// HasSecondsToNextExecutionTime returns a boolean if a field has been set.
func (o *SchedulerEntry) HasSecondsToNextExecutionTime() bool {
	if o != nil && !IsNil(o.SecondsToNextExecutionTime) {
		return true
	}

	return false
}

// SetSecondsToNextExecutionTime gets a reference to the given int64 and assigns it to the SecondsToNextExecutionTime field.
func (o *SchedulerEntry) SetSecondsToNextExecutionTime(v int64) {
	o.SecondsToNextExecutionTime = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *SchedulerEntry) GetSequence() []SchedulerEntryJob {
	if o == nil || IsNil(o.Sequence) {
		var ret []SchedulerEntryJob
		return ret
	}
	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerEntry) GetSequenceOk() ([]SchedulerEntryJob, bool) {
	if o == nil || IsNil(o.Sequence) {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *SchedulerEntry) HasSequence() bool {
	if o != nil && !IsNil(o.Sequence) {
		return true
	}

	return false
}

// SetSequence gets a reference to the given []SchedulerEntryJob and assigns it to the Sequence field.
func (o *SchedulerEntry) SetSequence(v []SchedulerEntryJob) {
	o.Sequence = v
}

// GetStatus returns the Status field value
func (o *SchedulerEntry) GetStatus() JobStatus {
	if o == nil {
		var ret JobStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SchedulerEntry) GetStatusOk() (*JobStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SchedulerEntry) SetStatus(v JobStatus) {
	o.Status = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SchedulerEntry) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerEntry) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SchedulerEntry) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SchedulerEntry) SetType(v string) {
	o.Type = &v
}

func (o SchedulerEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulerEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configurable) {
		toSerialize["configurable"] = o.Configurable
	}
	if !IsNil(o.CronExpression) {
		toSerialize["cronExpression"] = o.CronExpression
	}
	if !IsNil(o.Delay) {
		toSerialize["delay"] = o.Delay
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.MaxDelayedExecutionTime) {
		toSerialize["maxDelayedExecutionTime"] = o.MaxDelayedExecutionTime
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NextExecutionTime) {
		toSerialize["nextExecutionTime"] = o.NextExecutionTime
	}
	if !IsNil(o.SecondsToMaxDelayedExecutionTime) {
		toSerialize["secondsToMaxDelayedExecutionTime"] = o.SecondsToMaxDelayedExecutionTime
	}
	if !IsNil(o.SecondsToNextExecutionTime) {
		toSerialize["secondsToNextExecutionTime"] = o.SecondsToNextExecutionTime
	}
	if !IsNil(o.Sequence) {
		toSerialize["sequence"] = o.Sequence
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *SchedulerEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varSchedulerEntry := _SchedulerEntry{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchedulerEntry)

	if err != nil {
		return err
	}

	*o = SchedulerEntry(varSchedulerEntry)

	return err
}

type NullableSchedulerEntry struct {
	value *SchedulerEntry
	isSet bool
}

func (v NullableSchedulerEntry) Get() *SchedulerEntry {
	return v.value
}

func (v *NullableSchedulerEntry) Set(val *SchedulerEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulerEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulerEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulerEntry(val *SchedulerEntry) *NullableSchedulerEntry {
	return &NullableSchedulerEntry{value: val, isSet: true}
}

func (v NullableSchedulerEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulerEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
