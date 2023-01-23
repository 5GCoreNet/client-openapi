/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
	"fmt"
)

// NotificationEventTypeAnyOf the model 'NotificationEventTypeAnyOf'
type NotificationEventTypeAnyOf string

// List of NotificationEventType_anyOf
const (
	REGISTERED NotificationEventTypeAnyOf = "NF_REGISTERED"
	DEREGISTERED NotificationEventTypeAnyOf = "NF_DEREGISTERED"
	PROFILE_CHANGED NotificationEventTypeAnyOf = "NF_PROFILE_CHANGED"
)

// All allowed values of NotificationEventTypeAnyOf enum
var AllowedNotificationEventTypeAnyOfEnumValues = []NotificationEventTypeAnyOf{
	"NF_REGISTERED",
	"NF_DEREGISTERED",
	"NF_PROFILE_CHANGED",
}

func (v *NotificationEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationEventTypeAnyOf(value)
	for _, existing := range AllowedNotificationEventTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationEventTypeAnyOf", value)
}

// NewNotificationEventTypeAnyOfFromValue returns a pointer to a valid NotificationEventTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationEventTypeAnyOfFromValue(v string) (*NotificationEventTypeAnyOf, error) {
	ev := NotificationEventTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationEventTypeAnyOf: valid values are %v", v, AllowedNotificationEventTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationEventTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedNotificationEventTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationEventType_anyOf value
func (v NotificationEventTypeAnyOf) Ptr() *NotificationEventTypeAnyOf {
	return &v
}

type NullableNotificationEventTypeAnyOf struct {
	value *NotificationEventTypeAnyOf
	isSet bool
}

func (v NullableNotificationEventTypeAnyOf) Get() *NotificationEventTypeAnyOf {
	return v.value
}

func (v *NullableNotificationEventTypeAnyOf) Set(val *NotificationEventTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationEventTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationEventTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationEventTypeAnyOf(val *NotificationEventTypeAnyOf) *NullableNotificationEventTypeAnyOf {
	return &NullableNotificationEventTypeAnyOf{value: val, isSet: true}
}

func (v NullableNotificationEventTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

