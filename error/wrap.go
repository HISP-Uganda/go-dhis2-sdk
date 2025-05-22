package dhis2error

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized access to DHIS2")
	ErrNotFound     = errors.New("resource not found")
	ErrConflict     = errors.New("conflict when saving resource")
)

func WrapStatusCode(code int) error {
	switch code {
	case 401:
		return ErrUnauthorized
	case 404:
		return ErrNotFound
	case 409:
		return ErrConflict
	default:
		return nil
	}
}
