/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// PatchOperation Operations as defined in IETF RFC 6902.
type PatchOperation struct {
	PatchOperationAnyOf *PatchOperationAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PatchOperation) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PatchOperationAnyOf
	err = json.Unmarshal(data, &dst.PatchOperationAnyOf);
	if err == nil {
		jsonPatchOperationAnyOf, _ := json.Marshal(dst.PatchOperationAnyOf)
		if string(jsonPatchOperationAnyOf) == "{}" { // empty struct
			dst.PatchOperationAnyOf = nil
		} else {
			return nil // data stored in dst.PatchOperationAnyOf, return on the first match
		}
	} else {
		dst.PatchOperationAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PatchOperation)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PatchOperation) MarshalJSON() ([]byte, error) {
	if src.PatchOperationAnyOf != nil {
		return json.Marshal(&src.PatchOperationAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePatchOperation struct {
	value *PatchOperation
	isSet bool
}

func (v NullablePatchOperation) Get() *PatchOperation {
	return v.value
}

func (v *NullablePatchOperation) Set(val *PatchOperation) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchOperation) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchOperation(val *PatchOperation) *NullablePatchOperation {
	return &NullablePatchOperation{value: val, isSet: true}
}

func (v NullablePatchOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


