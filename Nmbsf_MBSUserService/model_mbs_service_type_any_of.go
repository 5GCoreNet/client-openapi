/*
nmbsf-mbs-us

API for MBS User Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbsf_MBSUserService

import (
	"encoding/json"
	"fmt"
)

// MbsServiceTypeAnyOf the model 'MbsServiceTypeAnyOf'
type MbsServiceTypeAnyOf string

// List of MbsServiceType_anyOf
const (
	MULTICAST MbsServiceTypeAnyOf = "MULTICAST"
	BROADCAST MbsServiceTypeAnyOf = "BROADCAST"
)

// All allowed values of MbsServiceTypeAnyOf enum
var AllowedMbsServiceTypeAnyOfEnumValues = []MbsServiceTypeAnyOf{
	"MULTICAST",
	"BROADCAST",
}

func (v *MbsServiceTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MbsServiceTypeAnyOf(value)
	for _, existing := range AllowedMbsServiceTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MbsServiceTypeAnyOf", value)
}

// NewMbsServiceTypeAnyOfFromValue returns a pointer to a valid MbsServiceTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMbsServiceTypeAnyOfFromValue(v string) (*MbsServiceTypeAnyOf, error) {
	ev := MbsServiceTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MbsServiceTypeAnyOf: valid values are %v", v, AllowedMbsServiceTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MbsServiceTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedMbsServiceTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MbsServiceType_anyOf value
func (v MbsServiceTypeAnyOf) Ptr() *MbsServiceTypeAnyOf {
	return &v
}

type NullableMbsServiceTypeAnyOf struct {
	value *MbsServiceTypeAnyOf
	isSet bool
}

func (v NullableMbsServiceTypeAnyOf) Get() *MbsServiceTypeAnyOf {
	return v.value
}

func (v *NullableMbsServiceTypeAnyOf) Set(val *MbsServiceTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsServiceTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsServiceTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsServiceTypeAnyOf(val *MbsServiceTypeAnyOf) *NullableMbsServiceTypeAnyOf {
	return &NullableMbsServiceTypeAnyOf{value: val, isSet: true}
}

func (v NullableMbsServiceTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsServiceTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

