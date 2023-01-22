/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// NotificationFlag Possible values are: - ACTIVATE: The event notification is activated. - DEACTIVATE: The event notification is deactivated and shall be muted. The available    event(s) shall be stored. - RETRIEVAL: The event notification shall be sent to the NF service consumer(s),   after that, is muted again.  
type NotificationFlag struct {
	NotificationFlagAnyOf *NotificationFlagAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NotificationFlag) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NotificationFlagAnyOf
	err = json.Unmarshal(data, &dst.NotificationFlagAnyOf);
	if err == nil {
		jsonNotificationFlagAnyOf, _ := json.Marshal(dst.NotificationFlagAnyOf)
		if string(jsonNotificationFlagAnyOf) == "{}" { // empty struct
			dst.NotificationFlagAnyOf = nil
		} else {
			return nil // data stored in dst.NotificationFlagAnyOf, return on the first match
		}
	} else {
		dst.NotificationFlagAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NotificationFlag)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NotificationFlag) MarshalJSON() ([]byte, error) {
	if src.NotificationFlagAnyOf != nil {
		return json.Marshal(&src.NotificationFlagAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNotificationFlag struct {
	value *NotificationFlag
	isSet bool
}

func (v NullableNotificationFlag) Get() *NotificationFlag {
	return v.value
}

func (v *NullableNotificationFlag) Set(val *NotificationFlag) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationFlag) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationFlag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationFlag(val *NotificationFlag) *NullableNotificationFlag {
	return &NullableNotificationFlag{value: val, isSet: true}
}

func (v NullableNotificationFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationFlag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


