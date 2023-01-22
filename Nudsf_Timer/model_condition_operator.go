/*
Nudsf_Timer

Nudsf Timer Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudsf_Timer

import (
	"encoding/json"
	"fmt"
)

// ConditionOperator TBD
type ConditionOperator struct {
	ConditionOperatorAnyOf *ConditionOperatorAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ConditionOperator) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ConditionOperatorAnyOf
	err = json.Unmarshal(data, &dst.ConditionOperatorAnyOf);
	if err == nil {
		jsonConditionOperatorAnyOf, _ := json.Marshal(dst.ConditionOperatorAnyOf)
		if string(jsonConditionOperatorAnyOf) == "{}" { // empty struct
			dst.ConditionOperatorAnyOf = nil
		} else {
			return nil // data stored in dst.ConditionOperatorAnyOf, return on the first match
		}
	} else {
		dst.ConditionOperatorAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ConditionOperator)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ConditionOperator) MarshalJSON() ([]byte, error) {
	if src.ConditionOperatorAnyOf != nil {
		return json.Marshal(&src.ConditionOperatorAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableConditionOperator struct {
	value *ConditionOperator
	isSet bool
}

func (v NullableConditionOperator) Get() *ConditionOperator {
	return v.value
}

func (v *NullableConditionOperator) Set(val *ConditionOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionOperator(val *ConditionOperator) *NullableConditionOperator {
	return &NullableConditionOperator{value: val, isSet: true}
}

func (v NullableConditionOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


