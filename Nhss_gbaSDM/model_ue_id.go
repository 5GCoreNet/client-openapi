/*
Nhss_gbaSDM

Nhss Subscriber Data Management Service for GBA.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_gbaSDM

import (
	"encoding/json"
	"fmt"
)

// UeId UE ID can be MSISDN, IMSI, IMPI or IMPU
type UeId struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UeId) UnmarshalJSON(data []byte) error {
	var err error
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

	return fmt.Errorf("data failed to match schemas in anyOf(UeId)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UeId) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUeId struct {
	value *UeId
	isSet bool
}

func (v NullableUeId) Get() *UeId {
	return v.value
}

func (v *NullableUeId) Set(val *UeId) {
	v.value = val
	v.isSet = true
}

func (v NullableUeId) IsSet() bool {
	return v.isSet
}

func (v *NullableUeId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeId(val *UeId) *NullableUeId {
	return &NullableUeId{value: val, isSet: true}
}

func (v NullableUeId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


