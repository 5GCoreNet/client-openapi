/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// AnNodeType Access Network Node Type (gNB, ng-eNB...)
type AnNodeType struct {
	AnNodeTypeAnyOf *AnNodeTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AnNodeType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AnNodeTypeAnyOf
	err = json.Unmarshal(data, &dst.AnNodeTypeAnyOf);
	if err == nil {
		jsonAnNodeTypeAnyOf, _ := json.Marshal(dst.AnNodeTypeAnyOf)
		if string(jsonAnNodeTypeAnyOf) == "{}" { // empty struct
			dst.AnNodeTypeAnyOf = nil
		} else {
			return nil // data stored in dst.AnNodeTypeAnyOf, return on the first match
		}
	} else {
		dst.AnNodeTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AnNodeType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AnNodeType) MarshalJSON() ([]byte, error) {
	if src.AnNodeTypeAnyOf != nil {
		return json.Marshal(&src.AnNodeTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAnNodeType struct {
	value *AnNodeType
	isSet bool
}

func (v NullableAnNodeType) Get() *AnNodeType {
	return v.value
}

func (v *NullableAnNodeType) Set(val *AnNodeType) {
	v.value = val
	v.isSet = true
}

func (v NullableAnNodeType) IsSet() bool {
	return v.isSet
}

func (v *NullableAnNodeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnNodeType(val *AnNodeType) *NullableAnNodeType {
	return &NullableAnNodeType{value: val, isSet: true}
}

func (v NullableAnNodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnNodeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


