/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// RequestedRuleDataType Possible values are - CH_ID: Indicates that the requested rule data is the charging identifier.  - MS_TIME_ZONE: Indicates that the requested access network info type is the UE's timezone. - USER_LOC_INFO: Indicates that the requested access network info type is the UE's location. - RES_RELEASE: Indicates that the requested rule data is the result of the release of resource. - SUCC_RES_ALLO: Indicates that the requested rule data is the successful resource allocation. - EPS_FALLBACK: Indicates that the requested rule data is the report of QoS flow rejection due to EPS fallback. 
type RequestedRuleDataType struct {
	RequestedRuleDataTypeAnyOf *RequestedRuleDataTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RequestedRuleDataType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RequestedRuleDataTypeAnyOf
	err = json.Unmarshal(data, &dst.RequestedRuleDataTypeAnyOf);
	if err == nil {
		jsonRequestedRuleDataTypeAnyOf, _ := json.Marshal(dst.RequestedRuleDataTypeAnyOf)
		if string(jsonRequestedRuleDataTypeAnyOf) == "{}" { // empty struct
			dst.RequestedRuleDataTypeAnyOf = nil
		} else {
			return nil // data stored in dst.RequestedRuleDataTypeAnyOf, return on the first match
		}
	} else {
		dst.RequestedRuleDataTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RequestedRuleDataType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RequestedRuleDataType) MarshalJSON() ([]byte, error) {
	if src.RequestedRuleDataTypeAnyOf != nil {
		return json.Marshal(&src.RequestedRuleDataTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRequestedRuleDataType struct {
	value *RequestedRuleDataType
	isSet bool
}

func (v NullableRequestedRuleDataType) Get() *RequestedRuleDataType {
	return v.value
}

func (v *NullableRequestedRuleDataType) Set(val *RequestedRuleDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestedRuleDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestedRuleDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestedRuleDataType(val *RequestedRuleDataType) *NullableRequestedRuleDataType {
	return &NullableRequestedRuleDataType{value: val, isSet: true}
}

func (v NullableRequestedRuleDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestedRuleDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


