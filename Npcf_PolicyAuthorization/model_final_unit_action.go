/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// FinalUnitAction struct for FinalUnitAction
type FinalUnitAction struct {
	FinalUnitActionAnyOf *FinalUnitActionAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FinalUnitAction) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FinalUnitActionAnyOf
	err = json.Unmarshal(data, &dst.FinalUnitActionAnyOf);
	if err == nil {
		jsonFinalUnitActionAnyOf, _ := json.Marshal(dst.FinalUnitActionAnyOf)
		if string(jsonFinalUnitActionAnyOf) == "{}" { // empty struct
			dst.FinalUnitActionAnyOf = nil
		} else {
			return nil // data stored in dst.FinalUnitActionAnyOf, return on the first match
		}
	} else {
		dst.FinalUnitActionAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(FinalUnitAction)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FinalUnitAction) MarshalJSON() ([]byte, error) {
	if src.FinalUnitActionAnyOf != nil {
		return json.Marshal(&src.FinalUnitActionAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFinalUnitAction struct {
	value *FinalUnitAction
	isSet bool
}

func (v NullableFinalUnitAction) Get() *FinalUnitAction {
	return v.value
}

func (v *NullableFinalUnitAction) Set(val *FinalUnitAction) {
	v.value = val
	v.isSet = true
}

func (v NullableFinalUnitAction) IsSet() bool {
	return v.isSet
}

func (v *NullableFinalUnitAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinalUnitAction(val *FinalUnitAction) *NullableFinalUnitAction {
	return &NullableFinalUnitAction{value: val, isSet: true}
}

func (v NullableFinalUnitAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinalUnitAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


