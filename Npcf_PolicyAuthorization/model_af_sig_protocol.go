/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// AfSigProtocol Possible values are - NO_INFORMATION: Indicate that no information about the AF signalling protocol is being provided.  - SIP: Indicate that the signalling protocol is Session Initiation Protocol. 
type AfSigProtocol struct {
	AfSigProtocolAnyOf *AfSigProtocolAnyOf
	NullValue *NullValue
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AfSigProtocol) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AfSigProtocolAnyOf
	err = json.Unmarshal(data, &dst.AfSigProtocolAnyOf);
	if err == nil {
		jsonAfSigProtocolAnyOf, _ := json.Marshal(dst.AfSigProtocolAnyOf)
		if string(jsonAfSigProtocolAnyOf) == "{}" { // empty struct
			dst.AfSigProtocolAnyOf = nil
		} else {
			return nil // data stored in dst.AfSigProtocolAnyOf, return on the first match
		}
	} else {
		dst.AfSigProtocolAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AfSigProtocol)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AfSigProtocol) MarshalJSON() ([]byte, error) {
	if src.AfSigProtocolAnyOf != nil {
		return json.Marshal(&src.AfSigProtocolAnyOf)
	}

	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAfSigProtocol struct {
	value *AfSigProtocol
	isSet bool
}

func (v NullableAfSigProtocol) Get() *AfSigProtocol {
	return v.value
}

func (v *NullableAfSigProtocol) Set(val *AfSigProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableAfSigProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableAfSigProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfSigProtocol(val *AfSigProtocol) *NullableAfSigProtocol {
	return &NullableAfSigProtocol{value: val, isSet: true}
}

func (v NullableAfSigProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfSigProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


