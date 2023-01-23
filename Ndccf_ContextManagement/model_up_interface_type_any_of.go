/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// UPInterfaceTypeAnyOf the model 'UPInterfaceTypeAnyOf'
type UPInterfaceTypeAnyOf string

// List of UPInterfaceType_anyOf
const (
	N3 UPInterfaceTypeAnyOf = "N3"
	N6 UPInterfaceTypeAnyOf = "N6"
	N9 UPInterfaceTypeAnyOf = "N9"
	DATA_FORWARDING UPInterfaceTypeAnyOf = "DATA_FORWARDING"
	N3_MB UPInterfaceTypeAnyOf = "N3MB"
	N6_MB UPInterfaceTypeAnyOf = "N6MB"
	N19_MB UPInterfaceTypeAnyOf = "N19MB"
	NMB9 UPInterfaceTypeAnyOf = "NMB9"
)

// All allowed values of UPInterfaceTypeAnyOf enum
var AllowedUPInterfaceTypeAnyOfEnumValues = []UPInterfaceTypeAnyOf{
	"N3",
	"N6",
	"N9",
	"DATA_FORWARDING",
	"N3MB",
	"N6MB",
	"N19MB",
	"NMB9",
}

func (v *UPInterfaceTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UPInterfaceTypeAnyOf(value)
	for _, existing := range AllowedUPInterfaceTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UPInterfaceTypeAnyOf", value)
}

// NewUPInterfaceTypeAnyOfFromValue returns a pointer to a valid UPInterfaceTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUPInterfaceTypeAnyOfFromValue(v string) (*UPInterfaceTypeAnyOf, error) {
	ev := UPInterfaceTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UPInterfaceTypeAnyOf: valid values are %v", v, AllowedUPInterfaceTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UPInterfaceTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedUPInterfaceTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UPInterfaceType_anyOf value
func (v UPInterfaceTypeAnyOf) Ptr() *UPInterfaceTypeAnyOf {
	return &v
}

type NullableUPInterfaceTypeAnyOf struct {
	value *UPInterfaceTypeAnyOf
	isSet bool
}

func (v NullableUPInterfaceTypeAnyOf) Get() *UPInterfaceTypeAnyOf {
	return v.value
}

func (v *NullableUPInterfaceTypeAnyOf) Set(val *UPInterfaceTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUPInterfaceTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUPInterfaceTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUPInterfaceTypeAnyOf(val *UPInterfaceTypeAnyOf) *NullableUPInterfaceTypeAnyOf {
	return &NullableUPInterfaceTypeAnyOf{value: val, isSet: true}
}

func (v NullableUPInterfaceTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUPInterfaceTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

