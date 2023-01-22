/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MdaNrm

import (
	"encoding/json"
	"fmt"
)

// UsageState the model 'UsageState'
type UsageState string

// List of UsageState
const (
	IDEL UsageState = "IDEL"
	ACTIVE UsageState = "ACTIVE"
	BUSY UsageState = "BUSY"
)

// All allowed values of UsageState enum
var AllowedUsageStateEnumValues = []UsageState{
	"IDEL",
	"ACTIVE",
	"BUSY",
}

func (v *UsageState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UsageState(value)
	for _, existing := range AllowedUsageStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UsageState", value)
}

// NewUsageStateFromValue returns a pointer to a valid UsageState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUsageStateFromValue(v string) (*UsageState, error) {
	ev := UsageState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UsageState: valid values are %v", v, AllowedUsageStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UsageState) IsValid() bool {
	for _, existing := range AllowedUsageStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageState value
func (v UsageState) Ptr() *UsageState {
	return &v
}

type NullableUsageState struct {
	value *UsageState
	isSet bool
}

func (v NullableUsageState) Get() *UsageState {
	return v.value
}

func (v *NullableUsageState) Set(val *UsageState) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageState) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageState(val *UsageState) *NullableUsageState {
	return &NullableUsageState{value: val, isSet: true}
}

func (v NullableUsageState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

