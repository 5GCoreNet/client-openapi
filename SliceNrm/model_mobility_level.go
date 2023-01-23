/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
	"fmt"
)

// MobilityLevel the model 'MobilityLevel'
type MobilityLevel string

// List of MobilityLevel
const (
	STATIONARY MobilityLevel = "STATIONARY"
	NOMADIC MobilityLevel = "NOMADIC"
	RESTRICTED_MOBILITY MobilityLevel = "RESTRICTED MOBILITY"
	FULLY_MOBILITY MobilityLevel = "FULLY MOBILITY"
)

// All allowed values of MobilityLevel enum
var AllowedMobilityLevelEnumValues = []MobilityLevel{
	"STATIONARY",
	"NOMADIC",
	"RESTRICTED MOBILITY",
	"FULLY MOBILITY",
}

func (v *MobilityLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MobilityLevel(value)
	for _, existing := range AllowedMobilityLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MobilityLevel", value)
}

// NewMobilityLevelFromValue returns a pointer to a valid MobilityLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMobilityLevelFromValue(v string) (*MobilityLevel, error) {
	ev := MobilityLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MobilityLevel: valid values are %v", v, AllowedMobilityLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MobilityLevel) IsValid() bool {
	for _, existing := range AllowedMobilityLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MobilityLevel value
func (v MobilityLevel) Ptr() *MobilityLevel {
	return &v
}

type NullableMobilityLevel struct {
	value *MobilityLevel
	isSet bool
}

func (v NullableMobilityLevel) Get() *MobilityLevel {
	return v.value
}

func (v *NullableMobilityLevel) Set(val *MobilityLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableMobilityLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableMobilityLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobilityLevel(val *MobilityLevel) *NullableMobilityLevel {
	return &NullableMobilityLevel{value: val, isSet: true}
}

func (v NullableMobilityLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobilityLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

