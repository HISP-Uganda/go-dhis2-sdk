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

// ApiTokenAttributesInner - struct for ApiTokenAttributesInner
type ApiTokenAttributesInner struct {
	IpAllowedList      *IpAllowedList
	MethodAllowedList  *MethodAllowedList
	RefererAllowedList *RefererAllowedList
}

// IpAllowedListAsApiTokenAttributesInner is a convenience function that returns IpAllowedList wrapped in ApiTokenAttributesInner
func IpAllowedListAsApiTokenAttributesInner(v *IpAllowedList) ApiTokenAttributesInner {
	return ApiTokenAttributesInner{
		IpAllowedList: v,
	}
}

// MethodAllowedListAsApiTokenAttributesInner is a convenience function that returns MethodAllowedList wrapped in ApiTokenAttributesInner
func MethodAllowedListAsApiTokenAttributesInner(v *MethodAllowedList) ApiTokenAttributesInner {
	return ApiTokenAttributesInner{
		MethodAllowedList: v,
	}
}

// RefererAllowedListAsApiTokenAttributesInner is a convenience function that returns RefererAllowedList wrapped in ApiTokenAttributesInner
func RefererAllowedListAsApiTokenAttributesInner(v *RefererAllowedList) ApiTokenAttributesInner {
	return ApiTokenAttributesInner{
		RefererAllowedList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApiTokenAttributesInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IpAllowedList
	err = newStrictDecoder(data).Decode(&dst.IpAllowedList)
	if err == nil {
		jsonIpAllowedList, _ := json.Marshal(dst.IpAllowedList)
		if string(jsonIpAllowedList) == "{}" { // empty struct
			dst.IpAllowedList = nil
		} else {
			match++
		}
	} else {
		dst.IpAllowedList = nil
	}

	// try to unmarshal data into MethodAllowedList
	err = newStrictDecoder(data).Decode(&dst.MethodAllowedList)
	if err == nil {
		jsonMethodAllowedList, _ := json.Marshal(dst.MethodAllowedList)
		if string(jsonMethodAllowedList) == "{}" { // empty struct
			dst.MethodAllowedList = nil
		} else {
			match++
		}
	} else {
		dst.MethodAllowedList = nil
	}

	// try to unmarshal data into RefererAllowedList
	err = newStrictDecoder(data).Decode(&dst.RefererAllowedList)
	if err == nil {
		jsonRefererAllowedList, _ := json.Marshal(dst.RefererAllowedList)
		if string(jsonRefererAllowedList) == "{}" { // empty struct
			dst.RefererAllowedList = nil
		} else {
			match++
		}
	} else {
		dst.RefererAllowedList = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IpAllowedList = nil
		dst.MethodAllowedList = nil
		dst.RefererAllowedList = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ApiTokenAttributesInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ApiTokenAttributesInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApiTokenAttributesInner) MarshalJSON() ([]byte, error) {
	if src.IpAllowedList != nil {
		return json.Marshal(&src.IpAllowedList)
	}

	if src.MethodAllowedList != nil {
		return json.Marshal(&src.MethodAllowedList)
	}

	if src.RefererAllowedList != nil {
		return json.Marshal(&src.RefererAllowedList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApiTokenAttributesInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IpAllowedList != nil {
		return obj.IpAllowedList
	}

	if obj.MethodAllowedList != nil {
		return obj.MethodAllowedList
	}

	if obj.RefererAllowedList != nil {
		return obj.RefererAllowedList
	}

	// all schemas are nil
	return nil
}

type NullableApiTokenAttributesInner struct {
	value *ApiTokenAttributesInner
	isSet bool
}

func (v NullableApiTokenAttributesInner) Get() *ApiTokenAttributesInner {
	return v.value
}

func (v *NullableApiTokenAttributesInner) Set(val *ApiTokenAttributesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTokenAttributesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTokenAttributesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTokenAttributesInner(val *ApiTokenAttributesInner) *NullableApiTokenAttributesInner {
	return &NullableApiTokenAttributesInner{value: val, isSet: true}
}

func (v NullableApiTokenAttributesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTokenAttributesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
