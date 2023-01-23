/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// NetLocAccessSupport Possible values are - ANR_NOT_SUPPORTED: Indicates that the access network does not support the report of access network information. - TZR_NOT_SUPPORTED: Indicates that the access network does not support the report of UE time zone. - LOC_NOT_SUPPORTED: Indicates that the access network does not support the report of UE Location (or PLMN Id). 
type NetLocAccessSupport struct {
	NetLocAccessSupportAnyOf *NetLocAccessSupportAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NetLocAccessSupport) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NetLocAccessSupportAnyOf
	err = json.Unmarshal(data, &dst.NetLocAccessSupportAnyOf);
	if err == nil {
		jsonNetLocAccessSupportAnyOf, _ := json.Marshal(dst.NetLocAccessSupportAnyOf)
		if string(jsonNetLocAccessSupportAnyOf) == "{}" { // empty struct
			dst.NetLocAccessSupportAnyOf = nil
		} else {
			return nil // data stored in dst.NetLocAccessSupportAnyOf, return on the first match
		}
	} else {
		dst.NetLocAccessSupportAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NetLocAccessSupport)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NetLocAccessSupport) MarshalJSON() ([]byte, error) {
	if src.NetLocAccessSupportAnyOf != nil {
		return json.Marshal(&src.NetLocAccessSupportAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNetLocAccessSupport struct {
	value *NetLocAccessSupport
	isSet bool
}

func (v NullableNetLocAccessSupport) Get() *NetLocAccessSupport {
	return v.value
}

func (v *NullableNetLocAccessSupport) Set(val *NetLocAccessSupport) {
	v.value = val
	v.isSet = true
}

func (v NullableNetLocAccessSupport) IsSet() bool {
	return v.isSet
}

func (v *NullableNetLocAccessSupport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetLocAccessSupport(val *NetLocAccessSupport) *NullableNetLocAccessSupport {
	return &NullableNetLocAccessSupport{value: val, isSet: true}
}

func (v NullableNetLocAccessSupport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetLocAccessSupport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


