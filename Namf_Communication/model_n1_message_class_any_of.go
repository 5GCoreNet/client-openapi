/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// N1MessageClassAnyOf the model 'N1MessageClassAnyOf'
type N1MessageClassAnyOf string

// List of N1MessageClass_anyOf
const (
	_5_GMM N1MessageClassAnyOf = "5GMM"
	SM N1MessageClassAnyOf = "SM"
	LPP N1MessageClassAnyOf = "LPP"
	SMS N1MessageClassAnyOf = "SMS"
	UPDP N1MessageClassAnyOf = "UPDP"
	LCS N1MessageClassAnyOf = "LCS"
)

// All allowed values of N1MessageClassAnyOf enum
var AllowedN1MessageClassAnyOfEnumValues = []N1MessageClassAnyOf{
	"5GMM",
	"SM",
	"LPP",
	"SMS",
	"UPDP",
	"LCS",
}

func (v *N1MessageClassAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := N1MessageClassAnyOf(value)
	for _, existing := range AllowedN1MessageClassAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid N1MessageClassAnyOf", value)
}

// NewN1MessageClassAnyOfFromValue returns a pointer to a valid N1MessageClassAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewN1MessageClassAnyOfFromValue(v string) (*N1MessageClassAnyOf, error) {
	ev := N1MessageClassAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for N1MessageClassAnyOf: valid values are %v", v, AllowedN1MessageClassAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v N1MessageClassAnyOf) IsValid() bool {
	for _, existing := range AllowedN1MessageClassAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to N1MessageClass_anyOf value
func (v N1MessageClassAnyOf) Ptr() *N1MessageClassAnyOf {
	return &v
}

type NullableN1MessageClassAnyOf struct {
	value *N1MessageClassAnyOf
	isSet bool
}

func (v NullableN1MessageClassAnyOf) Get() *N1MessageClassAnyOf {
	return v.value
}

func (v *NullableN1MessageClassAnyOf) Set(val *N1MessageClassAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableN1MessageClassAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableN1MessageClassAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN1MessageClassAnyOf(val *N1MessageClassAnyOf) *NullableN1MessageClassAnyOf {
	return &NullableN1MessageClassAnyOf{value: val, isSet: true}
}

func (v NullableN1MessageClassAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN1MessageClassAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

