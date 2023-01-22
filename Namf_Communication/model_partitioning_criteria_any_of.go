/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// PartitioningCriteriaAnyOf the model 'PartitioningCriteriaAnyOf'
type PartitioningCriteriaAnyOf string

// List of PartitioningCriteria_anyOf
const (
	TAC PartitioningCriteriaAnyOf = "TAC"
	SUBPLMN PartitioningCriteriaAnyOf = "SUBPLMN"
	GEOAREA PartitioningCriteriaAnyOf = "GEOAREA"
	SNSSAI PartitioningCriteriaAnyOf = "SNSSAI"
	DNN PartitioningCriteriaAnyOf = "DNN"
)

// All allowed values of PartitioningCriteriaAnyOf enum
var AllowedPartitioningCriteriaAnyOfEnumValues = []PartitioningCriteriaAnyOf{
	"TAC",
	"SUBPLMN",
	"GEOAREA",
	"SNSSAI",
	"DNN",
}

func (v *PartitioningCriteriaAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PartitioningCriteriaAnyOf(value)
	for _, existing := range AllowedPartitioningCriteriaAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PartitioningCriteriaAnyOf", value)
}

// NewPartitioningCriteriaAnyOfFromValue returns a pointer to a valid PartitioningCriteriaAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPartitioningCriteriaAnyOfFromValue(v string) (*PartitioningCriteriaAnyOf, error) {
	ev := PartitioningCriteriaAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PartitioningCriteriaAnyOf: valid values are %v", v, AllowedPartitioningCriteriaAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PartitioningCriteriaAnyOf) IsValid() bool {
	for _, existing := range AllowedPartitioningCriteriaAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PartitioningCriteria_anyOf value
func (v PartitioningCriteriaAnyOf) Ptr() *PartitioningCriteriaAnyOf {
	return &v
}

type NullablePartitioningCriteriaAnyOf struct {
	value *PartitioningCriteriaAnyOf
	isSet bool
}

func (v NullablePartitioningCriteriaAnyOf) Get() *PartitioningCriteriaAnyOf {
	return v.value
}

func (v *NullablePartitioningCriteriaAnyOf) Set(val *PartitioningCriteriaAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePartitioningCriteriaAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePartitioningCriteriaAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartitioningCriteriaAnyOf(val *PartitioningCriteriaAnyOf) *NullablePartitioningCriteriaAnyOf {
	return &NullablePartitioningCriteriaAnyOf{value: val, isSet: true}
}

func (v NullablePartitioningCriteriaAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartitioningCriteriaAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

