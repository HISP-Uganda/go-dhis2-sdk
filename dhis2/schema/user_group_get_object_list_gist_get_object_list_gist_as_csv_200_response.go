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

// UserGroupGetObjectListGistGetObjectListGistAsCsv200Response - struct for UserGroupGetObjectListGistGetObjectListGistAsCsv200Response
type UserGroupGetObjectListGistGetObjectListGistAsCsv200Response struct {
	UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf *UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	ArrayOfUserGroup                                                 *[]UserGroup
}

// UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsUserGroupGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf wrapped in UserGroupGetObjectListGistGetObjectListGistAsCsv200Response
func UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOfAsUserGroupGetObjectListGistGetObjectListGistAsCsv200Response(v *UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) UserGroupGetObjectListGistGetObjectListGistAsCsv200Response {
	return UserGroupGetObjectListGistGetObjectListGistAsCsv200Response{
		UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf: v,
	}
}

// []UserGroupAsUserGroupGetObjectListGistGetObjectListGistAsCsv200Response is a convenience function that returns []UserGroup wrapped in UserGroupGetObjectListGistGetObjectListGistAsCsv200Response
func ArrayOfUserGroupAsUserGroupGetObjectListGistGetObjectListGistAsCsv200Response(v *[]UserGroup) UserGroupGetObjectListGistGetObjectListGistAsCsv200Response {
	return UserGroupGetObjectListGistGetObjectListGistAsCsv200Response{
		ArrayOfUserGroup: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UserGroupGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	if err == nil {
		jsonUserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf, _ := json.Marshal(dst.UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
		if string(jsonUserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf) == "{}" { // empty struct
			dst.UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
	}

	// try to unmarshal data into ArrayOfUserGroup
	err = newStrictDecoder(data).Decode(&dst.ArrayOfUserGroup)
	if err == nil {
		jsonArrayOfUserGroup, _ := json.Marshal(dst.ArrayOfUserGroup)
		if string(jsonArrayOfUserGroup) == "{}" { // empty struct
			dst.ArrayOfUserGroup = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfUserGroup = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf = nil
		dst.ArrayOfUserGroup = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UserGroupGetObjectListGistGetObjectListGistAsCsv200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UserGroupGetObjectListGistGetObjectListGistAsCsv200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UserGroupGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	if src.UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return json.Marshal(&src.UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf)
	}

	if src.ArrayOfUserGroup != nil {
		return json.Marshal(&src.ArrayOfUserGroup)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UserGroupGetObjectListGistGetObjectListGistAsCsv200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf != nil {
		return obj.UserGroupGetObjectListGistGetObjectListGistAsCsv200ResponseOneOf
	}

	if obj.ArrayOfUserGroup != nil {
		return obj.ArrayOfUserGroup
	}

	// all schemas are nil
	return nil
}

type NullableUserGroupGetObjectListGistGetObjectListGistAsCsv200Response struct {
	value *UserGroupGetObjectListGistGetObjectListGistAsCsv200Response
	isSet bool
}

func (v NullableUserGroupGetObjectListGistGetObjectListGistAsCsv200Response) Get() *UserGroupGetObjectListGistGetObjectListGistAsCsv200Response {
	return v.value
}

func (v *NullableUserGroupGetObjectListGistGetObjectListGistAsCsv200Response) Set(val *UserGroupGetObjectListGistGetObjectListGistAsCsv200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupGetObjectListGistGetObjectListGistAsCsv200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupGetObjectListGistGetObjectListGistAsCsv200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupGetObjectListGistGetObjectListGistAsCsv200Response(val *UserGroupGetObjectListGistGetObjectListGistAsCsv200Response) *NullableUserGroupGetObjectListGistGetObjectListGistAsCsv200Response {
	return &NullableUserGroupGetObjectListGistGetObjectListGistAsCsv200Response{value: val, isSet: true}
}

func (v NullableUserGroupGetObjectListGistGetObjectListGistAsCsv200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupGetObjectListGistGetObjectListGistAsCsv200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
