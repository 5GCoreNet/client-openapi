/*
AUSF API

AUSF UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nausf_UEAuthentication

import (
	"encoding/json"
	"fmt"
)

// AuthType Indicates the authentication method used.
type AuthType struct {
	AuthTypeAnyOf *AuthTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AuthType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AuthTypeAnyOf
	err = json.Unmarshal(data, &dst.AuthTypeAnyOf);
	if err == nil {
		jsonAuthTypeAnyOf, _ := json.Marshal(dst.AuthTypeAnyOf)
		if string(jsonAuthTypeAnyOf) == "{}" { // empty struct
			dst.AuthTypeAnyOf = nil
		} else {
			return nil // data stored in dst.AuthTypeAnyOf, return on the first match
		}
	} else {
		dst.AuthTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AuthType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AuthType) MarshalJSON() ([]byte, error) {
	if src.AuthTypeAnyOf != nil {
		return json.Marshal(&src.AuthTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAuthType struct {
	value *AuthType
	isSet bool
}

func (v NullableAuthType) Get() *AuthType {
	return v.value
}

func (v *NullableAuthType) Set(val *AuthType) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthType) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthType(val *AuthType) *NullableAuthType {
	return &NullableAuthType{value: val, isSet: true}
}

func (v NullableAuthType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


