package tracker

import (
	"fmt"
	"strings"
	"time"
)

// CorrectLayout matches "YYYY-MM-DDThh:mm:ss.sss" (e.g., 2025-11-18T05:42:09.071)
const CorrectLayout = "2006-01-02T15:04:05.000"

// DHIS2Time is a wrapper around time.Time that provides custom JSON (un)marshalling
// for the DHIS2 DateTime format lacking a timezone offset.
type DHIS2Time struct {
	time.Time
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (dt *DHIS2Time) UnmarshalJSON(b []byte) (err error) {
	// 1. Convert byte slice to string and remove surrounding quotes
	s := strings.Trim(string(b), `"`)

	// Handle null or empty string case
	if s == "" || s == "null" {
		dt.Time = time.Time{}
		return nil
	}

	// 2. Parse the time using the custom layout
	t, err := time.Parse(CorrectLayout, s)
	if err != nil {
		return fmt.Errorf("parsing DHIS2 time %q as %q: %w", s, CorrectLayout, err)
	}

	dt.Time = t
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (dt DHIS2Time) MarshalJSON() ([]byte, error) {
	// If the time is zero (not set), marshal it as an empty JSON string.
	if dt.Time.IsZero() {
		return []byte(`""`), nil
	}

	// Format the time using the custom layout and wrap the result in quotes.
	formattedTime := dt.Time.Format(CorrectLayout)
	return []byte(`"` + formattedTime + `"`), nil
}
