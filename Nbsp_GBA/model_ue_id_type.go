/*
GBA BSF Nbsp_GBA Service

GBA BSF Nbsp_GBA Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsp_GBA

import (
	"encoding/json"
	"fmt"
)

// UeIdType Type of UE Identity (public or private)
type UeIdType struct {
	UeIdTypeAnyOf *UeIdTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UeIdType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UeIdTypeAnyOf
	err = json.Unmarshal(data, &dst.UeIdTypeAnyOf);
	if err == nil {
		jsonUeIdTypeAnyOf, _ := json.Marshal(dst.UeIdTypeAnyOf)
		if string(jsonUeIdTypeAnyOf) == "{}" { // empty struct
			dst.UeIdTypeAnyOf = nil
		} else {
			return nil // data stored in dst.UeIdTypeAnyOf, return on the first match
		}
	} else {
		dst.UeIdTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(UeIdType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UeIdType) MarshalJSON() ([]byte, error) {
	if src.UeIdTypeAnyOf != nil {
		return json.Marshal(&src.UeIdTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUeIdType struct {
	value *UeIdType
	isSet bool
}

func (v NullableUeIdType) Get() *UeIdType {
	return v.value
}

func (v *NullableUeIdType) Set(val *UeIdType) {
	v.value = val
	v.isSet = true
}

func (v NullableUeIdType) IsSet() bool {
	return v.isSet
}

func (v *NullableUeIdType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeIdType(val *UeIdType) *NullableUeIdType {
	return &NullableUeIdType{value: val, isSet: true}
}

func (v NullableUeIdType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeIdType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


