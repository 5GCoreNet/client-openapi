/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// NotificationEventType Types of events sent in notifications from NRF to subscribed NF Instances
type NotificationEventType struct {
	NotificationEventTypeAnyOf *NotificationEventTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NotificationEventType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NotificationEventTypeAnyOf
	err = json.Unmarshal(data, &dst.NotificationEventTypeAnyOf);
	if err == nil {
		jsonNotificationEventTypeAnyOf, _ := json.Marshal(dst.NotificationEventTypeAnyOf)
		if string(jsonNotificationEventTypeAnyOf) == "{}" { // empty struct
			dst.NotificationEventTypeAnyOf = nil
		} else {
			return nil // data stored in dst.NotificationEventTypeAnyOf, return on the first match
		}
	} else {
		dst.NotificationEventTypeAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NotificationEventType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NotificationEventType) MarshalJSON() ([]byte, error) {
	if src.NotificationEventTypeAnyOf != nil {
		return json.Marshal(&src.NotificationEventTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNotificationEventType struct {
	value *NotificationEventType
	isSet bool
}

func (v NullableNotificationEventType) Get() *NotificationEventType {
	return v.value
}

func (v *NullableNotificationEventType) Set(val *NotificationEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationEventType(val *NotificationEventType) *NullableNotificationEventType {
	return &NullableNotificationEventType{value: val, isSet: true}
}

func (v NullableNotificationEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


