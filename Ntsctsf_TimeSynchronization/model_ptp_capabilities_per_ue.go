/*
Ntsctsf_TimeSynchronization Service API

TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ntsctsf_TimeSynchronization

import (
	"encoding/json"
	"fmt"
)

// PtpCapabilitiesPerUe - Contains the supported PTP capabilities per UE.
type PtpCapabilitiesPerUe struct {
	Interface{} *interface{}
}

// interface{}AsPtpCapabilitiesPerUe is a convenience function that returns interface{} wrapped in PtpCapabilitiesPerUe
func Interface{}AsPtpCapabilitiesPerUe(v *interface{}) PtpCapabilitiesPerUe {
	return PtpCapabilitiesPerUe{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PtpCapabilitiesPerUe) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(PtpCapabilitiesPerUe)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PtpCapabilitiesPerUe)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PtpCapabilitiesPerUe) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PtpCapabilitiesPerUe) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullablePtpCapabilitiesPerUe struct {
	value *PtpCapabilitiesPerUe
	isSet bool
}

func (v NullablePtpCapabilitiesPerUe) Get() *PtpCapabilitiesPerUe {
	return v.value
}

func (v *NullablePtpCapabilitiesPerUe) Set(val *PtpCapabilitiesPerUe) {
	v.value = val
	v.isSet = true
}

func (v NullablePtpCapabilitiesPerUe) IsSet() bool {
	return v.isSet
}

func (v *NullablePtpCapabilitiesPerUe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePtpCapabilitiesPerUe(val *PtpCapabilitiesPerUe) *NullablePtpCapabilitiesPerUe {
	return &NullablePtpCapabilitiesPerUe{value: val, isSet: true}
}

func (v NullablePtpCapabilitiesPerUe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePtpCapabilitiesPerUe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


