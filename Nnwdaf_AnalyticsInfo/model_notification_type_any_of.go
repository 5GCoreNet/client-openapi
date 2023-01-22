/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// NotificationTypeAnyOf the model 'NotificationTypeAnyOf'
type NotificationTypeAnyOf string

// List of NotificationType_anyOf
const (
	N1_MESSAGES NotificationTypeAnyOf = "N1_MESSAGES"
	N2_INFORMATION NotificationTypeAnyOf = "N2_INFORMATION"
	LOCATION_NOTIFICATION NotificationTypeAnyOf = "LOCATION_NOTIFICATION"
	DATA_REMOVAL_NOTIFICATION NotificationTypeAnyOf = "DATA_REMOVAL_NOTIFICATION"
	DATA_CHANGE_NOTIFICATION NotificationTypeAnyOf = "DATA_CHANGE_NOTIFICATION"
	LOCATION_UPDATE_NOTIFICATION NotificationTypeAnyOf = "LOCATION_UPDATE_NOTIFICATION"
	NSSAA_REAUTH_NOTIFICATION NotificationTypeAnyOf = "NSSAA_REAUTH_NOTIFICATION"
	NSSAA_REVOC_NOTIFICATION NotificationTypeAnyOf = "NSSAA_REVOC_NOTIFICATION"
	MATCH_INFO_NOTIFICATION NotificationTypeAnyOf = "MATCH_INFO_NOTIFICATION"
	DATA_RESTORATION_NOTIFICATION NotificationTypeAnyOf = "DATA_RESTORATION_NOTIFICATION"
	TSCTS_NOTIFICATION NotificationTypeAnyOf = "TSCTS_NOTIFICATION"
)

// All allowed values of NotificationTypeAnyOf enum
var AllowedNotificationTypeAnyOfEnumValues = []NotificationTypeAnyOf{
	"N1_MESSAGES",
	"N2_INFORMATION",
	"LOCATION_NOTIFICATION",
	"DATA_REMOVAL_NOTIFICATION",
	"DATA_CHANGE_NOTIFICATION",
	"LOCATION_UPDATE_NOTIFICATION",
	"NSSAA_REAUTH_NOTIFICATION",
	"NSSAA_REVOC_NOTIFICATION",
	"MATCH_INFO_NOTIFICATION",
	"DATA_RESTORATION_NOTIFICATION",
	"TSCTS_NOTIFICATION",
}

func (v *NotificationTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationTypeAnyOf(value)
	for _, existing := range AllowedNotificationTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationTypeAnyOf", value)
}

// NewNotificationTypeAnyOfFromValue returns a pointer to a valid NotificationTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationTypeAnyOfFromValue(v string) (*NotificationTypeAnyOf, error) {
	ev := NotificationTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationTypeAnyOf: valid values are %v", v, AllowedNotificationTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedNotificationTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationType_anyOf value
func (v NotificationTypeAnyOf) Ptr() *NotificationTypeAnyOf {
	return &v
}

type NullableNotificationTypeAnyOf struct {
	value *NotificationTypeAnyOf
	isSet bool
}

func (v NullableNotificationTypeAnyOf) Get() *NotificationTypeAnyOf {
	return v.value
}

func (v *NullableNotificationTypeAnyOf) Set(val *NotificationTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationTypeAnyOf(val *NotificationTypeAnyOf) *NullableNotificationTypeAnyOf {
	return &NullableNotificationTypeAnyOf{value: val, isSet: true}
}

func (v NullableNotificationTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
