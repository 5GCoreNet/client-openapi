/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// NotificationControlIndicationAnyOf the model 'NotificationControlIndicationAnyOf'
type NotificationControlIndicationAnyOf string

// List of NotificationControlIndication_anyOf
const (
	DDN_FAILURE NotificationControlIndicationAnyOf = "DDN_FAILURE"
	DDD_STATUS NotificationControlIndicationAnyOf = "DDD_STATUS"
)

// All allowed values of NotificationControlIndicationAnyOf enum
var AllowedNotificationControlIndicationAnyOfEnumValues = []NotificationControlIndicationAnyOf{
	"DDN_FAILURE",
	"DDD_STATUS",
}

func (v *NotificationControlIndicationAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationControlIndicationAnyOf(value)
	for _, existing := range AllowedNotificationControlIndicationAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationControlIndicationAnyOf", value)
}

// NewNotificationControlIndicationAnyOfFromValue returns a pointer to a valid NotificationControlIndicationAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationControlIndicationAnyOfFromValue(v string) (*NotificationControlIndicationAnyOf, error) {
	ev := NotificationControlIndicationAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationControlIndicationAnyOf: valid values are %v", v, AllowedNotificationControlIndicationAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationControlIndicationAnyOf) IsValid() bool {
	for _, existing := range AllowedNotificationControlIndicationAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationControlIndication_anyOf value
func (v NotificationControlIndicationAnyOf) Ptr() *NotificationControlIndicationAnyOf {
	return &v
}

type NullableNotificationControlIndicationAnyOf struct {
	value *NotificationControlIndicationAnyOf
	isSet bool
}

func (v NullableNotificationControlIndicationAnyOf) Get() *NotificationControlIndicationAnyOf {
	return v.value
}

func (v *NullableNotificationControlIndicationAnyOf) Set(val *NotificationControlIndicationAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationControlIndicationAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationControlIndicationAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationControlIndicationAnyOf(val *NotificationControlIndicationAnyOf) *NullableNotificationControlIndicationAnyOf {
	return &NullableNotificationControlIndicationAnyOf{value: val, isSet: true}
}

func (v NullableNotificationControlIndicationAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationControlIndicationAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

