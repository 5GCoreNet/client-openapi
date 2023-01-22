/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// PdnConnectivityStatusAnyOf the model 'PdnConnectivityStatusAnyOf'
type PdnConnectivityStatusAnyOf string

// List of PdnConnectivityStatus_anyOf
const (
	ESTABLISHED PdnConnectivityStatusAnyOf = "ESTABLISHED"
	RELEASED PdnConnectivityStatusAnyOf = "RELEASED"
)

// All allowed values of PdnConnectivityStatusAnyOf enum
var AllowedPdnConnectivityStatusAnyOfEnumValues = []PdnConnectivityStatusAnyOf{
	"ESTABLISHED",
	"RELEASED",
}

func (v *PdnConnectivityStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PdnConnectivityStatusAnyOf(value)
	for _, existing := range AllowedPdnConnectivityStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PdnConnectivityStatusAnyOf", value)
}

// NewPdnConnectivityStatusAnyOfFromValue returns a pointer to a valid PdnConnectivityStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPdnConnectivityStatusAnyOfFromValue(v string) (*PdnConnectivityStatusAnyOf, error) {
	ev := PdnConnectivityStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PdnConnectivityStatusAnyOf: valid values are %v", v, AllowedPdnConnectivityStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PdnConnectivityStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedPdnConnectivityStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PdnConnectivityStatus_anyOf value
func (v PdnConnectivityStatusAnyOf) Ptr() *PdnConnectivityStatusAnyOf {
	return &v
}

type NullablePdnConnectivityStatusAnyOf struct {
	value *PdnConnectivityStatusAnyOf
	isSet bool
}

func (v NullablePdnConnectivityStatusAnyOf) Get() *PdnConnectivityStatusAnyOf {
	return v.value
}

func (v *NullablePdnConnectivityStatusAnyOf) Set(val *PdnConnectivityStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePdnConnectivityStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePdnConnectivityStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePdnConnectivityStatusAnyOf(val *PdnConnectivityStatusAnyOf) *NullablePdnConnectivityStatusAnyOf {
	return &NullablePdnConnectivityStatusAnyOf{value: val, isSet: true}
}

func (v NullablePdnConnectivityStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePdnConnectivityStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

