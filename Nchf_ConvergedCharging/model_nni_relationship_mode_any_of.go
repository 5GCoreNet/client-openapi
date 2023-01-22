/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// NNIRelationshipModeAnyOf the model 'NNIRelationshipModeAnyOf'
type NNIRelationshipModeAnyOf string

// List of NNIRelationshipMode_anyOf
const (
	TRUSTED NNIRelationshipModeAnyOf = "TRUSTED"
	NON_TRUSTED NNIRelationshipModeAnyOf = "NON_TRUSTED"
)

// All allowed values of NNIRelationshipModeAnyOf enum
var AllowedNNIRelationshipModeAnyOfEnumValues = []NNIRelationshipModeAnyOf{
	"TRUSTED",
	"NON_TRUSTED",
}

func (v *NNIRelationshipModeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NNIRelationshipModeAnyOf(value)
	for _, existing := range AllowedNNIRelationshipModeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NNIRelationshipModeAnyOf", value)
}

// NewNNIRelationshipModeAnyOfFromValue returns a pointer to a valid NNIRelationshipModeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNNIRelationshipModeAnyOfFromValue(v string) (*NNIRelationshipModeAnyOf, error) {
	ev := NNIRelationshipModeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NNIRelationshipModeAnyOf: valid values are %v", v, AllowedNNIRelationshipModeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NNIRelationshipModeAnyOf) IsValid() bool {
	for _, existing := range AllowedNNIRelationshipModeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NNIRelationshipMode_anyOf value
func (v NNIRelationshipModeAnyOf) Ptr() *NNIRelationshipModeAnyOf {
	return &v
}

type NullableNNIRelationshipModeAnyOf struct {
	value *NNIRelationshipModeAnyOf
	isSet bool
}

func (v NullableNNIRelationshipModeAnyOf) Get() *NNIRelationshipModeAnyOf {
	return v.value
}

func (v *NullableNNIRelationshipModeAnyOf) Set(val *NNIRelationshipModeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNNIRelationshipModeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNNIRelationshipModeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNNIRelationshipModeAnyOf(val *NNIRelationshipModeAnyOf) *NullableNNIRelationshipModeAnyOf {
	return &NullableNNIRelationshipModeAnyOf{value: val, isSet: true}
}

func (v NullableNNIRelationshipModeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNNIRelationshipModeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

