/*
Unified Data Repository Service API file for policy data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Policy_Data

import (
	"encoding/json"
	"fmt"
)

// PolicyDataSubsetAnyOf the model 'PolicyDataSubsetAnyOf'
type PolicyDataSubsetAnyOf string

// List of PolicyDataSubset_anyOf
const (
	AM_POLICY_DATA PolicyDataSubsetAnyOf = "AM_POLICY_DATA"
	SM_POLICY_DATA PolicyDataSubsetAnyOf = "SM_POLICY_DATA"
	UE_POLICY_DATA PolicyDataSubsetAnyOf = "UE_POLICY_DATA"
	UM_DATA PolicyDataSubsetAnyOf = "UM_DATA"
	OPERATOR_SPECIFIC_DATA PolicyDataSubsetAnyOf = "OPERATOR_SPECIFIC_DATA"
)

// All allowed values of PolicyDataSubsetAnyOf enum
var AllowedPolicyDataSubsetAnyOfEnumValues = []PolicyDataSubsetAnyOf{
	"AM_POLICY_DATA",
	"SM_POLICY_DATA",
	"UE_POLICY_DATA",
	"UM_DATA",
	"OPERATOR_SPECIFIC_DATA",
}

func (v *PolicyDataSubsetAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PolicyDataSubsetAnyOf(value)
	for _, existing := range AllowedPolicyDataSubsetAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PolicyDataSubsetAnyOf", value)
}

// NewPolicyDataSubsetAnyOfFromValue returns a pointer to a valid PolicyDataSubsetAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPolicyDataSubsetAnyOfFromValue(v string) (*PolicyDataSubsetAnyOf, error) {
	ev := PolicyDataSubsetAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PolicyDataSubsetAnyOf: valid values are %v", v, AllowedPolicyDataSubsetAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PolicyDataSubsetAnyOf) IsValid() bool {
	for _, existing := range AllowedPolicyDataSubsetAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PolicyDataSubset_anyOf value
func (v PolicyDataSubsetAnyOf) Ptr() *PolicyDataSubsetAnyOf {
	return &v
}

type NullablePolicyDataSubsetAnyOf struct {
	value *PolicyDataSubsetAnyOf
	isSet bool
}

func (v NullablePolicyDataSubsetAnyOf) Get() *PolicyDataSubsetAnyOf {
	return v.value
}

func (v *NullablePolicyDataSubsetAnyOf) Set(val *PolicyDataSubsetAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyDataSubsetAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyDataSubsetAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyDataSubsetAnyOf(val *PolicyDataSubsetAnyOf) *NullablePolicyDataSubsetAnyOf {
	return &NullablePolicyDataSubsetAnyOf{value: val, isSet: true}
}

func (v NullablePolicyDataSubsetAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyDataSubsetAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

