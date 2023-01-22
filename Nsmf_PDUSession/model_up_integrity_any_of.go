/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// UpIntegrityAnyOf the model 'UpIntegrityAnyOf'
type UpIntegrityAnyOf string

// List of UpIntegrity_anyOf
const (
	REQUIRED UpIntegrityAnyOf = "REQUIRED"
	PREFERRED UpIntegrityAnyOf = "PREFERRED"
	NOT_NEEDED UpIntegrityAnyOf = "NOT_NEEDED"
)

// All allowed values of UpIntegrityAnyOf enum
var AllowedUpIntegrityAnyOfEnumValues = []UpIntegrityAnyOf{
	"REQUIRED",
	"PREFERRED",
	"NOT_NEEDED",
}

func (v *UpIntegrityAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UpIntegrityAnyOf(value)
	for _, existing := range AllowedUpIntegrityAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UpIntegrityAnyOf", value)
}

// NewUpIntegrityAnyOfFromValue returns a pointer to a valid UpIntegrityAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUpIntegrityAnyOfFromValue(v string) (*UpIntegrityAnyOf, error) {
	ev := UpIntegrityAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UpIntegrityAnyOf: valid values are %v", v, AllowedUpIntegrityAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UpIntegrityAnyOf) IsValid() bool {
	for _, existing := range AllowedUpIntegrityAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpIntegrity_anyOf value
func (v UpIntegrityAnyOf) Ptr() *UpIntegrityAnyOf {
	return &v
}

type NullableUpIntegrityAnyOf struct {
	value *UpIntegrityAnyOf
	isSet bool
}

func (v NullableUpIntegrityAnyOf) Get() *UpIntegrityAnyOf {
	return v.value
}

func (v *NullableUpIntegrityAnyOf) Set(val *UpIntegrityAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUpIntegrityAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUpIntegrityAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpIntegrityAnyOf(val *UpIntegrityAnyOf) *NullableUpIntegrityAnyOf {
	return &NullableUpIntegrityAnyOf{value: val, isSet: true}
}

func (v NullableUpIntegrityAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpIntegrityAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

