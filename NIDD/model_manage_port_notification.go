/*
3gpp-nidd

API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NIDD

import (
	"encoding/json"
	"fmt"
)

// ManagePortNotification - Represents a ManagePort notification of port numbers that are reserved.
type ManagePortNotification struct {
	Interface{} *interface{}
}

// interface{}AsManagePortNotification is a convenience function that returns interface{} wrapped in ManagePortNotification
func Interface{}AsManagePortNotification(v *interface{}) ManagePortNotification {
	return ManagePortNotification{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ManagePortNotification) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(ManagePortNotification)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ManagePortNotification)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ManagePortNotification) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ManagePortNotification) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableManagePortNotification struct {
	value *ManagePortNotification
	isSet bool
}

func (v NullableManagePortNotification) Get() *ManagePortNotification {
	return v.value
}

func (v *NullableManagePortNotification) Set(val *ManagePortNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableManagePortNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableManagePortNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagePortNotification(val *ManagePortNotification) *NullableManagePortNotification {
	return &NullableManagePortNotification{value: val, isSet: true}
}

func (v NullableManagePortNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagePortNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


