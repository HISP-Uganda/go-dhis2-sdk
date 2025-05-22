package dhis2error

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

func ParseAPIError(resp *resty.Response) *APIError {
	var payload map[string]any
	_ = json.Unmarshal(resp.Body(), &payload)

	message := ""
	if msg, ok := payload["message"].(string); ok {
		message = msg
	}
	return NewAPIError(resp.StatusCode(), message, payload)
}

func ParseImportSummariesError(resp *resty.Response) *APIError {
	var payload struct {
		Response struct {
			Status          string              `json:"status"`
			Conflicts       []map[string]string `json:"conflicts"`
			ImportSummaries []map[string]any    `json:"importSummaries"`
			ImportCount     map[string]int      `json:"importCount"`
		} `json:"response"`
	}

	_ = json.Unmarshal(resp.Body(), &payload)

	msg := "DHIS2 import failed"
	if len(payload.Response.Conflicts) > 0 {
		msg = fmt.Sprintf("Conflicts: %v", payload.Response.Conflicts)
	}

	return NewAPIError(resp.StatusCode(), msg, payload.Response)
}
