/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// ExceptionTrendAnyOf the model 'ExceptionTrendAnyOf'
type ExceptionTrendAnyOf string

// List of ExceptionTrend_anyOf
const (
	UP ExceptionTrendAnyOf = "UP"
	DOWN ExceptionTrendAnyOf = "DOWN"
	UNKNOW ExceptionTrendAnyOf = "UNKNOW"
	STABLE ExceptionTrendAnyOf = "STABLE"
)

// All allowed values of ExceptionTrendAnyOf enum
var AllowedExceptionTrendAnyOfEnumValues = []ExceptionTrendAnyOf{
	"UP",
	"DOWN",
	"UNKNOW",
	"STABLE",
}

func (v *ExceptionTrendAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExceptionTrendAnyOf(value)
	for _, existing := range AllowedExceptionTrendAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExceptionTrendAnyOf", value)
}

// NewExceptionTrendAnyOfFromValue returns a pointer to a valid ExceptionTrendAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExceptionTrendAnyOfFromValue(v string) (*ExceptionTrendAnyOf, error) {
	ev := ExceptionTrendAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExceptionTrendAnyOf: valid values are %v", v, AllowedExceptionTrendAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExceptionTrendAnyOf) IsValid() bool {
	for _, existing := range AllowedExceptionTrendAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExceptionTrend_anyOf value
func (v ExceptionTrendAnyOf) Ptr() *ExceptionTrendAnyOf {
	return &v
}

type NullableExceptionTrendAnyOf struct {
	value *ExceptionTrendAnyOf
	isSet bool
}

func (v NullableExceptionTrendAnyOf) Get() *ExceptionTrendAnyOf {
	return v.value
}

func (v *NullableExceptionTrendAnyOf) Set(val *ExceptionTrendAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionTrendAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionTrendAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionTrendAnyOf(val *ExceptionTrendAnyOf) *NullableExceptionTrendAnyOf {
	return &NullableExceptionTrendAnyOf{value: val, isSet: true}
}

func (v NullableExceptionTrendAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExceptionTrendAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
