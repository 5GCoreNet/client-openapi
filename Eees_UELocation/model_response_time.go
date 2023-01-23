/*
EES UE Location Information_API

API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_UELocation

import (
	"encoding/json"
	"fmt"
)

// ResponseTime Indicates acceptable delay of location request.
type ResponseTime struct {
	ResponseTimeAnyOf *ResponseTimeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ResponseTime) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ResponseTimeAnyOf
	err = json.Unmarshal(data, &dst.ResponseTimeAnyOf);
	if err == nil {
		jsonResponseTimeAnyOf, _ := json.Marshal(dst.ResponseTimeAnyOf)
		if string(jsonResponseTimeAnyOf) == "{}" { // empty struct
			dst.ResponseTimeAnyOf = nil
		} else {
			return nil // data stored in dst.ResponseTimeAnyOf, return on the first match
		}
	} else {
		dst.ResponseTimeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ResponseTime)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ResponseTime) MarshalJSON() ([]byte, error) {
	if src.ResponseTimeAnyOf != nil {
		return json.Marshal(&src.ResponseTimeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableResponseTime struct {
	value *ResponseTime
	isSet bool
}

func (v NullableResponseTime) Get() *ResponseTime {
	return v.value
}

func (v *NullableResponseTime) Set(val *ResponseTime) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTime) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTime(val *ResponseTime) *NullableResponseTime {
	return &NullableResponseTime{value: val, isSet: true}
}

func (v NullableResponseTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


