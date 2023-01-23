/*
LMF Broadcast

LMF Broadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nlmf_Broadcast

import (
	"encoding/json"
	"fmt"
)

// DataAvailabilityAnyOf the model 'DataAvailabilityAnyOf'
type DataAvailabilityAnyOf string

// List of DataAvailability_anyOf
const (
	AVAILABLE DataAvailabilityAnyOf = "CIPHERING_KEY_DATA_AVAILABLE"
	NOT_AVAILABLE DataAvailabilityAnyOf = "CIPHERING_KEY_DATA_NOT_AVAILABLE"
)

// All allowed values of DataAvailabilityAnyOf enum
var AllowedDataAvailabilityAnyOfEnumValues = []DataAvailabilityAnyOf{
	"CIPHERING_KEY_DATA_AVAILABLE",
	"CIPHERING_KEY_DATA_NOT_AVAILABLE",
}

func (v *DataAvailabilityAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataAvailabilityAnyOf(value)
	for _, existing := range AllowedDataAvailabilityAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataAvailabilityAnyOf", value)
}

// NewDataAvailabilityAnyOfFromValue returns a pointer to a valid DataAvailabilityAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataAvailabilityAnyOfFromValue(v string) (*DataAvailabilityAnyOf, error) {
	ev := DataAvailabilityAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataAvailabilityAnyOf: valid values are %v", v, AllowedDataAvailabilityAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataAvailabilityAnyOf) IsValid() bool {
	for _, existing := range AllowedDataAvailabilityAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataAvailability_anyOf value
func (v DataAvailabilityAnyOf) Ptr() *DataAvailabilityAnyOf {
	return &v
}

type NullableDataAvailabilityAnyOf struct {
	value *DataAvailabilityAnyOf
	isSet bool
}

func (v NullableDataAvailabilityAnyOf) Get() *DataAvailabilityAnyOf {
	return v.value
}

func (v *NullableDataAvailabilityAnyOf) Set(val *DataAvailabilityAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDataAvailabilityAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDataAvailabilityAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataAvailabilityAnyOf(val *DataAvailabilityAnyOf) *NullableDataAvailabilityAnyOf {
	return &NullableDataAvailabilityAnyOf{value: val, isSet: true}
}

func (v NullableDataAvailabilityAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataAvailabilityAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

