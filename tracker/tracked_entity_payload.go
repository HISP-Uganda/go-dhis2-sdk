package tracker

import "github.com/HISP-Uganda/go-dhis2-sdk/dhis2/schema"

type TrackedEntityUpdatePayload struct {
	TrackedEntityInstance string                       `json:"trackedEntityInstance"`
	OrgUnit               string                       `json:"orgUnit"`
	TrackedEntityType     string                       `json:"trackedEntityType"`
	Attributes            []schema.AttributeValue      `json:"attributes"`
	Relationships         []schema.TrackerRelationship `json:"relationships,omitempty"`
}
