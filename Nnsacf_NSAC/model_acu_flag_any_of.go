/*
Nnsacf_NSAC

Nnsacf_NSAC Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnsacf_NSAC

import (
	"encoding/json"
	"fmt"
)

// AcuFlagAnyOf the model 'AcuFlagAnyOf'
type AcuFlagAnyOf string

// List of AcuFlag_anyOf
const (
	INCREASE AcuFlagAnyOf = "INCREASE"
	DECREASE AcuFlagAnyOf = "DECREASE"
	UPDATE AcuFlagAnyOf = "UPDATE"
)

// All allowed values of AcuFlagAnyOf enum
var AllowedAcuFlagAnyOfEnumValues = []AcuFlagAnyOf{
	"INCREASE",
	"DECREASE",
	"UPDATE",
}

func (v *AcuFlagAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AcuFlagAnyOf(value)
	for _, existing := range AllowedAcuFlagAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AcuFlagAnyOf", value)
}

// NewAcuFlagAnyOfFromValue returns a pointer to a valid AcuFlagAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAcuFlagAnyOfFromValue(v string) (*AcuFlagAnyOf, error) {
	ev := AcuFlagAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AcuFlagAnyOf: valid values are %v", v, AllowedAcuFlagAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AcuFlagAnyOf) IsValid() bool {
	for _, existing := range AllowedAcuFlagAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AcuFlag_anyOf value
func (v AcuFlagAnyOf) Ptr() *AcuFlagAnyOf {
	return &v
}

type NullableAcuFlagAnyOf struct {
	value *AcuFlagAnyOf
	isSet bool
}

func (v NullableAcuFlagAnyOf) Get() *AcuFlagAnyOf {
	return v.value
}

func (v *NullableAcuFlagAnyOf) Set(val *AcuFlagAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAcuFlagAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAcuFlagAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcuFlagAnyOf(val *AcuFlagAnyOf) *NullableAcuFlagAnyOf {
	return &NullableAcuFlagAnyOf{value: val, isSet: true}
}

func (v NullableAcuFlagAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcuFlagAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

