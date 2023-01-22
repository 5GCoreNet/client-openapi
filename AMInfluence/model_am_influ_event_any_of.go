/*
AMInfluence

AMInfluence API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package AMInfluence

import (
	"encoding/json"
	"fmt"
)

// AmInfluEventAnyOf the model 'AmInfluEventAnyOf'
type AmInfluEventAnyOf string

// List of AmInfluEvent_anyOf
const (
	SERVICE_AREA_COVRG_OUTCOME AmInfluEventAnyOf = "SERVICE_AREA_COVRG_OUTCOME"
)

// All allowed values of AmInfluEventAnyOf enum
var AllowedAmInfluEventAnyOfEnumValues = []AmInfluEventAnyOf{
	"SERVICE_AREA_COVRG_OUTCOME",
}

func (v *AmInfluEventAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AmInfluEventAnyOf(value)
	for _, existing := range AllowedAmInfluEventAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AmInfluEventAnyOf", value)
}

// NewAmInfluEventAnyOfFromValue returns a pointer to a valid AmInfluEventAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAmInfluEventAnyOfFromValue(v string) (*AmInfluEventAnyOf, error) {
	ev := AmInfluEventAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AmInfluEventAnyOf: valid values are %v", v, AllowedAmInfluEventAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AmInfluEventAnyOf) IsValid() bool {
	for _, existing := range AllowedAmInfluEventAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AmInfluEvent_anyOf value
func (v AmInfluEventAnyOf) Ptr() *AmInfluEventAnyOf {
	return &v
}

type NullableAmInfluEventAnyOf struct {
	value *AmInfluEventAnyOf
	isSet bool
}

func (v NullableAmInfluEventAnyOf) Get() *AmInfluEventAnyOf {
	return v.value
}

func (v *NullableAmInfluEventAnyOf) Set(val *AmInfluEventAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAmInfluEventAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAmInfluEventAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmInfluEventAnyOf(val *AmInfluEventAnyOf) *NullableAmInfluEventAnyOf {
	return &NullableAmInfluEventAnyOf{value: val, isSet: true}
}

func (v NullableAmInfluEventAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmInfluEventAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

