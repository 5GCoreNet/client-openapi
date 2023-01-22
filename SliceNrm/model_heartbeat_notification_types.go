/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SliceNrm

import (
	"encoding/json"
	"fmt"
)

// HeartbeatNotificationTypes the model 'HeartbeatNotificationTypes'
type HeartbeatNotificationTypes string

// List of HeartbeatNotificationTypes
const (
	NOTIFY_HEARTBEAT HeartbeatNotificationTypes = "notifyHeartbeat"
)

// All allowed values of HeartbeatNotificationTypes enum
var AllowedHeartbeatNotificationTypesEnumValues = []HeartbeatNotificationTypes{
	"notifyHeartbeat",
}

func (v *HeartbeatNotificationTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HeartbeatNotificationTypes(value)
	for _, existing := range AllowedHeartbeatNotificationTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HeartbeatNotificationTypes", value)
}

// NewHeartbeatNotificationTypesFromValue returns a pointer to a valid HeartbeatNotificationTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHeartbeatNotificationTypesFromValue(v string) (*HeartbeatNotificationTypes, error) {
	ev := HeartbeatNotificationTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HeartbeatNotificationTypes: valid values are %v", v, AllowedHeartbeatNotificationTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HeartbeatNotificationTypes) IsValid() bool {
	for _, existing := range AllowedHeartbeatNotificationTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeartbeatNotificationTypes value
func (v HeartbeatNotificationTypes) Ptr() *HeartbeatNotificationTypes {
	return &v
}

type NullableHeartbeatNotificationTypes struct {
	value *HeartbeatNotificationTypes
	isSet bool
}

func (v NullableHeartbeatNotificationTypes) Get() *HeartbeatNotificationTypes {
	return v.value
}

func (v *NullableHeartbeatNotificationTypes) Set(val *HeartbeatNotificationTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableHeartbeatNotificationTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableHeartbeatNotificationTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeartbeatNotificationTypes(val *HeartbeatNotificationTypes) *NullableHeartbeatNotificationTypes {
	return &NullableHeartbeatNotificationTypes{value: val, isSet: true}
}

func (v NullableHeartbeatNotificationTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeartbeatNotificationTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

