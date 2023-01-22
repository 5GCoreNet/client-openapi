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

// ServingNodeAddress struct for ServingNodeAddress
type ServingNodeAddress struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ServingNodeAddress) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ServingNodeAddress)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ServingNodeAddress) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableServingNodeAddress struct {
	value *ServingNodeAddress
	isSet bool
}

func (v NullableServingNodeAddress) Get() *ServingNodeAddress {
	return v.value
}

func (v *NullableServingNodeAddress) Set(val *ServingNodeAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableServingNodeAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableServingNodeAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServingNodeAddress(val *ServingNodeAddress) *NullableServingNodeAddress {
	return &NullableServingNodeAddress{value: val, isSet: true}
}

func (v NullableServingNodeAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServingNodeAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


