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

// SDPTypeAnyOf the model 'SDPTypeAnyOf'
type SDPTypeAnyOf string

// List of SDPType_anyOf
const (
	OFFER SDPTypeAnyOf = "OFFER"
	ANSWER SDPTypeAnyOf = "ANSWER"
)

// All allowed values of SDPTypeAnyOf enum
var AllowedSDPTypeAnyOfEnumValues = []SDPTypeAnyOf{
	"OFFER",
	"ANSWER",
}

func (v *SDPTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SDPTypeAnyOf(value)
	for _, existing := range AllowedSDPTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SDPTypeAnyOf", value)
}

// NewSDPTypeAnyOfFromValue returns a pointer to a valid SDPTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSDPTypeAnyOfFromValue(v string) (*SDPTypeAnyOf, error) {
	ev := SDPTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SDPTypeAnyOf: valid values are %v", v, AllowedSDPTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SDPTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedSDPTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SDPType_anyOf value
func (v SDPTypeAnyOf) Ptr() *SDPTypeAnyOf {
	return &v
}

type NullableSDPTypeAnyOf struct {
	value *SDPTypeAnyOf
	isSet bool
}

func (v NullableSDPTypeAnyOf) Get() *SDPTypeAnyOf {
	return v.value
}

func (v *NullableSDPTypeAnyOf) Set(val *SDPTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSDPTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSDPTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDPTypeAnyOf(val *SDPTypeAnyOf) *NullableSDPTypeAnyOf {
	return &NullableSDPTypeAnyOf{value: val, isSet: true}
}

func (v NullableSDPTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSDPTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

