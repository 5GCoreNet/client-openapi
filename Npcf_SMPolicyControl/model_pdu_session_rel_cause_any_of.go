/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// PduSessionRelCauseAnyOf the model 'PduSessionRelCauseAnyOf'
type PduSessionRelCauseAnyOf string

// List of PduSessionRelCause_anyOf
const (
	PS_TO_CS_HO PduSessionRelCauseAnyOf = "PS_TO_CS_HO"
	RULE_ERROR PduSessionRelCauseAnyOf = "RULE_ERROR"
)

// All allowed values of PduSessionRelCauseAnyOf enum
var AllowedPduSessionRelCauseAnyOfEnumValues = []PduSessionRelCauseAnyOf{
	"PS_TO_CS_HO",
	"RULE_ERROR",
}

func (v *PduSessionRelCauseAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PduSessionRelCauseAnyOf(value)
	for _, existing := range AllowedPduSessionRelCauseAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PduSessionRelCauseAnyOf", value)
}

// NewPduSessionRelCauseAnyOfFromValue returns a pointer to a valid PduSessionRelCauseAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPduSessionRelCauseAnyOfFromValue(v string) (*PduSessionRelCauseAnyOf, error) {
	ev := PduSessionRelCauseAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PduSessionRelCauseAnyOf: valid values are %v", v, AllowedPduSessionRelCauseAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PduSessionRelCauseAnyOf) IsValid() bool {
	for _, existing := range AllowedPduSessionRelCauseAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PduSessionRelCause_anyOf value
func (v PduSessionRelCauseAnyOf) Ptr() *PduSessionRelCauseAnyOf {
	return &v
}

type NullablePduSessionRelCauseAnyOf struct {
	value *PduSessionRelCauseAnyOf
	isSet bool
}

func (v NullablePduSessionRelCauseAnyOf) Get() *PduSessionRelCauseAnyOf {
	return v.value
}

func (v *NullablePduSessionRelCauseAnyOf) Set(val *PduSessionRelCauseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionRelCauseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionRelCauseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionRelCauseAnyOf(val *PduSessionRelCauseAnyOf) *NullablePduSessionRelCauseAnyOf {
	return &NullablePduSessionRelCauseAnyOf{value: val, isSet: true}
}

func (v NullablePduSessionRelCauseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionRelCauseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

