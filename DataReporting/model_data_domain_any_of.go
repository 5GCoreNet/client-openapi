/*
3gpp-data-reporting

API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package DataReporting

import (
	"encoding/json"
	"fmt"
)

// DataDomainAnyOf the model 'DataDomainAnyOf'
type DataDomainAnyOf string

// List of DataDomain_anyOf
const (
	SERVICE_EXPERIENCE DataDomainAnyOf = "SERVICE_EXPERIENCE"
	LOCATION DataDomainAnyOf = "LOCATION"
	COMMUNICATION DataDomainAnyOf = "COMMUNICATION"
	PERFORMANCE DataDomainAnyOf = "PERFORMANCE"
	APPLICATION_SPECIFIC DataDomainAnyOf = "APPLICATION_SPECIFIC"
	MS_ACCESS_ACTIVITY DataDomainAnyOf = "MS_ACCESS_ACTIVITY"
	PLANNED_TRIPS DataDomainAnyOf = "PLANNED_TRIPS"
)

// All allowed values of DataDomainAnyOf enum
var AllowedDataDomainAnyOfEnumValues = []DataDomainAnyOf{
	"SERVICE_EXPERIENCE",
	"LOCATION",
	"COMMUNICATION",
	"PERFORMANCE",
	"APPLICATION_SPECIFIC",
	"MS_ACCESS_ACTIVITY",
	"PLANNED_TRIPS",
}

func (v *DataDomainAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataDomainAnyOf(value)
	for _, existing := range AllowedDataDomainAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataDomainAnyOf", value)
}

// NewDataDomainAnyOfFromValue returns a pointer to a valid DataDomainAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataDomainAnyOfFromValue(v string) (*DataDomainAnyOf, error) {
	ev := DataDomainAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataDomainAnyOf: valid values are %v", v, AllowedDataDomainAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataDomainAnyOf) IsValid() bool {
	for _, existing := range AllowedDataDomainAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataDomain_anyOf value
func (v DataDomainAnyOf) Ptr() *DataDomainAnyOf {
	return &v
}

type NullableDataDomainAnyOf struct {
	value *DataDomainAnyOf
	isSet bool
}

func (v NullableDataDomainAnyOf) Get() *DataDomainAnyOf {
	return v.value
}

func (v *NullableDataDomainAnyOf) Set(val *DataDomainAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDataDomainAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDataDomainAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataDomainAnyOf(val *DataDomainAnyOf) *NullableDataDomainAnyOf {
	return &NullableDataDomainAnyOf{value: val, isSet: true}
}

func (v NullableDataDomainAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataDomainAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

