/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nlmf_Location

import (
	"encoding/json"
	"fmt"
)

// ResponseTimeAnyOf the model 'ResponseTimeAnyOf'
type ResponseTimeAnyOf string

// List of ResponseTime_anyOf
const (
	LOW_DELAY ResponseTimeAnyOf = "LOW_DELAY"
	DELAY_TOLERANT ResponseTimeAnyOf = "DELAY_TOLERANT"
	NO_DELAY ResponseTimeAnyOf = "NO_DELAY"
)

// All allowed values of ResponseTimeAnyOf enum
var AllowedResponseTimeAnyOfEnumValues = []ResponseTimeAnyOf{
	"LOW_DELAY",
	"DELAY_TOLERANT",
	"NO_DELAY",
}

func (v *ResponseTimeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResponseTimeAnyOf(value)
	for _, existing := range AllowedResponseTimeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResponseTimeAnyOf", value)
}

// NewResponseTimeAnyOfFromValue returns a pointer to a valid ResponseTimeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResponseTimeAnyOfFromValue(v string) (*ResponseTimeAnyOf, error) {
	ev := ResponseTimeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResponseTimeAnyOf: valid values are %v", v, AllowedResponseTimeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResponseTimeAnyOf) IsValid() bool {
	for _, existing := range AllowedResponseTimeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResponseTime_anyOf value
func (v ResponseTimeAnyOf) Ptr() *ResponseTimeAnyOf {
	return &v
}

type NullableResponseTimeAnyOf struct {
	value *ResponseTimeAnyOf
	isSet bool
}

func (v NullableResponseTimeAnyOf) Get() *ResponseTimeAnyOf {
	return v.value
}

func (v *NullableResponseTimeAnyOf) Set(val *ResponseTimeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTimeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTimeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTimeAnyOf(val *ResponseTimeAnyOf) *NullableResponseTimeAnyOf {
	return &NullableResponseTimeAnyOf{value: val, isSet: true}
}

func (v NullableResponseTimeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTimeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

