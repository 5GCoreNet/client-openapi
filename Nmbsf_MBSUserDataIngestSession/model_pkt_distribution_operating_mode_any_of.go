/*
nmbsf-mbs-ud-ingest

API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsf_MBSUserDataIngestSession

import (
	"encoding/json"
	"fmt"
)

// PktDistributionOperatingModeAnyOf the model 'PktDistributionOperatingModeAnyOf'
type PktDistributionOperatingModeAnyOf string

// List of PktDistributionOperatingMode_anyOf
const (
	PROXY PktDistributionOperatingModeAnyOf = "PACKET_PROXY"
	FORWARD_ONLY PktDistributionOperatingModeAnyOf = "PACKET_FORWARD_ONLY"
)

// All allowed values of PktDistributionOperatingModeAnyOf enum
var AllowedPktDistributionOperatingModeAnyOfEnumValues = []PktDistributionOperatingModeAnyOf{
	"PACKET_PROXY",
	"PACKET_FORWARD_ONLY",
}

func (v *PktDistributionOperatingModeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PktDistributionOperatingModeAnyOf(value)
	for _, existing := range AllowedPktDistributionOperatingModeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PktDistributionOperatingModeAnyOf", value)
}

// NewPktDistributionOperatingModeAnyOfFromValue returns a pointer to a valid PktDistributionOperatingModeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPktDistributionOperatingModeAnyOfFromValue(v string) (*PktDistributionOperatingModeAnyOf, error) {
	ev := PktDistributionOperatingModeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PktDistributionOperatingModeAnyOf: valid values are %v", v, AllowedPktDistributionOperatingModeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PktDistributionOperatingModeAnyOf) IsValid() bool {
	for _, existing := range AllowedPktDistributionOperatingModeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PktDistributionOperatingMode_anyOf value
func (v PktDistributionOperatingModeAnyOf) Ptr() *PktDistributionOperatingModeAnyOf {
	return &v
}

type NullablePktDistributionOperatingModeAnyOf struct {
	value *PktDistributionOperatingModeAnyOf
	isSet bool
}

func (v NullablePktDistributionOperatingModeAnyOf) Get() *PktDistributionOperatingModeAnyOf {
	return v.value
}

func (v *NullablePktDistributionOperatingModeAnyOf) Set(val *PktDistributionOperatingModeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePktDistributionOperatingModeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePktDistributionOperatingModeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePktDistributionOperatingModeAnyOf(val *PktDistributionOperatingModeAnyOf) *NullablePktDistributionOperatingModeAnyOf {
	return &NullablePktDistributionOperatingModeAnyOf{value: val, isSet: true}
}

func (v NullablePktDistributionOperatingModeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePktDistributionOperatingModeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

