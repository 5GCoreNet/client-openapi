/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_OfflineOnlyCharging

import (
	"encoding/json"
	"fmt"
)

// NullValue JSON's null value.
type NullValue string

// List of NullValue
const (
	NULL NullValue = "null"
)

// All allowed values of NullValue enum
var AllowedNullValueEnumValues = []NullValue{
	"null",
}

func (v *NullValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NullValue(value)
	for _, existing := range AllowedNullValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NullValue", value)
}

// NewNullValueFromValue returns a pointer to a valid NullValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNullValueFromValue(v string) (*NullValue, error) {
	ev := NullValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NullValue: valid values are %v", v, AllowedNullValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NullValue) IsValid() bool {
	for _, existing := range AllowedNullValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NullValue value
func (v NullValue) Ptr() *NullValue {
	return &v
}

type NullableNullValue struct {
	value *NullValue
	isSet bool
}

func (v NullableNullValue) Get() *NullValue {
	return v.value
}

func (v *NullableNullValue) Set(val *NullValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNullValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNullValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNullValue(val *NullValue) *NullableNullValue {
	return &NullableNullValue{value: val, isSet: true}
}

func (v NullableNullValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNullValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

