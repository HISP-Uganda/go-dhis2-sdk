package aggregate

import "github.com/HISP-Uganda/go-dhis2-sdk/dhis2/schema"

type ImportSummaryResponse struct {
	ResponseType  string               `json:"responseType"`
	Status        string               `json:"status"`
	ImportOptions schema.ImportOptions `json:"importOptions"`

	Description     string                  `json:"description"`
	ImportCount     schema.ImportCount      `json:"importCount"`
	Conflicts       []schema.ImportConflict `json:"conflicts,omitempty"`
	DataSetComplete string                  `json:"dataSetComplete,omitempty"`
}

// HasWarnings returns true if the status is WARNING or if there are any conflicts.
func (r *ImportSummaryResponse) HasWarnings() bool {
	return r.Status == "WARNING" || len(r.Conflicts) > 0
}

// HasErrors returns true if the status is ERROR.
func (r *ImportSummaryResponse) HasErrors() bool {
	return r.Status == "ERROR"
}

// IsSuccess returns true if the status is SUCCESS and there are no conflicts.
func (r *ImportSummaryResponse) IsSuccess() bool {
	return r.Status == "SUCCESS" && len(r.Conflicts) == 0
}

// ConflictMessages returns a slice of all conflict messages.
func (r *ImportSummaryResponse) ConflictMessages() []string {
	messages := make([]string, 0, len(r.Conflicts))
	for _, conflict := range r.Conflicts {
		messages = append(messages, *conflict.Value)
	}
	return messages
}
