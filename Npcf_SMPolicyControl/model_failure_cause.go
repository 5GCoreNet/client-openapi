/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// FailureCause Indicates the cause of the failure in a Partial Success Report.
type FailureCause struct {
	FailureCauseAnyOf *FailureCauseAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FailureCause) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FailureCauseAnyOf
	err = json.Unmarshal(data, &dst.FailureCauseAnyOf);
	if err == nil {
		jsonFailureCauseAnyOf, _ := json.Marshal(dst.FailureCauseAnyOf)
		if string(jsonFailureCauseAnyOf) == "{}" { // empty struct
			dst.FailureCauseAnyOf = nil
		} else {
			return nil // data stored in dst.FailureCauseAnyOf, return on the first match
		}
	} else {
		dst.FailureCauseAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(FailureCause)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FailureCause) MarshalJSON() ([]byte, error) {
	if src.FailureCauseAnyOf != nil {
		return json.Marshal(&src.FailureCauseAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFailureCause struct {
	value *FailureCause
	isSet bool
}

func (v NullableFailureCause) Get() *FailureCause {
	return v.value
}

func (v *NullableFailureCause) Set(val *FailureCause) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureCause) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureCause(val *FailureCause) *NullableFailureCause {
	return &NullableFailureCause{value: val, isSet: true}
}

func (v NullableFailureCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


