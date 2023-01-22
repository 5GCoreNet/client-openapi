/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnrf_NFDiscovery

import (
	"encoding/json"
	"fmt"
)

// AfEvent Represents Application Events.
type AfEvent struct {
	AfEventAnyOf *AfEventAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AfEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AfEventAnyOf
	err = json.Unmarshal(data, &dst.AfEventAnyOf);
	if err == nil {
		jsonAfEventAnyOf, _ := json.Marshal(dst.AfEventAnyOf)
		if string(jsonAfEventAnyOf) == "{}" { // empty struct
			dst.AfEventAnyOf = nil
		} else {
			return nil // data stored in dst.AfEventAnyOf, return on the first match
		}
	} else {
		dst.AfEventAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AfEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AfEvent) MarshalJSON() ([]byte, error) {
	if src.AfEventAnyOf != nil {
		return json.Marshal(&src.AfEventAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAfEvent struct {
	value *AfEvent
	isSet bool
}

func (v NullableAfEvent) Get() *AfEvent {
	return v.value
}

func (v *NullableAfEvent) Set(val *AfEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAfEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAfEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfEvent(val *AfEvent) *NullableAfEvent {
	return &NullableAfEvent{value: val, isSet: true}
}

func (v NullableAfEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


