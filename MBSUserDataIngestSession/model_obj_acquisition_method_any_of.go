/*
3gpp-mbs-ud-ingest

API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSUserDataIngestSession

import (
	"encoding/json"
	"fmt"
)

// ObjAcquisitionMethodAnyOf the model 'ObjAcquisitionMethodAnyOf'
type ObjAcquisitionMethodAnyOf string

// List of ObjAcquisitionMethod_anyOf
const (
	PULL ObjAcquisitionMethodAnyOf = "PULL"
	PUSH ObjAcquisitionMethodAnyOf = "PUSH"
)

// All allowed values of ObjAcquisitionMethodAnyOf enum
var AllowedObjAcquisitionMethodAnyOfEnumValues = []ObjAcquisitionMethodAnyOf{
	"PULL",
	"PUSH",
}

func (v *ObjAcquisitionMethodAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ObjAcquisitionMethodAnyOf(value)
	for _, existing := range AllowedObjAcquisitionMethodAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ObjAcquisitionMethodAnyOf", value)
}

// NewObjAcquisitionMethodAnyOfFromValue returns a pointer to a valid ObjAcquisitionMethodAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewObjAcquisitionMethodAnyOfFromValue(v string) (*ObjAcquisitionMethodAnyOf, error) {
	ev := ObjAcquisitionMethodAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ObjAcquisitionMethodAnyOf: valid values are %v", v, AllowedObjAcquisitionMethodAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ObjAcquisitionMethodAnyOf) IsValid() bool {
	for _, existing := range AllowedObjAcquisitionMethodAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObjAcquisitionMethod_anyOf value
func (v ObjAcquisitionMethodAnyOf) Ptr() *ObjAcquisitionMethodAnyOf {
	return &v
}

type NullableObjAcquisitionMethodAnyOf struct {
	value *ObjAcquisitionMethodAnyOf
	isSet bool
}

func (v NullableObjAcquisitionMethodAnyOf) Get() *ObjAcquisitionMethodAnyOf {
	return v.value
}

func (v *NullableObjAcquisitionMethodAnyOf) Set(val *ObjAcquisitionMethodAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableObjAcquisitionMethodAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableObjAcquisitionMethodAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjAcquisitionMethodAnyOf(val *ObjAcquisitionMethodAnyOf) *NullableObjAcquisitionMethodAnyOf {
	return &NullableObjAcquisitionMethodAnyOf{value: val, isSet: true}
}

func (v NullableObjAcquisitionMethodAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjAcquisitionMethodAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

