/*
Nucmf_Provisioning

UCMF_Provisioning Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nucmf_Provisioning

import (
	"encoding/json"
	"fmt"
)

// RacsFailureCodeAnyOf the model 'RacsFailureCodeAnyOf'
type RacsFailureCodeAnyOf string

// List of RacsFailureCode_anyOf
const (
	MALFUNCTION RacsFailureCodeAnyOf = "MALFUNCTION"
	RESOURCE_LIMITATION RacsFailureCodeAnyOf = "RESOURCE_LIMITATION"
	RACS_ID_DUPLICATED RacsFailureCodeAnyOf = "RACS_ID_DUPLICATED"
	OTHER_REASON RacsFailureCodeAnyOf = "OTHER_REASON"
)

// All allowed values of RacsFailureCodeAnyOf enum
var AllowedRacsFailureCodeAnyOfEnumValues = []RacsFailureCodeAnyOf{
	"MALFUNCTION",
	"RESOURCE_LIMITATION",
	"RACS_ID_DUPLICATED",
	"OTHER_REASON",
}

func (v *RacsFailureCodeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RacsFailureCodeAnyOf(value)
	for _, existing := range AllowedRacsFailureCodeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RacsFailureCodeAnyOf", value)
}

// NewRacsFailureCodeAnyOfFromValue returns a pointer to a valid RacsFailureCodeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRacsFailureCodeAnyOfFromValue(v string) (*RacsFailureCodeAnyOf, error) {
	ev := RacsFailureCodeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RacsFailureCodeAnyOf: valid values are %v", v, AllowedRacsFailureCodeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RacsFailureCodeAnyOf) IsValid() bool {
	for _, existing := range AllowedRacsFailureCodeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RacsFailureCode_anyOf value
func (v RacsFailureCodeAnyOf) Ptr() *RacsFailureCodeAnyOf {
	return &v
}

type NullableRacsFailureCodeAnyOf struct {
	value *RacsFailureCodeAnyOf
	isSet bool
}

func (v NullableRacsFailureCodeAnyOf) Get() *RacsFailureCodeAnyOf {
	return v.value
}

func (v *NullableRacsFailureCodeAnyOf) Set(val *RacsFailureCodeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRacsFailureCodeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRacsFailureCodeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRacsFailureCodeAnyOf(val *RacsFailureCodeAnyOf) *NullableRacsFailureCodeAnyOf {
	return &NullableRacsFailureCodeAnyOf{value: val, isSet: true}
}

func (v NullableRacsFailureCodeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRacsFailureCodeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

