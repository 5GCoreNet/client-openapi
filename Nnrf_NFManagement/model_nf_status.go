/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
	"fmt"
)

// NFStatus Status of a given NF Instance stored in NRF
type NFStatus struct {
	NFStatusAnyOf *NFStatusAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NFStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NFStatusAnyOf
	err = json.Unmarshal(data, &dst.NFStatusAnyOf);
	if err == nil {
		jsonNFStatusAnyOf, _ := json.Marshal(dst.NFStatusAnyOf)
		if string(jsonNFStatusAnyOf) == "{}" { // empty struct
			dst.NFStatusAnyOf = nil
		} else {
			return nil // data stored in dst.NFStatusAnyOf, return on the first match
		}
	} else {
		dst.NFStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NFStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NFStatus) MarshalJSON() ([]byte, error) {
	if src.NFStatusAnyOf != nil {
		return json.Marshal(&src.NFStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNFStatus struct {
	value *NFStatus
	isSet bool
}

func (v NullableNFStatus) Get() *NFStatus {
	return v.value
}

func (v *NullableNFStatus) Set(val *NFStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNFStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNFStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFStatus(val *NFStatus) *NullableNFStatus {
	return &NullableNFStatus{value: val, isSet: true}
}

func (v NullableNFStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


