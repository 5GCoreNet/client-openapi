/*
Nhss_UECM

HSS UE Context Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_UECM

import (
	"encoding/json"
	"fmt"
)

// DeregistrationReasonAnyOf the model 'DeregistrationReasonAnyOf'
type DeregistrationReasonAnyOf string

// List of DeregistrationReason_anyOf
const (
	UE_INITIAL_AND_SINGLE_REGISTRATION DeregistrationReasonAnyOf = "UE_INITIAL_AND_SINGLE_REGISTRATION"
	UE_INITIAL_AND_DUAL_REGISTRATION DeregistrationReasonAnyOf = "UE_INITIAL_AND_DUAL_REGISTRATION"
	EPS_TO_5_GS_MOBILITY DeregistrationReasonAnyOf = "EPS_TO_5GS_MOBILITY"
)

// All allowed values of DeregistrationReasonAnyOf enum
var AllowedDeregistrationReasonAnyOfEnumValues = []DeregistrationReasonAnyOf{
	"UE_INITIAL_AND_SINGLE_REGISTRATION",
	"UE_INITIAL_AND_DUAL_REGISTRATION",
	"EPS_TO_5GS_MOBILITY",
}

func (v *DeregistrationReasonAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeregistrationReasonAnyOf(value)
	for _, existing := range AllowedDeregistrationReasonAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeregistrationReasonAnyOf", value)
}

// NewDeregistrationReasonAnyOfFromValue returns a pointer to a valid DeregistrationReasonAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeregistrationReasonAnyOfFromValue(v string) (*DeregistrationReasonAnyOf, error) {
	ev := DeregistrationReasonAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeregistrationReasonAnyOf: valid values are %v", v, AllowedDeregistrationReasonAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeregistrationReasonAnyOf) IsValid() bool {
	for _, existing := range AllowedDeregistrationReasonAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeregistrationReason_anyOf value
func (v DeregistrationReasonAnyOf) Ptr() *DeregistrationReasonAnyOf {
	return &v
}

type NullableDeregistrationReasonAnyOf struct {
	value *DeregistrationReasonAnyOf
	isSet bool
}

func (v NullableDeregistrationReasonAnyOf) Get() *DeregistrationReasonAnyOf {
	return v.value
}

func (v *NullableDeregistrationReasonAnyOf) Set(val *DeregistrationReasonAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeregistrationReasonAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeregistrationReasonAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeregistrationReasonAnyOf(val *DeregistrationReasonAnyOf) *NullableDeregistrationReasonAnyOf {
	return &NullableDeregistrationReasonAnyOf{value: val, isSet: true}
}

func (v NullableDeregistrationReasonAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeregistrationReasonAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

