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

// DataSetGetObjectListGistGetObjectListGistAsCsv200Response - struct for DataSetGetObjectListGistGetObjectListGistAsCsv200Response
type DataSetGetObjectListGistGetObjectListGistAsCsv200Response struct {
	DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf *DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	ArrayOfDataSet                                                 *[]DataSet
}

// DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsDataSetGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf wrapped in DataSetGetObjectListGistGetObjectListGistAsCsv200Response
func DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsDataSetGetObjectListGistGetObjectListGistAsCsv200Response(v *DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) DataSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return DataSetGetObjectListGistGetObjectListGistAsCsv200Response{
		DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf: v,
	}
}

// []DataSetAsDataSetGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns []DataSet wrapped in DataSetGetObjectListGistGetObjectListGistAsCsv200Response
func ArrayOfDataSetAsDataSetGetObjectListGistGetObjectListGistAsCsv200Response(v *[]DataSet) DataSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return DataSetGetObjectListGistGetObjectListGistAsCsv200Response{
		ArrayOfDataSet: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DataSetGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	if err == nil {
		jsonDataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf, _ := json.Marshal(dst.DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
		if string(jsonDataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) == "{}" { // empty struct
			dst.DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
	}

	// try to unmarshal data into ArrayOfDataSet
	err = newStrictDecoder(data).Decode(&dst.ArrayOfDataSet)
	if err == nil {
		jsonArrayOfDataSet, _ := json.Marshal(dst.ArrayOfDataSet)
		if string(jsonArrayOfDataSet) == "{}" { // empty struct
			dst.ArrayOfDataSet = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfDataSet = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		dst.ArrayOfDataSet = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DataSetGetObjectListGistGetObjectListGistAsCsv200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DataSetGetObjectListGistGetObjectListGistAsCsv200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DataSetGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	if src.DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return json.Marshal(&src.DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	}

	if src.ArrayOfDataSet != nil {
		return json.Marshal(&src.ArrayOfDataSet)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DataSetGetObjectListGistGetObjectListGistAsCsv200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return obj.DataSetGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	}

	if obj.ArrayOfDataSet != nil {
		return obj.ArrayOfDataSet
	}

	// all schemas are nil
	return nil
}

type NullableDataSetGetObjectListGistGetObjectListGistAsCsv200Response struct {
	value *DataSetGetObjectListGistGetObjectListGistAsCsv200Response
	isSet bool
}

func (v NullableDataSetGetObjectListGistGetObjectListGistAsCsv200Response) Get() *DataSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return v.value
}

func (v *NullableDataSetGetObjectListGistGetObjectListGistAsCsv200Response) Set(val *DataSetGetObjectListGistGetObjectListGistAsCsv200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSetGetObjectListGistGetObjectListGistAsCsv200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSetGetObjectListGistGetObjectListGistAsCsv200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSetGetObjectListGistGetObjectListGistAsCsv200Response(val *DataSetGetObjectListGistGetObjectListGistAsCsv200Response) *NullableDataSetGetObjectListGistGetObjectListGistAsCsv200Response {
	return &NullableDataSetGetObjectListGistGetObjectListGistAsCsv200Response{value: val, isSet: true}
}

func (v NullableDataSetGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSetGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
