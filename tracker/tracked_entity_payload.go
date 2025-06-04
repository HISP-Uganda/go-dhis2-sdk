package tracker

import "github.com/HISP-Uganda/go-dhis2-sdk/dhis2/schema"

type TrackedEntityUpdatePayload struct {
	TrackedEntityInstance *string                      `json:"trackedEntityInstance,omitempty"` // legacy
	TrackedEntity         *string                      `json:"trackedEntity,omitempty"`
	OrgUnit               string                       `json:"orgUnit"`
	TrackedEntityType     string                       `json:"trackedEntityType"`
	Attributes            []TrackedEntityAttribute     `json:"attributes"`
	Relationships         []schema.TrackerRelationship `json:"relationships,omitempty"`
}

type TrackedEntityDeletePayload struct {
	TrackedEntities []schema.TrackerTrackedEntity `json:"trackedEntities"`
}
