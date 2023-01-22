/*
NSSF NS Selection

NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnssf_NSSelection

import (
	"encoding/json"
	"fmt"
)

// RoamingIndication Contains the indication on roaming
type RoamingIndication struct {
	RoamingIndicationAnyOf *RoamingIndicationAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RoamingIndication) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RoamingIndicationAnyOf
	err = json.Unmarshal(data, &dst.RoamingIndicationAnyOf);
	if err == nil {
		jsonRoamingIndicationAnyOf, _ := json.Marshal(dst.RoamingIndicationAnyOf)
		if string(jsonRoamingIndicationAnyOf) == "{}" { // empty struct
			dst.RoamingIndicationAnyOf = nil
		} else {
			return nil // data stored in dst.RoamingIndicationAnyOf, return on the first match
		}
	} else {
		dst.RoamingIndicationAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RoamingIndication)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RoamingIndication) MarshalJSON() ([]byte, error) {
	if src.RoamingIndicationAnyOf != nil {
		return json.Marshal(&src.RoamingIndicationAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRoamingIndication struct {
	value *RoamingIndication
	isSet bool
}

func (v NullableRoamingIndication) Get() *RoamingIndication {
	return v.value
}

func (v *NullableRoamingIndication) Set(val *RoamingIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableRoamingIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableRoamingIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoamingIndication(val *RoamingIndication) *NullableRoamingIndication {
	return &NullableRoamingIndication{value: val, isSet: true}
}

func (v NullableRoamingIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoamingIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


