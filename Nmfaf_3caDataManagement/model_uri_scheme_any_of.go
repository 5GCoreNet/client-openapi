/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmfaf_3caDataManagement

import (
	"encoding/json"
	"fmt"
)

// UriSchemeAnyOf the model 'UriSchemeAnyOf'
type UriSchemeAnyOf string

// List of UriScheme_anyOf
const (
	HTTP UriSchemeAnyOf = "http"
	HTTPS UriSchemeAnyOf = "https"
)

// All allowed values of UriSchemeAnyOf enum
var AllowedUriSchemeAnyOfEnumValues = []UriSchemeAnyOf{
	"http",
	"https",
}

func (v *UriSchemeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UriSchemeAnyOf(value)
	for _, existing := range AllowedUriSchemeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UriSchemeAnyOf", value)
}

// NewUriSchemeAnyOfFromValue returns a pointer to a valid UriSchemeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUriSchemeAnyOfFromValue(v string) (*UriSchemeAnyOf, error) {
	ev := UriSchemeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UriSchemeAnyOf: valid values are %v", v, AllowedUriSchemeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UriSchemeAnyOf) IsValid() bool {
	for _, existing := range AllowedUriSchemeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UriScheme_anyOf value
func (v UriSchemeAnyOf) Ptr() *UriSchemeAnyOf {
	return &v
}

type NullableUriSchemeAnyOf struct {
	value *UriSchemeAnyOf
	isSet bool
}

func (v NullableUriSchemeAnyOf) Get() *UriSchemeAnyOf {
	return v.value
}

func (v *NullableUriSchemeAnyOf) Set(val *UriSchemeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUriSchemeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUriSchemeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUriSchemeAnyOf(val *UriSchemeAnyOf) *NullableUriSchemeAnyOf {
	return &NullableUriSchemeAnyOf{value: val, isSet: true}
}

func (v NullableUriSchemeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUriSchemeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

