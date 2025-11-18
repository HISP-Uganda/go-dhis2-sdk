package tracker

import (
	"time"

	"github.com/HISP-Uganda/go-dhis2-sdk/dhis2/schema"
)

type LegacyAsyncResponse struct {
	HttpStatus     string `json:"httpStatus,omitempty"`
	HttpStatusCode int    `json:"httpStatusCode,omitempty"`
	Status         string `json:"status,omitempty"`
	Message        string `json:"message,omitempty"`
	Response       *struct {
		ID        *string    `json:"id"`
		Name      *string    `json:"name"`
		Created   *time.Time `json:"created"`
		JobType   *string    `json:"jobType"`
		JobStatus *string    `json:"jobStatus"`
	} `json:"response,omitempty"`
}

type LegacyTrackerResponse struct {
	SyncResponse  RootResponse
	AsyncResponse LegacyAsyncResponse
}

type RootResponse struct {
	HttpStatus     string   `json:"httpStatus,omitempty"`
	HttpStatusCode int      `json:"httpStatusCode,omitempty"`
	Status         string   `json:"status,omitempty"`
	Message        string   `json:"message,omitempty"`
	Response       Response `json:"response,omitempty"`
}

// Response represents the main response data.
type Response struct {
	ResponseType    string                  `json:"responseType,omitempty"`
	Status          string                  `json:"status"`
	Imported        int                     `json:"imported,omitempty"`
	Updated         int                     `json:"updated,omitempty"`
	Deleted         int                     `json:"deleted,omitempty"`
	Ignored         int                     `json:"ignored,omitempty"`
	ImportCount     schema.ImportCount      `json:"importCount,omitempty"`
	Conflicts       []schema.ImportConflict `json:"conflicts,omitempty"`
	ImportOptions   schema.ImportOptions    `json:"importOptions,omitempty"`
	ImportSummaries []schema.ImportSummary  `json:"importSummaries,omitempty"`
	Total           int                     `json:"total,omitempty"`
}
