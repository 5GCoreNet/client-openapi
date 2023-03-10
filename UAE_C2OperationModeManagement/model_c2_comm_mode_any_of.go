/*
UAE Server C2 Operation Mode Management Service

UAE Server C2 Operation Mode Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_UAE_C2OperationModeManagement

import (
	"encoding/json"
	"fmt"
)

// C2CommModeAnyOf the model 'C2CommModeAnyOf'
type C2CommModeAnyOf string

// List of C2CommMode_anyOf
const (
	DIRECT_C2_COMMUNICATION C2CommModeAnyOf = "DIRECT_C2_COMMUNICATION"
	NETWORK_ASSISTED_C2_COMMUNICATION C2CommModeAnyOf = "NETWORK_ASSISTED_C2_COMMUNICATION"
	UTM_NAVIGATED_C2_COMMUNICATION C2CommModeAnyOf = "UTM_NAVIGATED_C2_COMMUNICATION"
)

// All allowed values of C2CommModeAnyOf enum
var AllowedC2CommModeAnyOfEnumValues = []C2CommModeAnyOf{
	"DIRECT_C2_COMMUNICATION",
	"NETWORK_ASSISTED_C2_COMMUNICATION",
	"UTM_NAVIGATED_C2_COMMUNICATION",
}

func (v *C2CommModeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := C2CommModeAnyOf(value)
	for _, existing := range AllowedC2CommModeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid C2CommModeAnyOf", value)
}

// NewC2CommModeAnyOfFromValue returns a pointer to a valid C2CommModeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewC2CommModeAnyOfFromValue(v string) (*C2CommModeAnyOf, error) {
	ev := C2CommModeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for C2CommModeAnyOf: valid values are %v", v, AllowedC2CommModeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v C2CommModeAnyOf) IsValid() bool {
	for _, existing := range AllowedC2CommModeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to C2CommMode_anyOf value
func (v C2CommModeAnyOf) Ptr() *C2CommModeAnyOf {
	return &v
}

type NullableC2CommModeAnyOf struct {
	value *C2CommModeAnyOf
	isSet bool
}

func (v NullableC2CommModeAnyOf) Get() *C2CommModeAnyOf {
	return v.value
}

func (v *NullableC2CommModeAnyOf) Set(val *C2CommModeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableC2CommModeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableC2CommModeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC2CommModeAnyOf(val *C2CommModeAnyOf) *NullableC2CommModeAnyOf {
	return &NullableC2CommModeAnyOf{value: val, isSet: true}
}

func (v NullableC2CommModeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC2CommModeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

