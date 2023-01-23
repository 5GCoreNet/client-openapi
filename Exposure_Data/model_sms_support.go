/*
Unified Data Repository Service API file for structured data for exposure

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Exposure_Data

import (
	"encoding/json"
	"fmt"
)

// SmsSupport Indicates the supported SMS delivery of a UE
type SmsSupport struct {
	SmsSupportAnyOf *SmsSupportAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SmsSupport) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SmsSupportAnyOf
	err = json.Unmarshal(data, &dst.SmsSupportAnyOf);
	if err == nil {
		jsonSmsSupportAnyOf, _ := json.Marshal(dst.SmsSupportAnyOf)
		if string(jsonSmsSupportAnyOf) == "{}" { // empty struct
			dst.SmsSupportAnyOf = nil
		} else {
			return nil // data stored in dst.SmsSupportAnyOf, return on the first match
		}
	} else {
		dst.SmsSupportAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SmsSupport)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SmsSupport) MarshalJSON() ([]byte, error) {
	if src.SmsSupportAnyOf != nil {
		return json.Marshal(&src.SmsSupportAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSmsSupport struct {
	value *SmsSupport
	isSet bool
}

func (v NullableSmsSupport) Get() *SmsSupport {
	return v.value
}

func (v *NullableSmsSupport) Set(val *SmsSupport) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsSupport) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsSupport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsSupport(val *SmsSupport) *NullableSmsSupport {
	return &NullableSmsSupport{value: val, isSet: true}
}

func (v NullableSmsSupport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsSupport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


