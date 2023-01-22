/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// NotificationCause Cause for generating a notification. Possible values are - QOS_FULFILLED - QOS_NOT_FULFILLED - UP_SEC_FULFILLED - UP_SEC_NOT_FULFILLED 
type NotificationCause struct {
	NotificationCauseAnyOf *NotificationCauseAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NotificationCause) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NotificationCauseAnyOf
	err = json.Unmarshal(data, &dst.NotificationCauseAnyOf);
	if err == nil {
		jsonNotificationCauseAnyOf, _ := json.Marshal(dst.NotificationCauseAnyOf)
		if string(jsonNotificationCauseAnyOf) == "{}" { // empty struct
			dst.NotificationCauseAnyOf = nil
		} else {
			return nil // data stored in dst.NotificationCauseAnyOf, return on the first match
		}
	} else {
		dst.NotificationCauseAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NotificationCause)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NotificationCause) MarshalJSON() ([]byte, error) {
	if src.NotificationCauseAnyOf != nil {
		return json.Marshal(&src.NotificationCauseAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNotificationCause struct {
	value *NotificationCause
	isSet bool
}

func (v NullableNotificationCause) Get() *NotificationCause {
	return v.value
}

func (v *NullableNotificationCause) Set(val *NotificationCause) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationCause) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationCause(val *NotificationCause) *NullableNotificationCause {
	return &NullableNotificationCause{value: val, isSet: true}
}

func (v NullableNotificationCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


