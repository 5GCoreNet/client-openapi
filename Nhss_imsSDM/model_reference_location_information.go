/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// ReferenceLocationInformation Reference Location Information for the user in fixed access networks
type ReferenceLocationInformation struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReferenceLocationInformation) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.interface{});
	if err == nil {
		jsoninterface{}, _ := json.Marshal(dst.interface{})
		if string(jsoninterface{}) == "{}" { // empty struct
			dst.interface{} = nil
		} else {
			return nil // data stored in dst.interface{}, return on the first match
		}
	} else {
		dst.interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ReferenceLocationInformation)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReferenceLocationInformation) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReferenceLocationInformation struct {
	value *ReferenceLocationInformation
	isSet bool
}

func (v NullableReferenceLocationInformation) Get() *ReferenceLocationInformation {
	return v.value
}

func (v *NullableReferenceLocationInformation) Set(val *ReferenceLocationInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceLocationInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceLocationInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceLocationInformation(val *ReferenceLocationInformation) *NullableReferenceLocationInformation {
	return &NullableReferenceLocationInformation{value: val, isSet: true}
}

func (v NullableReferenceLocationInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceLocationInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


