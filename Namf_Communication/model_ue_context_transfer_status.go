/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// UeContextTransferStatus Describes the status of an individual ueContext resource in UE Context Transfer procedures
type UeContextTransferStatus struct {
	UeContextTransferStatusAnyOf *UeContextTransferStatusAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UeContextTransferStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UeContextTransferStatusAnyOf
	err = json.Unmarshal(data, &dst.UeContextTransferStatusAnyOf);
	if err == nil {
		jsonUeContextTransferStatusAnyOf, _ := json.Marshal(dst.UeContextTransferStatusAnyOf)
		if string(jsonUeContextTransferStatusAnyOf) == "{}" { // empty struct
			dst.UeContextTransferStatusAnyOf = nil
		} else {
			return nil // data stored in dst.UeContextTransferStatusAnyOf, return on the first match
		}
	} else {
		dst.UeContextTransferStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(UeContextTransferStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UeContextTransferStatus) MarshalJSON() ([]byte, error) {
	if src.UeContextTransferStatusAnyOf != nil {
		return json.Marshal(&src.UeContextTransferStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUeContextTransferStatus struct {
	value *UeContextTransferStatus
	isSet bool
}

func (v NullableUeContextTransferStatus) Get() *UeContextTransferStatus {
	return v.value
}

func (v *NullableUeContextTransferStatus) Set(val *UeContextTransferStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableUeContextTransferStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableUeContextTransferStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeContextTransferStatus(val *UeContextTransferStatus) *NullableUeContextTransferStatus {
	return &NullableUeContextTransferStatus{value: val, isSet: true}
}

func (v NullableUeContextTransferStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeContextTransferStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


