/*
N5g-ddnmf_Discovery API

N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_N5g-ddnmf_Discovery

import (
	"encoding/json"
	"fmt"
)

// ProseApplicationCodeSuffixPool Contains the Prose Application Code Suffix Pool.
type ProseApplicationCodeSuffixPool struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ProseApplicationCodeSuffixPool) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ProseApplicationCodeSuffixPool)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ProseApplicationCodeSuffixPool) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableProseApplicationCodeSuffixPool struct {
	value *ProseApplicationCodeSuffixPool
	isSet bool
}

func (v NullableProseApplicationCodeSuffixPool) Get() *ProseApplicationCodeSuffixPool {
	return v.value
}

func (v *NullableProseApplicationCodeSuffixPool) Set(val *ProseApplicationCodeSuffixPool) {
	v.value = val
	v.isSet = true
}

func (v NullableProseApplicationCodeSuffixPool) IsSet() bool {
	return v.isSet
}

func (v *NullableProseApplicationCodeSuffixPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProseApplicationCodeSuffixPool(val *ProseApplicationCodeSuffixPool) *NullableProseApplicationCodeSuffixPool {
	return &NullableProseApplicationCodeSuffixPool{value: val, isSet: true}
}

func (v NullableProseApplicationCodeSuffixPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProseApplicationCodeSuffixPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


