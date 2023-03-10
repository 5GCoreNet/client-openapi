/*
Namf_MBSBroadcast

AMF MBSBroadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MBSBroadcast

import (
	"encoding/json"
	"fmt"
)

// OperationStatusAnyOf the model 'OperationStatusAnyOf'
type OperationStatusAnyOf string

// List of OperationStatus_anyOf
const (
	START_COMPLETE OperationStatusAnyOf = "MBS_SESSION_START_COMPLETE"
	START_INCOMPLETE OperationStatusAnyOf = "MBS_SESSION_START_INCOMPLETE"
	UPDATE_COMPLETE OperationStatusAnyOf = "MBS_SESSION_UPDATE_COMPLETE"
	UPDATE_INCOMPLETE OperationStatusAnyOf = "MBS_SESSION_UPDATE_INCOMPLETE"
)

// All allowed values of OperationStatusAnyOf enum
var AllowedOperationStatusAnyOfEnumValues = []OperationStatusAnyOf{
	"MBS_SESSION_START_COMPLETE",
	"MBS_SESSION_START_INCOMPLETE",
	"MBS_SESSION_UPDATE_COMPLETE",
	"MBS_SESSION_UPDATE_INCOMPLETE",
}

func (v *OperationStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OperationStatusAnyOf(value)
	for _, existing := range AllowedOperationStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OperationStatusAnyOf", value)
}

// NewOperationStatusAnyOfFromValue returns a pointer to a valid OperationStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperationStatusAnyOfFromValue(v string) (*OperationStatusAnyOf, error) {
	ev := OperationStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperationStatusAnyOf: valid values are %v", v, AllowedOperationStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OperationStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedOperationStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OperationStatus_anyOf value
func (v OperationStatusAnyOf) Ptr() *OperationStatusAnyOf {
	return &v
}

type NullableOperationStatusAnyOf struct {
	value *OperationStatusAnyOf
	isSet bool
}

func (v NullableOperationStatusAnyOf) Get() *OperationStatusAnyOf {
	return v.value
}

func (v *NullableOperationStatusAnyOf) Set(val *OperationStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationStatusAnyOf(val *OperationStatusAnyOf) *NullableOperationStatusAnyOf {
	return &NullableOperationStatusAnyOf{value: val, isSet: true}
}

func (v NullableOperationStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

