/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// SmfSelectionTypeAnyOf the model 'SmfSelectionTypeAnyOf'
type SmfSelectionTypeAnyOf string

// List of SmfSelectionType_anyOf
const (
	CURRENT_PDU_SESSION SmfSelectionTypeAnyOf = "CURRENT_PDU_SESSION"
	NEXT_PDU_SESSION SmfSelectionTypeAnyOf = "NEXT_PDU_SESSION"
)

// All allowed values of SmfSelectionTypeAnyOf enum
var AllowedSmfSelectionTypeAnyOfEnumValues = []SmfSelectionTypeAnyOf{
	"CURRENT_PDU_SESSION",
	"NEXT_PDU_SESSION",
}

func (v *SmfSelectionTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SmfSelectionTypeAnyOf(value)
	for _, existing := range AllowedSmfSelectionTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SmfSelectionTypeAnyOf", value)
}

// NewSmfSelectionTypeAnyOfFromValue returns a pointer to a valid SmfSelectionTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSmfSelectionTypeAnyOfFromValue(v string) (*SmfSelectionTypeAnyOf, error) {
	ev := SmfSelectionTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SmfSelectionTypeAnyOf: valid values are %v", v, AllowedSmfSelectionTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SmfSelectionTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedSmfSelectionTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SmfSelectionType_anyOf value
func (v SmfSelectionTypeAnyOf) Ptr() *SmfSelectionTypeAnyOf {
	return &v
}

type NullableSmfSelectionTypeAnyOf struct {
	value *SmfSelectionTypeAnyOf
	isSet bool
}

func (v NullableSmfSelectionTypeAnyOf) Get() *SmfSelectionTypeAnyOf {
	return v.value
}

func (v *NullableSmfSelectionTypeAnyOf) Set(val *SmfSelectionTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSmfSelectionTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSmfSelectionTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmfSelectionTypeAnyOf(val *SmfSelectionTypeAnyOf) *NullableSmfSelectionTypeAnyOf {
	return &NullableSmfSelectionTypeAnyOf{value: val, isSet: true}
}

func (v NullableSmfSelectionTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmfSelectionTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

