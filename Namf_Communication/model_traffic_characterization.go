/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// TrafficCharacterization Identifies the detailed traffic characterization.
type TrafficCharacterization struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TrafficCharacterization) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(TrafficCharacterization)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TrafficCharacterization) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTrafficCharacterization struct {
	value *TrafficCharacterization
	isSet bool
}

func (v NullableTrafficCharacterization) Get() *TrafficCharacterization {
	return v.value
}

func (v *NullableTrafficCharacterization) Set(val *TrafficCharacterization) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficCharacterization) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficCharacterization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficCharacterization(val *TrafficCharacterization) *NullableTrafficCharacterization {
	return &NullableTrafficCharacterization{value: val, isSet: true}
}

func (v NullableTrafficCharacterization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficCharacterization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


