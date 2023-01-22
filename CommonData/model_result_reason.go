/*
TS 29.122 Common Data Types

Data types applicable to several APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CommonData

import (
	"encoding/json"
	"fmt"
)

// ResultReason Possible values are - ROAMING_NOT_ALLOWED: Identifies the configuration parameters are not allowed by roaming agreement. - OTHER_REASON: Identifies the configuration parameters are not configured due to other reason. 
type ResultReason struct {
	ResultReasonAnyOf *ResultReasonAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ResultReason) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ResultReasonAnyOf
	err = json.Unmarshal(data, &dst.ResultReasonAnyOf);
	if err == nil {
		jsonResultReasonAnyOf, _ := json.Marshal(dst.ResultReasonAnyOf)
		if string(jsonResultReasonAnyOf) == "{}" { // empty struct
			dst.ResultReasonAnyOf = nil
		} else {
			return nil // data stored in dst.ResultReasonAnyOf, return on the first match
		}
	} else {
		dst.ResultReasonAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ResultReason)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ResultReason) MarshalJSON() ([]byte, error) {
	if src.ResultReasonAnyOf != nil {
		return json.Marshal(&src.ResultReasonAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableResultReason struct {
	value *ResultReason
	isSet bool
}

func (v NullableResultReason) Get() *ResultReason {
	return v.value
}

func (v *NullableResultReason) Set(val *ResultReason) {
	v.value = val
	v.isSet = true
}

func (v NullableResultReason) IsSet() bool {
	return v.isSet
}

func (v *NullableResultReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultReason(val *ResultReason) *NullableResultReason {
	return &NullableResultReason{value: val, isSet: true}
}

func (v NullableResultReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


