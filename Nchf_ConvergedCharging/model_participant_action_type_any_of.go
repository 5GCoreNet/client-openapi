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

// ParticipantActionTypeAnyOf the model 'ParticipantActionTypeAnyOf'
type ParticipantActionTypeAnyOf string

// List of ParticipantActionType_anyOf
const (
	CREATE ParticipantActionTypeAnyOf = "CREATE"
	JOIN ParticipantActionTypeAnyOf = "JOIN"
	INVITE_INTO ParticipantActionTypeAnyOf = "INVITE_INTO"
	QUIT ParticipantActionTypeAnyOf = "QUIT"
)

// All allowed values of ParticipantActionTypeAnyOf enum
var AllowedParticipantActionTypeAnyOfEnumValues = []ParticipantActionTypeAnyOf{
	"CREATE",
	"JOIN",
	"INVITE_INTO",
	"QUIT",
}

func (v *ParticipantActionTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ParticipantActionTypeAnyOf(value)
	for _, existing := range AllowedParticipantActionTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ParticipantActionTypeAnyOf", value)
}

// NewParticipantActionTypeAnyOfFromValue returns a pointer to a valid ParticipantActionTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewParticipantActionTypeAnyOfFromValue(v string) (*ParticipantActionTypeAnyOf, error) {
	ev := ParticipantActionTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ParticipantActionTypeAnyOf: valid values are %v", v, AllowedParticipantActionTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ParticipantActionTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedParticipantActionTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ParticipantActionType_anyOf value
func (v ParticipantActionTypeAnyOf) Ptr() *ParticipantActionTypeAnyOf {
	return &v
}

type NullableParticipantActionTypeAnyOf struct {
	value *ParticipantActionTypeAnyOf
	isSet bool
}

func (v NullableParticipantActionTypeAnyOf) Get() *ParticipantActionTypeAnyOf {
	return v.value
}

func (v *NullableParticipantActionTypeAnyOf) Set(val *ParticipantActionTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableParticipantActionTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableParticipantActionTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParticipantActionTypeAnyOf(val *ParticipantActionTypeAnyOf) *NullableParticipantActionTypeAnyOf {
	return &NullableParticipantActionTypeAnyOf{value: val, isSet: true}
}

func (v NullableParticipantActionTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParticipantActionTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

