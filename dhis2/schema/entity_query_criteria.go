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

// checks if the EntityQueryCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityQueryCriteria{}

// EntityQueryCriteria struct for EntityQueryCriteria
type EntityQueryCriteria struct {
	AssignedUserMode       AssignedUserSelectionMode     `json:"assignedUserMode"`
	AssignedUsers          []string                      `json:"assignedUsers,omitempty"`
	AttributeValueFilters  []AttributeValueFilter        `json:"attributeValueFilters,omitempty"`
	DisplayColumnOrder     []string                      `json:"displayColumnOrder,omitempty"`
	EnrollmentCreatedDate  *DateFilterPeriod             `json:"enrollmentCreatedDate,omitempty"`
	EnrollmentIncidentDate *DateFilterPeriod             `json:"enrollmentIncidentDate,omitempty"`
	EnrollmentStatus       EnrollmentStatus              `json:"enrollmentStatus"`
	EventDate              *DateFilterPeriod             `json:"eventDate,omitempty"`
	EventStatus            EventStatus                   `json:"eventStatus"`
	FollowUp               *bool                         `json:"followUp,omitempty"`
	LastUpdatedDate        *DateFilterPeriod             `json:"lastUpdatedDate,omitempty"`
	Order                  *string                       `json:"order,omitempty"`
	OrganisationUnit       *string                       `json:"organisationUnit,omitempty"`
	OuMode                 OrganisationUnitSelectionMode `json:"ouMode"`
	ProgramStage           *string                       `json:"programStage,omitempty"`
	TrackedEntityInstances []string                      `json:"trackedEntityInstances,omitempty"`
	TrackedEntityType      *string                       `json:"trackedEntityType,omitempty"`
}

type _EntityQueryCriteria EntityQueryCriteria

// NewEntityQueryCriteria instantiates a new EntityQueryCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityQueryCriteria(assignedUserMode AssignedUserSelectionMode, enrollmentStatus EnrollmentStatus, eventStatus EventStatus, ouMode OrganisationUnitSelectionMode) *EntityQueryCriteria {
	this := EntityQueryCriteria{}
	this.AssignedUserMode = assignedUserMode
	this.EnrollmentStatus = enrollmentStatus
	this.EventStatus = eventStatus
	this.OuMode = ouMode
	return &this
}

// NewEntityQueryCriteriaWithDefaults instantiates a new EntityQueryCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityQueryCriteriaWithDefaults() *EntityQueryCriteria {
	this := EntityQueryCriteria{}
	return &this
}

// GetAssignedUserMode returns the AssignedUserMode field value
func (o *EntityQueryCriteria) GetAssignedUserMode() AssignedUserSelectionMode {
	if o == nil {
		var ret AssignedUserSelectionMode
		return ret
	}

	return o.AssignedUserMode
}

// GetAssignedUserModeOk returns a tuple with the AssignedUserMode field value
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetAssignedUserModeOk() (*AssignedUserSelectionMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedUserMode, true
}

// SetAssignedUserMode sets field value
func (o *EntityQueryCriteria) SetAssignedUserMode(v AssignedUserSelectionMode) {
	o.AssignedUserMode = v
}

// GetAssignedUsers returns the AssignedUsers field value if set, zero value otherwise.
func (o *EntityQueryCriteria) GetAssignedUsers() []string {
	if o == nil || IsNil(o.AssignedUsers) {
		var ret []string
		return ret
	}
	return o.AssignedUsers
}

// GetAssignedUsersOk returns a tuple with the AssignedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetAssignedUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.AssignedUsers) {
		return nil, false
	}
	return o.AssignedUsers, true
}

// HasAssignedUsers returns a boolean if a field has been set.
func (o *EntityQueryCriteria) HasAssignedUsers() bool {
	if o != nil && !IsNil(o.AssignedUsers) {
		return true
	}

	return false
}

// SetAssignedUsers gets a reference to the given []string and assigns it to the AssignedUsers field.
func (o *EntityQueryCriteria) SetAssignedUsers(v []string) {
	o.AssignedUsers = v
}

// GetAttributeValueFilters returns the AttributeValueFilters field value if set, zero value otherwise.
func (o *EntityQueryCriteria) GetAttributeValueFilters() []AttributeValueFilter {
	if o == nil || IsNil(o.AttributeValueFilters) {
		var ret []AttributeValueFilter
		return ret
	}
	return o.AttributeValueFilters
}

// GetAttributeValueFiltersOk returns a tuple with the AttributeValueFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetAttributeValueFiltersOk() ([]AttributeValueFilter, bool) {
	if o == nil || IsNil(o.AttributeValueFilters) {
		return nil, false
	}
	return o.AttributeValueFilters, true
}

// HasAttributeValueFilters returns a boolean if a field has been set.
func (o *EntityQueryCriteria) HasAttributeValueFilters() bool {
	if o != nil && !IsNil(o.AttributeValueFilters) {
		return true
	}

	return false
}

// SetAttributeValueFilters gets a reference to the given []AttributeValueFilter and assigns it to the AttributeValueFilters field.
func (o *EntityQueryCriteria) SetAttributeValueFilters(v []AttributeValueFilter) {
	o.AttributeValueFilters = v
}

// GetDisplayColumnOrder returns the DisplayColumnOrder field value if set, zero value otherwise.
func (o *EntityQueryCriteria) GetDisplayColumnOrder() []string {
	if o == nil || IsNil(o.DisplayColumnOrder) {
		var ret []string
		return ret
	}
	return o.DisplayColumnOrder
}

// GetDisplayColumnOrderOk returns a tuple with the DisplayColumnOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetDisplayColumnOrderOk() ([]string, bool) {
	if o == nil || IsNil(o.DisplayColumnOrder) {
		return nil, false
	}
	return o.DisplayColumnOrder, true
}

// HasDisplayColumnOrder returns a boolean if a field has been set.
func (o *EntityQueryCriteria) HasDisplayColumnOrder() bool {
	if o != nil && !IsNil(o.DisplayColumnOrder) {
		return true
	}

	return false
}

// SetDisplayColumnOrder gets a reference to the given []string and assigns it to the DisplayColumnOrder field.
func (o *EntityQueryCriteria) SetDisplayColumnOrder(v []string) {
	o.DisplayColumnOrder = v
}

// GetEnrollmentCreatedDate returns the EnrollmentCreatedDate field value if set, zero value otherwise.
func (o *EntityQueryCriteria) GetEnrollmentCreatedDate() DateFilterPeriod {
	if o == nil || IsNil(o.EnrollmentCreatedDate) {
		var ret DateFilterPeriod
		return ret
	}
	return *o.EnrollmentCreatedDate
}

// GetEnrollmentCreatedDateOk returns a tuple with the EnrollmentCreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetEnrollmentCreatedDateOk() (*DateFilterPeriod, bool) {
	if o == nil || IsNil(o.EnrollmentCreatedDate) {
		return nil, false
	}
	return o.EnrollmentCreatedDate, true
}

// HasEnrollmentCreatedDate returns a boolean if a field has been set.
func (o *EntityQueryCriteria) HasEnrollmentCreatedDate() bool {
	if o != nil && !IsNil(o.EnrollmentCreatedDate) {
		return true
	}

	return false
}

// SetEnrollmentCreatedDate gets a reference to the given DateFilterPeriod and assigns it to the EnrollmentCreatedDate field.
func (o *EntityQueryCriteria) SetEnrollmentCreatedDate(v DateFilterPeriod) {
	o.EnrollmentCreatedDate = &v
}

// GetEnrollmentIncidentDate returns the EnrollmentIncidentDate field value if set, zero value otherwise.
func (o *EntityQueryCriteria) GetEnrollmentIncidentDate() DateFilterPeriod {
	if o == nil || IsNil(o.EnrollmentIncidentDate) {
		var ret DateFilterPeriod
		return ret
	}
	return *o.EnrollmentIncidentDate
}

// GetEnrollmentIncidentDateOk returns a tuple with the EnrollmentIncidentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetEnrollmentIncidentDateOk() (*DateFilterPeriod, bool) {
	if o == nil || IsNil(o.EnrollmentIncidentDate) {
		return nil, false
	}
	return o.EnrollmentIncidentDate, true
}

// HasEnrollmentIncidentDate returns a boolean if a field has been set.
func (o *EntityQueryCriteria) HasEnrollmentIncidentDate() bool {
	if o != nil && !IsNil(o.EnrollmentIncidentDate) {
		return true
	}

	return false
}

// SetEnrollmentIncidentDate gets a reference to the given DateFilterPeriod and assigns it to the EnrollmentIncidentDate field.
func (o *EntityQueryCriteria) SetEnrollmentIncidentDate(v DateFilterPeriod) {
	o.EnrollmentIncidentDate = &v
}

// GetEnrollmentStatus returns the EnrollmentStatus field value
func (o *EntityQueryCriteria) GetEnrollmentStatus() EnrollmentStatus {
	if o == nil {
		var ret EnrollmentStatus
		return ret
	}

	return o.EnrollmentStatus
}

// GetEnrollmentStatusOk returns a tuple with the EnrollmentStatus field value
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetEnrollmentStatusOk() (*EnrollmentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnrollmentStatus, true
}

// SetEnrollmentStatus sets field value
func (o *EntityQueryCriteria) SetEnrollmentStatus(v EnrollmentStatus) {
	o.EnrollmentStatus = v
}

// GetEventDate returns the EventDate field value if set, zero value otherwise.
func (o *EntityQueryCriteria) GetEventDate() DateFilterPeriod {
	if o == nil || IsNil(o.EventDate) {
		var ret DateFilterPeriod
		return ret
	}
	return *o.EventDate
}

// GetEventDateOk returns a tuple with the EventDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetEventDateOk() (*DateFilterPeriod, bool) {
	if o == nil || IsNil(o.EventDate) {
		return nil, false
	}
	return o.EventDate, true
}

// HasEventDate returns a boolean if a field has been set.
func (o *EntityQueryCriteria) HasEventDate() bool {
	if o != nil && !IsNil(o.EventDate) {
		return true
	}

	return false
}

// SetEventDate gets a reference to the given DateFilterPeriod and assigns it to the EventDate field.
func (o *EntityQueryCriteria) SetEventDate(v DateFilterPeriod) {
	o.EventDate = &v
}

// GetEventStatus returns the EventStatus field value
func (o *EntityQueryCriteria) GetEventStatus() EventStatus {
	if o == nil {
		var ret EventStatus
		return ret
	}

	return o.EventStatus
}

// GetEventStatusOk returns a tuple with the EventStatus field value
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetEventStatusOk() (*EventStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventStatus, true
}

// SetEventStatus sets field value
func (o *EntityQueryCriteria) SetEventStatus(v EventStatus) {
	o.EventStatus = v
}

// GetFollowUp returns the FollowUp field value if set, zero value otherwise.
func (o *EntityQueryCriteria) GetFollowUp() bool {
	if o == nil || IsNil(o.FollowUp) {
		var ret bool
		return ret
	}
	return *o.FollowUp
}

// GetFollowUpOk returns a tuple with the FollowUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetFollowUpOk() (*bool, bool) {
	if o == nil || IsNil(o.FollowUp) {
		return nil, false
	}
	return o.FollowUp, true
}

// HasFollowUp returns a boolean if a field has been set.
func (o *EntityQueryCriteria) HasFollowUp() bool {
	if o != nil && !IsNil(o.FollowUp) {
		return true
	}

	return false
}

// SetFollowUp gets a reference to the given bool and assigns it to the FollowUp field.
func (o *EntityQueryCriteria) SetFollowUp(v bool) {
	o.FollowUp = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *EntityQueryCriteria) GetLastUpdatedDate() DateFilterPeriod {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret DateFilterPeriod
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetLastUpdatedDateOk() (*DateFilterPeriod, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *EntityQueryCriteria) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given DateFilterPeriod and assigns it to the LastUpdatedDate field.
func (o *EntityQueryCriteria) SetLastUpdatedDate(v DateFilterPeriod) {
	o.LastUpdatedDate = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *EntityQueryCriteria) GetOrder() string {
	if o == nil || IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetOrderOk() (*string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *EntityQueryCriteria) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *EntityQueryCriteria) SetOrder(v string) {
	o.Order = &v
}

// GetOrganisationUnit returns the OrganisationUnit field value if set, zero value otherwise.
func (o *EntityQueryCriteria) GetOrganisationUnit() string {
	if o == nil || IsNil(o.OrganisationUnit) {
		var ret string
		return ret
	}
	return *o.OrganisationUnit
}

// GetOrganisationUnitOk returns a tuple with the OrganisationUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetOrganisationUnitOk() (*string, bool) {
	if o == nil || IsNil(o.OrganisationUnit) {
		return nil, false
	}
	return o.OrganisationUnit, true
}

// HasOrganisationUnit returns a boolean if a field has been set.
func (o *EntityQueryCriteria) HasOrganisationUnit() bool {
	if o != nil && !IsNil(o.OrganisationUnit) {
		return true
	}

	return false
}

// SetOrganisationUnit gets a reference to the given string and assigns it to the OrganisationUnit field.
func (o *EntityQueryCriteria) SetOrganisationUnit(v string) {
	o.OrganisationUnit = &v
}

// GetOuMode returns the OuMode field value
func (o *EntityQueryCriteria) GetOuMode() OrganisationUnitSelectionMode {
	if o == nil {
		var ret OrganisationUnitSelectionMode
		return ret
	}

	return o.OuMode
}

// GetOuModeOk returns a tuple with the OuMode field value
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetOuModeOk() (*OrganisationUnitSelectionMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OuMode, true
}

// SetOuMode sets field value
func (o *EntityQueryCriteria) SetOuMode(v OrganisationUnitSelectionMode) {
	o.OuMode = v
}

// GetProgramStage returns the ProgramStage field value if set, zero value otherwise.
func (o *EntityQueryCriteria) GetProgramStage() string {
	if o == nil || IsNil(o.ProgramStage) {
		var ret string
		return ret
	}
	return *o.ProgramStage
}

// GetProgramStageOk returns a tuple with the ProgramStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetProgramStageOk() (*string, bool) {
	if o == nil || IsNil(o.ProgramStage) {
		return nil, false
	}
	return o.ProgramStage, true
}

// HasProgramStage returns a boolean if a field has been set.
func (o *EntityQueryCriteria) HasProgramStage() bool {
	if o != nil && !IsNil(o.ProgramStage) {
		return true
	}

	return false
}

// SetProgramStage gets a reference to the given string and assigns it to the ProgramStage field.
func (o *EntityQueryCriteria) SetProgramStage(v string) {
	o.ProgramStage = &v
}

// GetTrackedEntityInstances returns the TrackedEntityInstances field value if set, zero value otherwise.
func (o *EntityQueryCriteria) GetTrackedEntityInstances() []string {
	if o == nil || IsNil(o.TrackedEntityInstances) {
		var ret []string
		return ret
	}
	return o.TrackedEntityInstances
}

// GetTrackedEntityInstancesOk returns a tuple with the TrackedEntityInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetTrackedEntityInstancesOk() ([]string, bool) {
	if o == nil || IsNil(o.TrackedEntityInstances) {
		return nil, false
	}
	return o.TrackedEntityInstances, true
}

// HasTrackedEntityInstances returns a boolean if a field has been set.
func (o *EntityQueryCriteria) HasTrackedEntityInstances() bool {
	if o != nil && !IsNil(o.TrackedEntityInstances) {
		return true
	}

	return false
}

// SetTrackedEntityInstances gets a reference to the given []string and assigns it to the TrackedEntityInstances field.
func (o *EntityQueryCriteria) SetTrackedEntityInstances(v []string) {
	o.TrackedEntityInstances = v
}

// GetTrackedEntityType returns the TrackedEntityType field value if set, zero value otherwise.
func (o *EntityQueryCriteria) GetTrackedEntityType() string {
	if o == nil || IsNil(o.TrackedEntityType) {
		var ret string
		return ret
	}
	return *o.TrackedEntityType
}

// GetTrackedEntityTypeOk returns a tuple with the TrackedEntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityQueryCriteria) GetTrackedEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TrackedEntityType) {
		return nil, false
	}
	return o.TrackedEntityType, true
}

// HasTrackedEntityType returns a boolean if a field has been set.
func (o *EntityQueryCriteria) HasTrackedEntityType() bool {
	if o != nil && !IsNil(o.TrackedEntityType) {
		return true
	}

	return false
}

// SetTrackedEntityType gets a reference to the given string and assigns it to the TrackedEntityType field.
func (o *EntityQueryCriteria) SetTrackedEntityType(v string) {
	o.TrackedEntityType = &v
}

func (o EntityQueryCriteria) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityQueryCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assignedUserMode"] = o.AssignedUserMode
	if !IsNil(o.AssignedUsers) {
		toSerialize["assignedUsers"] = o.AssignedUsers
	}
	if !IsNil(o.AttributeValueFilters) {
		toSerialize["attributeValueFilters"] = o.AttributeValueFilters
	}
	if !IsNil(o.DisplayColumnOrder) {
		toSerialize["displayColumnOrder"] = o.DisplayColumnOrder
	}
	if !IsNil(o.EnrollmentCreatedDate) {
		toSerialize["enrollmentCreatedDate"] = o.EnrollmentCreatedDate
	}
	if !IsNil(o.EnrollmentIncidentDate) {
		toSerialize["enrollmentIncidentDate"] = o.EnrollmentIncidentDate
	}
	toSerialize["enrollmentStatus"] = o.EnrollmentStatus
	if !IsNil(o.EventDate) {
		toSerialize["eventDate"] = o.EventDate
	}
	toSerialize["eventStatus"] = o.EventStatus
	if !IsNil(o.FollowUp) {
		toSerialize["followUp"] = o.FollowUp
	}
	if !IsNil(o.LastUpdatedDate) {
		toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.OrganisationUnit) {
		toSerialize["organisationUnit"] = o.OrganisationUnit
	}
	toSerialize["ouMode"] = o.OuMode
	if !IsNil(o.ProgramStage) {
		toSerialize["programStage"] = o.ProgramStage
	}
	if !IsNil(o.TrackedEntityInstances) {
		toSerialize["trackedEntityInstances"] = o.TrackedEntityInstances
	}
	if !IsNil(o.TrackedEntityType) {
		toSerialize["trackedEntityType"] = o.TrackedEntityType
	}
	return toSerialize, nil
}

func (o *EntityQueryCriteria) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assignedUserMode",
		"enrollmentStatus",
		"eventStatus",
		"ouMode",
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

	varEntityQueryCriteria := _EntityQueryCriteria{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEntityQueryCriteria)

	if err != nil {
		return err
	}

	*o = EntityQueryCriteria(varEntityQueryCriteria)

	return err
}

type NullableEntityQueryCriteria struct {
	value *EntityQueryCriteria
	isSet bool
}

func (v NullableEntityQueryCriteria) Get() *EntityQueryCriteria {
	return v.value
}

func (v *NullableEntityQueryCriteria) Set(val *EntityQueryCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityQueryCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityQueryCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityQueryCriteria(val *EntityQueryCriteria) *NullableEntityQueryCriteria {
	return &NullableEntityQueryCriteria{value: val, isSet: true}
}

func (v NullableEntityQueryCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityQueryCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
