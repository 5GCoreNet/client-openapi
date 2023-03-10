/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// ExpectedAnalyticsTypeAnyOf the model 'ExpectedAnalyticsTypeAnyOf'
type ExpectedAnalyticsTypeAnyOf string

// List of ExpectedAnalyticsType_anyOf
const (
	MOBILITY ExpectedAnalyticsTypeAnyOf = "MOBILITY"
	COMMUN ExpectedAnalyticsTypeAnyOf = "COMMUN"
	MOBILITY_AND_COMMUN ExpectedAnalyticsTypeAnyOf = "MOBILITY_AND_COMMUN"
)

// All allowed values of ExpectedAnalyticsTypeAnyOf enum
var AllowedExpectedAnalyticsTypeAnyOfEnumValues = []ExpectedAnalyticsTypeAnyOf{
	"MOBILITY",
	"COMMUN",
	"MOBILITY_AND_COMMUN",
}

func (v *ExpectedAnalyticsTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExpectedAnalyticsTypeAnyOf(value)
	for _, existing := range AllowedExpectedAnalyticsTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExpectedAnalyticsTypeAnyOf", value)
}

// NewExpectedAnalyticsTypeAnyOfFromValue returns a pointer to a valid ExpectedAnalyticsTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExpectedAnalyticsTypeAnyOfFromValue(v string) (*ExpectedAnalyticsTypeAnyOf, error) {
	ev := ExpectedAnalyticsTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExpectedAnalyticsTypeAnyOf: valid values are %v", v, AllowedExpectedAnalyticsTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExpectedAnalyticsTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedExpectedAnalyticsTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExpectedAnalyticsType_anyOf value
func (v ExpectedAnalyticsTypeAnyOf) Ptr() *ExpectedAnalyticsTypeAnyOf {
	return &v
}

type NullableExpectedAnalyticsTypeAnyOf struct {
	value *ExpectedAnalyticsTypeAnyOf
	isSet bool
}

func (v NullableExpectedAnalyticsTypeAnyOf) Get() *ExpectedAnalyticsTypeAnyOf {
	return v.value
}

func (v *NullableExpectedAnalyticsTypeAnyOf) Set(val *ExpectedAnalyticsTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExpectedAnalyticsTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExpectedAnalyticsTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpectedAnalyticsTypeAnyOf(val *ExpectedAnalyticsTypeAnyOf) *NullableExpectedAnalyticsTypeAnyOf {
	return &NullableExpectedAnalyticsTypeAnyOf{value: val, isSet: true}
}

func (v NullableExpectedAnalyticsTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpectedAnalyticsTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

