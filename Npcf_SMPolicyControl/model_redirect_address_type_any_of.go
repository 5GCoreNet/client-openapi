/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// RedirectAddressTypeAnyOf the model 'RedirectAddressTypeAnyOf'
type RedirectAddressTypeAnyOf string

// List of RedirectAddressType_anyOf
const (
	IPV4_ADDR RedirectAddressTypeAnyOf = "IPV4_ADDR"
	IPV6_ADDR RedirectAddressTypeAnyOf = "IPV6_ADDR"
	URL RedirectAddressTypeAnyOf = "URL"
	SIP_URI RedirectAddressTypeAnyOf = "SIP_URI"
)

// All allowed values of RedirectAddressTypeAnyOf enum
var AllowedRedirectAddressTypeAnyOfEnumValues = []RedirectAddressTypeAnyOf{
	"IPV4_ADDR",
	"IPV6_ADDR",
	"URL",
	"SIP_URI",
}

func (v *RedirectAddressTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RedirectAddressTypeAnyOf(value)
	for _, existing := range AllowedRedirectAddressTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RedirectAddressTypeAnyOf", value)
}

// NewRedirectAddressTypeAnyOfFromValue returns a pointer to a valid RedirectAddressTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRedirectAddressTypeAnyOfFromValue(v string) (*RedirectAddressTypeAnyOf, error) {
	ev := RedirectAddressTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RedirectAddressTypeAnyOf: valid values are %v", v, AllowedRedirectAddressTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RedirectAddressTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedRedirectAddressTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RedirectAddressType_anyOf value
func (v RedirectAddressTypeAnyOf) Ptr() *RedirectAddressTypeAnyOf {
	return &v
}

type NullableRedirectAddressTypeAnyOf struct {
	value *RedirectAddressTypeAnyOf
	isSet bool
}

func (v NullableRedirectAddressTypeAnyOf) Get() *RedirectAddressTypeAnyOf {
	return v.value
}

func (v *NullableRedirectAddressTypeAnyOf) Set(val *RedirectAddressTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectAddressTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectAddressTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectAddressTypeAnyOf(val *RedirectAddressTypeAnyOf) *NullableRedirectAddressTypeAnyOf {
	return &NullableRedirectAddressTypeAnyOf{value: val, isSet: true}
}

func (v NullableRedirectAddressTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectAddressTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

