/*
3gpp-ms-event-exposure

API for Media Streaming Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MSEventExposure

import (
	"encoding/json"
	"fmt"
)

// DispersionCollection - Contains the dispersion information collected for an AF.
type DispersionCollection struct {
	Interface{} *interface{}
}

// interface{}AsDispersionCollection is a convenience function that returns interface{} wrapped in DispersionCollection
func Interface{}AsDispersionCollection(v *interface{}) DispersionCollection {
	return DispersionCollection{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DispersionCollection) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(DispersionCollection)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DispersionCollection)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DispersionCollection) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DispersionCollection) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableDispersionCollection struct {
	value *DispersionCollection
	isSet bool
}

func (v NullableDispersionCollection) Get() *DispersionCollection {
	return v.value
}

func (v *NullableDispersionCollection) Set(val *DispersionCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableDispersionCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableDispersionCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDispersionCollection(val *DispersionCollection) *NullableDispersionCollection {
	return &NullableDispersionCollection{value: val, isSet: true}
}

func (v NullableDispersionCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDispersionCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


