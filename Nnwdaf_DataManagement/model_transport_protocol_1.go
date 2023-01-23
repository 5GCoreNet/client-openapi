/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// TransportProtocol1 Types of transport protocol used in a given IP endpoint of an NF Service Instance
type TransportProtocol1 struct {
	TransportProtocol1AnyOf *TransportProtocol1AnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TransportProtocol1) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TransportProtocol1AnyOf
	err = json.Unmarshal(data, &dst.TransportProtocol1AnyOf);
	if err == nil {
		jsonTransportProtocol1AnyOf, _ := json.Marshal(dst.TransportProtocol1AnyOf)
		if string(jsonTransportProtocol1AnyOf) == "{}" { // empty struct
			dst.TransportProtocol1AnyOf = nil
		} else {
			return nil // data stored in dst.TransportProtocol1AnyOf, return on the first match
		}
	} else {
		dst.TransportProtocol1AnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(TransportProtocol1)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TransportProtocol1) MarshalJSON() ([]byte, error) {
	if src.TransportProtocol1AnyOf != nil {
		return json.Marshal(&src.TransportProtocol1AnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTransportProtocol1 struct {
	value *TransportProtocol1
	isSet bool
}

func (v NullableTransportProtocol1) Get() *TransportProtocol1 {
	return v.value
}

func (v *NullableTransportProtocol1) Set(val *TransportProtocol1) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportProtocol1) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportProtocol1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportProtocol1(val *TransportProtocol1) *NullableTransportProtocol1 {
	return &NullableTransportProtocol1{value: val, isSet: true}
}

func (v NullableTransportProtocol1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportProtocol1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


