/*
CAPIF_API_Invoker_Management_API

API for API invoker management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_API_Invoker_Management_API

import (
	"encoding/json"
	"fmt"
)

// InterfaceDescription - Represents the description of an API's interface.
type InterfaceDescription struct {
	Interface{} *interface{}
}

// interface{}AsInterfaceDescription is a convenience function that returns interface{} wrapped in InterfaceDescription
func Interface{}AsInterfaceDescription(v *interface{}) InterfaceDescription {
	return InterfaceDescription{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InterfaceDescription) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(InterfaceDescription)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InterfaceDescription)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InterfaceDescription) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InterfaceDescription) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableInterfaceDescription struct {
	value *InterfaceDescription
	isSet bool
}

func (v NullableInterfaceDescription) Get() *InterfaceDescription {
	return v.value
}

func (v *NullableInterfaceDescription) Set(val *InterfaceDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceDescription(val *InterfaceDescription) *NullableInterfaceDescription {
	return &NullableInterfaceDescription{value: val, isSet: true}
}

func (v NullableInterfaceDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


