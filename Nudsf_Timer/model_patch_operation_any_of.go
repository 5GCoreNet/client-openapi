/*
Nudsf_Timer

Nudsf Timer Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudsf_Timer

import (
	"encoding/json"
	"fmt"
)

// PatchOperationAnyOf the model 'PatchOperationAnyOf'
type PatchOperationAnyOf string

// List of PatchOperation_anyOf
const (
	ADD PatchOperationAnyOf = "add"
	COPY PatchOperationAnyOf = "copy"
	MOVE PatchOperationAnyOf = "move"
	REMOVE PatchOperationAnyOf = "remove"
	REPLACE PatchOperationAnyOf = "replace"
	TEST PatchOperationAnyOf = "test"
)

// All allowed values of PatchOperationAnyOf enum
var AllowedPatchOperationAnyOfEnumValues = []PatchOperationAnyOf{
	"add",
	"copy",
	"move",
	"remove",
	"replace",
	"test",
}

func (v *PatchOperationAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchOperationAnyOf(value)
	for _, existing := range AllowedPatchOperationAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchOperationAnyOf", value)
}

// NewPatchOperationAnyOfFromValue returns a pointer to a valid PatchOperationAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchOperationAnyOfFromValue(v string) (*PatchOperationAnyOf, error) {
	ev := PatchOperationAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchOperationAnyOf: valid values are %v", v, AllowedPatchOperationAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchOperationAnyOf) IsValid() bool {
	for _, existing := range AllowedPatchOperationAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchOperation_anyOf value
func (v PatchOperationAnyOf) Ptr() *PatchOperationAnyOf {
	return &v
}

type NullablePatchOperationAnyOf struct {
	value *PatchOperationAnyOf
	isSet bool
}

func (v NullablePatchOperationAnyOf) Get() *PatchOperationAnyOf {
	return v.value
}

func (v *NullablePatchOperationAnyOf) Set(val *PatchOperationAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchOperationAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchOperationAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchOperationAnyOf(val *PatchOperationAnyOf) *NullablePatchOperationAnyOf {
	return &NullablePatchOperationAnyOf{value: val, isSet: true}
}

func (v NullablePatchOperationAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchOperationAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

