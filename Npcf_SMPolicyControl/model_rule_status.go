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

// RuleStatus Possible values are - ACTIVE: Indicates that the PCC rule(s) are successfully installed (for those provisioned from PCF) or activated (for those pre-defined in SMF), or the session rule(s) are successfully installed  - INACTIVE: Indicates that the PCC rule(s) are removed (for those provisioned from PCF) or inactive (for those pre-defined in SMF) or the session rule(s) are removed. 
type RuleStatus struct {
	RuleStatusAnyOf *RuleStatusAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RuleStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RuleStatusAnyOf
	err = json.Unmarshal(data, &dst.RuleStatusAnyOf);
	if err == nil {
		jsonRuleStatusAnyOf, _ := json.Marshal(dst.RuleStatusAnyOf)
		if string(jsonRuleStatusAnyOf) == "{}" { // empty struct
			dst.RuleStatusAnyOf = nil
		} else {
			return nil // data stored in dst.RuleStatusAnyOf, return on the first match
		}
	} else {
		dst.RuleStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RuleStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RuleStatus) MarshalJSON() ([]byte, error) {
	if src.RuleStatusAnyOf != nil {
		return json.Marshal(&src.RuleStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRuleStatus struct {
	value *RuleStatus
	isSet bool
}

func (v NullableRuleStatus) Get() *RuleStatus {
	return v.value
}

func (v *NullableRuleStatus) Set(val *RuleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleStatus(val *RuleStatus) *NullableRuleStatus {
	return &NullableRuleStatus{value: val, isSet: true}
}

func (v NullableRuleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


