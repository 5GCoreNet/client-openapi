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

// PlmnIdRm This data type is defined in the same way as the 'PlmnId' data type, but with the OpenAPI 'nullable: true' property.  
type PlmnIdRm struct {
	NullValue *NullValue
	PlmnId *PlmnId
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PlmnIdRm) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into PlmnId
	err = json.Unmarshal(data, &dst.PlmnId);
	if err == nil {
		jsonPlmnId, _ := json.Marshal(dst.PlmnId)
		if string(jsonPlmnId) == "{}" { // empty struct
			dst.PlmnId = nil
		} else {
			return nil // data stored in dst.PlmnId, return on the first match
		}
	} else {
		dst.PlmnId = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PlmnIdRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PlmnIdRm) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.PlmnId != nil {
		return json.Marshal(&src.PlmnId)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePlmnIdRm struct {
	value *PlmnIdRm
	isSet bool
}

func (v NullablePlmnIdRm) Get() *PlmnIdRm {
	return v.value
}

func (v *NullablePlmnIdRm) Set(val *PlmnIdRm) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnIdRm) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnIdRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnIdRm(val *PlmnIdRm) *NullablePlmnIdRm {
	return &NullablePlmnIdRm{value: val, isSet: true}
}

func (v NullablePlmnIdRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnIdRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


