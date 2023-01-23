/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
	"fmt"
)

// ManagementData - struct for ManagementData
type ManagementData struct {
	ArrayOfString *[]string
}

// []stringAsManagementData is a convenience function that returns []string wrapped in ManagementData
func ArrayOfStringAsManagementData(v *[]string) ManagementData {
	return ManagementData{
		ArrayOfString: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ManagementData) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfString = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfString = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ManagementData)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ManagementData)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ManagementData) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ManagementData) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	// all schemas are nil
	return nil
}

type NullableManagementData struct {
	value *ManagementData
	isSet bool
}

func (v NullableManagementData) Get() *ManagementData {
	return v.value
}

func (v *NullableManagementData) Set(val *ManagementData) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementData) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementData(val *ManagementData) *NullableManagementData {
	return &NullableManagementData{value: val, isSet: true}
}

func (v NullableManagementData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


