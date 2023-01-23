/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
	"fmt"
)

// CollocatedNfTypeAnyOf the model 'CollocatedNfTypeAnyOf'
type CollocatedNfTypeAnyOf string

// List of CollocatedNfType_anyOf
const (
	UPF CollocatedNfTypeAnyOf = "UPF"
	SMF CollocatedNfTypeAnyOf = "SMF"
	MB_UPF CollocatedNfTypeAnyOf = "MB_UPF"
	MB_SMF CollocatedNfTypeAnyOf = "MB_SMF"
)

// All allowed values of CollocatedNfTypeAnyOf enum
var AllowedCollocatedNfTypeAnyOfEnumValues = []CollocatedNfTypeAnyOf{
	"UPF",
	"SMF",
	"MB_UPF",
	"MB_SMF",
}

func (v *CollocatedNfTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollocatedNfTypeAnyOf(value)
	for _, existing := range AllowedCollocatedNfTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollocatedNfTypeAnyOf", value)
}

// NewCollocatedNfTypeAnyOfFromValue returns a pointer to a valid CollocatedNfTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollocatedNfTypeAnyOfFromValue(v string) (*CollocatedNfTypeAnyOf, error) {
	ev := CollocatedNfTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollocatedNfTypeAnyOf: valid values are %v", v, AllowedCollocatedNfTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollocatedNfTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedCollocatedNfTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CollocatedNfType_anyOf value
func (v CollocatedNfTypeAnyOf) Ptr() *CollocatedNfTypeAnyOf {
	return &v
}

type NullableCollocatedNfTypeAnyOf struct {
	value *CollocatedNfTypeAnyOf
	isSet bool
}

func (v NullableCollocatedNfTypeAnyOf) Get() *CollocatedNfTypeAnyOf {
	return v.value
}

func (v *NullableCollocatedNfTypeAnyOf) Set(val *CollocatedNfTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCollocatedNfTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCollocatedNfTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollocatedNfTypeAnyOf(val *CollocatedNfTypeAnyOf) *NullableCollocatedNfTypeAnyOf {
	return &NullableCollocatedNfTypeAnyOf{value: val, isSet: true}
}

func (v NullableCollocatedNfTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollocatedNfTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

