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

// checks if the TrackerRelationshipItemTrackedEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackerRelationshipItemTrackedEntity{}

// TrackerRelationshipItemTrackedEntity struct for TrackerRelationshipItemTrackedEntity
type TrackerRelationshipItemTrackedEntity struct {
	Attributes      []TrackerAttribute                  `json:"attributes,omitempty"`
	CreatedAt       *Instant                            `json:"createdAt,omitempty"`
	CreatedAtClient *Instant                            `json:"createdAtClient,omitempty"`
	CreatedBy       *TrackerUser                        `json:"createdBy,omitempty"`
	Deleted         *bool                               `json:"deleted,omitempty"`
	Enrollments     []TrackerRelationshipItemEnrollment `json:"enrollments,omitempty"`
	Geometry        map[string]interface{}              `json:"geometry,omitempty"`
	Inactive        *bool                               `json:"inactive,omitempty"`
	// A UID for an OrganisationUnit object   (Java name `org.hisp.dhis.organisationunit.OrganisationUnit`)
	OrgUnit            *string               `json:"orgUnit,omitempty"`
	PotentialDuplicate *bool                 `json:"potentialDuplicate,omitempty"`
	ProgramOwners      []TrackerProgramOwner `json:"programOwners,omitempty"`
	// A UID for an User object   (Java name `org.hisp.dhis.user.User`)
	StoredBy *string `json:"storedBy,omitempty"`
	// A UID for an TrackedEntity object   (Java name `org.hisp.dhis.trackedentity.TrackedEntity`)
	TrackedEntity     *string      `json:"trackedEntity,omitempty"`
	TrackedEntityType *string      `json:"trackedEntityType,omitempty"`
	UpdatedAt         *Instant     `json:"updatedAt,omitempty"`
	UpdatedAtClient   *Instant     `json:"updatedAtClient,omitempty"`
	UpdatedBy         *TrackerUser `json:"updatedBy,omitempty"`
}

// NewTrackerRelationshipItemTrackedEntity instantiates a new TrackerRelationshipItemTrackedEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackerRelationshipItemTrackedEntity() *TrackerRelationshipItemTrackedEntity {
	this := TrackerRelationshipItemTrackedEntity{}
	return &this
}

// NewTrackerRelationshipItemTrackedEntityWithDefaults instantiates a new TrackerRelationshipItemTrackedEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackerRelationshipItemTrackedEntityWithDefaults() *TrackerRelationshipItemTrackedEntity {
	this := TrackerRelationshipItemTrackedEntity{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetAttributes() []TrackerAttribute {
	if o == nil || IsNil(o.Attributes) {
		var ret []TrackerAttribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetAttributesOk() ([]TrackerAttribute, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []TrackerAttribute and assigns it to the Attributes field.
func (o *TrackerRelationshipItemTrackedEntity) SetAttributes(v []TrackerAttribute) {
	o.Attributes = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetCreatedAt() Instant {
	if o == nil || IsNil(o.CreatedAt) {
		var ret Instant
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetCreatedAtOk() (*Instant, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given Instant and assigns it to the CreatedAt field.
func (o *TrackerRelationshipItemTrackedEntity) SetCreatedAt(v Instant) {
	o.CreatedAt = &v
}

// GetCreatedAtClient returns the CreatedAtClient field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetCreatedAtClient() Instant {
	if o == nil || IsNil(o.CreatedAtClient) {
		var ret Instant
		return ret
	}
	return *o.CreatedAtClient
}

// GetCreatedAtClientOk returns a tuple with the CreatedAtClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetCreatedAtClientOk() (*Instant, bool) {
	if o == nil || IsNil(o.CreatedAtClient) {
		return nil, false
	}
	return o.CreatedAtClient, true
}

// HasCreatedAtClient returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasCreatedAtClient() bool {
	if o != nil && !IsNil(o.CreatedAtClient) {
		return true
	}

	return false
}

// SetCreatedAtClient gets a reference to the given Instant and assigns it to the CreatedAtClient field.
func (o *TrackerRelationshipItemTrackedEntity) SetCreatedAtClient(v Instant) {
	o.CreatedAtClient = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetCreatedBy() TrackerUser {
	if o == nil || IsNil(o.CreatedBy) {
		var ret TrackerUser
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetCreatedByOk() (*TrackerUser, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given TrackerUser and assigns it to the CreatedBy field.
func (o *TrackerRelationshipItemTrackedEntity) SetCreatedBy(v TrackerUser) {
	o.CreatedBy = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *TrackerRelationshipItemTrackedEntity) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetEnrollments returns the Enrollments field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetEnrollments() []TrackerRelationshipItemEnrollment {
	if o == nil || IsNil(o.Enrollments) {
		var ret []TrackerRelationshipItemEnrollment
		return ret
	}
	return o.Enrollments
}

// GetEnrollmentsOk returns a tuple with the Enrollments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetEnrollmentsOk() ([]TrackerRelationshipItemEnrollment, bool) {
	if o == nil || IsNil(o.Enrollments) {
		return nil, false
	}
	return o.Enrollments, true
}

// HasEnrollments returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasEnrollments() bool {
	if o != nil && !IsNil(o.Enrollments) {
		return true
	}

	return false
}

// SetEnrollments gets a reference to the given []TrackerRelationshipItemEnrollment and assigns it to the Enrollments field.
func (o *TrackerRelationshipItemTrackedEntity) SetEnrollments(v []TrackerRelationshipItemEnrollment) {
	o.Enrollments = v
}

// GetGeometry returns the Geometry field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetGeometry() map[string]interface{} {
	if o == nil || IsNil(o.Geometry) {
		var ret map[string]interface{}
		return ret
	}
	return o.Geometry
}

// GetGeometryOk returns a tuple with the Geometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetGeometryOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Geometry) {
		return map[string]interface{}{}, false
	}
	return o.Geometry, true
}

// HasGeometry returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasGeometry() bool {
	if o != nil && !IsNil(o.Geometry) {
		return true
	}

	return false
}

// SetGeometry gets a reference to the given map[string]interface{} and assigns it to the Geometry field.
func (o *TrackerRelationshipItemTrackedEntity) SetGeometry(v map[string]interface{}) {
	o.Geometry = v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetInactive() bool {
	if o == nil || IsNil(o.Inactive) {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetInactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Inactive) {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasInactive() bool {
	if o != nil && !IsNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *TrackerRelationshipItemTrackedEntity) SetInactive(v bool) {
	o.Inactive = &v
}

// GetOrgUnit returns the OrgUnit field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetOrgUnit() string {
	if o == nil || IsNil(o.OrgUnit) {
		var ret string
		return ret
	}
	return *o.OrgUnit
}

// GetOrgUnitOk returns a tuple with the OrgUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetOrgUnitOk() (*string, bool) {
	if o == nil || IsNil(o.OrgUnit) {
		return nil, false
	}
	return o.OrgUnit, true
}

// HasOrgUnit returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasOrgUnit() bool {
	if o != nil && !IsNil(o.OrgUnit) {
		return true
	}

	return false
}

// SetOrgUnit gets a reference to the given string and assigns it to the OrgUnit field.
func (o *TrackerRelationshipItemTrackedEntity) SetOrgUnit(v string) {
	o.OrgUnit = &v
}

// GetPotentialDuplicate returns the PotentialDuplicate field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetPotentialDuplicate() bool {
	if o == nil || IsNil(o.PotentialDuplicate) {
		var ret bool
		return ret
	}
	return *o.PotentialDuplicate
}

// GetPotentialDuplicateOk returns a tuple with the PotentialDuplicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetPotentialDuplicateOk() (*bool, bool) {
	if o == nil || IsNil(o.PotentialDuplicate) {
		return nil, false
	}
	return o.PotentialDuplicate, true
}

// HasPotentialDuplicate returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasPotentialDuplicate() bool {
	if o != nil && !IsNil(o.PotentialDuplicate) {
		return true
	}

	return false
}

// SetPotentialDuplicate gets a reference to the given bool and assigns it to the PotentialDuplicate field.
func (o *TrackerRelationshipItemTrackedEntity) SetPotentialDuplicate(v bool) {
	o.PotentialDuplicate = &v
}

// GetProgramOwners returns the ProgramOwners field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetProgramOwners() []TrackerProgramOwner {
	if o == nil || IsNil(o.ProgramOwners) {
		var ret []TrackerProgramOwner
		return ret
	}
	return o.ProgramOwners
}

// GetProgramOwnersOk returns a tuple with the ProgramOwners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetProgramOwnersOk() ([]TrackerProgramOwner, bool) {
	if o == nil || IsNil(o.ProgramOwners) {
		return nil, false
	}
	return o.ProgramOwners, true
}

// HasProgramOwners returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasProgramOwners() bool {
	if o != nil && !IsNil(o.ProgramOwners) {
		return true
	}

	return false
}

// SetProgramOwners gets a reference to the given []TrackerProgramOwner and assigns it to the ProgramOwners field.
func (o *TrackerRelationshipItemTrackedEntity) SetProgramOwners(v []TrackerProgramOwner) {
	o.ProgramOwners = v
}

// GetStoredBy returns the StoredBy field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetStoredBy() string {
	if o == nil || IsNil(o.StoredBy) {
		var ret string
		return ret
	}
	return *o.StoredBy
}

// GetStoredByOk returns a tuple with the StoredBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetStoredByOk() (*string, bool) {
	if o == nil || IsNil(o.StoredBy) {
		return nil, false
	}
	return o.StoredBy, true
}

// HasStoredBy returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasStoredBy() bool {
	if o != nil && !IsNil(o.StoredBy) {
		return true
	}

	return false
}

// SetStoredBy gets a reference to the given string and assigns it to the StoredBy field.
func (o *TrackerRelationshipItemTrackedEntity) SetStoredBy(v string) {
	o.StoredBy = &v
}

// GetTrackedEntity returns the TrackedEntity field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetTrackedEntity() string {
	if o == nil || IsNil(o.TrackedEntity) {
		var ret string
		return ret
	}
	return *o.TrackedEntity
}

// GetTrackedEntityOk returns a tuple with the TrackedEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetTrackedEntityOk() (*string, bool) {
	if o == nil || IsNil(o.TrackedEntity) {
		return nil, false
	}
	return o.TrackedEntity, true
}

// HasTrackedEntity returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasTrackedEntity() bool {
	if o != nil && !IsNil(o.TrackedEntity) {
		return true
	}

	return false
}

// SetTrackedEntity gets a reference to the given string and assigns it to the TrackedEntity field.
func (o *TrackerRelationshipItemTrackedEntity) SetTrackedEntity(v string) {
	o.TrackedEntity = &v
}

// GetTrackedEntityType returns the TrackedEntityType field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetTrackedEntityType() string {
	if o == nil || IsNil(o.TrackedEntityType) {
		var ret string
		return ret
	}
	return *o.TrackedEntityType
}

// GetTrackedEntityTypeOk returns a tuple with the TrackedEntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetTrackedEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TrackedEntityType) {
		return nil, false
	}
	return o.TrackedEntityType, true
}

// HasTrackedEntityType returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasTrackedEntityType() bool {
	if o != nil && !IsNil(o.TrackedEntityType) {
		return true
	}

	return false
}

// SetTrackedEntityType gets a reference to the given string and assigns it to the TrackedEntityType field.
func (o *TrackerRelationshipItemTrackedEntity) SetTrackedEntityType(v string) {
	o.TrackedEntityType = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetUpdatedAt() Instant {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret Instant
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetUpdatedAtOk() (*Instant, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given Instant and assigns it to the UpdatedAt field.
func (o *TrackerRelationshipItemTrackedEntity) SetUpdatedAt(v Instant) {
	o.UpdatedAt = &v
}

// GetUpdatedAtClient returns the UpdatedAtClient field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetUpdatedAtClient() Instant {
	if o == nil || IsNil(o.UpdatedAtClient) {
		var ret Instant
		return ret
	}
	return *o.UpdatedAtClient
}

// GetUpdatedAtClientOk returns a tuple with the UpdatedAtClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetUpdatedAtClientOk() (*Instant, bool) {
	if o == nil || IsNil(o.UpdatedAtClient) {
		return nil, false
	}
	return o.UpdatedAtClient, true
}

// HasUpdatedAtClient returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasUpdatedAtClient() bool {
	if o != nil && !IsNil(o.UpdatedAtClient) {
		return true
	}

	return false
}

// SetUpdatedAtClient gets a reference to the given Instant and assigns it to the UpdatedAtClient field.
func (o *TrackerRelationshipItemTrackedEntity) SetUpdatedAtClient(v Instant) {
	o.UpdatedAtClient = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *TrackerRelationshipItemTrackedEntity) GetUpdatedBy() TrackerUser {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret TrackerUser
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackerRelationshipItemTrackedEntity) GetUpdatedByOk() (*TrackerUser, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *TrackerRelationshipItemTrackedEntity) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given TrackerUser and assigns it to the UpdatedBy field.
func (o *TrackerRelationshipItemTrackedEntity) SetUpdatedBy(v TrackerUser) {
	o.UpdatedBy = &v
}

func (o TrackerRelationshipItemTrackedEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrackerRelationshipItemTrackedEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.CreatedAtClient) {
		toSerialize["createdAtClient"] = o.CreatedAtClient
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !IsNil(o.Enrollments) {
		toSerialize["enrollments"] = o.Enrollments
	}
	if !IsNil(o.Geometry) {
		toSerialize["geometry"] = o.Geometry
	}
	if !IsNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	if !IsNil(o.OrgUnit) {
		toSerialize["orgUnit"] = o.OrgUnit
	}
	if !IsNil(o.PotentialDuplicate) {
		toSerialize["potentialDuplicate"] = o.PotentialDuplicate
	}
	if !IsNil(o.ProgramOwners) {
		toSerialize["programOwners"] = o.ProgramOwners
	}
	if !IsNil(o.StoredBy) {
		toSerialize["storedBy"] = o.StoredBy
	}
	if !IsNil(o.TrackedEntity) {
		toSerialize["trackedEntity"] = o.TrackedEntity
	}
	if !IsNil(o.TrackedEntityType) {
		toSerialize["trackedEntityType"] = o.TrackedEntityType
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedAtClient) {
		toSerialize["updatedAtClient"] = o.UpdatedAtClient
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	return toSerialize, nil
}

type NullableTrackerRelationshipItemTrackedEntity struct {
	value *TrackerRelationshipItemTrackedEntity
	isSet bool
}

func (v NullableTrackerRelationshipItemTrackedEntity) Get() *TrackerRelationshipItemTrackedEntity {
	return v.value
}

func (v *NullableTrackerRelationshipItemTrackedEntity) Set(val *TrackerRelationshipItemTrackedEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackerRelationshipItemTrackedEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackerRelationshipItemTrackedEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackerRelationshipItemTrackedEntity(val *TrackerRelationshipItemTrackedEntity) *NullableTrackerRelationshipItemTrackedEntity {
	return &NullableTrackerRelationshipItemTrackedEntity{value: val, isSet: true}
}

func (v NullableTrackerRelationshipItemTrackedEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackerRelationshipItemTrackedEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
