/*
3gpp-eas-deployment

API for AF provisioned EAS Deployment.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EASDeployment

import (
	"encoding/json"
	"fmt"
)

// EdiDeleteCriteria Contains criteria to be used for deleting EAS Deployment Information entries that match them.
type EdiDeleteCriteria struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EdiDeleteCriteria) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(EdiDeleteCriteria)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EdiDeleteCriteria) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEdiDeleteCriteria struct {
	value *EdiDeleteCriteria
	isSet bool
}

func (v NullableEdiDeleteCriteria) Get() *EdiDeleteCriteria {
	return v.value
}

func (v *NullableEdiDeleteCriteria) Set(val *EdiDeleteCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableEdiDeleteCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableEdiDeleteCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdiDeleteCriteria(val *EdiDeleteCriteria) *NullableEdiDeleteCriteria {
	return &NullableEdiDeleteCriteria{value: val, isSet: true}
}

func (v NullableEdiDeleteCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdiDeleteCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


