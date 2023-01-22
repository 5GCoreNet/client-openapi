/*
AMInfluence

AMInfluence API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package AMInfluence

import (
	"encoding/json"
	"fmt"
)

// AmInfluSub Represents an AM influence subscription.
type AmInfluSub struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AmInfluSub) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.interface{});
	if err == nil {
		jsoninterface{}, _ := json.Marshal(dst.interface{})
		if string(jsoninterface{}) == "{}" { // empty struct
			dst.interface{} = nil
		} else {
			return nil // data stored in dst.interface{}, return on the first match
		}
	} else {
		dst.interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AmInfluSub)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AmInfluSub) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAmInfluSub struct {
	value *AmInfluSub
	isSet bool
}

func (v NullableAmInfluSub) Get() *AmInfluSub {
	return v.value
}

func (v *NullableAmInfluSub) Set(val *AmInfluSub) {
	v.value = val
	v.isSet = true
}

func (v NullableAmInfluSub) IsSet() bool {
	return v.isSet
}

func (v *NullableAmInfluSub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmInfluSub(val *AmInfluSub) *NullableAmInfluSub {
	return &NullableAmInfluSub{value: val, isSet: true}
}

func (v NullableAmInfluSub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmInfluSub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


