/*
Nhss_gbaSDM

Nhss Subscriber Data Management Service for GBA.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_gbaSDM

import (
	"encoding/json"
	"fmt"
)

// UiccTypeAnyOf the model 'UiccTypeAnyOf'
type UiccTypeAnyOf string

// List of UiccType_anyOf
const (
	GBA UiccTypeAnyOf = "GBA"
	GBA_U UiccTypeAnyOf = "GBA_U"
)

// All allowed values of UiccTypeAnyOf enum
var AllowedUiccTypeAnyOfEnumValues = []UiccTypeAnyOf{
	"GBA",
	"GBA_U",
}

func (v *UiccTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UiccTypeAnyOf(value)
	for _, existing := range AllowedUiccTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UiccTypeAnyOf", value)
}

// NewUiccTypeAnyOfFromValue returns a pointer to a valid UiccTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUiccTypeAnyOfFromValue(v string) (*UiccTypeAnyOf, error) {
	ev := UiccTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UiccTypeAnyOf: valid values are %v", v, AllowedUiccTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UiccTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedUiccTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UiccType_anyOf value
func (v UiccTypeAnyOf) Ptr() *UiccTypeAnyOf {
	return &v
}

type NullableUiccTypeAnyOf struct {
	value *UiccTypeAnyOf
	isSet bool
}

func (v NullableUiccTypeAnyOf) Get() *UiccTypeAnyOf {
	return v.value
}

func (v *NullableUiccTypeAnyOf) Set(val *UiccTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUiccTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUiccTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiccTypeAnyOf(val *UiccTypeAnyOf) *NullableUiccTypeAnyOf {
	return &NullableUiccTypeAnyOf{value: val, isSet: true}
}

func (v NullableUiccTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiccTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

