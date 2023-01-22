/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmfaf_3caDataManagement

import (
	"encoding/json"
	"fmt"
)

// AfEvent Represents Application Events.
type AfEvent struct {
	NefEventAnyOf *NefEventAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AfEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NefEventAnyOf
	err = json.Unmarshal(data, &dst.NefEventAnyOf);
	if err == nil {
		jsonNefEventAnyOf, _ := json.Marshal(dst.NefEventAnyOf)
		if string(jsonNefEventAnyOf) == "{}" { // empty struct
			dst.NefEventAnyOf = nil
		} else {
			return nil // data stored in dst.NefEventAnyOf, return on the first match
		}
	} else {
		dst.NefEventAnyOf = nil
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
	if src.NefEventAnyOf != nil {
		return json.Marshal(&src.NefEventAnyOf)
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


