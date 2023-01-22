/*
3gpp-mbs-ud-ingest

API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MBSUserDataIngestSession

import (
	"encoding/json"
	"fmt"
)

// EventOneOf the model 'EventOneOf'
type EventOneOf string

// List of Event_oneOf
const (
	USER_DATA_ING_SESS_STARTING EventOneOf = "USER_DATA_ING_SESS_STARTING"
	USER_DATA_ING_SESS_STARTED EventOneOf = "USER_DATA_ING_SESS_STARTED"
	USER_DATA_ING_SESS_TERMINATED EventOneOf = "USER_DATA_ING_SESS_TERMINATED"
	DIST_SESS_STARTING EventOneOf = "DIST_SESS_STARTING"
	DIST_SESS_STARTED EventOneOf = "DIST_SESS_STARTED"
	DIST_SESS_TERMINATED EventOneOf = "DIST_SESS_TERMINATED"
	DIST_SESS_SERV_MNGT_FAILURE EventOneOf = "DIST_SESS_SERV_MNGT_FAILURE"
	DIST_SESS_POL_CRTL_FAILURE EventOneOf = "DIST_SESS_POL_CRTL_FAILURE"
	DATA_INGEST_FAILURE EventOneOf = "DATA_INGEST_FAILURE"
	DELIVERY_STARTED EventOneOf = "DELIVERY_STARTED"
	SESSION_TERMINATED EventOneOf = "SESSION_TERMINATED"
)

// All allowed values of EventOneOf enum
var AllowedEventOneOfEnumValues = []EventOneOf{
	"USER_DATA_ING_SESS_STARTING",
	"USER_DATA_ING_SESS_STARTED",
	"USER_DATA_ING_SESS_TERMINATED",
	"DIST_SESS_STARTING",
	"DIST_SESS_STARTED",
	"DIST_SESS_TERMINATED",
	"DIST_SESS_SERV_MNGT_FAILURE",
	"DIST_SESS_POL_CRTL_FAILURE",
	"DATA_INGEST_FAILURE",
	"DELIVERY_STARTED",
	"SESSION_TERMINATED",
}

func (v *EventOneOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventOneOf(value)
	for _, existing := range AllowedEventOneOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventOneOf", value)
}

// NewEventOneOfFromValue returns a pointer to a valid EventOneOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventOneOfFromValue(v string) (*EventOneOf, error) {
	ev := EventOneOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventOneOf: valid values are %v", v, AllowedEventOneOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventOneOf) IsValid() bool {
	for _, existing := range AllowedEventOneOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Event_oneOf value
func (v EventOneOf) Ptr() *EventOneOf {
	return &v
}

type NullableEventOneOf struct {
	value *EventOneOf
	isSet bool
}

func (v NullableEventOneOf) Get() *EventOneOf {
	return v.value
}

func (v *NullableEventOneOf) Set(val *EventOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEventOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEventOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventOneOf(val *EventOneOf) *NullableEventOneOf {
	return &NullableEventOneOf{value: val, isSet: true}
}

func (v NullableEventOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

