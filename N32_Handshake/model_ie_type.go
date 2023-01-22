/*
N32 Handshake API

N32-c Handshake Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package N32_Handshake

import (
	"encoding/json"
	"fmt"
)

// IeType Enumeration of types of IEs (i.e kind of IE) to specify the protection policy
type IeType struct {
	IeTypeAnyOf *IeTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *IeType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into IeTypeAnyOf
	err = json.Unmarshal(data, &dst.IeTypeAnyOf);
	if err == nil {
		jsonIeTypeAnyOf, _ := json.Marshal(dst.IeTypeAnyOf)
		if string(jsonIeTypeAnyOf) == "{}" { // empty struct
			dst.IeTypeAnyOf = nil
		} else {
			return nil // data stored in dst.IeTypeAnyOf, return on the first match
		}
	} else {
		dst.IeTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(IeType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *IeType) MarshalJSON() ([]byte, error) {
	if src.IeTypeAnyOf != nil {
		return json.Marshal(&src.IeTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableIeType struct {
	value *IeType
	isSet bool
}

func (v NullableIeType) Get() *IeType {
	return v.value
}

func (v *NullableIeType) Set(val *IeType) {
	v.value = val
	v.isSet = true
}

func (v NullableIeType) IsSet() bool {
	return v.isSet
}

func (v *NullableIeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIeType(val *IeType) *NullableIeType {
	return &NullableIeType{value: val, isSet: true}
}

func (v NullableIeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


