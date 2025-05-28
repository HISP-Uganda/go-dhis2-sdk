package tracker

import (
	"github.com/HISP-Uganda/go-dhis2-sdk/dhis2/schema"
	"github.com/HISP-Uganda/go-dhis2-sdk/tracker"
	"time"
)

// NewNestedTrackedEntity creates a tracked entity with basic identity fields.
func NewNestedTrackedEntity(teiType, orgUnit string) *tracker.NestedTrackedEntity {
	return &tracker.NestedTrackedEntity{
		TrackedEntityType: teiType,
		OrgUnit:           orgUnit,
		Attributes:        []schema.TrackerAttribute{},
		Enrollments:       []tracker.NestedEnrollment{},
	}
}

// NewNestedEnrollment builds a TrackerEnrollment with the required fields.
func NewNestedEnrollment(status, orgUnit, program, enrolledAt string) *tracker.NestedEnrollment {
	// Parse the enrollment date string into a time.Time object
	enrolledAtTime, err := time.Parse("2006-01-02", enrolledAt)
	if err != nil {
		// set enrolledAtTime to today's date if parsing fails
		enrolledAtTime = time.Now()
	}
	return &tracker.NestedEnrollment{
		Status:     status,
		OrgUnit:    orgUnit,
		Program:    program,
		EnrolledAt: enrolledAtTime,
		Events:     []tracker.NestedEvent{},
	}
}

// NewNestedEvent constructs a TrackerEvent with required fields.
func NewNestedEvent(program, programStage, orgUnit, occurredAt string) *tracker.NestedEvent {
	occurredAtAtTime, err := time.Parse("2006-01-02", occurredAt)
	if err != nil {
		// set occurredAtTime to today's date if parsing fails
		occurredAtAtTime = time.Now()
	}
	return &tracker.NestedEvent{
		Program:      program,
		ProgramStage: programStage,
		OrgUnit:      orgUnit,
		OccurredAt:   occurredAtAtTime,
		DataValues:   []schema.TrackerDataValue{},
	}
}
