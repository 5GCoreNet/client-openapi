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

// NNITypeAnyOf the model 'NNITypeAnyOf'
type NNITypeAnyOf string

// List of NNIType_anyOf
const (
	NON_ROAMING NNITypeAnyOf = "NON_ROAMING"
	ROAMING_NO_LOOPBACK NNITypeAnyOf = "ROAMING_NO_LOOPBACK"
	ROAMING_LOOPBACK NNITypeAnyOf = "ROAMING_LOOPBACK"
)

// All allowed values of NNITypeAnyOf enum
var AllowedNNITypeAnyOfEnumValues = []NNITypeAnyOf{
	"NON_ROAMING",
	"ROAMING_NO_LOOPBACK",
	"ROAMING_LOOPBACK",
}

func (v *NNITypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NNITypeAnyOf(value)
	for _, existing := range AllowedNNITypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NNITypeAnyOf", value)
}

// NewNNITypeAnyOfFromValue returns a pointer to a valid NNITypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNNITypeAnyOfFromValue(v string) (*NNITypeAnyOf, error) {
	ev := NNITypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NNITypeAnyOf: valid values are %v", v, AllowedNNITypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NNITypeAnyOf) IsValid() bool {
	for _, existing := range AllowedNNITypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NNIType_anyOf value
func (v NNITypeAnyOf) Ptr() *NNITypeAnyOf {
	return &v
}

type NullableNNITypeAnyOf struct {
	value *NNITypeAnyOf
	isSet bool
}

func (v NullableNNITypeAnyOf) Get() *NNITypeAnyOf {
	return v.value
}

func (v *NullableNNITypeAnyOf) Set(val *NNITypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNNITypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNNITypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNNITypeAnyOf(val *NNITypeAnyOf) *NullableNNITypeAnyOf {
	return &NullableNNITypeAnyOf{value: val, isSet: true}
}

func (v NullableNNITypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNNITypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
