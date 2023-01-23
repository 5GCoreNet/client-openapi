/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFDiscovery

import (
	"encoding/json"
	"fmt"
)

// NfServiceInstance - NF service instance
type NfServiceInstance struct {
	Interface{} *interface{}
}

// interface{}AsNfServiceInstance is a convenience function that returns interface{} wrapped in NfServiceInstance
func Interface{}AsNfServiceInstance(v *interface{}) NfServiceInstance {
	return NfServiceInstance{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NfServiceInstance) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(NfServiceInstance)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(NfServiceInstance)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NfServiceInstance) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NfServiceInstance) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableNfServiceInstance struct {
	value *NfServiceInstance
	isSet bool
}

func (v NullableNfServiceInstance) Get() *NfServiceInstance {
	return v.value
}

func (v *NullableNfServiceInstance) Set(val *NfServiceInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableNfServiceInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableNfServiceInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfServiceInstance(val *NfServiceInstance) *NullableNfServiceInstance {
	return &NullableNfServiceInstance{value: val, isSet: true}
}

func (v NullableNfServiceInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfServiceInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


