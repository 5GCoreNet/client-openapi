/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package AnalyticsExposure

import (
	"encoding/json"
	"fmt"
)

// CongestionTypeAnyOf the model 'CongestionTypeAnyOf'
type CongestionTypeAnyOf string

// List of CongestionType_anyOf
const (
	USER_PLANE CongestionTypeAnyOf = "USER_PLANE"
	CONTROL_PLANE CongestionTypeAnyOf = "CONTROL_PLANE"
	USER_AND_CONTROL_PLANE CongestionTypeAnyOf = "USER_AND_CONTROL_PLANE"
)

// All allowed values of CongestionTypeAnyOf enum
var AllowedCongestionTypeAnyOfEnumValues = []CongestionTypeAnyOf{
	"USER_PLANE",
	"CONTROL_PLANE",
	"USER_AND_CONTROL_PLANE",
}

func (v *CongestionTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CongestionTypeAnyOf(value)
	for _, existing := range AllowedCongestionTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CongestionTypeAnyOf", value)
}

// NewCongestionTypeAnyOfFromValue returns a pointer to a valid CongestionTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCongestionTypeAnyOfFromValue(v string) (*CongestionTypeAnyOf, error) {
	ev := CongestionTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CongestionTypeAnyOf: valid values are %v", v, AllowedCongestionTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CongestionTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedCongestionTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CongestionType_anyOf value
func (v CongestionTypeAnyOf) Ptr() *CongestionTypeAnyOf {
	return &v
}

type NullableCongestionTypeAnyOf struct {
	value *CongestionTypeAnyOf
	isSet bool
}

func (v NullableCongestionTypeAnyOf) Get() *CongestionTypeAnyOf {
	return v.value
}

func (v *NullableCongestionTypeAnyOf) Set(val *CongestionTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCongestionTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCongestionTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCongestionTypeAnyOf(val *CongestionTypeAnyOf) *NullableCongestionTypeAnyOf {
	return &NullableCongestionTypeAnyOf{value: val, isSet: true}
}

func (v NullableCongestionTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCongestionTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

