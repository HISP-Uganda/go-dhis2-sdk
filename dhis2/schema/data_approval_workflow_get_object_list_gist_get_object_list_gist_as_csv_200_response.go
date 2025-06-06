/*
DHIS2 API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.42
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
	"fmt"
)

// DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response - struct for DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response
type DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response struct {
	DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf *DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	ArrayOfDataApprovalWorkflow                                                 *[]DataApprovalWorkflow
}

// DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf wrapped in DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response
func DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response(v *DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response {
	return DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response{
		DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf: v,
	}
}

// []DataApprovalWorkflowAsDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns []DataApprovalWorkflow wrapped in DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response
func ArrayOfDataApprovalWorkflowAsDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response(v *[]DataApprovalWorkflow) DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response {
	return DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response{
		ArrayOfDataApprovalWorkflow: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	if err == nil {
		jsonDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf, _ := json.Marshal(dst.DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
		if string(jsonDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) == "{}" { // empty struct
			dst.DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
	}

	// try to unmarshal data into ArrayOfDataApprovalWorkflow
	err = newStrictDecoder(data).Decode(&dst.ArrayOfDataApprovalWorkflow)
	if err == nil {
		jsonArrayOfDataApprovalWorkflow, _ := json.Marshal(dst.ArrayOfDataApprovalWorkflow)
		if string(jsonArrayOfDataApprovalWorkflow) == "{}" { // empty struct
			dst.ArrayOfDataApprovalWorkflow = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfDataApprovalWorkflow = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		dst.ArrayOfDataApprovalWorkflow = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	if src.DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return json.Marshal(&src.DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	}

	if src.ArrayOfDataApprovalWorkflow != nil {
		return json.Marshal(&src.ArrayOfDataApprovalWorkflow)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return obj.DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	}

	if obj.ArrayOfDataApprovalWorkflow != nil {
		return obj.ArrayOfDataApprovalWorkflow
	}

	// all schemas are nil
	return nil
}

type NullableDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response struct {
	value *DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response
	isSet bool
}

func (v NullableDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response) Get() *DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response {
	return v.value
}

func (v *NullableDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response) Set(val *DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response(val *DataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response) *NullableDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response {
	return &NullableDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response{value: val, isSet: true}
}

func (v NullableDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataApprovalWorkflowGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
