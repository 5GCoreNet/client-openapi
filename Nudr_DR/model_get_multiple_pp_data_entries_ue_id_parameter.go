/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// GetMultiplePPDataEntriesUeIdParameter struct for GetMultiplePPDataEntriesUeIdParameter
type GetMultiplePPDataEntriesUeIdParameter struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GetMultiplePPDataEntriesUeIdParameter) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(GetMultiplePPDataEntriesUeIdParameter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GetMultiplePPDataEntriesUeIdParameter) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGetMultiplePPDataEntriesUeIdParameter struct {
	value *GetMultiplePPDataEntriesUeIdParameter
	isSet bool
}

func (v NullableGetMultiplePPDataEntriesUeIdParameter) Get() *GetMultiplePPDataEntriesUeIdParameter {
	return v.value
}

func (v *NullableGetMultiplePPDataEntriesUeIdParameter) Set(val *GetMultiplePPDataEntriesUeIdParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMultiplePPDataEntriesUeIdParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMultiplePPDataEntriesUeIdParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMultiplePPDataEntriesUeIdParameter(val *GetMultiplePPDataEntriesUeIdParameter) *NullableGetMultiplePPDataEntriesUeIdParameter {
	return &NullableGetMultiplePPDataEntriesUeIdParameter{value: val, isSet: true}
}

func (v NullableGetMultiplePPDataEntriesUeIdParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMultiplePPDataEntriesUeIdParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

