/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// TopApplication - Top application that contributes the most to the traffic.
type TopApplication struct {
	Interface{} *interface{}
}

// interface{}AsTopApplication is a convenience function that returns interface{} wrapped in TopApplication
func Interface{}AsTopApplication(v *interface{}) TopApplication {
	return TopApplication{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TopApplication) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(TopApplication)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TopApplication)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TopApplication) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TopApplication) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableTopApplication struct {
	value *TopApplication
	isSet bool
}

func (v NullableTopApplication) Get() *TopApplication {
	return v.value
}

func (v *NullableTopApplication) Set(val *TopApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableTopApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableTopApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopApplication(val *TopApplication) *NullableTopApplication {
	return &NullableTopApplication{value: val, isSet: true}
}

func (v NullableTopApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


