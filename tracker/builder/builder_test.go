package tracker

import (
	"encoding/json"
	"fmt"
	"github.com/HISP-Uganda/go-dhis2-sdk/tracker"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuilderFlow(t *testing.T) {
	// Simulated input
	tei := "W3kXO3uKf0p"
	orgUnit := "DiszpKrYNg8"
	program := "IpHINAT79UW"
	programStage := "EJhX3c6sJtl"
	date := time.Now().Format("2006-01-02")

	entity := NewNestedTrackedEntity("MCPQUTHX1Ze", orgUnit).
		AddAttribute("w75KJ2mc4zz", "Jane Doe")

	enrollment := NewNestedEnrollment(tei, orgUnit, program, date)

	event := NewNestedEvent(program, programStage, orgUnit, date).
		AddDataValue("de001", "Positive")
	enrollment.AddEvent(*event)

	entity.AddEnrollment(*enrollment)

	payload := tracker.NestedPayload{
		TrackedEntities: []tracker.NestedTrackedEntity{*entity},
	}

	// Marshal to check output
	out, err := json.MarshalIndent(payload, "", "  ")
	assert.NoError(t, err)
	fmt.Println(string(out))

	// Write to file for inspection
	err = os.WriteFile("tracker_payload_output.json", out, 0644)
	assert.NoError(t, err)

	// Basic assertions
	assert.Equal(t, 1, len(payload.TrackedEntities))
	assert.Equal(t, 1, len(payload.TrackedEntities[0].Enrollments))
	assert.Equal(t, 1, len(payload.TrackedEntities[0].Enrollments[0].Events))
}
