package dhis2error

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

// ParseAPIError
// example
// resp, err := client.Post("/api/dataValueSets", payload)
// if err != nil {
// 	return nil, err
// }
// if resp.IsError() {
// 	return nil, dhis2error.ParseAPIError(resp)
// }

func ParseAPIError(resp *resty.Response) *APIError {
	var payload map[string]any
	_ = json.Unmarshal(resp.Body(), &payload)

	message := ""
	if msg, ok := payload["message"].(string); ok {
		message = msg
	}
	return NewAPIError(resp.StatusCode(), message, payload)
}
