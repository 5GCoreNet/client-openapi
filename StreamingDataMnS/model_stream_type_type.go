/*
TS 28.532 Streaming data reporting service

OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package StreamingDataMnS

import (
	"encoding/json"
	"fmt"
)

// StreamTypeType the model 'StreamTypeType'
type StreamTypeType string

// List of streamType-Type
const (
	TRACE StreamTypeType = "TRACE"
	PERFORMANCE StreamTypeType = "PERFORMANCE"
	ANALYTICS StreamTypeType = "ANALYTICS"
	PROPRIETARY StreamTypeType = "PROPRIETARY"
)

// All allowed values of StreamTypeType enum
var AllowedStreamTypeTypeEnumValues = []StreamTypeType{
	"TRACE",
	"PERFORMANCE",
	"ANALYTICS",
	"PROPRIETARY",
}

func (v *StreamTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StreamTypeType(value)
	for _, existing := range AllowedStreamTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StreamTypeType", value)
}

// NewStreamTypeTypeFromValue returns a pointer to a valid StreamTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStreamTypeTypeFromValue(v string) (*StreamTypeType, error) {
	ev := StreamTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StreamTypeType: valid values are %v", v, AllowedStreamTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StreamTypeType) IsValid() bool {
	for _, existing := range AllowedStreamTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to streamType-Type value
func (v StreamTypeType) Ptr() *StreamTypeType {
	return &v
}

type NullableStreamTypeType struct {
	value *StreamTypeType
	isSet bool
}

func (v NullableStreamTypeType) Get() *StreamTypeType {
	return v.value
}

func (v *NullableStreamTypeType) Set(val *StreamTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamTypeType(val *StreamTypeType) *NullableStreamTypeType {
	return &NullableStreamTypeType{value: val, isSet: true}
}

func (v NullableStreamTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

