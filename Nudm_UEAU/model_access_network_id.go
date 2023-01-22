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

// AccessNetworkId struct for AccessNetworkId
type AccessNetworkId struct {
	AccessNetworkIdAnyOf *AccessNetworkIdAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AccessNetworkId) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AccessNetworkIdAnyOf
	err = json.Unmarshal(data, &dst.AccessNetworkIdAnyOf);
	if err == nil {
		jsonAccessNetworkIdAnyOf, _ := json.Marshal(dst.AccessNetworkIdAnyOf)
		if string(jsonAccessNetworkIdAnyOf) == "{}" { // empty struct
			dst.AccessNetworkIdAnyOf = nil
		} else {
			return nil // data stored in dst.AccessNetworkIdAnyOf, return on the first match
		}
	} else {
		dst.AccessNetworkIdAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AccessNetworkId)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AccessNetworkId) MarshalJSON() ([]byte, error) {
	if src.AccessNetworkIdAnyOf != nil {
		return json.Marshal(&src.AccessNetworkIdAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAccessNetworkId struct {
	value *AccessNetworkId
	isSet bool
}

func (v NullableAccessNetworkId) Get() *AccessNetworkId {
	return v.value
}

func (v *NullableAccessNetworkId) Set(val *AccessNetworkId) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessNetworkId) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessNetworkId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessNetworkId(val *AccessNetworkId) *NullableAccessNetworkId {
	return &NullableAccessNetworkId{value: val, isSet: true}
}

func (v NullableAccessNetworkId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessNetworkId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


