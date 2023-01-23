/*
CAPIF_Security_API

API for CAPIF security management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Security_API

import (
	"encoding/json"
	"fmt"
)

// SecurityInformation - Represents the interface details and the security method.
type SecurityInformation struct {
	Interface{} *interface{}
}

// interface{}AsSecurityInformation is a convenience function that returns interface{} wrapped in SecurityInformation
func Interface{}AsSecurityInformation(v *interface{}) SecurityInformation {
	return SecurityInformation{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SecurityInformation) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(SecurityInformation)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SecurityInformation)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SecurityInformation) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SecurityInformation) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableSecurityInformation struct {
	value *SecurityInformation
	isSet bool
}

func (v NullableSecurityInformation) Get() *SecurityInformation {
	return v.value
}

func (v *NullableSecurityInformation) Set(val *SecurityInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityInformation(val *SecurityInformation) *NullableSecurityInformation {
	return &NullableSecurityInformation{value: val, isSet: true}
}

func (v NullableSecurityInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


