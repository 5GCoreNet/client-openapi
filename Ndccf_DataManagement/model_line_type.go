/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// LineType Possible values are: - DSL: Identifies a DSL line - PON: Identifies a PON line 
type LineType struct {
	LineTypeAnyOf *LineTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LineType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LineTypeAnyOf
	err = json.Unmarshal(data, &dst.LineTypeAnyOf);
	if err == nil {
		jsonLineTypeAnyOf, _ := json.Marshal(dst.LineTypeAnyOf)
		if string(jsonLineTypeAnyOf) == "{}" { // empty struct
			dst.LineTypeAnyOf = nil
		} else {
			return nil // data stored in dst.LineTypeAnyOf, return on the first match
		}
	} else {
		dst.LineTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(LineType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LineType) MarshalJSON() ([]byte, error) {
	if src.LineTypeAnyOf != nil {
		return json.Marshal(&src.LineTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLineType struct {
	value *LineType
	isSet bool
}

func (v NullableLineType) Get() *LineType {
	return v.value
}

func (v *NullableLineType) Set(val *LineType) {
	v.value = val
	v.isSet = true
}

func (v NullableLineType) IsSet() bool {
	return v.isSet
}

func (v *NullableLineType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineType(val *LineType) *NullableLineType {
	return &NullableLineType{value: val, isSet: true}
}

func (v NullableLineType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


