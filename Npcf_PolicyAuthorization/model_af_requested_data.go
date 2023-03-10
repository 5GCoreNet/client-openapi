/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// AfRequestedData Represents the information that the AF requested to be exposed.
type AfRequestedData struct {
	AfRequestedDataAnyOf *AfRequestedDataAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AfRequestedData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AfRequestedDataAnyOf
	err = json.Unmarshal(data, &dst.AfRequestedDataAnyOf);
	if err == nil {
		jsonAfRequestedDataAnyOf, _ := json.Marshal(dst.AfRequestedDataAnyOf)
		if string(jsonAfRequestedDataAnyOf) == "{}" { // empty struct
			dst.AfRequestedDataAnyOf = nil
		} else {
			return nil // data stored in dst.AfRequestedDataAnyOf, return on the first match
		}
	} else {
		dst.AfRequestedDataAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AfRequestedData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AfRequestedData) MarshalJSON() ([]byte, error) {
	if src.AfRequestedDataAnyOf != nil {
		return json.Marshal(&src.AfRequestedDataAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAfRequestedData struct {
	value *AfRequestedData
	isSet bool
}

func (v NullableAfRequestedData) Get() *AfRequestedData {
	return v.value
}

func (v *NullableAfRequestedData) Set(val *AfRequestedData) {
	v.value = val
	v.isSet = true
}

func (v NullableAfRequestedData) IsSet() bool {
	return v.isSet
}

func (v *NullableAfRequestedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfRequestedData(val *AfRequestedData) *NullableAfRequestedData {
	return &NullableAfRequestedData{value: val, isSet: true}
}

func (v NullableAfRequestedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfRequestedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


