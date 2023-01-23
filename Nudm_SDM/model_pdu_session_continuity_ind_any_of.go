/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
	"fmt"
)

// PduSessionContinuityIndAnyOf the model 'PduSessionContinuityIndAnyOf'
type PduSessionContinuityIndAnyOf string

// List of PduSessionContinuityInd_anyOf
const (
	MAINTAIN_PDUSESSION PduSessionContinuityIndAnyOf = "MAINTAIN_PDUSESSION"
	RECONNECT_PDUSESSION PduSessionContinuityIndAnyOf = "RECONNECT_PDUSESSION"
	RELEASE_PDUSESSION PduSessionContinuityIndAnyOf = "RELEASE_PDUSESSION"
)

// All allowed values of PduSessionContinuityIndAnyOf enum
var AllowedPduSessionContinuityIndAnyOfEnumValues = []PduSessionContinuityIndAnyOf{
	"MAINTAIN_PDUSESSION",
	"RECONNECT_PDUSESSION",
	"RELEASE_PDUSESSION",
}

func (v *PduSessionContinuityIndAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PduSessionContinuityIndAnyOf(value)
	for _, existing := range AllowedPduSessionContinuityIndAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PduSessionContinuityIndAnyOf", value)
}

// NewPduSessionContinuityIndAnyOfFromValue returns a pointer to a valid PduSessionContinuityIndAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPduSessionContinuityIndAnyOfFromValue(v string) (*PduSessionContinuityIndAnyOf, error) {
	ev := PduSessionContinuityIndAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PduSessionContinuityIndAnyOf: valid values are %v", v, AllowedPduSessionContinuityIndAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PduSessionContinuityIndAnyOf) IsValid() bool {
	for _, existing := range AllowedPduSessionContinuityIndAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PduSessionContinuityInd_anyOf value
func (v PduSessionContinuityIndAnyOf) Ptr() *PduSessionContinuityIndAnyOf {
	return &v
}

type NullablePduSessionContinuityIndAnyOf struct {
	value *PduSessionContinuityIndAnyOf
	isSet bool
}

func (v NullablePduSessionContinuityIndAnyOf) Get() *PduSessionContinuityIndAnyOf {
	return v.value
}

func (v *NullablePduSessionContinuityIndAnyOf) Set(val *PduSessionContinuityIndAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionContinuityIndAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionContinuityIndAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionContinuityIndAnyOf(val *PduSessionContinuityIndAnyOf) *NullablePduSessionContinuityIndAnyOf {
	return &NullablePduSessionContinuityIndAnyOf{value: val, isSet: true}
}

func (v NullablePduSessionContinuityIndAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionContinuityIndAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

