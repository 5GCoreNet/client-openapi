/*
NSSF NSSAI Availability

NSSF NSSAI Availability Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnssf_NSSAIAvailability

import (
	"encoding/json"
	"fmt"
)

// NssfEventType This contains the event for the subscription
type NssfEventType struct {
	NssfEventTypeAnyOf *NssfEventTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NssfEventType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NssfEventTypeAnyOf
	err = json.Unmarshal(data, &dst.NssfEventTypeAnyOf);
	if err == nil {
		jsonNssfEventTypeAnyOf, _ := json.Marshal(dst.NssfEventTypeAnyOf)
		if string(jsonNssfEventTypeAnyOf) == "{}" { // empty struct
			dst.NssfEventTypeAnyOf = nil
		} else {
			return nil // data stored in dst.NssfEventTypeAnyOf, return on the first match
		}
	} else {
		dst.NssfEventTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NssfEventType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NssfEventType) MarshalJSON() ([]byte, error) {
	if src.NssfEventTypeAnyOf != nil {
		return json.Marshal(&src.NssfEventTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNssfEventType struct {
	value *NssfEventType
	isSet bool
}

func (v NullableNssfEventType) Get() *NssfEventType {
	return v.value
}

func (v *NullableNssfEventType) Set(val *NssfEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableNssfEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableNssfEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNssfEventType(val *NssfEventType) *NullableNssfEventType {
	return &NullableNssfEventType{value: val, isSet: true}
}

func (v NullableNssfEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNssfEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


