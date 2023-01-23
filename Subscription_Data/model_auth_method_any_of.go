/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// AuthMethodAnyOf the model 'AuthMethodAnyOf'
type AuthMethodAnyOf string

// List of AuthMethod_anyOf
const (
	_5_G_AKA AuthMethodAnyOf = "5G_AKA"
	EAP_AKA_PRIME AuthMethodAnyOf = "EAP_AKA_PRIME"
	EAP_TLS AuthMethodAnyOf = "EAP_TLS"
	EAP_TTLS AuthMethodAnyOf = "EAP_TTLS"
	NONE AuthMethodAnyOf = "NONE"
)

// All allowed values of AuthMethodAnyOf enum
var AllowedAuthMethodAnyOfEnumValues = []AuthMethodAnyOf{
	"5G_AKA",
	"EAP_AKA_PRIME",
	"EAP_TLS",
	"EAP_TTLS",
	"NONE",
}

func (v *AuthMethodAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthMethodAnyOf(value)
	for _, existing := range AllowedAuthMethodAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthMethodAnyOf", value)
}

// NewAuthMethodAnyOfFromValue returns a pointer to a valid AuthMethodAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthMethodAnyOfFromValue(v string) (*AuthMethodAnyOf, error) {
	ev := AuthMethodAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthMethodAnyOf: valid values are %v", v, AllowedAuthMethodAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthMethodAnyOf) IsValid() bool {
	for _, existing := range AllowedAuthMethodAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthMethod_anyOf value
func (v AuthMethodAnyOf) Ptr() *AuthMethodAnyOf {
	return &v
}

type NullableAuthMethodAnyOf struct {
	value *AuthMethodAnyOf
	isSet bool
}

func (v NullableAuthMethodAnyOf) Get() *AuthMethodAnyOf {
	return v.value
}

func (v *NullableAuthMethodAnyOf) Set(val *AuthMethodAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthMethodAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthMethodAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthMethodAnyOf(val *AuthMethodAnyOf) *NullableAuthMethodAnyOf {
	return &NullableAuthMethodAnyOf{value: val, isSet: true}
}

func (v NullableAuthMethodAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthMethodAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

