/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// SessionFailoverAnyOf the model 'SessionFailoverAnyOf'
type SessionFailoverAnyOf string

// List of SessionFailover_anyOf
const (
	NOT_SUPPORTED SessionFailoverAnyOf = "FAILOVER_NOT_SUPPORTED"
	SUPPORTED SessionFailoverAnyOf = "FAILOVER_SUPPORTED"
)

// All allowed values of SessionFailoverAnyOf enum
var AllowedSessionFailoverAnyOfEnumValues = []SessionFailoverAnyOf{
	"FAILOVER_NOT_SUPPORTED",
	"FAILOVER_SUPPORTED",
}

func (v *SessionFailoverAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SessionFailoverAnyOf(value)
	for _, existing := range AllowedSessionFailoverAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SessionFailoverAnyOf", value)
}

// NewSessionFailoverAnyOfFromValue returns a pointer to a valid SessionFailoverAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSessionFailoverAnyOfFromValue(v string) (*SessionFailoverAnyOf, error) {
	ev := SessionFailoverAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SessionFailoverAnyOf: valid values are %v", v, AllowedSessionFailoverAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SessionFailoverAnyOf) IsValid() bool {
	for _, existing := range AllowedSessionFailoverAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SessionFailover_anyOf value
func (v SessionFailoverAnyOf) Ptr() *SessionFailoverAnyOf {
	return &v
}

type NullableSessionFailoverAnyOf struct {
	value *SessionFailoverAnyOf
	isSet bool
}

func (v NullableSessionFailoverAnyOf) Get() *SessionFailoverAnyOf {
	return v.value
}

func (v *NullableSessionFailoverAnyOf) Set(val *SessionFailoverAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionFailoverAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionFailoverAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionFailoverAnyOf(val *SessionFailoverAnyOf) *NullableSessionFailoverAnyOf {
	return &NullableSessionFailoverAnyOf{value: val, isSet: true}
}

func (v NullableSessionFailoverAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionFailoverAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

