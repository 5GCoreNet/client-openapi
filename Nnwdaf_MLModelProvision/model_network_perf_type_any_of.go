/*
Nnwdaf_MLModelProvision

Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_MLModelProvision

import (
	"encoding/json"
	"fmt"
)

// NetworkPerfTypeAnyOf the model 'NetworkPerfTypeAnyOf'
type NetworkPerfTypeAnyOf string

// List of NetworkPerfType_anyOf
const (
	GNB_ACTIVE_RATIO NetworkPerfTypeAnyOf = "GNB_ACTIVE_RATIO"
	GNB_COMPUTING_USAGE NetworkPerfTypeAnyOf = "GNB_COMPUTING_USAGE"
	GNB_MEMORY_USAGE NetworkPerfTypeAnyOf = "GNB_MEMORY_USAGE"
	GNB_DISK_USAGE NetworkPerfTypeAnyOf = "GNB_DISK_USAGE"
	NUM_OF_UE NetworkPerfTypeAnyOf = "NUM_OF_UE"
	SESS_SUCC_RATIO NetworkPerfTypeAnyOf = "SESS_SUCC_RATIO"
	HO_SUCC_RATIO NetworkPerfTypeAnyOf = "HO_SUCC_RATIO"
)

// All allowed values of NetworkPerfTypeAnyOf enum
var AllowedNetworkPerfTypeAnyOfEnumValues = []NetworkPerfTypeAnyOf{
	"GNB_ACTIVE_RATIO",
	"GNB_COMPUTING_USAGE",
	"GNB_MEMORY_USAGE",
	"GNB_DISK_USAGE",
	"NUM_OF_UE",
	"SESS_SUCC_RATIO",
	"HO_SUCC_RATIO",
}

func (v *NetworkPerfTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkPerfTypeAnyOf(value)
	for _, existing := range AllowedNetworkPerfTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkPerfTypeAnyOf", value)
}

// NewNetworkPerfTypeAnyOfFromValue returns a pointer to a valid NetworkPerfTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkPerfTypeAnyOfFromValue(v string) (*NetworkPerfTypeAnyOf, error) {
	ev := NetworkPerfTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkPerfTypeAnyOf: valid values are %v", v, AllowedNetworkPerfTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkPerfTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedNetworkPerfTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkPerfType_anyOf value
func (v NetworkPerfTypeAnyOf) Ptr() *NetworkPerfTypeAnyOf {
	return &v
}

type NullableNetworkPerfTypeAnyOf struct {
	value *NetworkPerfTypeAnyOf
	isSet bool
}

func (v NullableNetworkPerfTypeAnyOf) Get() *NetworkPerfTypeAnyOf {
	return v.value
}

func (v *NullableNetworkPerfTypeAnyOf) Set(val *NetworkPerfTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkPerfTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkPerfTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkPerfTypeAnyOf(val *NetworkPerfTypeAnyOf) *NullableNetworkPerfTypeAnyOf {
	return &NullableNetworkPerfTypeAnyOf{value: val, isSet: true}
}

func (v NullableNetworkPerfTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkPerfTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

