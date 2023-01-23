/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
	"fmt"
)

// DnnSmfInfoItemDnn struct for DnnSmfInfoItemDnn
type DnnSmfInfoItemDnn struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DnnSmfInfoItemDnn) UnmarshalJSON(data []byte) error {
	var err error
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

	return fmt.Errorf("data failed to match schemas in anyOf(DnnSmfInfoItemDnn)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DnnSmfInfoItemDnn) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDnnSmfInfoItemDnn struct {
	value *DnnSmfInfoItemDnn
	isSet bool
}

func (v NullableDnnSmfInfoItemDnn) Get() *DnnSmfInfoItemDnn {
	return v.value
}

func (v *NullableDnnSmfInfoItemDnn) Set(val *DnnSmfInfoItemDnn) {
	v.value = val
	v.isSet = true
}

func (v NullableDnnSmfInfoItemDnn) IsSet() bool {
	return v.isSet
}

func (v *NullableDnnSmfInfoItemDnn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnnSmfInfoItemDnn(val *DnnSmfInfoItemDnn) *NullableDnnSmfInfoItemDnn {
	return &NullableDnnSmfInfoItemDnn{value: val, isSet: true}
}

func (v NullableDnnSmfInfoItemDnn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnnSmfInfoItemDnn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


