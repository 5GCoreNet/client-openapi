/*
MSGS_MSGDelivery

API for MSGG MSGin5G Server Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MSGS_MSGDelivery

import (
	"encoding/json"
	"fmt"
)

// PriorityAnyOf the model 'PriorityAnyOf'
type PriorityAnyOf string

// List of Priority_anyOf
const (
	HIGH PriorityAnyOf = "HIGH"
	MIDDLE PriorityAnyOf = "MIDDLE"
	LOW PriorityAnyOf = "LOW"
)

// All allowed values of PriorityAnyOf enum
var AllowedPriorityAnyOfEnumValues = []PriorityAnyOf{
	"HIGH",
	"MIDDLE",
	"LOW",
}

func (v *PriorityAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PriorityAnyOf(value)
	for _, existing := range AllowedPriorityAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PriorityAnyOf", value)
}

// NewPriorityAnyOfFromValue returns a pointer to a valid PriorityAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPriorityAnyOfFromValue(v string) (*PriorityAnyOf, error) {
	ev := PriorityAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PriorityAnyOf: valid values are %v", v, AllowedPriorityAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PriorityAnyOf) IsValid() bool {
	for _, existing := range AllowedPriorityAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Priority_anyOf value
func (v PriorityAnyOf) Ptr() *PriorityAnyOf {
	return &v
}

type NullablePriorityAnyOf struct {
	value *PriorityAnyOf
	isSet bool
}

func (v NullablePriorityAnyOf) Get() *PriorityAnyOf {
	return v.value
}

func (v *NullablePriorityAnyOf) Set(val *PriorityAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePriorityAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePriorityAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriorityAnyOf(val *PriorityAnyOf) *NullablePriorityAnyOf {
	return &NullablePriorityAnyOf{value: val, isSet: true}
}

func (v NullablePriorityAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriorityAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
