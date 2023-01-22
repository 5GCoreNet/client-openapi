/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ngmlc_Location

import (
	"encoding/json"
	"fmt"
)

// GnssIdAnyOf the model 'GnssIdAnyOf'
type GnssIdAnyOf string

// List of GnssId_anyOf
const (
	GPS GnssIdAnyOf = "GPS"
	GALILEO GnssIdAnyOf = "GALILEO"
	SBAS GnssIdAnyOf = "SBAS"
	MODERNIZED_GPS GnssIdAnyOf = "MODERNIZED_GPS"
	QZSS GnssIdAnyOf = "QZSS"
	GLONASS GnssIdAnyOf = "GLONASS"
	BDS GnssIdAnyOf = "BDS"
	NAVIC GnssIdAnyOf = "NAVIC"
)

// All allowed values of GnssIdAnyOf enum
var AllowedGnssIdAnyOfEnumValues = []GnssIdAnyOf{
	"GPS",
	"GALILEO",
	"SBAS",
	"MODERNIZED_GPS",
	"QZSS",
	"GLONASS",
	"BDS",
	"NAVIC",
}

func (v *GnssIdAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GnssIdAnyOf(value)
	for _, existing := range AllowedGnssIdAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GnssIdAnyOf", value)
}

// NewGnssIdAnyOfFromValue returns a pointer to a valid GnssIdAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGnssIdAnyOfFromValue(v string) (*GnssIdAnyOf, error) {
	ev := GnssIdAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GnssIdAnyOf: valid values are %v", v, AllowedGnssIdAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GnssIdAnyOf) IsValid() bool {
	for _, existing := range AllowedGnssIdAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GnssId_anyOf value
func (v GnssIdAnyOf) Ptr() *GnssIdAnyOf {
	return &v
}

type NullableGnssIdAnyOf struct {
	value *GnssIdAnyOf
	isSet bool
}

func (v NullableGnssIdAnyOf) Get() *GnssIdAnyOf {
	return v.value
}

func (v *NullableGnssIdAnyOf) Set(val *GnssIdAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGnssIdAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGnssIdAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGnssIdAnyOf(val *GnssIdAnyOf) *NullableGnssIdAnyOf {
	return &NullableGnssIdAnyOf{value: val, isSet: true}
}

func (v NullableGnssIdAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGnssIdAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

