/*
NRF OAuth2

NRF OAuth2 Authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_AccessToken

import (
	"encoding/json"
	"fmt"
)

// NFType NF types known to NRF
type NFType struct {
	NFTypeAnyOf *NFTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NFType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NFTypeAnyOf
	err = json.Unmarshal(data, &dst.NFTypeAnyOf);
	if err == nil {
		jsonNFTypeAnyOf, _ := json.Marshal(dst.NFTypeAnyOf)
		if string(jsonNFTypeAnyOf) == "{}" { // empty struct
			dst.NFTypeAnyOf = nil
		} else {
			return nil // data stored in dst.NFTypeAnyOf, return on the first match
		}
	} else {
		dst.NFTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NFType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NFType) MarshalJSON() ([]byte, error) {
	if src.NFTypeAnyOf != nil {
		return json.Marshal(&src.NFTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNFType struct {
	value *NFType
	isSet bool
}

func (v NullableNFType) Get() *NFType {
	return v.value
}

func (v *NullableNFType) Set(val *NFType) {
	v.value = val
	v.isSet = true
}

func (v NullableNFType) IsSet() bool {
	return v.isSet
}

func (v *NullableNFType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFType(val *NFType) *NullableNFType {
	return &NullableNFType{value: val, isSet: true}
}

func (v NullableNFType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


