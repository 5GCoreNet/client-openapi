/*
3gpp-ecr-control

API for enhanced converage restriction control.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ECRControl

import (
	"encoding/json"
	"fmt"
)

// ECRControl - Represents the parameters to request Enhanced Coverage Restriction control.
type ECRControl struct {
	Interface{} *interface{}
}

// interface{}AsECRControl is a convenience function that returns interface{} wrapped in ECRControl
func Interface{}AsECRControl(v *interface{}) ECRControl {
	return ECRControl{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ECRControl) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(ECRControl)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ECRControl)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ECRControl) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ECRControl) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableECRControl struct {
	value *ECRControl
	isSet bool
}

func (v NullableECRControl) Get() *ECRControl {
	return v.value
}

func (v *NullableECRControl) Set(val *ECRControl) {
	v.value = val
	v.isSet = true
}

func (v NullableECRControl) IsSet() bool {
	return v.isSet
}

func (v *NullableECRControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECRControl(val *ECRControl) *NullableECRControl {
	return &NullableECRControl{value: val, isSet: true}
}

func (v NullableECRControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECRControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

