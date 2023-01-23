/*
Nmfaf_3daDataManagement

MFAF 3GPP DCCF Adaptor (3DA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3daDataManagement

import (
	"encoding/json"
	"fmt"
)

// SACEventTriggerAnyOf the model 'SACEventTriggerAnyOf'
type SACEventTriggerAnyOf string

// List of SACEventTrigger_anyOf
const (
	THRESHOLD SACEventTriggerAnyOf = "THRESHOLD"
	PERIODIC SACEventTriggerAnyOf = "PERIODIC"
)

// All allowed values of SACEventTriggerAnyOf enum
var AllowedSACEventTriggerAnyOfEnumValues = []SACEventTriggerAnyOf{
	"THRESHOLD",
	"PERIODIC",
}

func (v *SACEventTriggerAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SACEventTriggerAnyOf(value)
	for _, existing := range AllowedSACEventTriggerAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SACEventTriggerAnyOf", value)
}

// NewSACEventTriggerAnyOfFromValue returns a pointer to a valid SACEventTriggerAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSACEventTriggerAnyOfFromValue(v string) (*SACEventTriggerAnyOf, error) {
	ev := SACEventTriggerAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SACEventTriggerAnyOf: valid values are %v", v, AllowedSACEventTriggerAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SACEventTriggerAnyOf) IsValid() bool {
	for _, existing := range AllowedSACEventTriggerAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SACEventTrigger_anyOf value
func (v SACEventTriggerAnyOf) Ptr() *SACEventTriggerAnyOf {
	return &v
}

type NullableSACEventTriggerAnyOf struct {
	value *SACEventTriggerAnyOf
	isSet bool
}

func (v NullableSACEventTriggerAnyOf) Get() *SACEventTriggerAnyOf {
	return v.value
}

func (v *NullableSACEventTriggerAnyOf) Set(val *SACEventTriggerAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSACEventTriggerAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSACEventTriggerAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSACEventTriggerAnyOf(val *SACEventTriggerAnyOf) *NullableSACEventTriggerAnyOf {
	return &NullableSACEventTriggerAnyOf{value: val, isSet: true}
}

func (v NullableSACEventTriggerAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSACEventTriggerAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

