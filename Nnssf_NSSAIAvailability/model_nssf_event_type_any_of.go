/*
NSSF NSSAI Availability

NSSF NSSAI Availability Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnssf_NSSAIAvailability

import (
	"encoding/json"
	"fmt"
)

// NssfEventTypeAnyOf the model 'NssfEventTypeAnyOf'
type NssfEventTypeAnyOf string

// List of NssfEventType_anyOf
const (
	SNSSAI_STATUS_CHANGE_REPORT NssfEventTypeAnyOf = "SNSSAI_STATUS_CHANGE_REPORT"
)

// All allowed values of NssfEventTypeAnyOf enum
var AllowedNssfEventTypeAnyOfEnumValues = []NssfEventTypeAnyOf{
	"SNSSAI_STATUS_CHANGE_REPORT",
}

func (v *NssfEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NssfEventTypeAnyOf(value)
	for _, existing := range AllowedNssfEventTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NssfEventTypeAnyOf", value)
}

// NewNssfEventTypeAnyOfFromValue returns a pointer to a valid NssfEventTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNssfEventTypeAnyOfFromValue(v string) (*NssfEventTypeAnyOf, error) {
	ev := NssfEventTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NssfEventTypeAnyOf: valid values are %v", v, AllowedNssfEventTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NssfEventTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedNssfEventTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NssfEventType_anyOf value
func (v NssfEventTypeAnyOf) Ptr() *NssfEventTypeAnyOf {
	return &v
}

type NullableNssfEventTypeAnyOf struct {
	value *NssfEventTypeAnyOf
	isSet bool
}

func (v NullableNssfEventTypeAnyOf) Get() *NssfEventTypeAnyOf {
	return v.value
}

func (v *NullableNssfEventTypeAnyOf) Set(val *NssfEventTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNssfEventTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNssfEventTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNssfEventTypeAnyOf(val *NssfEventTypeAnyOf) *NullableNssfEventTypeAnyOf {
	return &NullableNssfEventTypeAnyOf{value: val, isSet: true}
}

func (v NullableNssfEventTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNssfEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

