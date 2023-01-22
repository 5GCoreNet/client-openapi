/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// BdtPolicyStatusAnyOf the model 'BdtPolicyStatusAnyOf'
type BdtPolicyStatusAnyOf string

// List of BdtPolicyStatus_anyOf
const (
	INVALID BdtPolicyStatusAnyOf = "INVALID"
	VALID BdtPolicyStatusAnyOf = "VALID"
)

// All allowed values of BdtPolicyStatusAnyOf enum
var AllowedBdtPolicyStatusAnyOfEnumValues = []BdtPolicyStatusAnyOf{
	"INVALID",
	"VALID",
}

func (v *BdtPolicyStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BdtPolicyStatusAnyOf(value)
	for _, existing := range AllowedBdtPolicyStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BdtPolicyStatusAnyOf", value)
}

// NewBdtPolicyStatusAnyOfFromValue returns a pointer to a valid BdtPolicyStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBdtPolicyStatusAnyOfFromValue(v string) (*BdtPolicyStatusAnyOf, error) {
	ev := BdtPolicyStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BdtPolicyStatusAnyOf: valid values are %v", v, AllowedBdtPolicyStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BdtPolicyStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedBdtPolicyStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BdtPolicyStatus_anyOf value
func (v BdtPolicyStatusAnyOf) Ptr() *BdtPolicyStatusAnyOf {
	return &v
}

type NullableBdtPolicyStatusAnyOf struct {
	value *BdtPolicyStatusAnyOf
	isSet bool
}

func (v NullableBdtPolicyStatusAnyOf) Get() *BdtPolicyStatusAnyOf {
	return v.value
}

func (v *NullableBdtPolicyStatusAnyOf) Set(val *BdtPolicyStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBdtPolicyStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBdtPolicyStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBdtPolicyStatusAnyOf(val *BdtPolicyStatusAnyOf) *NullableBdtPolicyStatusAnyOf {
	return &NullableBdtPolicyStatusAnyOf{value: val, isSet: true}
}

func (v NullableBdtPolicyStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBdtPolicyStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

