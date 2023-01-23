/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFDiscovery

import (
	"encoding/json"
	"fmt"
)

// LocalityType Type of locality description. An operator may define custom locality type values other  than those listed in this enumeration.  
type LocalityType struct {
	LocalityTypeAnyOf *LocalityTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LocalityType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LocalityTypeAnyOf
	err = json.Unmarshal(data, &dst.LocalityTypeAnyOf);
	if err == nil {
		jsonLocalityTypeAnyOf, _ := json.Marshal(dst.LocalityTypeAnyOf)
		if string(jsonLocalityTypeAnyOf) == "{}" { // empty struct
			dst.LocalityTypeAnyOf = nil
		} else {
			return nil // data stored in dst.LocalityTypeAnyOf, return on the first match
		}
	} else {
		dst.LocalityTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(LocalityType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LocalityType) MarshalJSON() ([]byte, error) {
	if src.LocalityTypeAnyOf != nil {
		return json.Marshal(&src.LocalityTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLocalityType struct {
	value *LocalityType
	isSet bool
}

func (v NullableLocalityType) Get() *LocalityType {
	return v.value
}

func (v *NullableLocalityType) Set(val *LocalityType) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalityType) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalityType(val *LocalityType) *NullableLocalityType {
	return &NullableLocalityType{value: val, isSet: true}
}

func (v NullableLocalityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


