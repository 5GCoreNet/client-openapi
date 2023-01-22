/*
TS 28.550 Performance Measurement Job Control Service

OAS 3.0.1 specification of the Performance Measurement Job Control Service @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 16.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package PerfMeasJobCtrlMnS

import (
	"encoding/json"
	"fmt"
)

// DayOfWeekType the model 'DayOfWeekType'
type DayOfWeekType string

// List of dayOfWeek-Type
const (
	MONDAY DayOfWeekType = "Monday"
	TUESDAY DayOfWeekType = "Tuesday"
	WEDNESDAY DayOfWeekType = "Wednesday"
	THURSDAY DayOfWeekType = "Thursday"
	FRIDAY DayOfWeekType = "Friday"
	SATURDAY DayOfWeekType = "Saturday"
	SUNDAY DayOfWeekType = "Sunday"
)

// All allowed values of DayOfWeekType enum
var AllowedDayOfWeekTypeEnumValues = []DayOfWeekType{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

func (v *DayOfWeekType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DayOfWeekType(value)
	for _, existing := range AllowedDayOfWeekTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DayOfWeekType", value)
}

// NewDayOfWeekTypeFromValue returns a pointer to a valid DayOfWeekType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDayOfWeekTypeFromValue(v string) (*DayOfWeekType, error) {
	ev := DayOfWeekType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DayOfWeekType: valid values are %v", v, AllowedDayOfWeekTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DayOfWeekType) IsValid() bool {
	for _, existing := range AllowedDayOfWeekTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dayOfWeek-Type value
func (v DayOfWeekType) Ptr() *DayOfWeekType {
	return &v
}

type NullableDayOfWeekType struct {
	value *DayOfWeekType
	isSet bool
}

func (v NullableDayOfWeekType) Get() *DayOfWeekType {
	return v.value
}

func (v *NullableDayOfWeekType) Set(val *DayOfWeekType) {
	v.value = val
	v.isSet = true
}

func (v NullableDayOfWeekType) IsSet() bool {
	return v.isSet
}

func (v *NullableDayOfWeekType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDayOfWeekType(val *DayOfWeekType) *NullableDayOfWeekType {
	return &NullableDayOfWeekType{value: val, isSet: true}
}

func (v NullableDayOfWeekType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDayOfWeekType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
