/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// RequestTypeAnyOf the model 'RequestTypeAnyOf'
type RequestTypeAnyOf string

// List of RequestType_anyOf
const (
	INITIAL_REQUEST RequestTypeAnyOf = "INITIAL_REQUEST"
	EXISTING_PDU_SESSION RequestTypeAnyOf = "EXISTING_PDU_SESSION"
	INITIAL_EMERGENCY_REQUEST RequestTypeAnyOf = "INITIAL_EMERGENCY_REQUEST"
	EXISTING_EMERGENCY_PDU_SESSION RequestTypeAnyOf = "EXISTING_EMERGENCY_PDU_SESSION"
)

// All allowed values of RequestTypeAnyOf enum
var AllowedRequestTypeAnyOfEnumValues = []RequestTypeAnyOf{
	"INITIAL_REQUEST",
	"EXISTING_PDU_SESSION",
	"INITIAL_EMERGENCY_REQUEST",
	"EXISTING_EMERGENCY_PDU_SESSION",
}

func (v *RequestTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestTypeAnyOf(value)
	for _, existing := range AllowedRequestTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestTypeAnyOf", value)
}

// NewRequestTypeAnyOfFromValue returns a pointer to a valid RequestTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestTypeAnyOfFromValue(v string) (*RequestTypeAnyOf, error) {
	ev := RequestTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestTypeAnyOf: valid values are %v", v, AllowedRequestTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedRequestTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequestType_anyOf value
func (v RequestTypeAnyOf) Ptr() *RequestTypeAnyOf {
	return &v
}

type NullableRequestTypeAnyOf struct {
	value *RequestTypeAnyOf
	isSet bool
}

func (v NullableRequestTypeAnyOf) Get() *RequestTypeAnyOf {
	return v.value
}

func (v *NullableRequestTypeAnyOf) Set(val *RequestTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeAnyOf(val *RequestTypeAnyOf) *NullableRequestTypeAnyOf {
	return &NullableRequestTypeAnyOf{value: val, isSet: true}
}

func (v NullableRequestTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

