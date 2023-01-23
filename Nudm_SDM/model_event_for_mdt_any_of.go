/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
	"fmt"
)

// EventForMdtAnyOf the model 'EventForMdtAnyOf'
type EventForMdtAnyOf string

// List of EventForMdt_anyOf
const (
	OUT_OF_COVERAG EventForMdtAnyOf = "OUT_OF_COVERAG"
	A2_EVENT EventForMdtAnyOf = "A2_EVENT"
)

// All allowed values of EventForMdtAnyOf enum
var AllowedEventForMdtAnyOfEnumValues = []EventForMdtAnyOf{
	"OUT_OF_COVERAG",
	"A2_EVENT",
}

func (v *EventForMdtAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventForMdtAnyOf(value)
	for _, existing := range AllowedEventForMdtAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventForMdtAnyOf", value)
}

// NewEventForMdtAnyOfFromValue returns a pointer to a valid EventForMdtAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventForMdtAnyOfFromValue(v string) (*EventForMdtAnyOf, error) {
	ev := EventForMdtAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventForMdtAnyOf: valid values are %v", v, AllowedEventForMdtAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventForMdtAnyOf) IsValid() bool {
	for _, existing := range AllowedEventForMdtAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventForMdt_anyOf value
func (v EventForMdtAnyOf) Ptr() *EventForMdtAnyOf {
	return &v
}

type NullableEventForMdtAnyOf struct {
	value *EventForMdtAnyOf
	isSet bool
}

func (v NullableEventForMdtAnyOf) Get() *EventForMdtAnyOf {
	return v.value
}

func (v *NullableEventForMdtAnyOf) Set(val *EventForMdtAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEventForMdtAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEventForMdtAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventForMdtAnyOf(val *EventForMdtAnyOf) *NullableEventForMdtAnyOf {
	return &NullableEventForMdtAnyOf{value: val, isSet: true}
}

func (v NullableEventForMdtAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventForMdtAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

