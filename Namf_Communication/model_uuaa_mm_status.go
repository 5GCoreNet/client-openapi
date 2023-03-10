/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// UuaaMmStatus Indicates the UUAA-MM status
type UuaaMmStatus struct {
	UuaaMmStatusAnyOf *UuaaMmStatusAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UuaaMmStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UuaaMmStatusAnyOf
	err = json.Unmarshal(data, &dst.UuaaMmStatusAnyOf);
	if err == nil {
		jsonUuaaMmStatusAnyOf, _ := json.Marshal(dst.UuaaMmStatusAnyOf)
		if string(jsonUuaaMmStatusAnyOf) == "{}" { // empty struct
			dst.UuaaMmStatusAnyOf = nil
		} else {
			return nil // data stored in dst.UuaaMmStatusAnyOf, return on the first match
		}
	} else {
		dst.UuaaMmStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(UuaaMmStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UuaaMmStatus) MarshalJSON() ([]byte, error) {
	if src.UuaaMmStatusAnyOf != nil {
		return json.Marshal(&src.UuaaMmStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUuaaMmStatus struct {
	value *UuaaMmStatus
	isSet bool
}

func (v NullableUuaaMmStatus) Get() *UuaaMmStatus {
	return v.value
}

func (v *NullableUuaaMmStatus) Set(val *UuaaMmStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableUuaaMmStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableUuaaMmStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUuaaMmStatus(val *UuaaMmStatus) *NullableUuaaMmStatus {
	return &NullableUuaaMmStatus{value: val, isSet: true}
}

func (v NullableUuaaMmStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUuaaMmStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


