/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// OdbPacketServices The enumeration OdbPacketServices defines the Barring of Packet Oriented Services. See 3GPP TS 23.015 for further description. It shall comply with the provisions defined in table 5.7.3.2-1 
type OdbPacketServices struct {
	AnyOfstringstring *AnyOfstringstring
	NullValue *NullValue
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *OdbPacketServices) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AnyOfstringstring
	err = json.Unmarshal(data, &dst.AnyOfstringstring);
	if err == nil {
		jsonAnyOfstringstring, _ := json.Marshal(dst.AnyOfstringstring)
		if string(jsonAnyOfstringstring) == "{}" { // empty struct
			dst.AnyOfstringstring = nil
		} else {
			return nil // data stored in dst.AnyOfstringstring, return on the first match
		}
	} else {
		dst.AnyOfstringstring = nil
	}

	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue);
	if err == nil {
		jsonNullValue, _ := json.Marshal(dst.NullValue)
		if string(jsonNullValue) == "{}" { // empty struct
			dst.NullValue = nil
		} else {
			return nil // data stored in dst.NullValue, return on the first match
		}
	} else {
		dst.NullValue = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(OdbPacketServices)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *OdbPacketServices) MarshalJSON() ([]byte, error) {
	if src.AnyOfstringstring != nil {
		return json.Marshal(&src.AnyOfstringstring)
	}

	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableOdbPacketServices struct {
	value *OdbPacketServices
	isSet bool
}

func (v NullableOdbPacketServices) Get() *OdbPacketServices {
	return v.value
}

func (v *NullableOdbPacketServices) Set(val *OdbPacketServices) {
	v.value = val
	v.isSet = true
}

func (v NullableOdbPacketServices) IsSet() bool {
	return v.isSet
}

func (v *NullableOdbPacketServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOdbPacketServices(val *OdbPacketServices) *NullableOdbPacketServices {
	return &NullableOdbPacketServices{value: val, isSet: true}
}

func (v NullableOdbPacketServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOdbPacketServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


