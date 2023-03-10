/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// N2InfoNotifyReason N2 Information Notify Reason
type N2InfoNotifyReason struct {
	N2InfoNotifyReasonAnyOf *N2InfoNotifyReasonAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *N2InfoNotifyReason) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into N2InfoNotifyReasonAnyOf
	err = json.Unmarshal(data, &dst.N2InfoNotifyReasonAnyOf);
	if err == nil {
		jsonN2InfoNotifyReasonAnyOf, _ := json.Marshal(dst.N2InfoNotifyReasonAnyOf)
		if string(jsonN2InfoNotifyReasonAnyOf) == "{}" { // empty struct
			dst.N2InfoNotifyReasonAnyOf = nil
		} else {
			return nil // data stored in dst.N2InfoNotifyReasonAnyOf, return on the first match
		}
	} else {
		dst.N2InfoNotifyReasonAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(N2InfoNotifyReason)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *N2InfoNotifyReason) MarshalJSON() ([]byte, error) {
	if src.N2InfoNotifyReasonAnyOf != nil {
		return json.Marshal(&src.N2InfoNotifyReasonAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableN2InfoNotifyReason struct {
	value *N2InfoNotifyReason
	isSet bool
}

func (v NullableN2InfoNotifyReason) Get() *N2InfoNotifyReason {
	return v.value
}

func (v *NullableN2InfoNotifyReason) Set(val *N2InfoNotifyReason) {
	v.value = val
	v.isSet = true
}

func (v NullableN2InfoNotifyReason) IsSet() bool {
	return v.isSet
}

func (v *NullableN2InfoNotifyReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2InfoNotifyReason(val *N2InfoNotifyReason) *NullableN2InfoNotifyReason {
	return &NullableN2InfoNotifyReason{value: val, isSet: true}
}

func (v NullableN2InfoNotifyReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2InfoNotifyReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


