/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// SscMode represents the service and session continuity mode It shall comply with the provisions defined in table 5.4.3.6-1.  
type SscMode struct {
	SscModeAnyOf *SscModeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SscMode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SscModeAnyOf
	err = json.Unmarshal(data, &dst.SscModeAnyOf);
	if err == nil {
		jsonSscModeAnyOf, _ := json.Marshal(dst.SscModeAnyOf)
		if string(jsonSscModeAnyOf) == "{}" { // empty struct
			dst.SscModeAnyOf = nil
		} else {
			return nil // data stored in dst.SscModeAnyOf, return on the first match
		}
	} else {
		dst.SscModeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SscMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SscMode) MarshalJSON() ([]byte, error) {
	if src.SscModeAnyOf != nil {
		return json.Marshal(&src.SscModeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSscMode struct {
	value *SscMode
	isSet bool
}

func (v NullableSscMode) Get() *SscMode {
	return v.value
}

func (v *NullableSscMode) Set(val *SscMode) {
	v.value = val
	v.isSet = true
}

func (v NullableSscMode) IsSet() bool {
	return v.isSet
}

func (v *NullableSscMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSscMode(val *SscMode) *NullableSscMode {
	return &NullableSscMode{value: val, isSet: true}
}

func (v NullableSscMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSscMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


