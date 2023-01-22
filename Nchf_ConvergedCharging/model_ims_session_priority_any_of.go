/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// IMSSessionPriorityAnyOf the model 'IMSSessionPriorityAnyOf'
type IMSSessionPriorityAnyOf string

// List of IMSSessionPriority_anyOf
const (
	_0 IMSSessionPriorityAnyOf = "PRIORITY_0"
	_1 IMSSessionPriorityAnyOf = "PRIORITY_1"
	_2 IMSSessionPriorityAnyOf = "PRIORITY_2"
	_3 IMSSessionPriorityAnyOf = "PRIORITY_3"
	_4 IMSSessionPriorityAnyOf = "PRIORITY_4"
)

// All allowed values of IMSSessionPriorityAnyOf enum
var AllowedIMSSessionPriorityAnyOfEnumValues = []IMSSessionPriorityAnyOf{
	"PRIORITY_0",
	"PRIORITY_1",
	"PRIORITY_2",
	"PRIORITY_3",
	"PRIORITY_4",
}

func (v *IMSSessionPriorityAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IMSSessionPriorityAnyOf(value)
	for _, existing := range AllowedIMSSessionPriorityAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IMSSessionPriorityAnyOf", value)
}

// NewIMSSessionPriorityAnyOfFromValue returns a pointer to a valid IMSSessionPriorityAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIMSSessionPriorityAnyOfFromValue(v string) (*IMSSessionPriorityAnyOf, error) {
	ev := IMSSessionPriorityAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IMSSessionPriorityAnyOf: valid values are %v", v, AllowedIMSSessionPriorityAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IMSSessionPriorityAnyOf) IsValid() bool {
	for _, existing := range AllowedIMSSessionPriorityAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IMSSessionPriority_anyOf value
func (v IMSSessionPriorityAnyOf) Ptr() *IMSSessionPriorityAnyOf {
	return &v
}

type NullableIMSSessionPriorityAnyOf struct {
	value *IMSSessionPriorityAnyOf
	isSet bool
}

func (v NullableIMSSessionPriorityAnyOf) Get() *IMSSessionPriorityAnyOf {
	return v.value
}

func (v *NullableIMSSessionPriorityAnyOf) Set(val *IMSSessionPriorityAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIMSSessionPriorityAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIMSSessionPriorityAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIMSSessionPriorityAnyOf(val *IMSSessionPriorityAnyOf) *NullableIMSSessionPriorityAnyOf {
	return &NullableIMSSessionPriorityAnyOf{value: val, isSet: true}
}

func (v NullableIMSSessionPriorityAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIMSSessionPriorityAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

