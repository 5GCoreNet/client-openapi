/*
3gpp-mo-lcs-notify

API for UE updated location information notification.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MoLcsNotify

import (
	"encoding/json"
	"fmt"
)

// LdrTypeAnyOf the model 'LdrTypeAnyOf'
type LdrTypeAnyOf string

// List of LdrType_anyOf
const (
	UE_AVAILABLE LdrTypeAnyOf = "UE_AVAILABLE"
	PERIODIC LdrTypeAnyOf = "PERIODIC"
	ENTERING_INTO_AREA LdrTypeAnyOf = "ENTERING_INTO_AREA"
	LEAVING_FROM_AREA LdrTypeAnyOf = "LEAVING_FROM_AREA"
	BEING_INSIDE_AREA LdrTypeAnyOf = "BEING_INSIDE_AREA"
	MOTION LdrTypeAnyOf = "MOTION"
)

// All allowed values of LdrTypeAnyOf enum
var AllowedLdrTypeAnyOfEnumValues = []LdrTypeAnyOf{
	"UE_AVAILABLE",
	"PERIODIC",
	"ENTERING_INTO_AREA",
	"LEAVING_FROM_AREA",
	"BEING_INSIDE_AREA",
	"MOTION",
}

func (v *LdrTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LdrTypeAnyOf(value)
	for _, existing := range AllowedLdrTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LdrTypeAnyOf", value)
}

// NewLdrTypeAnyOfFromValue returns a pointer to a valid LdrTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLdrTypeAnyOfFromValue(v string) (*LdrTypeAnyOf, error) {
	ev := LdrTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LdrTypeAnyOf: valid values are %v", v, AllowedLdrTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LdrTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedLdrTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LdrType_anyOf value
func (v LdrTypeAnyOf) Ptr() *LdrTypeAnyOf {
	return &v
}

type NullableLdrTypeAnyOf struct {
	value *LdrTypeAnyOf
	isSet bool
}

func (v NullableLdrTypeAnyOf) Get() *LdrTypeAnyOf {
	return v.value
}

func (v *NullableLdrTypeAnyOf) Set(val *LdrTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLdrTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLdrTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdrTypeAnyOf(val *LdrTypeAnyOf) *NullableLdrTypeAnyOf {
	return &NullableLdrTypeAnyOf{value: val, isSet: true}
}

func (v NullableLdrTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdrTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

