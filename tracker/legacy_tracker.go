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
	ResponseType    string          `json:"responseType,omitempty"`
	Status          string          `json:"status"`
	Imported        int             `json:"imported,omitempty"`
	Updated         int             `json:"updated,omitempty"`
	Deleted         int             `json:"deleted,omitempty"`
	Ignored         int             `json:"ignored,omitempty"`
	ImportSummaries []ImportSummary `json:"importSummaries,omitempty"`
	Total           int             `json:"total,omitempty"`
}

type ImportOptions struct {
	IdSchemes                   map[string]interface{} `json:"idSchemes,omitempty"`
	DryRun                      bool                   `json:"dryRun,omitempty"`
	Async                       bool                   `json:"async,omitempty"`
	ImportStrategy              string                 `json:"importStrategy,omitempty"`
	MergeMode                   string                 `json:"mergeMode,omitempty"`
	ReportMode                  string                 `json:"reportMode,omitempty"`
	SkipExistingCheck           bool                   `json:"skipExistingCheck,omitempty"`
	Sharing                     bool                   `json:"sharing,omitempty"`
	SkipNotifications           bool                   `json:"skipNotifications,omitempty"`
	SkipAudit                   bool                   `json:"skipAudit,omitempty"`
	DatasetAllowsPeriods        bool                   `json:"datasetAllowsPeriods,omitempty"`
	StrictPeriods               bool                   `json:"strictPeriods,omitempty"`
	StrictDataElements          bool                   `json:"strictDataElements,omitempty"`
	StrictCategoryOptionCombos  bool                   `json:"strictCategoryOptionCombos,omitempty"`
	StrictAttributeOptionCombos bool                   `json:"strictAttributeOptionCombos,omitempty"`
	StrictOrganisationUnits     bool                   `json:"strictOrganisationUnits,omitempty"`
	StrictDataSetApproval       bool                   `json:"strictDataSetApproval,omitempty"`
	StrictDataSetLocking        bool                   `json:"strictDataSetLocking,omitempty"`
	StrictDataSetInputPeriods   bool                   `json:"strictDataSetInputPeriods,omitempty"`
	RequireCategoryOptionCombo  bool                   `json:"requireCategoryOptionCombo,omitempty"`
	RequireAttributeOptionCombo bool                   `json:"requireAttributeOptionCombo,omitempty"`
	SkipPatternValidation       bool                   `json:"skipPatternValidation,omitempty"`
	IgnoreEmptyCollection       bool                   `json:"ignoreEmptyCollection,omitempty"`
	Force                       bool                   `json:"force,omitempty"`
	FirstRowIsHeader            bool                   `json:"firstRowIsHeader,omitempty"`
	SkipLastUpdated             bool                   `json:"skipLastUpdated,omitempty"`
	MergeDataValues             bool                   `json:"mergeDataValues,omitempty"`
	SkipCache                   bool                   `json:"skipCache,omitempty"`
}

type EventResponse struct {
	ResponseType    string `json:"responseType,omitempty"`
	Status          string `json:"status,omitempty"`
	Imported        int    `json:"imported,omitempty"`
	Updated         int    `json:"updated,omitempty"`
	Deleted         int    `json:"deleted,omitempty"`
	Ignored         int    `json:"ignored,omitempty"`
	ImportSummaries []struct {
		ResponseType string         `json:"responseType,omitempty"`
		Status       string         `json:"status,omitempty"`
		ImportCount  ImportCount    `json:"importCount,omitempty"`
		Conflicts    ImportConflict `json:"importConflict,omitempty"`
	} `json:"importSummaries,omitempty"`
}

type EnrollmentResponse struct {
	ResponseType    string `json:"responseType,omitempty"`
	Status          string `json:"status,omitempty"`
	Imported        int    `json:"imported,omitempty"`
	Updated         int    `json:"updated,omitempty"`
	Deleted         int    `json:"deleted,omitempty"`
	Ignored         int    `json:"ignored,omitempty"`
	ImportSummaries []struct {
		ResponseType string         `json:"responseType,omitempty"`
		Status       string         `json:"status,omitempty"`
		ImportCount  ImportCount    `json:"importCount,omitempty"`
		Conflicts    ImportConflict `json:"importConflict,omitempty"`
		Events       EventResponse
	} `json:"importSummaries,omitempty"`
}

var t schema.ImportSummary

type ImportSummary struct {
	ResponseType    string              `json:"responseType,omitempty"`
	Status          string              `json:"status,omitempty"`
	ImportOptions   ImportOptions       `json:"importOptions,omitempty"`
	ImportCount     ImportCount         `json:"importCount,omitempty"`
	Conflicts       []ImportConflict    `json:"conflicts,omitempty"`
	RejectedIndexes []any               `json:"rejectedIndexes,omitempty"`
	Reference       string              `json:"reference,omitempty"`
	Href            string              `json:"href,omitempty"`
	Enrollments     *EnrollmentResponse `json:"enrollments,omitempty"`
}

// ImportCount represents the count of imported, updated, ignored, and deleted items.
type ImportCount struct {
	Imported int `json:"imported,omitempty"`
	Updated  int `json:"updated,omitempty"`
	Ignored  int `json:"ignored,omitempty"`
	Deleted  int `json:"deleted,omitempty"`
}

type ImportConflict struct {
	Indexes  []int32            `json:"indexes,omitempty"`
	Object   *string            `json:"object,omitempty"`
	Objects  *map[string]string `json:"objects,omitempty"`
	Property *string            `json:"property,omitempty"`
	Value    *string            `json:"value,omitempty"`
}
