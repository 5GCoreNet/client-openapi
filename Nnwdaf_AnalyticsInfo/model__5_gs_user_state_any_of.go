/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// 5GsUserStateAnyOf the model '5GsUserStateAnyOf'
type 5GsUserStateAnyOf string

// List of _5GsUserState_anyOf
const (
	DEREGISTERED 5GsUserStateAnyOf = "DEREGISTERED"
	CONNECTED_NOT_REACHABLE_FOR_PAGING 5GsUserStateAnyOf = "CONNECTED_NOT_REACHABLE_FOR_PAGING"
	CONNECTED_REACHABLE_FOR_PAGING 5GsUserStateAnyOf = "CONNECTED_REACHABLE_FOR_PAGING"
	NOT_PROVIDED_FROM_AMF 5GsUserStateAnyOf = "NOT_PROVIDED_FROM_AMF"
)

// All allowed values of 5GsUserStateAnyOf enum
var Allowed5GsUserStateAnyOfEnumValues = []5GsUserStateAnyOf{
	"DEREGISTERED",
	"CONNECTED_NOT_REACHABLE_FOR_PAGING",
	"CONNECTED_REACHABLE_FOR_PAGING",
	"NOT_PROVIDED_FROM_AMF",
}

func (v *5GsUserStateAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := 5GsUserStateAnyOf(value)
	for _, existing := range Allowed5GsUserStateAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid 5GsUserStateAnyOf", value)
}

// New5GsUserStateAnyOfFromValue returns a pointer to a valid 5GsUserStateAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func New5GsUserStateAnyOfFromValue(v string) (*5GsUserStateAnyOf, error) {
	ev := 5GsUserStateAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for 5GsUserStateAnyOf: valid values are %v", v, Allowed5GsUserStateAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v 5GsUserStateAnyOf) IsValid() bool {
	for _, existing := range Allowed5GsUserStateAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to _5GsUserState_anyOf value
func (v 5GsUserStateAnyOf) Ptr() *5GsUserStateAnyOf {
	return &v
}

type Nullable5GsUserStateAnyOf struct {
	value *5GsUserStateAnyOf
	isSet bool
}

func (v Nullable5GsUserStateAnyOf) Get() *5GsUserStateAnyOf {
	return v.value
}

func (v *Nullable5GsUserStateAnyOf) Set(val *5GsUserStateAnyOf) {
	v.value = val
	v.isSet = true
}

func (v Nullable5GsUserStateAnyOf) IsSet() bool {
	return v.isSet
}

func (v *Nullable5GsUserStateAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullable5GsUserStateAnyOf(val *5GsUserStateAnyOf) *Nullable5GsUserStateAnyOf {
	return &Nullable5GsUserStateAnyOf{value: val, isSet: true}
}

func (v Nullable5GsUserStateAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *Nullable5GsUserStateAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

