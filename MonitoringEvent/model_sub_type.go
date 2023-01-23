/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
	"fmt"
)

// SubType Possible values are - AERIAL_UE: The UE has Aerial subscription. 
type SubType struct {
	SubTypeAnyOf *SubTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SubType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SubTypeAnyOf
	err = json.Unmarshal(data, &dst.SubTypeAnyOf);
	if err == nil {
		jsonSubTypeAnyOf, _ := json.Marshal(dst.SubTypeAnyOf)
		if string(jsonSubTypeAnyOf) == "{}" { // empty struct
			dst.SubTypeAnyOf = nil
		} else {
			return nil // data stored in dst.SubTypeAnyOf, return on the first match
		}
	} else {
		dst.SubTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SubType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SubType) MarshalJSON() ([]byte, error) {
	if src.SubTypeAnyOf != nil {
		return json.Marshal(&src.SubTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSubType struct {
	value *SubType
	isSet bool
}

func (v NullableSubType) Get() *SubType {
	return v.value
}

func (v *NullableSubType) Set(val *SubType) {
	v.value = val
	v.isSet = true
}

func (v NullableSubType) IsSet() bool {
	return v.isSet
}

func (v *NullableSubType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubType(val *SubType) *NullableSubType {
	return &NullableSubType{value: val, isSet: true}
}

func (v NullableSubType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


