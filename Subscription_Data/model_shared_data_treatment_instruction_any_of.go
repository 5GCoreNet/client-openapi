/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// SharedDataTreatmentInstructionAnyOf the model 'SharedDataTreatmentInstructionAnyOf'
type SharedDataTreatmentInstructionAnyOf string

// List of SharedDataTreatmentInstruction_anyOf
const (
	USE_IF_NO_CLASH SharedDataTreatmentInstructionAnyOf = "USE_IF_NO_CLASH"
	OVERWRITE SharedDataTreatmentInstructionAnyOf = "OVERWRITE"
	MAX SharedDataTreatmentInstructionAnyOf = "MAX"
	MIN SharedDataTreatmentInstructionAnyOf = "MIN"
)

// All allowed values of SharedDataTreatmentInstructionAnyOf enum
var AllowedSharedDataTreatmentInstructionAnyOfEnumValues = []SharedDataTreatmentInstructionAnyOf{
	"USE_IF_NO_CLASH",
	"OVERWRITE",
	"MAX",
	"MIN",
}

func (v *SharedDataTreatmentInstructionAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SharedDataTreatmentInstructionAnyOf(value)
	for _, existing := range AllowedSharedDataTreatmentInstructionAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SharedDataTreatmentInstructionAnyOf", value)
}

// NewSharedDataTreatmentInstructionAnyOfFromValue returns a pointer to a valid SharedDataTreatmentInstructionAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSharedDataTreatmentInstructionAnyOfFromValue(v string) (*SharedDataTreatmentInstructionAnyOf, error) {
	ev := SharedDataTreatmentInstructionAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SharedDataTreatmentInstructionAnyOf: valid values are %v", v, AllowedSharedDataTreatmentInstructionAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SharedDataTreatmentInstructionAnyOf) IsValid() bool {
	for _, existing := range AllowedSharedDataTreatmentInstructionAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SharedDataTreatmentInstruction_anyOf value
func (v SharedDataTreatmentInstructionAnyOf) Ptr() *SharedDataTreatmentInstructionAnyOf {
	return &v
}

type NullableSharedDataTreatmentInstructionAnyOf struct {
	value *SharedDataTreatmentInstructionAnyOf
	isSet bool
}

func (v NullableSharedDataTreatmentInstructionAnyOf) Get() *SharedDataTreatmentInstructionAnyOf {
	return v.value
}

func (v *NullableSharedDataTreatmentInstructionAnyOf) Set(val *SharedDataTreatmentInstructionAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedDataTreatmentInstructionAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedDataTreatmentInstructionAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedDataTreatmentInstructionAnyOf(val *SharedDataTreatmentInstructionAnyOf) *NullableSharedDataTreatmentInstructionAnyOf {
	return &NullableSharedDataTreatmentInstructionAnyOf{value: val, isSet: true}
}

func (v NullableSharedDataTreatmentInstructionAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedDataTreatmentInstructionAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

