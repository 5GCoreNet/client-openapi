/*
Nhss_UECM

HSS UE Context Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nhss_UECM

import (
	"encoding/json"
	"fmt"
)

// DeregistrationReason The reason that triggers that the serving node needs to be deregistered by HSS
type DeregistrationReason struct {
	DeregistrationReasonAnyOf *DeregistrationReasonAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DeregistrationReason) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DeregistrationReasonAnyOf
	err = json.Unmarshal(data, &dst.DeregistrationReasonAnyOf);
	if err == nil {
		jsonDeregistrationReasonAnyOf, _ := json.Marshal(dst.DeregistrationReasonAnyOf)
		if string(jsonDeregistrationReasonAnyOf) == "{}" { // empty struct
			dst.DeregistrationReasonAnyOf = nil
		} else {
			return nil // data stored in dst.DeregistrationReasonAnyOf, return on the first match
		}
	} else {
		dst.DeregistrationReasonAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DeregistrationReason)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DeregistrationReason) MarshalJSON() ([]byte, error) {
	if src.DeregistrationReasonAnyOf != nil {
		return json.Marshal(&src.DeregistrationReasonAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDeregistrationReason struct {
	value *DeregistrationReason
	isSet bool
}

func (v NullableDeregistrationReason) Get() *DeregistrationReason {
	return v.value
}

func (v *NullableDeregistrationReason) Set(val *DeregistrationReason) {
	v.value = val
	v.isSet = true
}

func (v NullableDeregistrationReason) IsSet() bool {
	return v.isSet
}

func (v *NullableDeregistrationReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeregistrationReason(val *DeregistrationReason) *NullableDeregistrationReason {
	return &NullableDeregistrationReason{value: val, isSet: true}
}

func (v NullableDeregistrationReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeregistrationReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


