/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// ReflectiveQoSAttributeAnyOf the model 'ReflectiveQoSAttributeAnyOf'
type ReflectiveQoSAttributeAnyOf string

// List of ReflectiveQoSAttribute_anyOf
const (
	RQOS ReflectiveQoSAttributeAnyOf = "RQOS"
	NO_RQOS ReflectiveQoSAttributeAnyOf = "NO_RQOS"
)

// All allowed values of ReflectiveQoSAttributeAnyOf enum
var AllowedReflectiveQoSAttributeAnyOfEnumValues = []ReflectiveQoSAttributeAnyOf{
	"RQOS",
	"NO_RQOS",
}

func (v *ReflectiveQoSAttributeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReflectiveQoSAttributeAnyOf(value)
	for _, existing := range AllowedReflectiveQoSAttributeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReflectiveQoSAttributeAnyOf", value)
}

// NewReflectiveQoSAttributeAnyOfFromValue returns a pointer to a valid ReflectiveQoSAttributeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReflectiveQoSAttributeAnyOfFromValue(v string) (*ReflectiveQoSAttributeAnyOf, error) {
	ev := ReflectiveQoSAttributeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReflectiveQoSAttributeAnyOf: valid values are %v", v, AllowedReflectiveQoSAttributeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReflectiveQoSAttributeAnyOf) IsValid() bool {
	for _, existing := range AllowedReflectiveQoSAttributeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReflectiveQoSAttribute_anyOf value
func (v ReflectiveQoSAttributeAnyOf) Ptr() *ReflectiveQoSAttributeAnyOf {
	return &v
}

type NullableReflectiveQoSAttributeAnyOf struct {
	value *ReflectiveQoSAttributeAnyOf
	isSet bool
}

func (v NullableReflectiveQoSAttributeAnyOf) Get() *ReflectiveQoSAttributeAnyOf {
	return v.value
}

func (v *NullableReflectiveQoSAttributeAnyOf) Set(val *ReflectiveQoSAttributeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReflectiveQoSAttributeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReflectiveQoSAttributeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReflectiveQoSAttributeAnyOf(val *ReflectiveQoSAttributeAnyOf) *NullableReflectiveQoSAttributeAnyOf {
	return &NullableReflectiveQoSAttributeAnyOf{value: val, isSet: true}
}

func (v NullableReflectiveQoSAttributeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReflectiveQoSAttributeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

