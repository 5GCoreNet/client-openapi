/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package AnalyticsExposure

import (
	"encoding/json"
	"fmt"
)

// ScheduledCommunicationTypeAnyOf the model 'ScheduledCommunicationTypeAnyOf'
type ScheduledCommunicationTypeAnyOf string

// List of ScheduledCommunicationType_anyOf
const (
	DOWNLINK_ONLY ScheduledCommunicationTypeAnyOf = "DOWNLINK_ONLY"
	UPLINK_ONLY ScheduledCommunicationTypeAnyOf = "UPLINK_ONLY"
	BIDIRECTIONAL ScheduledCommunicationTypeAnyOf = "BIDIRECTIONAL"
)

// All allowed values of ScheduledCommunicationTypeAnyOf enum
var AllowedScheduledCommunicationTypeAnyOfEnumValues = []ScheduledCommunicationTypeAnyOf{
	"DOWNLINK_ONLY",
	"UPLINK_ONLY",
	"BIDIRECTIONAL",
}

func (v *ScheduledCommunicationTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScheduledCommunicationTypeAnyOf(value)
	for _, existing := range AllowedScheduledCommunicationTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScheduledCommunicationTypeAnyOf", value)
}

// NewScheduledCommunicationTypeAnyOfFromValue returns a pointer to a valid ScheduledCommunicationTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScheduledCommunicationTypeAnyOfFromValue(v string) (*ScheduledCommunicationTypeAnyOf, error) {
	ev := ScheduledCommunicationTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScheduledCommunicationTypeAnyOf: valid values are %v", v, AllowedScheduledCommunicationTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScheduledCommunicationTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedScheduledCommunicationTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduledCommunicationType_anyOf value
func (v ScheduledCommunicationTypeAnyOf) Ptr() *ScheduledCommunicationTypeAnyOf {
	return &v
}

type NullableScheduledCommunicationTypeAnyOf struct {
	value *ScheduledCommunicationTypeAnyOf
	isSet bool
}

func (v NullableScheduledCommunicationTypeAnyOf) Get() *ScheduledCommunicationTypeAnyOf {
	return v.value
}

func (v *NullableScheduledCommunicationTypeAnyOf) Set(val *ScheduledCommunicationTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledCommunicationTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledCommunicationTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledCommunicationTypeAnyOf(val *ScheduledCommunicationTypeAnyOf) *NullableScheduledCommunicationTypeAnyOf {
	return &NullableScheduledCommunicationTypeAnyOf{value: val, isSet: true}
}

func (v NullableScheduledCommunicationTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledCommunicationTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

