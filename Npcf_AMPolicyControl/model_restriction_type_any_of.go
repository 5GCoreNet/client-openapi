/*
Npcf_AMPolicyControl

Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_AMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// RestrictionTypeAnyOf the model 'RestrictionTypeAnyOf'
type RestrictionTypeAnyOf string

// List of RestrictionType_anyOf
const (
	ALLOWED_AREAS RestrictionTypeAnyOf = "ALLOWED_AREAS"
	NOT_ALLOWED_AREAS RestrictionTypeAnyOf = "NOT_ALLOWED_AREAS"
)

// All allowed values of RestrictionTypeAnyOf enum
var AllowedRestrictionTypeAnyOfEnumValues = []RestrictionTypeAnyOf{
	"ALLOWED_AREAS",
	"NOT_ALLOWED_AREAS",
}

func (v *RestrictionTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RestrictionTypeAnyOf(value)
	for _, existing := range AllowedRestrictionTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RestrictionTypeAnyOf", value)
}

// NewRestrictionTypeAnyOfFromValue returns a pointer to a valid RestrictionTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRestrictionTypeAnyOfFromValue(v string) (*RestrictionTypeAnyOf, error) {
	ev := RestrictionTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RestrictionTypeAnyOf: valid values are %v", v, AllowedRestrictionTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RestrictionTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedRestrictionTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RestrictionType_anyOf value
func (v RestrictionTypeAnyOf) Ptr() *RestrictionTypeAnyOf {
	return &v
}

type NullableRestrictionTypeAnyOf struct {
	value *RestrictionTypeAnyOf
	isSet bool
}

func (v NullableRestrictionTypeAnyOf) Get() *RestrictionTypeAnyOf {
	return v.value
}

func (v *NullableRestrictionTypeAnyOf) Set(val *RestrictionTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRestrictionTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRestrictionTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestrictionTypeAnyOf(val *RestrictionTypeAnyOf) *NullableRestrictionTypeAnyOf {
	return &NullableRestrictionTypeAnyOf{value: val, isSet: true}
}

func (v NullableRestrictionTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestrictionTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

