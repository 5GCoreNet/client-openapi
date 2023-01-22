/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CommonData

import (
	"encoding/json"
	"fmt"
)

// LineTypeRm This data type is defined in the same way as the 'LineType' data type, but with the OpenAPI 'nullable: true' property.  
type LineTypeRm struct {
	LineType *LineType
	NullValue *NullValue
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LineTypeRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LineType
	err = json.Unmarshal(data, &dst.LineType);
	if err == nil {
		jsonLineType, _ := json.Marshal(dst.LineType)
		if string(jsonLineType) == "{}" { // empty struct
			dst.LineType = nil
		} else {
			return nil // data stored in dst.LineType, return on the first match
		}
	} else {
		dst.LineType = nil
	}

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

	return fmt.Errorf("data failed to match schemas in anyOf(LineTypeRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LineTypeRm) MarshalJSON() ([]byte, error) {
	if src.LineType != nil {
		return json.Marshal(&src.LineType)
	}

	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLineTypeRm struct {
	value *LineTypeRm
	isSet bool
}

func (v NullableLineTypeRm) Get() *LineTypeRm {
	return v.value
}

func (v *NullableLineTypeRm) Set(val *LineTypeRm) {
	v.value = val
	v.isSet = true
}

func (v NullableLineTypeRm) IsSet() bool {
	return v.isSet
}

func (v *NullableLineTypeRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineTypeRm(val *LineTypeRm) *NullableLineTypeRm {
	return &NullableLineTypeRm{value: val, isSet: true}
}

func (v NullableLineTypeRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineTypeRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

