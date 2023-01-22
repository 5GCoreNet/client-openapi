/*
CAPIF_API_Invoker_Management_API

API for API invoker management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CAPIF_API_Invoker_Management_API

import (
	"encoding/json"
	"fmt"
)

// DataFormatAnyOf the model 'DataFormatAnyOf'
type DataFormatAnyOf string

// List of DataFormat_anyOf
const (
	JSON DataFormatAnyOf = "JSON"
)

// All allowed values of DataFormatAnyOf enum
var AllowedDataFormatAnyOfEnumValues = []DataFormatAnyOf{
	"JSON",
}

func (v *DataFormatAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataFormatAnyOf(value)
	for _, existing := range AllowedDataFormatAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataFormatAnyOf", value)
}

// NewDataFormatAnyOfFromValue returns a pointer to a valid DataFormatAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataFormatAnyOfFromValue(v string) (*DataFormatAnyOf, error) {
	ev := DataFormatAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataFormatAnyOf: valid values are %v", v, AllowedDataFormatAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataFormatAnyOf) IsValid() bool {
	for _, existing := range AllowedDataFormatAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataFormat_anyOf value
func (v DataFormatAnyOf) Ptr() *DataFormatAnyOf {
	return &v
}

type NullableDataFormatAnyOf struct {
	value *DataFormatAnyOf
	isSet bool
}

func (v NullableDataFormatAnyOf) Get() *DataFormatAnyOf {
	return v.value
}

func (v *NullableDataFormatAnyOf) Set(val *DataFormatAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDataFormatAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDataFormatAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataFormatAnyOf(val *DataFormatAnyOf) *NullableDataFormatAnyOf {
	return &NullableDataFormatAnyOf{value: val, isSet: true}
}

func (v NullableDataFormatAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataFormatAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

