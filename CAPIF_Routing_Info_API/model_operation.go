/*
CAPIF_Routing_Info_API

API for Routing information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CAPIF_Routing_Info_API

import (
	"encoding/json"
	"fmt"
)

// Operation Possible values are: - GET: HTTP GET method - POST: HTTP POST method - PUT: HTTP PUT method - PATCH: HTTP PATCH method - DELETE: HTTP DELETE method 
type Operation struct {
	OperationAnyOf *OperationAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Operation) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into OperationAnyOf
	err = json.Unmarshal(data, &dst.OperationAnyOf);
	if err == nil {
		jsonOperationAnyOf, _ := json.Marshal(dst.OperationAnyOf)
		if string(jsonOperationAnyOf) == "{}" { // empty struct
			dst.OperationAnyOf = nil
		} else {
			return nil // data stored in dst.OperationAnyOf, return on the first match
		}
	} else {
		dst.OperationAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(Operation)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Operation) MarshalJSON() ([]byte, error) {
	if src.OperationAnyOf != nil {
		return json.Marshal(&src.OperationAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableOperation struct {
	value *Operation
	isSet bool
}

func (v NullableOperation) Get() *Operation {
	return v.value
}

func (v *NullableOperation) Set(val *Operation) {
	v.value = val
	v.isSet = true
}

func (v NullableOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperation(val *Operation) *NullableOperation {
	return &NullableOperation{value: val, isSet: true}
}

func (v NullableOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


