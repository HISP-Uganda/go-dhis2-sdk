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

// checks if the DataValueDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataValueDto{}

// DataValueDto struct for DataValueDto
type DataValueDto struct {
	Attribute *DataValueCategoryDto `json:"attribute,omitempty"`
	// A UID for an CategoryOptionCombo object   (Java name `org.hisp.dhis.category.CategoryOptionCombo`)
	CategoryOptionCombo *string    `json:"categoryOptionCombo,omitempty"`
	Comment             *string    `json:"comment,omitempty"`
	Created             *time.Time `json:"created,omitempty"`
	// A UID for an DataElement object   (Java name `org.hisp.dhis.dataelement.DataElement`)
	DataElement *string `json:"dataElement,omitempty"`
	// A UID for an DataSet object   (Java name `org.hisp.dhis.dataset.DataSet`)
	DataSet     *string    `json:"dataSet,omitempty"`
	FollowUp    *bool      `json:"followUp,omitempty"`
	Force       *bool      `json:"force,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// A UID for an OrganisationUnit object   (Java name `org.hisp.dhis.organisationunit.OrganisationUnit`)
	OrgUnit  *string `json:"orgUnit,omitempty"`
	Period   *string `json:"period,omitempty"`
	StoredBy *string `json:"storedBy,omitempty"`
	Value    *string `json:"value,omitempty"`
}

// NewDataValueDto instantiates a new DataValueDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataValueDto() *DataValueDto {
	this := DataValueDto{}
	return &this
}

// NewDataValueDtoWithDefaults instantiates a new DataValueDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataValueDtoWithDefaults() *DataValueDto {
	this := DataValueDto{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *DataValueDto) GetAttribute() DataValueCategoryDto {
	if o == nil || IsNil(o.Attribute) {
		var ret DataValueCategoryDto
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataValueDto) GetAttributeOk() (*DataValueCategoryDto, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *DataValueDto) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given DataValueCategoryDto and assigns it to the Attribute field.
func (o *DataValueDto) SetAttribute(v DataValueCategoryDto) {
	o.Attribute = &v
}

// GetCategoryOptionCombo returns the CategoryOptionCombo field value if set, zero value otherwise.
func (o *DataValueDto) GetCategoryOptionCombo() string {
	if o == nil || IsNil(o.CategoryOptionCombo) {
		var ret string
		return ret
	}
	return *o.CategoryOptionCombo
}

// GetCategoryOptionComboOk returns a tuple with the CategoryOptionCombo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataValueDto) GetCategoryOptionComboOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryOptionCombo) {
		return nil, false
	}
	return o.CategoryOptionCombo, true
}

// HasCategoryOptionCombo returns a boolean if a field has been set.
func (o *DataValueDto) HasCategoryOptionCombo() bool {
	if o != nil && !IsNil(o.CategoryOptionCombo) {
		return true
	}

	return false
}

// SetCategoryOptionCombo gets a reference to the given string and assigns it to the CategoryOptionCombo field.
func (o *DataValueDto) SetCategoryOptionCombo(v string) {
	o.CategoryOptionCombo = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *DataValueDto) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataValueDto) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *DataValueDto) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *DataValueDto) SetComment(v string) {
	o.Comment = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *DataValueDto) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataValueDto) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *DataValueDto) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *DataValueDto) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDataElement returns the DataElement field value if set, zero value otherwise.
func (o *DataValueDto) GetDataElement() string {
	if o == nil || IsNil(o.DataElement) {
		var ret string
		return ret
	}
	return *o.DataElement
}

// GetDataElementOk returns a tuple with the DataElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataValueDto) GetDataElementOk() (*string, bool) {
	if o == nil || IsNil(o.DataElement) {
		return nil, false
	}
	return o.DataElement, true
}

// HasDataElement returns a boolean if a field has been set.
func (o *DataValueDto) HasDataElement() bool {
	if o != nil && !IsNil(o.DataElement) {
		return true
	}

	return false
}

// SetDataElement gets a reference to the given string and assigns it to the DataElement field.
func (o *DataValueDto) SetDataElement(v string) {
	o.DataElement = &v
}

// GetDataSet returns the DataSet field value if set, zero value otherwise.
func (o *DataValueDto) GetDataSet() string {
	if o == nil || IsNil(o.DataSet) {
		var ret string
		return ret
	}
	return *o.DataSet
}

// GetDataSetOk returns a tuple with the DataSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataValueDto) GetDataSetOk() (*string, bool) {
	if o == nil || IsNil(o.DataSet) {
		return nil, false
	}
	return o.DataSet, true
}

// HasDataSet returns a boolean if a field has been set.
func (o *DataValueDto) HasDataSet() bool {
	if o != nil && !IsNil(o.DataSet) {
		return true
	}

	return false
}

// SetDataSet gets a reference to the given string and assigns it to the DataSet field.
func (o *DataValueDto) SetDataSet(v string) {
	o.DataSet = &v
}

// GetFollowUp returns the FollowUp field value if set, zero value otherwise.
func (o *DataValueDto) GetFollowUp() bool {
	if o == nil || IsNil(o.FollowUp) {
		var ret bool
		return ret
	}
	return *o.FollowUp
}

// GetFollowUpOk returns a tuple with the FollowUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataValueDto) GetFollowUpOk() (*bool, bool) {
	if o == nil || IsNil(o.FollowUp) {
		return nil, false
	}
	return o.FollowUp, true
}

// HasFollowUp returns a boolean if a field has been set.
func (o *DataValueDto) HasFollowUp() bool {
	if o != nil && !IsNil(o.FollowUp) {
		return true
	}

	return false
}

// SetFollowUp gets a reference to the given bool and assigns it to the FollowUp field.
func (o *DataValueDto) SetFollowUp(v bool) {
	o.FollowUp = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *DataValueDto) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataValueDto) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *DataValueDto) HasForce() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *DataValueDto) SetForce(v bool) {
	o.Force = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *DataValueDto) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataValueDto) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *DataValueDto) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *DataValueDto) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetOrgUnit returns the OrgUnit field value if set, zero value otherwise.
func (o *DataValueDto) GetOrgUnit() string {
	if o == nil || IsNil(o.OrgUnit) {
		var ret string
		return ret
	}
	return *o.OrgUnit
}

// GetOrgUnitOk returns a tuple with the OrgUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataValueDto) GetOrgUnitOk() (*string, bool) {
	if o == nil || IsNil(o.OrgUnit) {
		return nil, false
	}
	return o.OrgUnit, true
}

// HasOrgUnit returns a boolean if a field has been set.
func (o *DataValueDto) HasOrgUnit() bool {
	if o != nil && !IsNil(o.OrgUnit) {
		return true
	}

	return false
}

// SetOrgUnit gets a reference to the given string and assigns it to the OrgUnit field.
func (o *DataValueDto) SetOrgUnit(v string) {
	o.OrgUnit = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *DataValueDto) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataValueDto) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *DataValueDto) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *DataValueDto) SetPeriod(v string) {
	o.Period = &v
}

// GetStoredBy returns the StoredBy field value if set, zero value otherwise.
func (o *DataValueDto) GetStoredBy() string {
	if o == nil || IsNil(o.StoredBy) {
		var ret string
		return ret
	}
	return *o.StoredBy
}

// GetStoredByOk returns a tuple with the StoredBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataValueDto) GetStoredByOk() (*string, bool) {
	if o == nil || IsNil(o.StoredBy) {
		return nil, false
	}
	return o.StoredBy, true
}

// HasStoredBy returns a boolean if a field has been set.
func (o *DataValueDto) HasStoredBy() bool {
	if o != nil && !IsNil(o.StoredBy) {
		return true
	}

	return false
}

// SetStoredBy gets a reference to the given string and assigns it to the StoredBy field.
func (o *DataValueDto) SetStoredBy(v string) {
	o.StoredBy = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DataValueDto) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataValueDto) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DataValueDto) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DataValueDto) SetValue(v string) {
	o.Value = &v
}

func (o DataValueDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataValueDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !IsNil(o.CategoryOptionCombo) {
		toSerialize["categoryOptionCombo"] = o.CategoryOptionCombo
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.DataElement) {
		toSerialize["dataElement"] = o.DataElement
	}
	if !IsNil(o.DataSet) {
		toSerialize["dataSet"] = o.DataSet
	}
	if !IsNil(o.FollowUp) {
		toSerialize["followUp"] = o.FollowUp
	}
	if !IsNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.OrgUnit) {
		toSerialize["orgUnit"] = o.OrgUnit
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.StoredBy) {
		toSerialize["storedBy"] = o.StoredBy
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableDataValueDto struct {
	value *DataValueDto
	isSet bool
}

func (v NullableDataValueDto) Get() *DataValueDto {
	return v.value
}

func (v *NullableDataValueDto) Set(val *DataValueDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDataValueDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDataValueDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataValueDto(val *DataValueDto) *NullableDataValueDto {
	return &NullableDataValueDto{value: val, isSet: true}
}

func (v NullableDataValueDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataValueDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
