/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// SteerModeIndicator Contains Autonomous load-balance indicator or UE-assistance indicator.
type SteerModeIndicator struct {
	SteerModeIndicatorAnyOf *SteerModeIndicatorAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SteerModeIndicator) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SteerModeIndicatorAnyOf
	err = json.Unmarshal(data, &dst.SteerModeIndicatorAnyOf);
	if err == nil {
		jsonSteerModeIndicatorAnyOf, _ := json.Marshal(dst.SteerModeIndicatorAnyOf)
		if string(jsonSteerModeIndicatorAnyOf) == "{}" { // empty struct
			dst.SteerModeIndicatorAnyOf = nil
		} else {
			return nil // data stored in dst.SteerModeIndicatorAnyOf, return on the first match
		}
	} else {
		dst.SteerModeIndicatorAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SteerModeIndicator)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SteerModeIndicator) MarshalJSON() ([]byte, error) {
	if src.SteerModeIndicatorAnyOf != nil {
		return json.Marshal(&src.SteerModeIndicatorAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSteerModeIndicator struct {
	value *SteerModeIndicator
	isSet bool
}

func (v NullableSteerModeIndicator) Get() *SteerModeIndicator {
	return v.value
}

func (v *NullableSteerModeIndicator) Set(val *SteerModeIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableSteerModeIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableSteerModeIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSteerModeIndicator(val *SteerModeIndicator) *NullableSteerModeIndicator {
	return &NullableSteerModeIndicator{value: val, isSet: true}
}

func (v NullableSteerModeIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSteerModeIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


