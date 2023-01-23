/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// Sign Sign of the DIF value.
type Sign string

// List of Sign
const (
	POSITIVE Sign = "POSITIVE"
	NEGATIVE Sign = "NEGATIVE"
)

// All allowed values of Sign enum
var AllowedSignEnumValues = []Sign{
	"POSITIVE",
	"NEGATIVE",
}

func (v *Sign) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Sign(value)
	for _, existing := range AllowedSignEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Sign", value)
}

// NewSignFromValue returns a pointer to a valid Sign
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSignFromValue(v string) (*Sign, error) {
	ev := Sign(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Sign: valid values are %v", v, AllowedSignEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Sign) IsValid() bool {
	for _, existing := range AllowedSignEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Sign value
func (v Sign) Ptr() *Sign {
	return &v
}

type NullableSign struct {
	value *Sign
	isSet bool
}

func (v NullableSign) Get() *Sign {
	return v.value
}

func (v *NullableSign) Set(val *Sign) {
	v.value = val
	v.isSet = true
}

func (v NullableSign) IsSet() bool {
	return v.isSet
}

func (v *NullableSign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSign(val *Sign) *NullableSign {
	return &NullableSign{value: val, isSet: true}
}

func (v NullableSign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

