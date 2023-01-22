/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// RetainabilityThreshold - Represents a QoS flow retainability threshold.
type RetainabilityThreshold struct {
	Interface{} *interface{}
}

// interface{}AsRetainabilityThreshold is a convenience function that returns interface{} wrapped in RetainabilityThreshold
func Interface{}AsRetainabilityThreshold(v *interface{}) RetainabilityThreshold {
	return RetainabilityThreshold{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RetainabilityThreshold) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface{}
	err = newStrictDecoder(data).Decode(&dst.Interface{})
	if err == nil {
		jsonInterface{}, _ := json.Marshal(dst.Interface{})
		if string(jsonInterface{}) == "{}" { // empty struct
			dst.Interface{} = nil
		} else {
			match++
		}
	} else {
		dst.Interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface{} = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RetainabilityThreshold)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RetainabilityThreshold)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RetainabilityThreshold) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RetainabilityThreshold) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableRetainabilityThreshold struct {
	value *RetainabilityThreshold
	isSet bool
}

func (v NullableRetainabilityThreshold) Get() *RetainabilityThreshold {
	return v.value
}

func (v *NullableRetainabilityThreshold) Set(val *RetainabilityThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableRetainabilityThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableRetainabilityThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetainabilityThreshold(val *RetainabilityThreshold) *NullableRetainabilityThreshold {
	return &NullableRetainabilityThreshold{value: val, isSet: true}
}

func (v NullableRetainabilityThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetainabilityThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


