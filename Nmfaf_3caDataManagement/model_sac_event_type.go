/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmfaf_3caDataManagement

import (
	"encoding/json"
	"fmt"
)

// SACEventType Describes the supported event types
type SACEventType struct {
	SACEventTypeAnyOf *SACEventTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SACEventType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SACEventTypeAnyOf
	err = json.Unmarshal(data, &dst.SACEventTypeAnyOf);
	if err == nil {
		jsonSACEventTypeAnyOf, _ := json.Marshal(dst.SACEventTypeAnyOf)
		if string(jsonSACEventTypeAnyOf) == "{}" { // empty struct
			dst.SACEventTypeAnyOf = nil
		} else {
			return nil // data stored in dst.SACEventTypeAnyOf, return on the first match
		}
	} else {
		dst.SACEventTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SACEventType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SACEventType) MarshalJSON() ([]byte, error) {
	if src.SACEventTypeAnyOf != nil {
		return json.Marshal(&src.SACEventTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSACEventType struct {
	value *SACEventType
	isSet bool
}

func (v NullableSACEventType) Get() *SACEventType {
	return v.value
}

func (v *NullableSACEventType) Set(val *SACEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableSACEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableSACEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSACEventType(val *SACEventType) *NullableSACEventType {
	return &NullableSACEventType{value: val, isSet: true}
}

func (v NullableSACEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSACEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


