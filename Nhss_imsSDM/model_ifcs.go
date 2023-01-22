/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// Ifcs List of IFCs associated to the IMS public Identity
type Ifcs struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Ifcs) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(Ifcs)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Ifcs) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableIfcs struct {
	value *Ifcs
	isSet bool
}

func (v NullableIfcs) Get() *Ifcs {
	return v.value
}

func (v *NullableIfcs) Set(val *Ifcs) {
	v.value = val
	v.isSet = true
}

func (v NullableIfcs) IsSet() bool {
	return v.isSet
}

func (v *NullableIfcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIfcs(val *Ifcs) *NullableIfcs {
	return &NullableIfcs{value: val, isSet: true}
}

func (v NullableIfcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIfcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


