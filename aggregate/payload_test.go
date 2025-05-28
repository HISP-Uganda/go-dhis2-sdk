package aggregate

import (
	"github.com/HISP-Uganda/go-dhis2-sdk/dhis2/schema"
	"github.com/HISP-Uganda/go-dhis2-sdk/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDataValueSetPayload_Valid(t *testing.T) {
	payload := &DataValueSetPayload{
		DataSet:      "pBOMPrpg1QX",
		Period:       "202404",
		OrgUnit:      "DiszpKrYNg8",
		CompleteDate: "2024-04-30",
		DataValues: []schema.DataValue{
			{
				DataElement:         utils.StringPtr("f7n9E0hX8qk"),
				CategoryOptionCombo: utils.StringPtr("HllvX50cXC0"),
				Value:               utils.StringPtr("42"),
			},
		},
	}

	err := payload.Validate()
	assert.NoError(t, err)
}

func TestValidateDataValueSetPayload_MissingRequired(t *testing.T) {
	payload := &DataValueSetPayload{
		// Missing DataSet, Period, OrgUnit, CompleteDate
		DataValues: []schema.DataValue{
			{
				DataElement: utils.StringPtr("abc123"),
				Value:       utils.StringPtr("99"),
			},
		},
	}

	err := payload.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Payload validation errors")
}

func TestValidateDataValueSetPayload_EmptyDataValues(t *testing.T) {
	payload := &DataValueSetPayload{
		DataSet:      "datasetUID",
		Period:       "202404",
		OrgUnit:      "orgUnitUID",
		CompleteDate: "2024-04-30",
		DataValues:   []schema.DataValue{}, // Empty dataValues
	}

	err := payload.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Payload validation errors")
}

func TestValidateDataValueSetPayload_InvalidDateFormat(t *testing.T) {
	payload := &DataValueSetPayload{
		DataSet:      "datasetUID",
		Period:       "202404",
		OrgUnit:      "orgUnitUID",
		CompleteDate: "30-04-2024", // Invalid format (should be yyyy-MM-dd)
		DataValues: []schema.DataValue{
			{
				DataElement: utils.StringPtr("deUID"),
				Value:       utils.StringPtr("17"),
			},
		},
	}

	err := payload.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Payload validation errors")
}
