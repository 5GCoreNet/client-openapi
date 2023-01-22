/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// SmfChangeIndicationAnyOf the model 'SmfChangeIndicationAnyOf'
type SmfChangeIndicationAnyOf string

// List of SmfChangeIndication_anyOf
const (
	CHANGED SmfChangeIndicationAnyOf = "CHANGED"
	REMOVED SmfChangeIndicationAnyOf = "REMOVED"
)

// All allowed values of SmfChangeIndicationAnyOf enum
var AllowedSmfChangeIndicationAnyOfEnumValues = []SmfChangeIndicationAnyOf{
	"CHANGED",
	"REMOVED",
}

func (v *SmfChangeIndicationAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SmfChangeIndicationAnyOf(value)
	for _, existing := range AllowedSmfChangeIndicationAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SmfChangeIndicationAnyOf", value)
}

// NewSmfChangeIndicationAnyOfFromValue returns a pointer to a valid SmfChangeIndicationAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSmfChangeIndicationAnyOfFromValue(v string) (*SmfChangeIndicationAnyOf, error) {
	ev := SmfChangeIndicationAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SmfChangeIndicationAnyOf: valid values are %v", v, AllowedSmfChangeIndicationAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SmfChangeIndicationAnyOf) IsValid() bool {
	for _, existing := range AllowedSmfChangeIndicationAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SmfChangeIndication_anyOf value
func (v SmfChangeIndicationAnyOf) Ptr() *SmfChangeIndicationAnyOf {
	return &v
}

type NullableSmfChangeIndicationAnyOf struct {
	value *SmfChangeIndicationAnyOf
	isSet bool
}

func (v NullableSmfChangeIndicationAnyOf) Get() *SmfChangeIndicationAnyOf {
	return v.value
}

func (v *NullableSmfChangeIndicationAnyOf) Set(val *SmfChangeIndicationAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSmfChangeIndicationAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSmfChangeIndicationAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmfChangeIndicationAnyOf(val *SmfChangeIndicationAnyOf) *NullableSmfChangeIndicationAnyOf {
	return &NullableSmfChangeIndicationAnyOf{value: val, isSet: true}
}

func (v NullableSmfChangeIndicationAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmfChangeIndicationAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
