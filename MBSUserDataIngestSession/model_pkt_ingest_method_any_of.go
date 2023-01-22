/*
3gpp-mbs-ud-ingest

API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MBSUserDataIngestSession

import (
	"encoding/json"
	"fmt"
)

// PktIngestMethodAnyOf the model 'PktIngestMethodAnyOf'
type PktIngestMethodAnyOf string

// List of PktIngestMethod_anyOf
const (
	MULTICAST PktIngestMethodAnyOf = "MULTICAST"
	UNICAST PktIngestMethodAnyOf = "UNICAST"
)

// All allowed values of PktIngestMethodAnyOf enum
var AllowedPktIngestMethodAnyOfEnumValues = []PktIngestMethodAnyOf{
	"MULTICAST",
	"UNICAST",
}

func (v *PktIngestMethodAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PktIngestMethodAnyOf(value)
	for _, existing := range AllowedPktIngestMethodAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PktIngestMethodAnyOf", value)
}

// NewPktIngestMethodAnyOfFromValue returns a pointer to a valid PktIngestMethodAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPktIngestMethodAnyOfFromValue(v string) (*PktIngestMethodAnyOf, error) {
	ev := PktIngestMethodAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PktIngestMethodAnyOf: valid values are %v", v, AllowedPktIngestMethodAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PktIngestMethodAnyOf) IsValid() bool {
	for _, existing := range AllowedPktIngestMethodAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PktIngestMethod_anyOf value
func (v PktIngestMethodAnyOf) Ptr() *PktIngestMethodAnyOf {
	return &v
}

type NullablePktIngestMethodAnyOf struct {
	value *PktIngestMethodAnyOf
	isSet bool
}

func (v NullablePktIngestMethodAnyOf) Get() *PktIngestMethodAnyOf {
	return v.value
}

func (v *NullablePktIngestMethodAnyOf) Set(val *PktIngestMethodAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePktIngestMethodAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePktIngestMethodAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePktIngestMethodAnyOf(val *PktIngestMethodAnyOf) *NullablePktIngestMethodAnyOf {
	return &NullablePktIngestMethodAnyOf{value: val, isSet: true}
}

func (v NullablePktIngestMethodAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePktIngestMethodAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

