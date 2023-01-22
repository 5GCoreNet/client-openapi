/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnrf_NFManagement

import (
	"encoding/json"
	"fmt"
)

// ExternalClientType Indicates types of External Clients.
type ExternalClientType struct {
	ExternalClientTypeAnyOf *ExternalClientTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ExternalClientType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ExternalClientTypeAnyOf
	err = json.Unmarshal(data, &dst.ExternalClientTypeAnyOf);
	if err == nil {
		jsonExternalClientTypeAnyOf, _ := json.Marshal(dst.ExternalClientTypeAnyOf)
		if string(jsonExternalClientTypeAnyOf) == "{}" { // empty struct
			dst.ExternalClientTypeAnyOf = nil
		} else {
			return nil // data stored in dst.ExternalClientTypeAnyOf, return on the first match
		}
	} else {
		dst.ExternalClientTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ExternalClientType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ExternalClientType) MarshalJSON() ([]byte, error) {
	if src.ExternalClientTypeAnyOf != nil {
		return json.Marshal(&src.ExternalClientTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableExternalClientType struct {
	value *ExternalClientType
	isSet bool
}

func (v NullableExternalClientType) Get() *ExternalClientType {
	return v.value
}

func (v *NullableExternalClientType) Set(val *ExternalClientType) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalClientType) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalClientType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalClientType(val *ExternalClientType) *NullableExternalClientType {
	return &NullableExternalClientType{value: val, isSet: true}
}

func (v NullableExternalClientType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalClientType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


