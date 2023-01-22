/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_UEAU

import (
	"encoding/json"
	"fmt"
)

// ProSeAuthenticationVectors - struct for ProSeAuthenticationVectors
type ProSeAuthenticationVectors struct {
	ArrayOfAvEapAkaPrime *[]AvEapAkaPrime
}

// []AvEapAkaPrimeAsProSeAuthenticationVectors is a convenience function that returns []AvEapAkaPrime wrapped in ProSeAuthenticationVectors
func ArrayOfAvEapAkaPrimeAsProSeAuthenticationVectors(v *[]AvEapAkaPrime) ProSeAuthenticationVectors {
	return ProSeAuthenticationVectors{
		ArrayOfAvEapAkaPrime: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ProSeAuthenticationVectors) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfAvEapAkaPrime
	err = newStrictDecoder(data).Decode(&dst.ArrayOfAvEapAkaPrime)
	if err == nil {
		jsonArrayOfAvEapAkaPrime, _ := json.Marshal(dst.ArrayOfAvEapAkaPrime)
		if string(jsonArrayOfAvEapAkaPrime) == "{}" { // empty struct
			dst.ArrayOfAvEapAkaPrime = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfAvEapAkaPrime = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfAvEapAkaPrime = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ProSeAuthenticationVectors)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ProSeAuthenticationVectors)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ProSeAuthenticationVectors) MarshalJSON() ([]byte, error) {
	if src.ArrayOfAvEapAkaPrime != nil {
		return json.Marshal(&src.ArrayOfAvEapAkaPrime)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ProSeAuthenticationVectors) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfAvEapAkaPrime != nil {
		return obj.ArrayOfAvEapAkaPrime
	}

	// all schemas are nil
	return nil
}

type NullableProSeAuthenticationVectors struct {
	value *ProSeAuthenticationVectors
	isSet bool
}

func (v NullableProSeAuthenticationVectors) Get() *ProSeAuthenticationVectors {
	return v.value
}

func (v *NullableProSeAuthenticationVectors) Set(val *ProSeAuthenticationVectors) {
	v.value = val
	v.isSet = true
}

func (v NullableProSeAuthenticationVectors) IsSet() bool {
	return v.isSet
}

func (v *NullableProSeAuthenticationVectors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProSeAuthenticationVectors(val *ProSeAuthenticationVectors) *NullableProSeAuthenticationVectors {
	return &NullableProSeAuthenticationVectors{value: val, isSet: true}
}

func (v NullableProSeAuthenticationVectors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProSeAuthenticationVectors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


