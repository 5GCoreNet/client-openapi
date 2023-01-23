/*
Nspaf_SecuredPacket

Nspaf Secured Packet Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nspaf_SecuredPacket

import (
	"encoding/json"
	"fmt"
)

// UiccConfigurationParameter - Represents UICC Configuration Parameters.
type UiccConfigurationParameter struct {
	Interface{} *interface{}
}

// interface{}AsUiccConfigurationParameter is a convenience function that returns interface{} wrapped in UiccConfigurationParameter
func Interface{}AsUiccConfigurationParameter(v *interface{}) UiccConfigurationParameter {
	return UiccConfigurationParameter{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *UiccConfigurationParameter) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(UiccConfigurationParameter)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UiccConfigurationParameter)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UiccConfigurationParameter) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UiccConfigurationParameter) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableUiccConfigurationParameter struct {
	value *UiccConfigurationParameter
	isSet bool
}

func (v NullableUiccConfigurationParameter) Get() *UiccConfigurationParameter {
	return v.value
}

func (v *NullableUiccConfigurationParameter) Set(val *UiccConfigurationParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableUiccConfigurationParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableUiccConfigurationParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiccConfigurationParameter(val *UiccConfigurationParameter) *NullableUiccConfigurationParameter {
	return &NullableUiccConfigurationParameter{value: val, isSet: true}
}

func (v NullableUiccConfigurationParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiccConfigurationParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


