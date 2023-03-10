/*
EES ACR Status Update Service

EES ACR Status Update Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRStatusUpdate

import (
	"encoding/json"
	"fmt"
)

// ACTResultAnyOf the model 'ACTResultAnyOf'
type ACTResultAnyOf string

// List of ACTResult_anyOf
const (
	SUCCESSFUL ACTResultAnyOf = "SUCCESSFUL"
	FAILED ACTResultAnyOf = "FAILED"
)

// All allowed values of ACTResultAnyOf enum
var AllowedACTResultAnyOfEnumValues = []ACTResultAnyOf{
	"SUCCESSFUL",
	"FAILED",
}

func (v *ACTResultAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ACTResultAnyOf(value)
	for _, existing := range AllowedACTResultAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ACTResultAnyOf", value)
}

// NewACTResultAnyOfFromValue returns a pointer to a valid ACTResultAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewACTResultAnyOfFromValue(v string) (*ACTResultAnyOf, error) {
	ev := ACTResultAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ACTResultAnyOf: valid values are %v", v, AllowedACTResultAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ACTResultAnyOf) IsValid() bool {
	for _, existing := range AllowedACTResultAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ACTResult_anyOf value
func (v ACTResultAnyOf) Ptr() *ACTResultAnyOf {
	return &v
}

type NullableACTResultAnyOf struct {
	value *ACTResultAnyOf
	isSet bool
}

func (v NullableACTResultAnyOf) Get() *ACTResultAnyOf {
	return v.value
}

func (v *NullableACTResultAnyOf) Set(val *ACTResultAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableACTResultAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableACTResultAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACTResultAnyOf(val *ACTResultAnyOf) *NullableACTResultAnyOf {
	return &NullableACTResultAnyOf{value: val, isSet: true}
}

func (v NullableACTResultAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACTResultAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

