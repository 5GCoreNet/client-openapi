/*
3gpp-service-parameter

API for AF service paramter   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ServiceParameter

import (
	"encoding/json"
	"fmt"
)

// AuthorizationResultAnyOf the model 'AuthorizationResultAnyOf'
type AuthorizationResultAnyOf string

// List of AuthorizationResult_anyOf
const (
	AUTH_REVOKED AuthorizationResultAnyOf = "AUTH_REVOKED"
)

// All allowed values of AuthorizationResultAnyOf enum
var AllowedAuthorizationResultAnyOfEnumValues = []AuthorizationResultAnyOf{
	"AUTH_REVOKED",
}

func (v *AuthorizationResultAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthorizationResultAnyOf(value)
	for _, existing := range AllowedAuthorizationResultAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthorizationResultAnyOf", value)
}

// NewAuthorizationResultAnyOfFromValue returns a pointer to a valid AuthorizationResultAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthorizationResultAnyOfFromValue(v string) (*AuthorizationResultAnyOf, error) {
	ev := AuthorizationResultAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthorizationResultAnyOf: valid values are %v", v, AllowedAuthorizationResultAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthorizationResultAnyOf) IsValid() bool {
	for _, existing := range AllowedAuthorizationResultAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthorizationResult_anyOf value
func (v AuthorizationResultAnyOf) Ptr() *AuthorizationResultAnyOf {
	return &v
}

type NullableAuthorizationResultAnyOf struct {
	value *AuthorizationResultAnyOf
	isSet bool
}

func (v NullableAuthorizationResultAnyOf) Get() *AuthorizationResultAnyOf {
	return v.value
}

func (v *NullableAuthorizationResultAnyOf) Set(val *AuthorizationResultAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationResultAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationResultAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationResultAnyOf(val *AuthorizationResultAnyOf) *NullableAuthorizationResultAnyOf {
	return &NullableAuthorizationResultAnyOf{value: val, isSet: true}
}

func (v NullableAuthorizationResultAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationResultAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

