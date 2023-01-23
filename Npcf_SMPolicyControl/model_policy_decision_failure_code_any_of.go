/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// PolicyDecisionFailureCodeAnyOf the model 'PolicyDecisionFailureCodeAnyOf'
type PolicyDecisionFailureCodeAnyOf string

// List of PolicyDecisionFailureCode_anyOf
const (
	TRA_CTRL_DECS_ERR PolicyDecisionFailureCodeAnyOf = "TRA_CTRL_DECS_ERR"
	QOS_DECS_ERR PolicyDecisionFailureCodeAnyOf = "QOS_DECS_ERR"
	CHG_DECS_ERR PolicyDecisionFailureCodeAnyOf = "CHG_DECS_ERR"
	USA_MON_DECS_ERR PolicyDecisionFailureCodeAnyOf = "USA_MON_DECS_ERR"
	QOS_MON_DECS_ERR PolicyDecisionFailureCodeAnyOf = "QOS_MON_DECS_ERR"
	CON_DATA_ERR PolicyDecisionFailureCodeAnyOf = "CON_DATA_ERR"
	POLICY_PARAM_ERR PolicyDecisionFailureCodeAnyOf = "POLICY_PARAM_ERR"
)

// All allowed values of PolicyDecisionFailureCodeAnyOf enum
var AllowedPolicyDecisionFailureCodeAnyOfEnumValues = []PolicyDecisionFailureCodeAnyOf{
	"TRA_CTRL_DECS_ERR",
	"QOS_DECS_ERR",
	"CHG_DECS_ERR",
	"USA_MON_DECS_ERR",
	"QOS_MON_DECS_ERR",
	"CON_DATA_ERR",
	"POLICY_PARAM_ERR",
}

func (v *PolicyDecisionFailureCodeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PolicyDecisionFailureCodeAnyOf(value)
	for _, existing := range AllowedPolicyDecisionFailureCodeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PolicyDecisionFailureCodeAnyOf", value)
}

// NewPolicyDecisionFailureCodeAnyOfFromValue returns a pointer to a valid PolicyDecisionFailureCodeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPolicyDecisionFailureCodeAnyOfFromValue(v string) (*PolicyDecisionFailureCodeAnyOf, error) {
	ev := PolicyDecisionFailureCodeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PolicyDecisionFailureCodeAnyOf: valid values are %v", v, AllowedPolicyDecisionFailureCodeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PolicyDecisionFailureCodeAnyOf) IsValid() bool {
	for _, existing := range AllowedPolicyDecisionFailureCodeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PolicyDecisionFailureCode_anyOf value
func (v PolicyDecisionFailureCodeAnyOf) Ptr() *PolicyDecisionFailureCodeAnyOf {
	return &v
}

type NullablePolicyDecisionFailureCodeAnyOf struct {
	value *PolicyDecisionFailureCodeAnyOf
	isSet bool
}

func (v NullablePolicyDecisionFailureCodeAnyOf) Get() *PolicyDecisionFailureCodeAnyOf {
	return v.value
}

func (v *NullablePolicyDecisionFailureCodeAnyOf) Set(val *PolicyDecisionFailureCodeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyDecisionFailureCodeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyDecisionFailureCodeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyDecisionFailureCodeAnyOf(val *PolicyDecisionFailureCodeAnyOf) *NullablePolicyDecisionFailureCodeAnyOf {
	return &NullablePolicyDecisionFailureCodeAnyOf{value: val, isSet: true}
}

func (v NullablePolicyDecisionFailureCodeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyDecisionFailureCodeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

