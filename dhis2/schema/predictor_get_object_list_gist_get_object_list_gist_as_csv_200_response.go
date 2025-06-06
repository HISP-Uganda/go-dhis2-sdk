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

// PredictorGetObjectListGistGetObjectListGistAsCsv200Response - struct for PredictorGetObjectListGistGetObjectListGistAsCsv200Response
type PredictorGetObjectListGistGetObjectListGistAsCsv200Response struct {
	PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf *PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	ArrayOfPredictor                                                 *[]Predictor
}

// PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsPredictorGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf wrapped in PredictorGetObjectListGistGetObjectListGistAsCsv200Response
func PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsPredictorGetObjectListGistGetObjectListGistAsCsv200Response(v *PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) PredictorGetObjectListGistGetObjectListGistAsCsv200Response {
	return PredictorGetObjectListGistGetObjectListGistAsCsv200Response{
		PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf: v,
	}
}

// []PredictorAsPredictorGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns []Predictor wrapped in PredictorGetObjectListGistGetObjectListGistAsCsv200Response
func ArrayOfPredictorAsPredictorGetObjectListGistGetObjectListGistAsCsv200Response(v *[]Predictor) PredictorGetObjectListGistGetObjectListGistAsCsv200Response {
	return PredictorGetObjectListGistGetObjectListGistAsCsv200Response{
		ArrayOfPredictor: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PredictorGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	if err == nil {
		jsonPredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf, _ := json.Marshal(dst.PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
		if string(jsonPredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) == "{}" { // empty struct
			dst.PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
	}

	// try to unmarshal data into ArrayOfPredictor
	err = newStrictDecoder(data).Decode(&dst.ArrayOfPredictor)
	if err == nil {
		jsonArrayOfPredictor, _ := json.Marshal(dst.ArrayOfPredictor)
		if string(jsonArrayOfPredictor) == "{}" { // empty struct
			dst.ArrayOfPredictor = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfPredictor = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		dst.ArrayOfPredictor = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PredictorGetObjectListGistGetObjectListGistAsCsv200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PredictorGetObjectListGistGetObjectListGistAsCsv200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PredictorGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	if src.PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return json.Marshal(&src.PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	}

	if src.ArrayOfPredictor != nil {
		return json.Marshal(&src.ArrayOfPredictor)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PredictorGetObjectListGistGetObjectListGistAsCsv200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return obj.PredictorGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	}

	if obj.ArrayOfPredictor != nil {
		return obj.ArrayOfPredictor
	}

	// all schemas are nil
	return nil
}

type NullablePredictorGetObjectListGistGetObjectListGistAsCsv200Response struct {
	value *PredictorGetObjectListGistGetObjectListGistAsCsv200Response
	isSet bool
}

func (v NullablePredictorGetObjectListGistGetObjectListGistAsCsv200Response) Get() *PredictorGetObjectListGistGetObjectListGistAsCsv200Response {
	return v.value
}

func (v *NullablePredictorGetObjectListGistGetObjectListGistAsCsv200Response) Set(val *PredictorGetObjectListGistGetObjectListGistAsCsv200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePredictorGetObjectListGistGetObjectListGistAsCsv200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePredictorGetObjectListGistGetObjectListGistAsCsv200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredictorGetObjectListGistGetObjectListGistAsCsv200Response(val *PredictorGetObjectListGistGetObjectListGistAsCsv200Response) *NullablePredictorGetObjectListGistGetObjectListGistAsCsv200Response {
	return &NullablePredictorGetObjectListGistGetObjectListGistAsCsv200Response{value: val, isSet: true}
}

func (v NullablePredictorGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredictorGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
