/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"fmt"
)

// LoggingDurationType See details in 3GPP TS 32.422 clause 5.10.9.
type LoggingDurationType string

// List of loggingDuration-Type
const (
	_600S LoggingDurationType = "600s"
	_1200S LoggingDurationType = "1200s"
	_2400S LoggingDurationType = "2400s"
	_3600S LoggingDurationType = "3600s"
	_5400S LoggingDurationType = "5400s"
	_7200S LoggingDurationType = "7200s"
)

// All allowed values of LoggingDurationType enum
var AllowedLoggingDurationTypeEnumValues = []LoggingDurationType{
	"600s",
	"1200s",
	"2400s",
	"3600s",
	"5400s",
	"7200s",
}

func (v *LoggingDurationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LoggingDurationType(value)
	for _, existing := range AllowedLoggingDurationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LoggingDurationType", value)
}

// NewLoggingDurationTypeFromValue returns a pointer to a valid LoggingDurationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoggingDurationTypeFromValue(v string) (*LoggingDurationType, error) {
	ev := LoggingDurationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LoggingDurationType: valid values are %v", v, AllowedLoggingDurationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoggingDurationType) IsValid() bool {
	for _, existing := range AllowedLoggingDurationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to loggingDuration-Type value
func (v LoggingDurationType) Ptr() *LoggingDurationType {
	return &v
}

type NullableLoggingDurationType struct {
	value *LoggingDurationType
	isSet bool
}

func (v NullableLoggingDurationType) Get() *LoggingDurationType {
	return v.value
}

func (v *NullableLoggingDurationType) Set(val *LoggingDurationType) {
	v.value = val
	v.isSet = true
}

func (v NullableLoggingDurationType) IsSet() bool {
	return v.isSet
}

func (v *NullableLoggingDurationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoggingDurationType(val *LoggingDurationType) *NullableLoggingDurationType {
	return &NullableLoggingDurationType{value: val, isSet: true}
}

func (v NullableLoggingDurationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoggingDurationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

