/*
SS_Events

API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SS_Events

import (
	"encoding/json"
	"fmt"
)

// LocDevNotificationAnyOf the model 'LocDevNotificationAnyOf'
type LocDevNotificationAnyOf string

// List of LocDevNotification_anyOf
const (
	MISMATCH_LOCATION LocDevNotificationAnyOf = "NOTIFY_MISMATCH_LOCATION"
	ABSENCE LocDevNotificationAnyOf = "NOTIFY_ABSENCE"
	PRESENCE LocDevNotificationAnyOf = "NOTIFY_PRESENCE"
)

// All allowed values of LocDevNotificationAnyOf enum
var AllowedLocDevNotificationAnyOfEnumValues = []LocDevNotificationAnyOf{
	"NOTIFY_MISMATCH_LOCATION",
	"NOTIFY_ABSENCE",
	"NOTIFY_PRESENCE",
}

func (v *LocDevNotificationAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocDevNotificationAnyOf(value)
	for _, existing := range AllowedLocDevNotificationAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LocDevNotificationAnyOf", value)
}

// NewLocDevNotificationAnyOfFromValue returns a pointer to a valid LocDevNotificationAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocDevNotificationAnyOfFromValue(v string) (*LocDevNotificationAnyOf, error) {
	ev := LocDevNotificationAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocDevNotificationAnyOf: valid values are %v", v, AllowedLocDevNotificationAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocDevNotificationAnyOf) IsValid() bool {
	for _, existing := range AllowedLocDevNotificationAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LocDevNotification_anyOf value
func (v LocDevNotificationAnyOf) Ptr() *LocDevNotificationAnyOf {
	return &v
}

type NullableLocDevNotificationAnyOf struct {
	value *LocDevNotificationAnyOf
	isSet bool
}

func (v NullableLocDevNotificationAnyOf) Get() *LocDevNotificationAnyOf {
	return v.value
}

func (v *NullableLocDevNotificationAnyOf) Set(val *LocDevNotificationAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLocDevNotificationAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLocDevNotificationAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocDevNotificationAnyOf(val *LocDevNotificationAnyOf) *NullableLocDevNotificationAnyOf {
	return &NullableLocDevNotificationAnyOf{value: val, isSet: true}
}

func (v NullableLocDevNotificationAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocDevNotificationAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

