/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// RequestDirectionAnyOf the model 'RequestDirectionAnyOf'
type RequestDirectionAnyOf string

// List of RequestDirection_anyOf
const (
	ORIGINATING_REGISTERED RequestDirectionAnyOf = "ORIGINATING_REGISTERED"
	ORIGINATING_UNREGISTERED RequestDirectionAnyOf = "ORIGINATING_UNREGISTERED"
	ORIGINATING_CDIV RequestDirectionAnyOf = "ORIGINATING_CDIV"
	TERMINATING_REGISTERED RequestDirectionAnyOf = "TERMINATING_REGISTERED"
	TERMINATING_UNREGISTERED RequestDirectionAnyOf = "TERMINATING_UNREGISTERED"
)

// All allowed values of RequestDirectionAnyOf enum
var AllowedRequestDirectionAnyOfEnumValues = []RequestDirectionAnyOf{
	"ORIGINATING_REGISTERED",
	"ORIGINATING_UNREGISTERED",
	"ORIGINATING_CDIV",
	"TERMINATING_REGISTERED",
	"TERMINATING_UNREGISTERED",
}

func (v *RequestDirectionAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestDirectionAnyOf(value)
	for _, existing := range AllowedRequestDirectionAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestDirectionAnyOf", value)
}

// NewRequestDirectionAnyOfFromValue returns a pointer to a valid RequestDirectionAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestDirectionAnyOfFromValue(v string) (*RequestDirectionAnyOf, error) {
	ev := RequestDirectionAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestDirectionAnyOf: valid values are %v", v, AllowedRequestDirectionAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestDirectionAnyOf) IsValid() bool {
	for _, existing := range AllowedRequestDirectionAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequestDirection_anyOf value
func (v RequestDirectionAnyOf) Ptr() *RequestDirectionAnyOf {
	return &v
}

type NullableRequestDirectionAnyOf struct {
	value *RequestDirectionAnyOf
	isSet bool
}

func (v NullableRequestDirectionAnyOf) Get() *RequestDirectionAnyOf {
	return v.value
}

func (v *NullableRequestDirectionAnyOf) Set(val *RequestDirectionAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestDirectionAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestDirectionAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestDirectionAnyOf(val *RequestDirectionAnyOf) *NullableRequestDirectionAnyOf {
	return &NullableRequestDirectionAnyOf{value: val, isSet: true}
}

func (v NullableRequestDirectionAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestDirectionAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

