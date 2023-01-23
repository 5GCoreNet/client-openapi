/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_OfflineOnlyCharging

import (
	"encoding/json"
	"fmt"
)

// SscModeAnyOf the model 'SscModeAnyOf'
type SscModeAnyOf string

// List of SscMode_anyOf
const (
	_1 SscModeAnyOf = "SSC_MODE_1"
	_2 SscModeAnyOf = "SSC_MODE_2"
	_3 SscModeAnyOf = "SSC_MODE_3"
)

// All allowed values of SscModeAnyOf enum
var AllowedSscModeAnyOfEnumValues = []SscModeAnyOf{
	"SSC_MODE_1",
	"SSC_MODE_2",
	"SSC_MODE_3",
}

func (v *SscModeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SscModeAnyOf(value)
	for _, existing := range AllowedSscModeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SscModeAnyOf", value)
}

// NewSscModeAnyOfFromValue returns a pointer to a valid SscModeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSscModeAnyOfFromValue(v string) (*SscModeAnyOf, error) {
	ev := SscModeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SscModeAnyOf: valid values are %v", v, AllowedSscModeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SscModeAnyOf) IsValid() bool {
	for _, existing := range AllowedSscModeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SscMode_anyOf value
func (v SscModeAnyOf) Ptr() *SscModeAnyOf {
	return &v
}

type NullableSscModeAnyOf struct {
	value *SscModeAnyOf
	isSet bool
}

func (v NullableSscModeAnyOf) Get() *SscModeAnyOf {
	return v.value
}

func (v *NullableSscModeAnyOf) Set(val *SscModeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSscModeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSscModeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSscModeAnyOf(val *SscModeAnyOf) *NullableSscModeAnyOf {
	return &NullableSscModeAnyOf{value: val, isSet: true}
}

func (v NullableSscModeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSscModeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

