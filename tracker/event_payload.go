package tracker

import (
	"fmt"
	"github.com/HISP-Uganda/go-dhis2-sdk/dhis2/schema"
	"github.com/HISP-Uganda/go-dhis2-sdk/utils"
	"time"
)

// Legacy payloads for DHIS2 Tracker events.
func ParseStringToInstant(dateStr string) (*schema.Instant, error) {
	if dateStr == "" {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date format (expected yyyy-MM-dd): %w", err)
	}
	return &schema.Instant{
		TimeTime: &t,
	}, nil
}

// EventCreatePayload represents the payload for creating a new event in DHIS2 Tracker.
type EventCreatePayload struct {
	TrackedEntityInstance string                    `json:"trackedEntityInstance,omitempty"` // legacy
	Program               string                    `json:"program"`
	ProgramStage          string                    `json:"programStage"`
	Enrollment            string                    `json:"enrollment"`
	OrgUnit               string                    `json:"orgUnit"`
	Status                string                    `json:"status"`
	EventDate             string                    `json:"eventDate"`
	DataValues            []schema.TrackerDataValue `json:"dataValues"`
}

type EventUpdatePayload struct {
	Event string `json:"event"`
	//EventDate     string      `json:"eventDate"`
	Program       string                    `json:"program"`
	OrgUnit       string                    `json:"orgUnit"`
	ProgramStage  string                    `json:"programStage"`
	Status        string                    `json:"status,omitempty"`
	TrackedEntity string                    `json:"trackedEntityInstance,omitempty"`
	DataValues    []schema.TrackerDataValue `json:"dataValues"`
}

func (e *EventCreatePayload) ToTrackerEvent() (*schema.TrackerEvent, error) {
	inst, err := ParseStringToInstant(e.EventDate)
	if err != nil {
		return nil, err
	}

	return &schema.TrackerEvent{
		TrackedEntity: utils.StringPtr(e.TrackedEntityInstance),
		Program:       utils.StringPtr(e.Program),
		ProgramStage:  utils.StringPtr(e.ProgramStage),
		OrgUnit:       utils.StringPtr(e.OrgUnit),
		OccurredAt:    inst,
		Status:        schema.EventStatus(e.Status),
		DataValues:    e.DataValues,
		// Map DataValues if needed
	}, nil
}
