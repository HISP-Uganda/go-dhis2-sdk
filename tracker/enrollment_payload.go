package tracker

type EnrollmentPayload struct {
	Program               string `json:"program"`
	Status                string `json:"status"`
	OrgUnit               string `json:"orgUnit"`
	EnrollmentDate        string `json:"enrollmentDate"`
	IncidentDate          string `json:"incidentDate"`
	TrackedEntityInstance string `json:"trackedEntityInstance"`
}
