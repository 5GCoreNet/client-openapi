/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// ReachabilityForSmsConfigurationAnyOf the model 'ReachabilityForSmsConfigurationAnyOf'
type ReachabilityForSmsConfigurationAnyOf string

// List of ReachabilityForSmsConfiguration_anyOf
const (
	NAS ReachabilityForSmsConfigurationAnyOf = "REACHABILITY_FOR_SMS_OVER_NAS"
	IP ReachabilityForSmsConfigurationAnyOf = "REACHABILITY_FOR_SMS_OVER_IP"
)

// All allowed values of ReachabilityForSmsConfigurationAnyOf enum
var AllowedReachabilityForSmsConfigurationAnyOfEnumValues = []ReachabilityForSmsConfigurationAnyOf{
	"REACHABILITY_FOR_SMS_OVER_NAS",
	"REACHABILITY_FOR_SMS_OVER_IP",
}

func (v *ReachabilityForSmsConfigurationAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReachabilityForSmsConfigurationAnyOf(value)
	for _, existing := range AllowedReachabilityForSmsConfigurationAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReachabilityForSmsConfigurationAnyOf", value)
}

// NewReachabilityForSmsConfigurationAnyOfFromValue returns a pointer to a valid ReachabilityForSmsConfigurationAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReachabilityForSmsConfigurationAnyOfFromValue(v string) (*ReachabilityForSmsConfigurationAnyOf, error) {
	ev := ReachabilityForSmsConfigurationAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReachabilityForSmsConfigurationAnyOf: valid values are %v", v, AllowedReachabilityForSmsConfigurationAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReachabilityForSmsConfigurationAnyOf) IsValid() bool {
	for _, existing := range AllowedReachabilityForSmsConfigurationAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReachabilityForSmsConfiguration_anyOf value
func (v ReachabilityForSmsConfigurationAnyOf) Ptr() *ReachabilityForSmsConfigurationAnyOf {
	return &v
}

type NullableReachabilityForSmsConfigurationAnyOf struct {
	value *ReachabilityForSmsConfigurationAnyOf
	isSet bool
}

func (v NullableReachabilityForSmsConfigurationAnyOf) Get() *ReachabilityForSmsConfigurationAnyOf {
	return v.value
}

func (v *NullableReachabilityForSmsConfigurationAnyOf) Set(val *ReachabilityForSmsConfigurationAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReachabilityForSmsConfigurationAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReachabilityForSmsConfigurationAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachabilityForSmsConfigurationAnyOf(val *ReachabilityForSmsConfigurationAnyOf) *NullableReachabilityForSmsConfigurationAnyOf {
	return &NullableReachabilityForSmsConfigurationAnyOf{value: val, isSet: true}
}

func (v NullableReachabilityForSmsConfigurationAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachabilityForSmsConfigurationAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
