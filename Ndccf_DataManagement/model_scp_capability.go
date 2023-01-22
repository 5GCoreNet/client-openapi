/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// ScpCapability Indicates the capabilities supported by an SCP
type ScpCapability struct {
	ScpCapabilityAnyOf *ScpCapabilityAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ScpCapability) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ScpCapabilityAnyOf
	err = json.Unmarshal(data, &dst.ScpCapabilityAnyOf);
	if err == nil {
		jsonScpCapabilityAnyOf, _ := json.Marshal(dst.ScpCapabilityAnyOf)
		if string(jsonScpCapabilityAnyOf) == "{}" { // empty struct
			dst.ScpCapabilityAnyOf = nil
		} else {
			return nil // data stored in dst.ScpCapabilityAnyOf, return on the first match
		}
	} else {
		dst.ScpCapabilityAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ScpCapability)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ScpCapability) MarshalJSON() ([]byte, error) {
	if src.ScpCapabilityAnyOf != nil {
		return json.Marshal(&src.ScpCapabilityAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableScpCapability struct {
	value *ScpCapability
	isSet bool
}

func (v NullableScpCapability) Get() *ScpCapability {
	return v.value
}

func (v *NullableScpCapability) Set(val *ScpCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableScpCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableScpCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScpCapability(val *ScpCapability) *NullableScpCapability {
	return &NullableScpCapability{value: val, isSet: true}
}

func (v NullableScpCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScpCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


