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

// StatusChangeAnyOf the model 'StatusChangeAnyOf'
type StatusChangeAnyOf string

// List of StatusChange_anyOf
const (
	UNAVAILABLE StatusChangeAnyOf = "AMF_UNAVAILABLE"
	AVAILABLE StatusChangeAnyOf = "AMF_AVAILABLE"
)

// All allowed values of StatusChangeAnyOf enum
var AllowedStatusChangeAnyOfEnumValues = []StatusChangeAnyOf{
	"AMF_UNAVAILABLE",
	"AMF_AVAILABLE",
}

func (v *StatusChangeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StatusChangeAnyOf(value)
	for _, existing := range AllowedStatusChangeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StatusChangeAnyOf", value)
}

// NewStatusChangeAnyOfFromValue returns a pointer to a valid StatusChangeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatusChangeAnyOfFromValue(v string) (*StatusChangeAnyOf, error) {
	ev := StatusChangeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StatusChangeAnyOf: valid values are %v", v, AllowedStatusChangeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StatusChangeAnyOf) IsValid() bool {
	for _, existing := range AllowedStatusChangeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusChange_anyOf value
func (v StatusChangeAnyOf) Ptr() *StatusChangeAnyOf {
	return &v
}

type NullableStatusChangeAnyOf struct {
	value *StatusChangeAnyOf
	isSet bool
}

func (v NullableStatusChangeAnyOf) Get() *StatusChangeAnyOf {
	return v.value
}

func (v *NullableStatusChangeAnyOf) Set(val *StatusChangeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusChangeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusChangeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusChangeAnyOf(val *StatusChangeAnyOf) *NullableStatusChangeAnyOf {
	return &NullableStatusChangeAnyOf{value: val, isSet: true}
}

func (v NullableStatusChangeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusChangeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

