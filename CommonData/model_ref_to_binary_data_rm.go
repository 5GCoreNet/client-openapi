/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
	"fmt"
)

// RefToBinaryDataRm This data type is defined in the same way as the ' RefToBinaryData ' data type, but with the OpenAPI 'nullable: true' property.  
type RefToBinaryDataRm struct {
	NullValue *NullValue
	RefToBinaryData *RefToBinaryData
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RefToBinaryDataRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue);
	if err == nil {
		jsonNullValue, _ := json.Marshal(dst.NullValue)
		if string(jsonNullValue) == "{}" { // empty struct
			dst.NullValue = nil
		} else {
			return nil // data stored in dst.NullValue, return on the first match
		}
	} else {
		dst.NullValue = nil
	}

	// try to unmarshal JSON data into RefToBinaryData
	err = json.Unmarshal(data, &dst.RefToBinaryData);
	if err == nil {
		jsonRefToBinaryData, _ := json.Marshal(dst.RefToBinaryData)
		if string(jsonRefToBinaryData) == "{}" { // empty struct
			dst.RefToBinaryData = nil
		} else {
			return nil // data stored in dst.RefToBinaryData, return on the first match
		}
	} else {
		dst.RefToBinaryData = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RefToBinaryDataRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RefToBinaryDataRm) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.RefToBinaryData != nil {
		return json.Marshal(&src.RefToBinaryData)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRefToBinaryDataRm struct {
	value *RefToBinaryDataRm
	isSet bool
}

func (v NullableRefToBinaryDataRm) Get() *RefToBinaryDataRm {
	return v.value
}

func (v *NullableRefToBinaryDataRm) Set(val *RefToBinaryDataRm) {
	v.value = val
	v.isSet = true
}

func (v NullableRefToBinaryDataRm) IsSet() bool {
	return v.isSet
}

func (v *NullableRefToBinaryDataRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefToBinaryDataRm(val *RefToBinaryDataRm) *NullableRefToBinaryDataRm {
	return &NullableRefToBinaryDataRm{value: val, isSet: true}
}

func (v NullableRefToBinaryDataRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefToBinaryDataRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


