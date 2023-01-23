/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// AmfEventTriggerAnyOf the model 'AmfEventTriggerAnyOf'
type AmfEventTriggerAnyOf string

// List of AmfEventTrigger_anyOf
const (
	ONE_TIME AmfEventTriggerAnyOf = "ONE_TIME"
	CONTINUOUS AmfEventTriggerAnyOf = "CONTINUOUS"
	PERIODIC AmfEventTriggerAnyOf = "PERIODIC"
)

// All allowed values of AmfEventTriggerAnyOf enum
var AllowedAmfEventTriggerAnyOfEnumValues = []AmfEventTriggerAnyOf{
	"ONE_TIME",
	"CONTINUOUS",
	"PERIODIC",
}

func (v *AmfEventTriggerAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AmfEventTriggerAnyOf(value)
	for _, existing := range AllowedAmfEventTriggerAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AmfEventTriggerAnyOf", value)
}

// NewAmfEventTriggerAnyOfFromValue returns a pointer to a valid AmfEventTriggerAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAmfEventTriggerAnyOfFromValue(v string) (*AmfEventTriggerAnyOf, error) {
	ev := AmfEventTriggerAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AmfEventTriggerAnyOf: valid values are %v", v, AllowedAmfEventTriggerAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AmfEventTriggerAnyOf) IsValid() bool {
	for _, existing := range AllowedAmfEventTriggerAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AmfEventTrigger_anyOf value
func (v AmfEventTriggerAnyOf) Ptr() *AmfEventTriggerAnyOf {
	return &v
}

type NullableAmfEventTriggerAnyOf struct {
	value *AmfEventTriggerAnyOf
	isSet bool
}

func (v NullableAmfEventTriggerAnyOf) Get() *AmfEventTriggerAnyOf {
	return v.value
}

func (v *NullableAmfEventTriggerAnyOf) Set(val *AmfEventTriggerAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventTriggerAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventTriggerAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventTriggerAnyOf(val *AmfEventTriggerAnyOf) *NullableAmfEventTriggerAnyOf {
	return &NullableAmfEventTriggerAnyOf{value: val, isSet: true}
}

func (v NullableAmfEventTriggerAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventTriggerAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

