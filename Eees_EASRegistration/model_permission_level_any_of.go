/*
EES EAS Registration_API

API for EAS Registration.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_EASRegistration

import (
	"encoding/json"
	"fmt"
)

// PermissionLevelAnyOf the model 'PermissionLevelAnyOf'
type PermissionLevelAnyOf string

// List of PermissionLevel_anyOf
const (
	TRIAL PermissionLevelAnyOf = "TRIAL"
	GOLD PermissionLevelAnyOf = "GOLD"
	SILVER PermissionLevelAnyOf = "SILVER"
	OTHER PermissionLevelAnyOf = "OTHER"
)

// All allowed values of PermissionLevelAnyOf enum
var AllowedPermissionLevelAnyOfEnumValues = []PermissionLevelAnyOf{
	"TRIAL",
	"GOLD",
	"SILVER",
	"OTHER",
}

func (v *PermissionLevelAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PermissionLevelAnyOf(value)
	for _, existing := range AllowedPermissionLevelAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PermissionLevelAnyOf", value)
}

// NewPermissionLevelAnyOfFromValue returns a pointer to a valid PermissionLevelAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPermissionLevelAnyOfFromValue(v string) (*PermissionLevelAnyOf, error) {
	ev := PermissionLevelAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PermissionLevelAnyOf: valid values are %v", v, AllowedPermissionLevelAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PermissionLevelAnyOf) IsValid() bool {
	for _, existing := range AllowedPermissionLevelAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PermissionLevel_anyOf value
func (v PermissionLevelAnyOf) Ptr() *PermissionLevelAnyOf {
	return &v
}

type NullablePermissionLevelAnyOf struct {
	value *PermissionLevelAnyOf
	isSet bool
}

func (v NullablePermissionLevelAnyOf) Get() *PermissionLevelAnyOf {
	return v.value
}

func (v *NullablePermissionLevelAnyOf) Set(val *PermissionLevelAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionLevelAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionLevelAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionLevelAnyOf(val *PermissionLevelAnyOf) *NullablePermissionLevelAnyOf {
	return &NullablePermissionLevelAnyOf{value: val, isSet: true}
}

func (v NullablePermissionLevelAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionLevelAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

