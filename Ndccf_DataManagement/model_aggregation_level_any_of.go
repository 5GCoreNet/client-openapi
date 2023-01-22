/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// AggregationLevelAnyOf the model 'AggregationLevelAnyOf'
type AggregationLevelAnyOf string

// List of AggregationLevel_anyOf
const (
	UE AggregationLevelAnyOf = "UE"
	AOI AggregationLevelAnyOf = "AOI"
)

// All allowed values of AggregationLevelAnyOf enum
var AllowedAggregationLevelAnyOfEnumValues = []AggregationLevelAnyOf{
	"UE",
	"AOI",
}

func (v *AggregationLevelAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AggregationLevelAnyOf(value)
	for _, existing := range AllowedAggregationLevelAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AggregationLevelAnyOf", value)
}

// NewAggregationLevelAnyOfFromValue returns a pointer to a valid AggregationLevelAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAggregationLevelAnyOfFromValue(v string) (*AggregationLevelAnyOf, error) {
	ev := AggregationLevelAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AggregationLevelAnyOf: valid values are %v", v, AllowedAggregationLevelAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AggregationLevelAnyOf) IsValid() bool {
	for _, existing := range AllowedAggregationLevelAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregationLevel_anyOf value
func (v AggregationLevelAnyOf) Ptr() *AggregationLevelAnyOf {
	return &v
}

type NullableAggregationLevelAnyOf struct {
	value *AggregationLevelAnyOf
	isSet bool
}

func (v NullableAggregationLevelAnyOf) Get() *AggregationLevelAnyOf {
	return v.value
}

func (v *NullableAggregationLevelAnyOf) Set(val *AggregationLevelAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregationLevelAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregationLevelAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregationLevelAnyOf(val *AggregationLevelAnyOf) *NullableAggregationLevelAnyOf {
	return &NullableAggregationLevelAnyOf{value: val, isSet: true}
}

func (v NullableAggregationLevelAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregationLevelAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

