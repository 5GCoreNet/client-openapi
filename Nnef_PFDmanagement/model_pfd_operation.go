/*
Nnef_PFDmanagement Service API

Packet Flow Description Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_PFDmanagement

import (
	"encoding/json"
	"fmt"
)

// PfdOperation Indicates the operation to be applied on PFD(s).
type PfdOperation struct {
	PfdOperationAnyOf *PfdOperationAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PfdOperation) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PfdOperationAnyOf
	err = json.Unmarshal(data, &dst.PfdOperationAnyOf);
	if err == nil {
		jsonPfdOperationAnyOf, _ := json.Marshal(dst.PfdOperationAnyOf)
		if string(jsonPfdOperationAnyOf) == "{}" { // empty struct
			dst.PfdOperationAnyOf = nil
		} else {
			return nil // data stored in dst.PfdOperationAnyOf, return on the first match
		}
	} else {
		dst.PfdOperationAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PfdOperation)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PfdOperation) MarshalJSON() ([]byte, error) {
	if src.PfdOperationAnyOf != nil {
		return json.Marshal(&src.PfdOperationAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePfdOperation struct {
	value *PfdOperation
	isSet bool
}

func (v NullablePfdOperation) Get() *PfdOperation {
	return v.value
}

func (v *NullablePfdOperation) Set(val *PfdOperation) {
	v.value = val
	v.isSet = true
}

func (v NullablePfdOperation) IsSet() bool {
	return v.isSet
}

func (v *NullablePfdOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePfdOperation(val *PfdOperation) *NullablePfdOperation {
	return &NullablePfdOperation{value: val, isSet: true}
}

func (v NullablePfdOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePfdOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


