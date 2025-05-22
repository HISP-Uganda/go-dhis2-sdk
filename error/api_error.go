package dhis2error

import "fmt"

type APIError struct {
	StatusCode int
	Message    string
	Details    any
}

func (e *APIError) Error() string {
	return fmt.Sprintf("DHIS2 API error: %d - %s", e.StatusCode, e.Message)
}

func NewAPIError(status int, msg string, details any) *APIError {
	return &APIError{
		StatusCode: status,
		Message:    msg,
		Details:    details,
	}
}

func (e *APIError) LogFields() map[string]any {
	return map[string]any{
		"status_code": e.StatusCode,
		"message":     e.Message,
		"details":     e.Details,
	}
}
