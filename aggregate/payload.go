package aggregate

import "github.com/HISP-Uganda/go-dhis2-sdk/dhis2/schema"

// DataValueSetPayload is the full payload submitted to DHIS2 for aggregate data.
type DataValueSetPayload struct {
	DataSet              string             `json:"dataSet"`
	CompleteDate         string             `json:"completeDate"` // Format: yyyy-MM-dd
	Period               string             `json:"period"`
	OrgUnit              string             `json:"orgUnit"`
	AttributeOptionCombo string             `json:"attributeOptionCombo,omitempty"`
	DataValues           []schema.DataValue `json:"dataValues"`
}

func (p *DataValueSetPayload) Validate() error {
	return ValidateDataValueSetPayload(p)
}
