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
