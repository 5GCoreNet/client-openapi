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

// SACEventTypeAnyOf the model 'SACEventTypeAnyOf'
type SACEventTypeAnyOf string

// List of SACEventType_anyOf
const (
	REGD_UES SACEventTypeAnyOf = "NUM_OF_REGD_UES"
	ESTD_PDU_SESSIONS SACEventTypeAnyOf = "NUM_OF_ESTD_PDU_SESSIONS"
)

// All allowed values of SACEventTypeAnyOf enum
var AllowedSACEventTypeAnyOfEnumValues = []SACEventTypeAnyOf{
	"NUM_OF_REGD_UES",
	"NUM_OF_ESTD_PDU_SESSIONS",
}

func (v *SACEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SACEventTypeAnyOf(value)
	for _, existing := range AllowedSACEventTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SACEventTypeAnyOf", value)
}

// NewSACEventTypeAnyOfFromValue returns a pointer to a valid SACEventTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSACEventTypeAnyOfFromValue(v string) (*SACEventTypeAnyOf, error) {
	ev := SACEventTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SACEventTypeAnyOf: valid values are %v", v, AllowedSACEventTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SACEventTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedSACEventTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SACEventType_anyOf value
func (v SACEventTypeAnyOf) Ptr() *SACEventTypeAnyOf {
	return &v
}

type NullableSACEventTypeAnyOf struct {
	value *SACEventTypeAnyOf
	isSet bool
}

func (v NullableSACEventTypeAnyOf) Get() *SACEventTypeAnyOf {
	return v.value
}

func (v *NullableSACEventTypeAnyOf) Set(val *SACEventTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSACEventTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSACEventTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSACEventTypeAnyOf(val *SACEventTypeAnyOf) *NullableSACEventTypeAnyOf {
	return &NullableSACEventTypeAnyOf{value: val, isSet: true}
}

func (v NullableSACEventTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSACEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

