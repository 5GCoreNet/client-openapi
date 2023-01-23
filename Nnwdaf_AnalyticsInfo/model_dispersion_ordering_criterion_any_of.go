/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// DispersionOrderingCriterionAnyOf the model 'DispersionOrderingCriterionAnyOf'
type DispersionOrderingCriterionAnyOf string

// List of DispersionOrderingCriterion_anyOf
const (
	TIME_SLOT_START DispersionOrderingCriterionAnyOf = "TIME_SLOT_START"
	DISPERSION DispersionOrderingCriterionAnyOf = "DISPERSION"
	CLASSIFICATION DispersionOrderingCriterionAnyOf = "CLASSIFICATION"
	RANKING DispersionOrderingCriterionAnyOf = "RANKING"
	PERCENTILE_RANKING DispersionOrderingCriterionAnyOf = "PERCENTILE_RANKING"
)

// All allowed values of DispersionOrderingCriterionAnyOf enum
var AllowedDispersionOrderingCriterionAnyOfEnumValues = []DispersionOrderingCriterionAnyOf{
	"TIME_SLOT_START",
	"DISPERSION",
	"CLASSIFICATION",
	"RANKING",
	"PERCENTILE_RANKING",
}

func (v *DispersionOrderingCriterionAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DispersionOrderingCriterionAnyOf(value)
	for _, existing := range AllowedDispersionOrderingCriterionAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DispersionOrderingCriterionAnyOf", value)
}

// NewDispersionOrderingCriterionAnyOfFromValue returns a pointer to a valid DispersionOrderingCriterionAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDispersionOrderingCriterionAnyOfFromValue(v string) (*DispersionOrderingCriterionAnyOf, error) {
	ev := DispersionOrderingCriterionAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DispersionOrderingCriterionAnyOf: valid values are %v", v, AllowedDispersionOrderingCriterionAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DispersionOrderingCriterionAnyOf) IsValid() bool {
	for _, existing := range AllowedDispersionOrderingCriterionAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DispersionOrderingCriterion_anyOf value
func (v DispersionOrderingCriterionAnyOf) Ptr() *DispersionOrderingCriterionAnyOf {
	return &v
}

type NullableDispersionOrderingCriterionAnyOf struct {
	value *DispersionOrderingCriterionAnyOf
	isSet bool
}

func (v NullableDispersionOrderingCriterionAnyOf) Get() *DispersionOrderingCriterionAnyOf {
	return v.value
}

func (v *NullableDispersionOrderingCriterionAnyOf) Set(val *DispersionOrderingCriterionAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDispersionOrderingCriterionAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDispersionOrderingCriterionAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDispersionOrderingCriterionAnyOf(val *DispersionOrderingCriterionAnyOf) *NullableDispersionOrderingCriterionAnyOf {
	return &NullableDispersionOrderingCriterionAnyOf{value: val, isSet: true}
}

func (v NullableDispersionOrderingCriterionAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDispersionOrderingCriterionAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

