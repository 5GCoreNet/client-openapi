/*
Nhss_imsUECM

Nhss UE Context Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nhss_imsUECM

import (
	"encoding/json"
	"fmt"
)

// ImsRegistrationTypeAnyOf the model 'ImsRegistrationTypeAnyOf'
type ImsRegistrationTypeAnyOf string

// List of ImsRegistrationType_anyOf
const (
	INITIAL_REGISTRATION ImsRegistrationTypeAnyOf = "INITIAL_REGISTRATION"
	RE_REGISTRATION ImsRegistrationTypeAnyOf = "RE_REGISTRATION"
	TIMEOUT_DEREGISTRATION ImsRegistrationTypeAnyOf = "TIMEOUT_DEREGISTRATION"
	USER_DEREGISTRATION ImsRegistrationTypeAnyOf = "USER_DEREGISTRATION"
	ADMINISTRATIVE_DEREGISTRATION ImsRegistrationTypeAnyOf = "ADMINISTRATIVE_DEREGISTRATION"
	AUTHENTICATION_FAILURE ImsRegistrationTypeAnyOf = "AUTHENTICATION_FAILURE"
	AUTHENTICATION_TIMEOUT ImsRegistrationTypeAnyOf = "AUTHENTICATION_TIMEOUT"
	UNREGISTERED_USER ImsRegistrationTypeAnyOf = "UNREGISTERED_USER"
)

// All allowed values of ImsRegistrationTypeAnyOf enum
var AllowedImsRegistrationTypeAnyOfEnumValues = []ImsRegistrationTypeAnyOf{
	"INITIAL_REGISTRATION",
	"RE_REGISTRATION",
	"TIMEOUT_DEREGISTRATION",
	"USER_DEREGISTRATION",
	"ADMINISTRATIVE_DEREGISTRATION",
	"AUTHENTICATION_FAILURE",
	"AUTHENTICATION_TIMEOUT",
	"UNREGISTERED_USER",
}

func (v *ImsRegistrationTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImsRegistrationTypeAnyOf(value)
	for _, existing := range AllowedImsRegistrationTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImsRegistrationTypeAnyOf", value)
}

// NewImsRegistrationTypeAnyOfFromValue returns a pointer to a valid ImsRegistrationTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImsRegistrationTypeAnyOfFromValue(v string) (*ImsRegistrationTypeAnyOf, error) {
	ev := ImsRegistrationTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImsRegistrationTypeAnyOf: valid values are %v", v, AllowedImsRegistrationTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImsRegistrationTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedImsRegistrationTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImsRegistrationType_anyOf value
func (v ImsRegistrationTypeAnyOf) Ptr() *ImsRegistrationTypeAnyOf {
	return &v
}

type NullableImsRegistrationTypeAnyOf struct {
	value *ImsRegistrationTypeAnyOf
	isSet bool
}

func (v NullableImsRegistrationTypeAnyOf) Get() *ImsRegistrationTypeAnyOf {
	return v.value
}

func (v *NullableImsRegistrationTypeAnyOf) Set(val *ImsRegistrationTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableImsRegistrationTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableImsRegistrationTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImsRegistrationTypeAnyOf(val *ImsRegistrationTypeAnyOf) *NullableImsRegistrationTypeAnyOf {
	return &NullableImsRegistrationTypeAnyOf{value: val, isSet: true}
}

func (v NullableImsRegistrationTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImsRegistrationTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
