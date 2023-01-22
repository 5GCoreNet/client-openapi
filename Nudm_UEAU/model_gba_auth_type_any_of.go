/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_UEAU

import (
	"encoding/json"
	"fmt"
)

// GbaAuthTypeAnyOf the model 'GbaAuthTypeAnyOf'
type GbaAuthTypeAnyOf string

// List of GbaAuthType_anyOf
const (
	DIGEST_AKAV1_MD5 GbaAuthTypeAnyOf = "DIGEST_AKAV1_MD5"
)

// All allowed values of GbaAuthTypeAnyOf enum
var AllowedGbaAuthTypeAnyOfEnumValues = []GbaAuthTypeAnyOf{
	"DIGEST_AKAV1_MD5",
}

func (v *GbaAuthTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GbaAuthTypeAnyOf(value)
	for _, existing := range AllowedGbaAuthTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GbaAuthTypeAnyOf", value)
}

// NewGbaAuthTypeAnyOfFromValue returns a pointer to a valid GbaAuthTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGbaAuthTypeAnyOfFromValue(v string) (*GbaAuthTypeAnyOf, error) {
	ev := GbaAuthTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GbaAuthTypeAnyOf: valid values are %v", v, AllowedGbaAuthTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GbaAuthTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedGbaAuthTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GbaAuthType_anyOf value
func (v GbaAuthTypeAnyOf) Ptr() *GbaAuthTypeAnyOf {
	return &v
}

type NullableGbaAuthTypeAnyOf struct {
	value *GbaAuthTypeAnyOf
	isSet bool
}

func (v NullableGbaAuthTypeAnyOf) Get() *GbaAuthTypeAnyOf {
	return v.value
}

func (v *NullableGbaAuthTypeAnyOf) Set(val *GbaAuthTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGbaAuthTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGbaAuthTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGbaAuthTypeAnyOf(val *GbaAuthTypeAnyOf) *NullableGbaAuthTypeAnyOf {
	return &NullableGbaAuthTypeAnyOf{value: val, isSet: true}
}

func (v NullableGbaAuthTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGbaAuthTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

