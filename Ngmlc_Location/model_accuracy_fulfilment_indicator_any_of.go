/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ngmlc_Location

import (
	"encoding/json"
	"fmt"
)

// AccuracyFulfilmentIndicatorAnyOf the model 'AccuracyFulfilmentIndicatorAnyOf'
type AccuracyFulfilmentIndicatorAnyOf string

// List of AccuracyFulfilmentIndicator_anyOf
const (
	FULFILLED AccuracyFulfilmentIndicatorAnyOf = "REQUESTED_ACCURACY_FULFILLED"
	NOT_FULFILLED AccuracyFulfilmentIndicatorAnyOf = "REQUESTED_ACCURACY_NOT_FULFILLED"
)

// All allowed values of AccuracyFulfilmentIndicatorAnyOf enum
var AllowedAccuracyFulfilmentIndicatorAnyOfEnumValues = []AccuracyFulfilmentIndicatorAnyOf{
	"REQUESTED_ACCURACY_FULFILLED",
	"REQUESTED_ACCURACY_NOT_FULFILLED",
}

func (v *AccuracyFulfilmentIndicatorAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccuracyFulfilmentIndicatorAnyOf(value)
	for _, existing := range AllowedAccuracyFulfilmentIndicatorAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccuracyFulfilmentIndicatorAnyOf", value)
}

// NewAccuracyFulfilmentIndicatorAnyOfFromValue returns a pointer to a valid AccuracyFulfilmentIndicatorAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccuracyFulfilmentIndicatorAnyOfFromValue(v string) (*AccuracyFulfilmentIndicatorAnyOf, error) {
	ev := AccuracyFulfilmentIndicatorAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccuracyFulfilmentIndicatorAnyOf: valid values are %v", v, AllowedAccuracyFulfilmentIndicatorAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccuracyFulfilmentIndicatorAnyOf) IsValid() bool {
	for _, existing := range AllowedAccuracyFulfilmentIndicatorAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccuracyFulfilmentIndicator_anyOf value
func (v AccuracyFulfilmentIndicatorAnyOf) Ptr() *AccuracyFulfilmentIndicatorAnyOf {
	return &v
}

type NullableAccuracyFulfilmentIndicatorAnyOf struct {
	value *AccuracyFulfilmentIndicatorAnyOf
	isSet bool
}

func (v NullableAccuracyFulfilmentIndicatorAnyOf) Get() *AccuracyFulfilmentIndicatorAnyOf {
	return v.value
}

func (v *NullableAccuracyFulfilmentIndicatorAnyOf) Set(val *AccuracyFulfilmentIndicatorAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccuracyFulfilmentIndicatorAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccuracyFulfilmentIndicatorAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccuracyFulfilmentIndicatorAnyOf(val *AccuracyFulfilmentIndicatorAnyOf) *NullableAccuracyFulfilmentIndicatorAnyOf {
	return &NullableAccuracyFulfilmentIndicatorAnyOf{value: val, isSet: true}
}

func (v NullableAccuracyFulfilmentIndicatorAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccuracyFulfilmentIndicatorAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

