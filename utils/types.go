package utils

import "github.com/HISP-Uganda/go-dhis2-sdk/dhis2/schema"

func StringPtr(s string) *string {
	return &s
}

func IntPtr(i int) *int {
	return &i
}

func BoolPtr(b bool) *bool {
	return &b
}

// InstantToString returns the Instant as a yyyy-MM-dd formatted string.
func InstantToString(i *schema.Instant) string {
	if i == nil || i.TimeTime == nil {
		return ""
	}
	return i.TimeTime.Format("2006-01-02")
}
