package tracker

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

//go:embed schema/tracker_payload_example.json
var trackerPayloadExample []byte

//go:embed schema/tracker_nested_schema.json
var trackerNestedSchema []byte

func TestTrackerNestedPayload_Valid(t *testing.T) {
	schemaLoader := gojsonschema.NewBytesLoader(trackerNestedSchema)
	documentLoader := gojsonschema.NewBytesLoader(trackerPayloadExample)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	assert.NoError(t, err)
	assert.True(t, result.Valid(), "Expected the tracker payload to be valid")

	if !result.Valid() {
		for _, e := range result.Errors() {
			t.Logf("Validation error: %s", e)
		}
		t.FailNow()
	}
}

func TestTrackerNestedPayload_Invalid(t *testing.T) {
	// Modify a copy of the example to omit required field
	var data map[string]interface{}
	err := json.Unmarshal(trackerPayloadExample, &data)
	assert.NoError(t, err)

	// Remove occurredAt from event
	entities := data["trackedEntities"].([]interface{})
	enrollments := entities[0].(map[string]interface{})["enrollments"].([]interface{})
	events := enrollments[0].(map[string]interface{})["events"].([]interface{})
	delete(events[0].(map[string]interface{}), "occurredAt")

	invalidBytes, err := json.Marshal(data)
	assert.NoError(t, err)

	schemaLoader := gojsonschema.NewBytesLoader(trackerNestedSchema)
	documentLoader := gojsonschema.NewBytesLoader(invalidBytes)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	assert.NoError(t, err)
	assert.False(t, result.Valid(), "Expected the modified payload to be invalid")
}
