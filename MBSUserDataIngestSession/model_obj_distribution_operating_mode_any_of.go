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

// ObjDistributionOperatingModeAnyOf the model 'ObjDistributionOperatingModeAnyOf'
type ObjDistributionOperatingModeAnyOf string

// List of ObjDistributionOperatingMode_anyOf
const (
	SINGLE ObjDistributionOperatingModeAnyOf = "SINGLE"
	COLLECTION ObjDistributionOperatingModeAnyOf = "COLLECTION"
	CAROUSEL ObjDistributionOperatingModeAnyOf = "CAROUSEL"
	STREAMING ObjDistributionOperatingModeAnyOf = "STREAMING"
)

// All allowed values of ObjDistributionOperatingModeAnyOf enum
var AllowedObjDistributionOperatingModeAnyOfEnumValues = []ObjDistributionOperatingModeAnyOf{
	"SINGLE",
	"COLLECTION",
	"CAROUSEL",
	"STREAMING",
}

func (v *ObjDistributionOperatingModeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ObjDistributionOperatingModeAnyOf(value)
	for _, existing := range AllowedObjDistributionOperatingModeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ObjDistributionOperatingModeAnyOf", value)
}

// NewObjDistributionOperatingModeAnyOfFromValue returns a pointer to a valid ObjDistributionOperatingModeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewObjDistributionOperatingModeAnyOfFromValue(v string) (*ObjDistributionOperatingModeAnyOf, error) {
	ev := ObjDistributionOperatingModeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ObjDistributionOperatingModeAnyOf: valid values are %v", v, AllowedObjDistributionOperatingModeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ObjDistributionOperatingModeAnyOf) IsValid() bool {
	for _, existing := range AllowedObjDistributionOperatingModeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObjDistributionOperatingMode_anyOf value
func (v ObjDistributionOperatingModeAnyOf) Ptr() *ObjDistributionOperatingModeAnyOf {
	return &v
}

type NullableObjDistributionOperatingModeAnyOf struct {
	value *ObjDistributionOperatingModeAnyOf
	isSet bool
}

func (v NullableObjDistributionOperatingModeAnyOf) Get() *ObjDistributionOperatingModeAnyOf {
	return v.value
}

func (v *NullableObjDistributionOperatingModeAnyOf) Set(val *ObjDistributionOperatingModeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableObjDistributionOperatingModeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableObjDistributionOperatingModeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjDistributionOperatingModeAnyOf(val *ObjDistributionOperatingModeAnyOf) *NullableObjDistributionOperatingModeAnyOf {
	return &NullableObjDistributionOperatingModeAnyOf{value: val, isSet: true}
}

func (v NullableObjDistributionOperatingModeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjDistributionOperatingModeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

