/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
	"fmt"
)

// SharingLevel the model 'SharingLevel'
type SharingLevel string

// List of SharingLevel
const (
	SHARED SharingLevel = "SHARED"
	NON_SHARED SharingLevel = "NON-SHARED"
)

// All allowed values of SharingLevel enum
var AllowedSharingLevelEnumValues = []SharingLevel{
	"SHARED",
	"NON-SHARED",
}

func (v *SharingLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SharingLevel(value)
	for _, existing := range AllowedSharingLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SharingLevel", value)
}

// NewSharingLevelFromValue returns a pointer to a valid SharingLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSharingLevelFromValue(v string) (*SharingLevel, error) {
	ev := SharingLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SharingLevel: valid values are %v", v, AllowedSharingLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SharingLevel) IsValid() bool {
	for _, existing := range AllowedSharingLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SharingLevel value
func (v SharingLevel) Ptr() *SharingLevel {
	return &v
}

type NullableSharingLevel struct {
	value *SharingLevel
	isSet bool
}

func (v NullableSharingLevel) Get() *SharingLevel {
	return v.value
}

func (v *NullableSharingLevel) Set(val *SharingLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableSharingLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableSharingLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharingLevel(val *SharingLevel) *NullableSharingLevel {
	return &NullableSharingLevel{value: val, isSet: true}
}

func (v NullableSharingLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharingLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

