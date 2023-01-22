/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// N2InformationClassAnyOf the model 'N2InformationClassAnyOf'
type N2InformationClassAnyOf string

// List of N2InformationClass_anyOf
const (
	SM N2InformationClassAnyOf = "SM"
	NRPPA N2InformationClassAnyOf = "NRPPa"
	PWS N2InformationClassAnyOf = "PWS"
	PWS_BCAL N2InformationClassAnyOf = "PWS-BCAL"
	PWS_RF N2InformationClassAnyOf = "PWS-RF"
	RAN N2InformationClassAnyOf = "RAN"
	V2_X N2InformationClassAnyOf = "V2X"
	PROSE N2InformationClassAnyOf = "PROSE"
)

// All allowed values of N2InformationClassAnyOf enum
var AllowedN2InformationClassAnyOfEnumValues = []N2InformationClassAnyOf{
	"SM",
	"NRPPa",
	"PWS",
	"PWS-BCAL",
	"PWS-RF",
	"RAN",
	"V2X",
	"PROSE",
}

func (v *N2InformationClassAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := N2InformationClassAnyOf(value)
	for _, existing := range AllowedN2InformationClassAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid N2InformationClassAnyOf", value)
}

// NewN2InformationClassAnyOfFromValue returns a pointer to a valid N2InformationClassAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewN2InformationClassAnyOfFromValue(v string) (*N2InformationClassAnyOf, error) {
	ev := N2InformationClassAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for N2InformationClassAnyOf: valid values are %v", v, AllowedN2InformationClassAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v N2InformationClassAnyOf) IsValid() bool {
	for _, existing := range AllowedN2InformationClassAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to N2InformationClass_anyOf value
func (v N2InformationClassAnyOf) Ptr() *N2InformationClassAnyOf {
	return &v
}

type NullableN2InformationClassAnyOf struct {
	value *N2InformationClassAnyOf
	isSet bool
}

func (v NullableN2InformationClassAnyOf) Get() *N2InformationClassAnyOf {
	return v.value
}

func (v *NullableN2InformationClassAnyOf) Set(val *N2InformationClassAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableN2InformationClassAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableN2InformationClassAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2InformationClassAnyOf(val *N2InformationClassAnyOf) *NullableN2InformationClassAnyOf {
	return &NullableN2InformationClassAnyOf{value: val, isSet: true}
}

func (v NullableN2InformationClassAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2InformationClassAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

