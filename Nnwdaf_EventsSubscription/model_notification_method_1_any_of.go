/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_EventsSubscription

import (
	"encoding/json"
	"fmt"
)

// NotificationMethod1AnyOf the model 'NotificationMethod1AnyOf'
type NotificationMethod1AnyOf string

// List of NotificationMethod_1_anyOf
const (
	PERIODIC NotificationMethod1AnyOf = "PERIODIC"
	ONE_TIME NotificationMethod1AnyOf = "ONE_TIME"
	ON_EVENT_DETECTION NotificationMethod1AnyOf = "ON_EVENT_DETECTION"
)

// All allowed values of NotificationMethod1AnyOf enum
var AllowedNotificationMethod1AnyOfEnumValues = []NotificationMethod1AnyOf{
	"PERIODIC",
	"ONE_TIME",
	"ON_EVENT_DETECTION",
}

func (v *NotificationMethod1AnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationMethod1AnyOf(value)
	for _, existing := range AllowedNotificationMethod1AnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationMethod1AnyOf", value)
}

// NewNotificationMethod1AnyOfFromValue returns a pointer to a valid NotificationMethod1AnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationMethod1AnyOfFromValue(v string) (*NotificationMethod1AnyOf, error) {
	ev := NotificationMethod1AnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationMethod1AnyOf: valid values are %v", v, AllowedNotificationMethod1AnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationMethod1AnyOf) IsValid() bool {
	for _, existing := range AllowedNotificationMethod1AnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationMethod_1_anyOf value
func (v NotificationMethod1AnyOf) Ptr() *NotificationMethod1AnyOf {
	return &v
}

type NullableNotificationMethod1AnyOf struct {
	value *NotificationMethod1AnyOf
	isSet bool
}

func (v NullableNotificationMethod1AnyOf) Get() *NotificationMethod1AnyOf {
	return v.value
}

func (v *NullableNotificationMethod1AnyOf) Set(val *NotificationMethod1AnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationMethod1AnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationMethod1AnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationMethod1AnyOf(val *NotificationMethod1AnyOf) *NullableNotificationMethod1AnyOf {
	return &NullableNotificationMethod1AnyOf{value: val, isSet: true}
}

func (v NullableNotificationMethod1AnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationMethod1AnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

