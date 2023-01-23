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

// AuthMethod Contains the Authentication Method.
type AuthMethod struct {
	AuthMethodAnyOf *AuthMethodAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AuthMethod) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AuthMethodAnyOf
	err = json.Unmarshal(data, &dst.AuthMethodAnyOf);
	if err == nil {
		jsonAuthMethodAnyOf, _ := json.Marshal(dst.AuthMethodAnyOf)
		if string(jsonAuthMethodAnyOf) == "{}" { // empty struct
			dst.AuthMethodAnyOf = nil
		} else {
			return nil // data stored in dst.AuthMethodAnyOf, return on the first match
		}
	} else {
		dst.AuthMethodAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AuthMethod)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AuthMethod) MarshalJSON() ([]byte, error) {
	if src.AuthMethodAnyOf != nil {
		return json.Marshal(&src.AuthMethodAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAuthMethod struct {
	value *AuthMethod
	isSet bool
}

func (v NullableAuthMethod) Get() *AuthMethod {
	return v.value
}

func (v *NullableAuthMethod) Set(val *AuthMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthMethod(val *AuthMethod) *NullableAuthMethod {
	return &NullableAuthMethod{value: val, isSet: true}
}

func (v NullableAuthMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


