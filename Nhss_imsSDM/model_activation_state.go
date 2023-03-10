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

// ActivationState Represents the activation state of a PSI or DSAI tag
type ActivationState struct {
	ActivationStateAnyOf *ActivationStateAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ActivationState) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ActivationStateAnyOf
	err = json.Unmarshal(data, &dst.ActivationStateAnyOf);
	if err == nil {
		jsonActivationStateAnyOf, _ := json.Marshal(dst.ActivationStateAnyOf)
		if string(jsonActivationStateAnyOf) == "{}" { // empty struct
			dst.ActivationStateAnyOf = nil
		} else {
			return nil // data stored in dst.ActivationStateAnyOf, return on the first match
		}
	} else {
		dst.ActivationStateAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ActivationState)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ActivationState) MarshalJSON() ([]byte, error) {
	if src.ActivationStateAnyOf != nil {
		return json.Marshal(&src.ActivationStateAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableActivationState struct {
	value *ActivationState
	isSet bool
}

func (v NullableActivationState) Get() *ActivationState {
	return v.value
}

func (v *NullableActivationState) Set(val *ActivationState) {
	v.value = val
	v.isSet = true
}

func (v NullableActivationState) IsSet() bool {
	return v.isSet
}

func (v *NullableActivationState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivationState(val *ActivationState) *NullableActivationState {
	return &NullableActivationState{value: val, isSet: true}
}

func (v NullableActivationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivationState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


