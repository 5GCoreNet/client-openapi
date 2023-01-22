/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// UcPurposeAnyOf the model 'UcPurposeAnyOf'
type UcPurposeAnyOf string

// List of UcPurpose_anyOf
const (
	ANALYTICS UcPurposeAnyOf = "ANALYTICS"
	MODEL_TRAINING UcPurposeAnyOf = "MODEL_TRAINING"
	NW_CAP_EXPOSURE UcPurposeAnyOf = "NW_CAP_EXPOSURE"
	EDGEAPP_UE_LOCATION UcPurposeAnyOf = "EDGEAPP_UE_LOCATION"
)

// All allowed values of UcPurposeAnyOf enum
var AllowedUcPurposeAnyOfEnumValues = []UcPurposeAnyOf{
	"ANALYTICS",
	"MODEL_TRAINING",
	"NW_CAP_EXPOSURE",
	"EDGEAPP_UE_LOCATION",
}

func (v *UcPurposeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UcPurposeAnyOf(value)
	for _, existing := range AllowedUcPurposeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UcPurposeAnyOf", value)
}

// NewUcPurposeAnyOfFromValue returns a pointer to a valid UcPurposeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUcPurposeAnyOfFromValue(v string) (*UcPurposeAnyOf, error) {
	ev := UcPurposeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UcPurposeAnyOf: valid values are %v", v, AllowedUcPurposeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UcPurposeAnyOf) IsValid() bool {
	for _, existing := range AllowedUcPurposeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UcPurpose_anyOf value
func (v UcPurposeAnyOf) Ptr() *UcPurposeAnyOf {
	return &v
}

type NullableUcPurposeAnyOf struct {
	value *UcPurposeAnyOf
	isSet bool
}

func (v NullableUcPurposeAnyOf) Get() *UcPurposeAnyOf {
	return v.value
}

func (v *NullableUcPurposeAnyOf) Set(val *UcPurposeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUcPurposeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUcPurposeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUcPurposeAnyOf(val *UcPurposeAnyOf) *NullableUcPurposeAnyOf {
	return &NullableUcPurposeAnyOf{value: val, isSet: true}
}

func (v NullableUcPurposeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUcPurposeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

