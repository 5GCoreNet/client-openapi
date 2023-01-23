/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// FailureHandlingAnyOf the model 'FailureHandlingAnyOf'
type FailureHandlingAnyOf string

// List of FailureHandling_anyOf
const (
	TERMINATE FailureHandlingAnyOf = "TERMINATE"
	CONTINUE FailureHandlingAnyOf = "CONTINUE"
	RETRY_AND_TERMINATE FailureHandlingAnyOf = "RETRY_AND_TERMINATE"
)

// All allowed values of FailureHandlingAnyOf enum
var AllowedFailureHandlingAnyOfEnumValues = []FailureHandlingAnyOf{
	"TERMINATE",
	"CONTINUE",
	"RETRY_AND_TERMINATE",
}

func (v *FailureHandlingAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FailureHandlingAnyOf(value)
	for _, existing := range AllowedFailureHandlingAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FailureHandlingAnyOf", value)
}

// NewFailureHandlingAnyOfFromValue returns a pointer to a valid FailureHandlingAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFailureHandlingAnyOfFromValue(v string) (*FailureHandlingAnyOf, error) {
	ev := FailureHandlingAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FailureHandlingAnyOf: valid values are %v", v, AllowedFailureHandlingAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FailureHandlingAnyOf) IsValid() bool {
	for _, existing := range AllowedFailureHandlingAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FailureHandling_anyOf value
func (v FailureHandlingAnyOf) Ptr() *FailureHandlingAnyOf {
	return &v
}

type NullableFailureHandlingAnyOf struct {
	value *FailureHandlingAnyOf
	isSet bool
}

func (v NullableFailureHandlingAnyOf) Get() *FailureHandlingAnyOf {
	return v.value
}

func (v *NullableFailureHandlingAnyOf) Set(val *FailureHandlingAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureHandlingAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureHandlingAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureHandlingAnyOf(val *FailureHandlingAnyOf) *NullableFailureHandlingAnyOf {
	return &NullableFailureHandlingAnyOf{value: val, isSet: true}
}

func (v NullableFailureHandlingAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureHandlingAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

