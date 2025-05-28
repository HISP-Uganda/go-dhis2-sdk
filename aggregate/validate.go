package aggregate

import (
	"encoding/json"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

// ValidateDataValueSetPayload validates the DataValueSetPayload Go struct against the JSON Schema.
func ValidateDataValueSetPayload(payload *DataValueSetPayload) error {
	// 1. Marshal the Go struct to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	// 2. Load the schema and JSON as gojsonschema loaders
	schemaLoader := gojsonschema.NewBytesLoader(dataValueSetSchema)
	documentLoader := gojsonschema.NewBytesLoader(payloadBytes)

	// 3. Perform validation
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("schema validation error: %w", err)
	}

	// 4. Report any validation errors
	if !result.Valid() {
		msg := "Payload validation errors:\n"
		for _, e := range result.Errors() {
			msg += fmt.Sprintf(" - %s\n", e)
		}
		return fmt.Errorf(msg)
	}

	return nil
}
