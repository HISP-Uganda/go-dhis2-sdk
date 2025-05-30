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
)

// checks if the ImportSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportSummary{}

// ImportSummary struct for ImportSummary
type ImportSummary struct {
	Conflicts       []ImportConflict `json:"conflicts,omitempty"`
	DataSetComplete *string          `json:"dataSetComplete,omitempty"`
	Description     *string          `json:"description,omitempty"`
	Href            *string          `json:"href,omitempty"`
	ImportCount     *ImportCount     `json:"importCount,omitempty"`
	ImportOptions   *ImportOptions   `json:"importOptions,omitempty"`
	Reference       *string          `json:"reference,omitempty"`
	RejectedIndexes []int32          `json:"rejectedIndexes,omitempty"`
	ResponseType    *string          `json:"responseType,omitempty"`
	Status          ImportStatus     `json:"status"`
}

type _ImportSummary ImportSummary

// NewImportSummary instantiates a new ImportSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportSummary(status ImportStatus) *ImportSummary {
	this := ImportSummary{}
	this.Status = status
	return &this
}

// NewImportSummaryWithDefaults instantiates a new ImportSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportSummaryWithDefaults() *ImportSummary {
	this := ImportSummary{}
	return &this
}

// GetConflicts returns the Conflicts field value if set, zero value otherwise.
func (o *ImportSummary) GetConflicts() []ImportConflict {
	if o == nil || IsNil(o.Conflicts) {
		var ret []ImportConflict
		return ret
	}
	return o.Conflicts
}

// GetConflictsOk returns a tuple with the Conflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetConflictsOk() ([]ImportConflict, bool) {
	if o == nil || IsNil(o.Conflicts) {
		return nil, false
	}
	return o.Conflicts, true
}

// HasConflicts returns a boolean if a field has been set.
func (o *ImportSummary) HasConflicts() bool {
	if o != nil && !IsNil(o.Conflicts) {
		return true
	}

	return false
}

// SetConflicts gets a reference to the given []ImportConflict and assigns it to the Conflicts field.
func (o *ImportSummary) SetConflicts(v []ImportConflict) {
	o.Conflicts = v
}

// GetDataSetComplete returns the DataSetComplete field value if set, zero value otherwise.
func (o *ImportSummary) GetDataSetComplete() string {
	if o == nil || IsNil(o.DataSetComplete) {
		var ret string
		return ret
	}
	return *o.DataSetComplete
}

// GetDataSetCompleteOk returns a tuple with the DataSetComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetDataSetCompleteOk() (*string, bool) {
	if o == nil || IsNil(o.DataSetComplete) {
		return nil, false
	}
	return o.DataSetComplete, true
}

// HasDataSetComplete returns a boolean if a field has been set.
func (o *ImportSummary) HasDataSetComplete() bool {
	if o != nil && !IsNil(o.DataSetComplete) {
		return true
	}

	return false
}

// SetDataSetComplete gets a reference to the given string and assigns it to the DataSetComplete field.
func (o *ImportSummary) SetDataSetComplete(v string) {
	o.DataSetComplete = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ImportSummary) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ImportSummary) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ImportSummary) SetDescription(v string) {
	o.Description = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ImportSummary) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ImportSummary) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ImportSummary) SetHref(v string) {
	o.Href = &v
}

// GetImportCount returns the ImportCount field value if set, zero value otherwise.
func (o *ImportSummary) GetImportCount() ImportCount {
	if o == nil || IsNil(o.ImportCount) {
		var ret ImportCount
		return ret
	}
	return *o.ImportCount
}

// GetImportCountOk returns a tuple with the ImportCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetImportCountOk() (*ImportCount, bool) {
	if o == nil || IsNil(o.ImportCount) {
		return nil, false
	}
	return o.ImportCount, true
}

// HasImportCount returns a boolean if a field has been set.
func (o *ImportSummary) HasImportCount() bool {
	if o != nil && !IsNil(o.ImportCount) {
		return true
	}

	return false
}

// SetImportCount gets a reference to the given ImportCount and assigns it to the ImportCount field.
func (o *ImportSummary) SetImportCount(v ImportCount) {
	o.ImportCount = &v
}

// GetImportOptions returns the ImportOptions field value if set, zero value otherwise.
func (o *ImportSummary) GetImportOptions() ImportOptions {
	if o == nil || IsNil(o.ImportOptions) {
		var ret ImportOptions
		return ret
	}
	return *o.ImportOptions
}

// GetImportOptionsOk returns a tuple with the ImportOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetImportOptionsOk() (*ImportOptions, bool) {
	if o == nil || IsNil(o.ImportOptions) {
		return nil, false
	}
	return o.ImportOptions, true
}

// HasImportOptions returns a boolean if a field has been set.
func (o *ImportSummary) HasImportOptions() bool {
	if o != nil && !IsNil(o.ImportOptions) {
		return true
	}

	return false
}

// SetImportOptions gets a reference to the given ImportOptions and assigns it to the ImportOptions field.
func (o *ImportSummary) SetImportOptions(v ImportOptions) {
	o.ImportOptions = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ImportSummary) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ImportSummary) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ImportSummary) SetReference(v string) {
	o.Reference = &v
}

// GetRejectedIndexes returns the RejectedIndexes field value if set, zero value otherwise.
func (o *ImportSummary) GetRejectedIndexes() []int32 {
	if o == nil || IsNil(o.RejectedIndexes) {
		var ret []int32
		return ret
	}
	return o.RejectedIndexes
}

// GetRejectedIndexesOk returns a tuple with the RejectedIndexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetRejectedIndexesOk() ([]int32, bool) {
	if o == nil || IsNil(o.RejectedIndexes) {
		return nil, false
	}
	return o.RejectedIndexes, true
}

// HasRejectedIndexes returns a boolean if a field has been set.
func (o *ImportSummary) HasRejectedIndexes() bool {
	if o != nil && !IsNil(o.RejectedIndexes) {
		return true
	}

	return false
}

// SetRejectedIndexes gets a reference to the given []int32 and assigns it to the RejectedIndexes field.
func (o *ImportSummary) SetRejectedIndexes(v []int32) {
	o.RejectedIndexes = v
}

// GetResponseType returns the ResponseType field value if set, zero value otherwise.
func (o *ImportSummary) GetResponseType() string {
	if o == nil || IsNil(o.ResponseType) {
		var ret string
		return ret
	}
	return *o.ResponseType
}

// GetResponseTypeOk returns a tuple with the ResponseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetResponseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseType) {
		return nil, false
	}
	return o.ResponseType, true
}

// HasResponseType returns a boolean if a field has been set.
func (o *ImportSummary) HasResponseType() bool {
	if o != nil && !IsNil(o.ResponseType) {
		return true
	}

	return false
}

// SetResponseType gets a reference to the given string and assigns it to the ResponseType field.
func (o *ImportSummary) SetResponseType(v string) {
	o.ResponseType = &v
}

// GetStatus returns the Status field value
func (o *ImportSummary) GetStatus() ImportStatus {
	if o == nil {
		var ret ImportStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetStatusOk() (*ImportStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ImportSummary) SetStatus(v ImportStatus) {
	o.Status = v
}

func (o ImportSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Conflicts) {
		toSerialize["conflicts"] = o.Conflicts
	}
	if !IsNil(o.DataSetComplete) {
		toSerialize["dataSetComplete"] = o.DataSetComplete
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.ImportCount) {
		toSerialize["importCount"] = o.ImportCount
	}
	if !IsNil(o.ImportOptions) {
		toSerialize["importOptions"] = o.ImportOptions
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.RejectedIndexes) {
		toSerialize["rejectedIndexes"] = o.RejectedIndexes
	}
	if !IsNil(o.ResponseType) {
		toSerialize["responseType"] = o.ResponseType
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ImportSummary) UnmarshalJSON(data []byte) (err error) {
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

	varImportSummary := _ImportSummary{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImportSummary)

	if err != nil {
		return err
	}

	*o = ImportSummary(varImportSummary)

	return err
}

type NullableImportSummary struct {
	value *ImportSummary
	isSet bool
}

func (v NullableImportSummary) Get() *ImportSummary {
	return v.value
}

func (v *NullableImportSummary) Set(val *ImportSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableImportSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableImportSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportSummary(val *ImportSummary) *NullableImportSummary {
	return &NullableImportSummary{value: val, isSet: true}
}

func (v NullableImportSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
