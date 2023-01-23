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

// LocationPrivacyIndAnyOf the model 'LocationPrivacyIndAnyOf'
type LocationPrivacyIndAnyOf string

// List of LocationPrivacyInd_anyOf
const (
	DISALLOWED LocationPrivacyIndAnyOf = "LOCATION_DISALLOWED"
	ALLOWED LocationPrivacyIndAnyOf = "LOCATION_ALLOWED"
)

// All allowed values of LocationPrivacyIndAnyOf enum
var AllowedLocationPrivacyIndAnyOfEnumValues = []LocationPrivacyIndAnyOf{
	"LOCATION_DISALLOWED",
	"LOCATION_ALLOWED",
}

func (v *LocationPrivacyIndAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocationPrivacyIndAnyOf(value)
	for _, existing := range AllowedLocationPrivacyIndAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LocationPrivacyIndAnyOf", value)
}

// NewLocationPrivacyIndAnyOfFromValue returns a pointer to a valid LocationPrivacyIndAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocationPrivacyIndAnyOfFromValue(v string) (*LocationPrivacyIndAnyOf, error) {
	ev := LocationPrivacyIndAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocationPrivacyIndAnyOf: valid values are %v", v, AllowedLocationPrivacyIndAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocationPrivacyIndAnyOf) IsValid() bool {
	for _, existing := range AllowedLocationPrivacyIndAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LocationPrivacyInd_anyOf value
func (v LocationPrivacyIndAnyOf) Ptr() *LocationPrivacyIndAnyOf {
	return &v
}

type NullableLocationPrivacyIndAnyOf struct {
	value *LocationPrivacyIndAnyOf
	isSet bool
}

func (v NullableLocationPrivacyIndAnyOf) Get() *LocationPrivacyIndAnyOf {
	return v.value
}

func (v *NullableLocationPrivacyIndAnyOf) Set(val *LocationPrivacyIndAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationPrivacyIndAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationPrivacyIndAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationPrivacyIndAnyOf(val *LocationPrivacyIndAnyOf) *NullableLocationPrivacyIndAnyOf {
	return &NullableLocationPrivacyIndAnyOf{value: val, isSet: true}
}

func (v NullableLocationPrivacyIndAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationPrivacyIndAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

