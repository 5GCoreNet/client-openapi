/*
Unified Data Repository Service API file for Application Data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Application_Data

import (
	"encoding/json"
	"fmt"
)

// AccessRightStatus Possible values are: - FULLY_ALLOWED: The User is fully allowed to access to the channel. - PREVIEW_ALLOWED: The User is preview allowed to access to the channel. - NO_ALLOWED: The User is not allowed to access to the channel. 
type AccessRightStatus struct {
	AccessRightStatusAnyOf *AccessRightStatusAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AccessRightStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AccessRightStatusAnyOf
	err = json.Unmarshal(data, &dst.AccessRightStatusAnyOf);
	if err == nil {
		jsonAccessRightStatusAnyOf, _ := json.Marshal(dst.AccessRightStatusAnyOf)
		if string(jsonAccessRightStatusAnyOf) == "{}" { // empty struct
			dst.AccessRightStatusAnyOf = nil
		} else {
			return nil // data stored in dst.AccessRightStatusAnyOf, return on the first match
		}
	} else {
		dst.AccessRightStatusAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AccessRightStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AccessRightStatus) MarshalJSON() ([]byte, error) {
	if src.AccessRightStatusAnyOf != nil {
		return json.Marshal(&src.AccessRightStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAccessRightStatus struct {
	value *AccessRightStatus
	isSet bool
}

func (v NullableAccessRightStatus) Get() *AccessRightStatus {
	return v.value
}

func (v *NullableAccessRightStatus) Set(val *AccessRightStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRightStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRightStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRightStatus(val *AccessRightStatus) *NullableAccessRightStatus {
	return &NullableAccessRightStatus{value: val, isSet: true}
}

func (v NullableAccessRightStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRightStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


