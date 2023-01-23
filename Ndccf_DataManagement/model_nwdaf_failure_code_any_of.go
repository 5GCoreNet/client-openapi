/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// NwdafFailureCodeAnyOf the model 'NwdafFailureCodeAnyOf'
type NwdafFailureCodeAnyOf string

// List of NwdafFailureCode_anyOf
const (
	UNAVAILABLE_DATA NwdafFailureCodeAnyOf = "UNAVAILABLE_DATA"
	BOTH_STAT_PRED_NOT_ALLOWED NwdafFailureCodeAnyOf = "BOTH_STAT_PRED_NOT_ALLOWED"
	UNSATISFIED_REQUESTED_ANALYTICS_TIME NwdafFailureCodeAnyOf = "UNSATISFIED_REQUESTED_ANALYTICS_TIME"
	OTHER NwdafFailureCodeAnyOf = "OTHER"
)

// All allowed values of NwdafFailureCodeAnyOf enum
var AllowedNwdafFailureCodeAnyOfEnumValues = []NwdafFailureCodeAnyOf{
	"UNAVAILABLE_DATA",
	"BOTH_STAT_PRED_NOT_ALLOWED",
	"UNSATISFIED_REQUESTED_ANALYTICS_TIME",
	"OTHER",
}

func (v *NwdafFailureCodeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NwdafFailureCodeAnyOf(value)
	for _, existing := range AllowedNwdafFailureCodeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NwdafFailureCodeAnyOf", value)
}

// NewNwdafFailureCodeAnyOfFromValue returns a pointer to a valid NwdafFailureCodeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNwdafFailureCodeAnyOfFromValue(v string) (*NwdafFailureCodeAnyOf, error) {
	ev := NwdafFailureCodeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NwdafFailureCodeAnyOf: valid values are %v", v, AllowedNwdafFailureCodeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NwdafFailureCodeAnyOf) IsValid() bool {
	for _, existing := range AllowedNwdafFailureCodeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NwdafFailureCode_anyOf value
func (v NwdafFailureCodeAnyOf) Ptr() *NwdafFailureCodeAnyOf {
	return &v
}

type NullableNwdafFailureCodeAnyOf struct {
	value *NwdafFailureCodeAnyOf
	isSet bool
}

func (v NullableNwdafFailureCodeAnyOf) Get() *NwdafFailureCodeAnyOf {
	return v.value
}

func (v *NullableNwdafFailureCodeAnyOf) Set(val *NwdafFailureCodeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafFailureCodeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafFailureCodeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafFailureCodeAnyOf(val *NwdafFailureCodeAnyOf) *NullableNwdafFailureCodeAnyOf {
	return &NullableNwdafFailureCodeAnyOf{value: val, isSet: true}
}

func (v NullableNwdafFailureCodeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafFailureCodeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

