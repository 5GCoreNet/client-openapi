/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// TrafficInformation Traffic information including UL/DL data rate and/or Traffic volume.
type TrafficInformation struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TrafficInformation) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(TrafficInformation)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TrafficInformation) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTrafficInformation struct {
	value *TrafficInformation
	isSet bool
}

func (v NullableTrafficInformation) Get() *TrafficInformation {
	return v.value
}

func (v *NullableTrafficInformation) Set(val *TrafficInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficInformation(val *TrafficInformation) *NullableTrafficInformation {
	return &NullableTrafficInformation{value: val, isSet: true}
}

func (v NullableTrafficInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


