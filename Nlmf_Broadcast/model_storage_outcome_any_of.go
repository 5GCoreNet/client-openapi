/*
LMF Broadcast

LMF Broadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nlmf_Broadcast

import (
	"encoding/json"
	"fmt"
)

// StorageOutcomeAnyOf the model 'StorageOutcomeAnyOf'
type StorageOutcomeAnyOf string

// List of StorageOutcome_anyOf
const (
	SUCCESSFUL StorageOutcomeAnyOf = "STORAGE_SUCCESSFUL"
	FAILED StorageOutcomeAnyOf = "STORAGE_FAILED"
)

// All allowed values of StorageOutcomeAnyOf enum
var AllowedStorageOutcomeAnyOfEnumValues = []StorageOutcomeAnyOf{
	"STORAGE_SUCCESSFUL",
	"STORAGE_FAILED",
}

func (v *StorageOutcomeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StorageOutcomeAnyOf(value)
	for _, existing := range AllowedStorageOutcomeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StorageOutcomeAnyOf", value)
}

// NewStorageOutcomeAnyOfFromValue returns a pointer to a valid StorageOutcomeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageOutcomeAnyOfFromValue(v string) (*StorageOutcomeAnyOf, error) {
	ev := StorageOutcomeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StorageOutcomeAnyOf: valid values are %v", v, AllowedStorageOutcomeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageOutcomeAnyOf) IsValid() bool {
	for _, existing := range AllowedStorageOutcomeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StorageOutcome_anyOf value
func (v StorageOutcomeAnyOf) Ptr() *StorageOutcomeAnyOf {
	return &v
}

type NullableStorageOutcomeAnyOf struct {
	value *StorageOutcomeAnyOf
	isSet bool
}

func (v NullableStorageOutcomeAnyOf) Get() *StorageOutcomeAnyOf {
	return v.value
}

func (v *NullableStorageOutcomeAnyOf) Set(val *StorageOutcomeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageOutcomeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageOutcomeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageOutcomeAnyOf(val *StorageOutcomeAnyOf) *NullableStorageOutcomeAnyOf {
	return &NullableStorageOutcomeAnyOf{value: val, isSet: true}
}

func (v NullableStorageOutcomeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageOutcomeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

