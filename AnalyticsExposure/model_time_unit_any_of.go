/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AnalyticsExposure

import (
	"encoding/json"
	"fmt"
)

// TimeUnitAnyOf the model 'TimeUnitAnyOf'
type TimeUnitAnyOf string

// List of TimeUnit_anyOf
const (
	MINUTE TimeUnitAnyOf = "MINUTE"
	HOUR TimeUnitAnyOf = "HOUR"
	DAY TimeUnitAnyOf = "DAY"
)

// All allowed values of TimeUnitAnyOf enum
var AllowedTimeUnitAnyOfEnumValues = []TimeUnitAnyOf{
	"MINUTE",
	"HOUR",
	"DAY",
}

func (v *TimeUnitAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TimeUnitAnyOf(value)
	for _, existing := range AllowedTimeUnitAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TimeUnitAnyOf", value)
}

// NewTimeUnitAnyOfFromValue returns a pointer to a valid TimeUnitAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTimeUnitAnyOfFromValue(v string) (*TimeUnitAnyOf, error) {
	ev := TimeUnitAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TimeUnitAnyOf: valid values are %v", v, AllowedTimeUnitAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TimeUnitAnyOf) IsValid() bool {
	for _, existing := range AllowedTimeUnitAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeUnit_anyOf value
func (v TimeUnitAnyOf) Ptr() *TimeUnitAnyOf {
	return &v
}

type NullableTimeUnitAnyOf struct {
	value *TimeUnitAnyOf
	isSet bool
}

func (v NullableTimeUnitAnyOf) Get() *TimeUnitAnyOf {
	return v.value
}

func (v *NullableTimeUnitAnyOf) Set(val *TimeUnitAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeUnitAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeUnitAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeUnitAnyOf(val *TimeUnitAnyOf) *NullableTimeUnitAnyOf {
	return &NullableTimeUnitAnyOf{value: val, isSet: true}
}

func (v NullableTimeUnitAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeUnitAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

