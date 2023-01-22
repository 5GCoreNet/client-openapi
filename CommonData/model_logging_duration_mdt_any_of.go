/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CommonData

import (
	"encoding/json"
	"fmt"
)

// LoggingDurationMdtAnyOf the model 'LoggingDurationMdtAnyOf'
type LoggingDurationMdtAnyOf string

// List of LoggingDurationMdt_anyOf
const (
	_600 LoggingDurationMdtAnyOf = "600"
	_1200 LoggingDurationMdtAnyOf = "1200"
	_2400 LoggingDurationMdtAnyOf = "2400"
	_3600 LoggingDurationMdtAnyOf = "3600"
	_5400 LoggingDurationMdtAnyOf = "5400"
	_7200 LoggingDurationMdtAnyOf = "7200"
)

// All allowed values of LoggingDurationMdtAnyOf enum
var AllowedLoggingDurationMdtAnyOfEnumValues = []LoggingDurationMdtAnyOf{
	"600",
	"1200",
	"2400",
	"3600",
	"5400",
	"7200",
}

func (v *LoggingDurationMdtAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LoggingDurationMdtAnyOf(value)
	for _, existing := range AllowedLoggingDurationMdtAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LoggingDurationMdtAnyOf", value)
}

// NewLoggingDurationMdtAnyOfFromValue returns a pointer to a valid LoggingDurationMdtAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoggingDurationMdtAnyOfFromValue(v string) (*LoggingDurationMdtAnyOf, error) {
	ev := LoggingDurationMdtAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LoggingDurationMdtAnyOf: valid values are %v", v, AllowedLoggingDurationMdtAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoggingDurationMdtAnyOf) IsValid() bool {
	for _, existing := range AllowedLoggingDurationMdtAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LoggingDurationMdt_anyOf value
func (v LoggingDurationMdtAnyOf) Ptr() *LoggingDurationMdtAnyOf {
	return &v
}

type NullableLoggingDurationMdtAnyOf struct {
	value *LoggingDurationMdtAnyOf
	isSet bool
}

func (v NullableLoggingDurationMdtAnyOf) Get() *LoggingDurationMdtAnyOf {
	return v.value
}

func (v *NullableLoggingDurationMdtAnyOf) Set(val *LoggingDurationMdtAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLoggingDurationMdtAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLoggingDurationMdtAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoggingDurationMdtAnyOf(val *LoggingDurationMdtAnyOf) *NullableLoggingDurationMdtAnyOf {
	return &NullableLoggingDurationMdtAnyOf{value: val, isSet: true}
}

func (v NullableLoggingDurationMdtAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoggingDurationMdtAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
