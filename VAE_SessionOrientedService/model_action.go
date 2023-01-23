/*
VAE_SessionOrientedService

API for VAE_SessionOrientedService   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_VAE_SessionOrientedService

import (
	"encoding/json"
	"fmt"
)

// Action Indicate the action to the session-oriented service.
type Action struct {
	ActionAnyOf *ActionAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Action) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ActionAnyOf
	err = json.Unmarshal(data, &dst.ActionAnyOf);
	if err == nil {
		jsonActionAnyOf, _ := json.Marshal(dst.ActionAnyOf)
		if string(jsonActionAnyOf) == "{}" { // empty struct
			dst.ActionAnyOf = nil
		} else {
			return nil // data stored in dst.ActionAnyOf, return on the first match
		}
	} else {
		dst.ActionAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(Action)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Action) MarshalJSON() ([]byte, error) {
	if src.ActionAnyOf != nil {
		return json.Marshal(&src.ActionAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAction struct {
	value *Action
	isSet bool
}

func (v NullableAction) Get() *Action {
	return v.value
}

func (v *NullableAction) Set(val *Action) {
	v.value = val
	v.isSet = true
}

func (v NullableAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAction(val *Action) *NullableAction {
	return &NullableAction{value: val, isSet: true}
}

func (v NullableAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


