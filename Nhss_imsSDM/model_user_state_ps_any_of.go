/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// UserStatePsAnyOf the model 'UserStatePsAnyOf'
type UserStatePsAnyOf string

// List of UserStatePs_anyOf
const (
	DETACHED UserStatePsAnyOf = "DETACHED"
	ATTACHED_NOT_REACHABLE_FOR_PAGING UserStatePsAnyOf = "ATTACHED_NOT_REACHABLE_FOR_PAGING"
	ATTACHED_REACHABLE_FOR_PAGING UserStatePsAnyOf = "ATTACHED_REACHABLE_FOR_PAGING"
	CONNECTED_NOT_REACHABLE_FOR_PAGING UserStatePsAnyOf = "CONNECTED_NOT_REACHABLE_FOR_PAGING"
	CONNECTED_REACHABLE_FOR_PAGING UserStatePsAnyOf = "CONNECTED_REACHABLE_FOR_PAGING"
	NOT_PROVIDED_FROM_SGSN_OR_MME_OR_AMF UserStatePsAnyOf = "NOT_PROVIDED_FROM_SGSN_OR_MME_OR_AMF"
	NETWORK_DETERMINED_NOT_REACHABLE UserStatePsAnyOf = "NETWORK_DETERMINED_NOT_REACHABLE"
)

// All allowed values of UserStatePsAnyOf enum
var AllowedUserStatePsAnyOfEnumValues = []UserStatePsAnyOf{
	"DETACHED",
	"ATTACHED_NOT_REACHABLE_FOR_PAGING",
	"ATTACHED_REACHABLE_FOR_PAGING",
	"CONNECTED_NOT_REACHABLE_FOR_PAGING",
	"CONNECTED_REACHABLE_FOR_PAGING",
	"NOT_PROVIDED_FROM_SGSN_OR_MME_OR_AMF",
	"NETWORK_DETERMINED_NOT_REACHABLE",
}

func (v *UserStatePsAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserStatePsAnyOf(value)
	for _, existing := range AllowedUserStatePsAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserStatePsAnyOf", value)
}

// NewUserStatePsAnyOfFromValue returns a pointer to a valid UserStatePsAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserStatePsAnyOfFromValue(v string) (*UserStatePsAnyOf, error) {
	ev := UserStatePsAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserStatePsAnyOf: valid values are %v", v, AllowedUserStatePsAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserStatePsAnyOf) IsValid() bool {
	for _, existing := range AllowedUserStatePsAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserStatePs_anyOf value
func (v UserStatePsAnyOf) Ptr() *UserStatePsAnyOf {
	return &v
}

type NullableUserStatePsAnyOf struct {
	value *UserStatePsAnyOf
	isSet bool
}

func (v NullableUserStatePsAnyOf) Get() *UserStatePsAnyOf {
	return v.value
}

func (v *NullableUserStatePsAnyOf) Set(val *UserStatePsAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStatePsAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStatePsAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStatePsAnyOf(val *UserStatePsAnyOf) *NullableUserStatePsAnyOf {
	return &NullableUserStatePsAnyOf{value: val, isSet: true}
}

func (v NullableUserStatePsAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStatePsAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

