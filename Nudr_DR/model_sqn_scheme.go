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

// SqnScheme Scheme for generation of Sequence Numbers.
type SqnScheme struct {
	SqnSchemeAnyOf *SqnSchemeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SqnScheme) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SqnSchemeAnyOf
	err = json.Unmarshal(data, &dst.SqnSchemeAnyOf);
	if err == nil {
		jsonSqnSchemeAnyOf, _ := json.Marshal(dst.SqnSchemeAnyOf)
		if string(jsonSqnSchemeAnyOf) == "{}" { // empty struct
			dst.SqnSchemeAnyOf = nil
		} else {
			return nil // data stored in dst.SqnSchemeAnyOf, return on the first match
		}
	} else {
		dst.SqnSchemeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SqnScheme)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SqnScheme) MarshalJSON() ([]byte, error) {
	if src.SqnSchemeAnyOf != nil {
		return json.Marshal(&src.SqnSchemeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSqnScheme struct {
	value *SqnScheme
	isSet bool
}

func (v NullableSqnScheme) Get() *SqnScheme {
	return v.value
}

func (v *NullableSqnScheme) Set(val *SqnScheme) {
	v.value = val
	v.isSet = true
}

func (v NullableSqnScheme) IsSet() bool {
	return v.isSet
}

func (v *NullableSqnScheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSqnScheme(val *SqnScheme) *NullableSqnScheme {
	return &NullableSqnScheme{value: val, isSet: true}
}

func (v NullableSqnScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSqnScheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


