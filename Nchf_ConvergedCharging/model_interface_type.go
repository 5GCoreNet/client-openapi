/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// InterfaceType struct for InterfaceType
type InterfaceType struct {
	InterfaceTypeAnyOf *InterfaceTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *InterfaceType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into InterfaceTypeAnyOf
	err = json.Unmarshal(data, &dst.InterfaceTypeAnyOf);
	if err == nil {
		jsonInterfaceTypeAnyOf, _ := json.Marshal(dst.InterfaceTypeAnyOf)
		if string(jsonInterfaceTypeAnyOf) == "{}" { // empty struct
			dst.InterfaceTypeAnyOf = nil
		} else {
			return nil // data stored in dst.InterfaceTypeAnyOf, return on the first match
		}
	} else {
		dst.InterfaceTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(InterfaceType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *InterfaceType) MarshalJSON() ([]byte, error) {
	if src.InterfaceTypeAnyOf != nil {
		return json.Marshal(&src.InterfaceTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableInterfaceType struct {
	value *InterfaceType
	isSet bool
}

func (v NullableInterfaceType) Get() *InterfaceType {
	return v.value
}

func (v *NullableInterfaceType) Set(val *InterfaceType) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceType) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceType(val *InterfaceType) *NullableInterfaceType {
	return &NullableInterfaceType{value: val, isSet: true}
}

func (v NullableInterfaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


