/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
	"fmt"
)

// UserConsentAnyOf the model 'UserConsentAnyOf'
type UserConsentAnyOf string

// List of UserConsent_anyOf
const (
	NOT_GIVEN UserConsentAnyOf = "CONSENT_NOT_GIVEN"
	GIVEN UserConsentAnyOf = "CONSENT_GIVEN"
)

// All allowed values of UserConsentAnyOf enum
var AllowedUserConsentAnyOfEnumValues = []UserConsentAnyOf{
	"CONSENT_NOT_GIVEN",
	"CONSENT_GIVEN",
}

func (v *UserConsentAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserConsentAnyOf(value)
	for _, existing := range AllowedUserConsentAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserConsentAnyOf", value)
}

// NewUserConsentAnyOfFromValue returns a pointer to a valid UserConsentAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserConsentAnyOfFromValue(v string) (*UserConsentAnyOf, error) {
	ev := UserConsentAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserConsentAnyOf: valid values are %v", v, AllowedUserConsentAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserConsentAnyOf) IsValid() bool {
	for _, existing := range AllowedUserConsentAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserConsent_anyOf value
func (v UserConsentAnyOf) Ptr() *UserConsentAnyOf {
	return &v
}

type NullableUserConsentAnyOf struct {
	value *UserConsentAnyOf
	isSet bool
}

func (v NullableUserConsentAnyOf) Get() *UserConsentAnyOf {
	return v.value
}

func (v *NullableUserConsentAnyOf) Set(val *UserConsentAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserConsentAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserConsentAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserConsentAnyOf(val *UserConsentAnyOf) *NullableUserConsentAnyOf {
	return &NullableUserConsentAnyOf{value: val, isSet: true}
}

func (v NullableUserConsentAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserConsentAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

