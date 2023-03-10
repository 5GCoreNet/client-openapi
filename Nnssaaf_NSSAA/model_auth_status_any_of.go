/*
Nnssaaf_NSSAA

Network Slice-Specific Authentication and Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnssaaf_NSSAA

import (
	"encoding/json"
	"fmt"
)

// AuthStatusAnyOf the model 'AuthStatusAnyOf'
type AuthStatusAnyOf string

// List of AuthStatus_anyOf
const (
	EAP_SUCCESS AuthStatusAnyOf = "EAP_SUCCESS"
	EAP_FAILURE AuthStatusAnyOf = "EAP_FAILURE"
	PENDING AuthStatusAnyOf = "PENDING"
)

// All allowed values of AuthStatusAnyOf enum
var AllowedAuthStatusAnyOfEnumValues = []AuthStatusAnyOf{
	"EAP_SUCCESS",
	"EAP_FAILURE",
	"PENDING",
}

func (v *AuthStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthStatusAnyOf(value)
	for _, existing := range AllowedAuthStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthStatusAnyOf", value)
}

// NewAuthStatusAnyOfFromValue returns a pointer to a valid AuthStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthStatusAnyOfFromValue(v string) (*AuthStatusAnyOf, error) {
	ev := AuthStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthStatusAnyOf: valid values are %v", v, AllowedAuthStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedAuthStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthStatus_anyOf value
func (v AuthStatusAnyOf) Ptr() *AuthStatusAnyOf {
	return &v
}

type NullableAuthStatusAnyOf struct {
	value *AuthStatusAnyOf
	isSet bool
}

func (v NullableAuthStatusAnyOf) Get() *AuthStatusAnyOf {
	return v.value
}

func (v *NullableAuthStatusAnyOf) Set(val *AuthStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthStatusAnyOf(val *AuthStatusAnyOf) *NullableAuthStatusAnyOf {
	return &NullableAuthStatusAnyOf{value: val, isSet: true}
}

func (v NullableAuthStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

