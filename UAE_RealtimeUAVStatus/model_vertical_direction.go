/*
UAE Server Real-time UAV Status Service

UAE Server Real-time UAV Status Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package UAE_RealtimeUAVStatus

import (
	"encoding/json"
	"fmt"
)

// VerticalDirection Indicates direction of vertical speed.
type VerticalDirection string

// List of VerticalDirection
const (
	UPWARD VerticalDirection = "UPWARD"
	DOWNWARD VerticalDirection = "DOWNWARD"
)

// All allowed values of VerticalDirection enum
var AllowedVerticalDirectionEnumValues = []VerticalDirection{
	"UPWARD",
	"DOWNWARD",
}

func (v *VerticalDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VerticalDirection(value)
	for _, existing := range AllowedVerticalDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VerticalDirection", value)
}

// NewVerticalDirectionFromValue returns a pointer to a valid VerticalDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVerticalDirectionFromValue(v string) (*VerticalDirection, error) {
	ev := VerticalDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VerticalDirection: valid values are %v", v, AllowedVerticalDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VerticalDirection) IsValid() bool {
	for _, existing := range AllowedVerticalDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VerticalDirection value
func (v VerticalDirection) Ptr() *VerticalDirection {
	return &v
}

type NullableVerticalDirection struct {
	value *VerticalDirection
	isSet bool
}

func (v NullableVerticalDirection) Get() *VerticalDirection {
	return v.value
}

func (v *NullableVerticalDirection) Set(val *VerticalDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableVerticalDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableVerticalDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerticalDirection(val *VerticalDirection) *NullableVerticalDirection {
	return &NullableVerticalDirection{value: val, isSet: true}
}

func (v NullableVerticalDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerticalDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

