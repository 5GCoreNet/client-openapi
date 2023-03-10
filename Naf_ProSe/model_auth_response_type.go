/*
Naf_ProSe API

Naf_ProSe Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naf_ProSe

import (
	"encoding/json"
	"fmt"
)

// AuthResponseType Possible values are: - OPEN_DISCOVERY_EXTENSION_ANNOUNCE_ACK: Indicates that the Authorization Response Type is  open discovery with application-controlled extension/announce ack. - RESTRICTED_DISCOVERY_ANNOUNCE_ACK: Indicates that the Authorization Response Type is  restricted discovery/announce ack. - RESTRICTED_DISCOVERY_EXTENSION_ANNOUNCE_ACK: Indicates that the Authorization Response Type  is restricted discovery with application-controlled extension/announce ack. - OPEN_DISCOVERY_EXTENSION_MONITOR_ACK: Indicates that the Authorization Response Type is  open discovery with application-controlled extension/monitor ack. - RESTRICTED_DISCOVERY_MONITOR_ACK: Indicates that the Authorization Response Type is  restricted discovery/monitor ack. - RESTRICTED_DISCOVERY_EXTENSION_MONITOR_ACK: Indicates that the Authorization Response Type  is restricted discovery with application-controlled extension/monitor ack. - RESTRICTED_DISCOVERY_PERMISSION_ACK: Indicates that the Authorization Response Type is  restricted discovery/permission ack. - RESTRICTED_DISCOVERY_RESPONSE_ACK: Indicates that the Authorization Response Type is  restricted discovery/response ack. - RESTRICTED_DISCOVERY_QUERY_ACK: Indicates that the Authorization Response Type is  restricted discovery/query ack. - RESTRICTED_DISCOVERY_MATCH_ACK: Indicates that the Authorization Response Type is  restricted discovery/match ack. 
type AuthResponseType struct {
	AuthResponseTypeAnyOf *AuthResponseTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AuthResponseType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AuthResponseTypeAnyOf
	err = json.Unmarshal(data, &dst.AuthResponseTypeAnyOf);
	if err == nil {
		jsonAuthResponseTypeAnyOf, _ := json.Marshal(dst.AuthResponseTypeAnyOf)
		if string(jsonAuthResponseTypeAnyOf) == "{}" { // empty struct
			dst.AuthResponseTypeAnyOf = nil
		} else {
			return nil // data stored in dst.AuthResponseTypeAnyOf, return on the first match
		}
	} else {
		dst.AuthResponseTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AuthResponseType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AuthResponseType) MarshalJSON() ([]byte, error) {
	if src.AuthResponseTypeAnyOf != nil {
		return json.Marshal(&src.AuthResponseTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAuthResponseType struct {
	value *AuthResponseType
	isSet bool
}

func (v NullableAuthResponseType) Get() *AuthResponseType {
	return v.value
}

func (v *NullableAuthResponseType) Set(val *AuthResponseType) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthResponseType) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthResponseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthResponseType(val *AuthResponseType) *NullableAuthResponseType {
	return &NullableAuthResponseType{value: val, isSet: true}
}

func (v NullableAuthResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthResponseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


