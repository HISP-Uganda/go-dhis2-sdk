package tracker

import (
	"github.com/HISP-Uganda/go-dhis2-sdk/dhis2/schema"
	"github.com/HISP-Uganda/go-dhis2-sdk/utils"
	"time"
)

type Attribute struct {
	Attribute   *string   `json:"attribute"`
	Code        *string   `json:"code,omitempty"`
	DisplayName *string   `json:"displayName,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	StoredBy    *string   `json:"storedBy,omitempty"`
	ValueType   *string   `json:"valueType,omitempty"`
	Value       *string   `json:"value"`
}

type AsyncResponse struct {
	Response *struct {
		ID string `json:"id"`
	} `json:"response,omitempty"`
	Name                     *string `json:"name,omitempty"`
	ID                       *string `json:"id,omitempty"`
	Created                  *string `json:"created,omitempty"`
	JobType                  *string `json:"jobType,omitempty"`
	RelativeNotifierEndpoint *string `json:"relativeNotifierEndpoint,omitempty"`
}

// GetJobID returns the job ID from the async response.
func (r *AsyncResponse) GetJobID() string {
	if r.Response != nil && r.Response.ID != "" {
		return r.Response.ID
	} else if r.ID != nil {
		return *r.ID
	}
	return ""
}

type TrackerResponse struct {
	SyncResp  *schema.TrackerImportReport
	AsyncResp *AsyncResponse
}

type LegacyNestedPayload struct {
	TrackedEntities []LegacyNestedTrackedEntity `json:"trackedEntityInstances,omitempty"`
}

// LegacyNestedTrackedEntity represents a tracked entity with nested enrollments.
type LegacyNestedTrackedEntity struct {
	Attributes        []schema.TrackerAttribute `json:"attributes,omitempty"`
	Enrollments       []LegacyNestedEnrollment  `json:"enrollments,omitempty"`
	OrgUnit           string                    `json:"orgUnit"`
	TrackedEntityType string                    `json:"trackedEntityType"`
}

// LegacyNestedEnrollment represents an enrollment with nested attributes and events.
type LegacyNestedEnrollment struct {
	Attributes []schema.TrackerAttribute `json:"attributes,omitempty"`
	EnrolledAt time.Time                 `json:"enrollmentDate"`
	Events     []LegacyNestedEvent       `json:"events,omitempty"`
	OccurredAt time.Time                 `json:"incidentDate"`
	OrgUnit    string                    `json:"orgUnit"`
	Program    string                    `json:"program"`
	Status     string                    `json:"status"`
}

// LegacyNestedEvent represents an event with nested data values and notes.
type LegacyNestedEvent struct {
	AttributeCategoryOptions string               `json:"attributeCategoryOptions,omitempty"`
	AttributeOptionCombo     string               `json:"attributeOptionCombo,omitempty"`
	DataValues               []schema.DataValue   `json:"dataValues,omitempty"`
	EnrollmentStatus         string               `json:"enrollmentStatus,omitempty"`
	Notes                    []schema.TrackerNote `json:"notes,omitempty"`
	OccurredAt               time.Time            `json:"eventDate"`
	OrgUnit                  string               `json:"orgUnit"`
	Program                  string               `json:"program"`
	ProgramStage             string               `json:"programStage"`
	ScheduledAt              *time.Time           `json:"dueDate,omitempty"`
	Status                   string               `json:"status"`
}

// NestedPayload represents the top-level structure of a nested payload for DHIS2 tracker data import.
type NestedPayload struct {
	TrackedEntities []NestedTrackedEntity `json:"trackedEntities,omitempty"`
}

// NestedTrackedEntity represents a tracked entity with nested enrollments.
type NestedTrackedEntity struct {
	Attributes        []Attribute        `json:"attributes,omitempty"`
	Enrollments       []NestedEnrollment `json:"enrollments,omitempty"`
	OrgUnit           string             `json:"orgUnit"`
	TrackedEntityType string             `json:"trackedEntityType"`
}

// NestedEnrollment represents an enrollment with nested attributes and events.
type NestedEnrollment struct {
	Attributes []schema.TrackerAttribute `json:"attributes,omitempty"`
	EnrolledAt time.Time                 `json:"enrollmentDate"`
	Events     []NestedEvent             `json:"events,omitempty"`
	OccurredAt time.Time                 `json:"incidentDate"`
	OrgUnit    string                    `json:"orgUnit"`
	Program    string                    `json:"program"`
	Status     string                    `json:"status"`
	// TrackedEntityType string            `json:"trackedEntityType"`
}

// NestedEvent represents an event with nested data values and notes.
type NestedEvent struct {
	AttributeCategoryOptions string                    `json:"attributeCategoryOptions,omitempty"`
	AttributeOptionCombo     string                    `json:"attributeOptionCombo,omitempty"`
	DataValues               []schema.TrackerDataValue `json:"dataValues,omitempty"`
	EnrollmentStatus         string                    `json:"enrollmentStatus,omitempty"`
	Notes                    []schema.TrackerNote      `json:"notes,omitempty"`
	OccurredAt               time.Time                 `json:"eventDate"`
	OrgUnit                  string                    `json:"orgUnit"`
	Program                  string                    `json:"program"`
	ProgramStage             string                    `json:"programStage"`
	ScheduledAt              *time.Time                `json:"dueDate,omitempty"`
	Status                   string                    `json:"status"`
}

// AddAttribute appends an attribute to the tracked entity.
func (te *NestedTrackedEntity) AddAttribute(attr, value string) *NestedTrackedEntity {
	te.Attributes = append(te.Attributes, Attribute{
		Attribute: utils.StringPtr(attr),
		Value:     utils.StringPtr(value),
	})
	return te
}

// AddEnrollment appends an enrollment to the tracked entity.
func (te *NestedTrackedEntity) AddEnrollment(e NestedEnrollment) *NestedTrackedEntity {
	te.Enrollments = append(te.Enrollments, e)
	return te
}

// AddEvent appends an event to an enrollment.
func (e *NestedEnrollment) AddEvent(event NestedEvent) *NestedEnrollment {
	e.Events = append(e.Events, event)
	return e
}

// AddDataValue appends a data value to the event.
func (ev *NestedEvent) AddDataValue(de, value string) *NestedEvent {
	ev.DataValues = append(ev.DataValues, schema.TrackerDataValue{
		DataElement: utils.StringPtr(de),
		Value:       utils.StringPtr(value),
	})
	return ev
}
