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

// FileNotificationTypes the model 'FileNotificationTypes'
type FileNotificationTypes string

// List of FileNotificationTypes
const (
	NOTIFY_FILE_READY FileNotificationTypes = "notifyFileReady"
	NOTIFY_FILE_PREPARATION_ERROR FileNotificationTypes = "notifyFilePreparationError"
)

// All allowed values of FileNotificationTypes enum
var AllowedFileNotificationTypesEnumValues = []FileNotificationTypes{
	"notifyFileReady",
	"notifyFilePreparationError",
}

func (v *FileNotificationTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileNotificationTypes(value)
	for _, existing := range AllowedFileNotificationTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileNotificationTypes", value)
}

// NewFileNotificationTypesFromValue returns a pointer to a valid FileNotificationTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileNotificationTypesFromValue(v string) (*FileNotificationTypes, error) {
	ev := FileNotificationTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileNotificationTypes: valid values are %v", v, AllowedFileNotificationTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileNotificationTypes) IsValid() bool {
	for _, existing := range AllowedFileNotificationTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileNotificationTypes value
func (v FileNotificationTypes) Ptr() *FileNotificationTypes {
	return &v
}

type NullableFileNotificationTypes struct {
	value *FileNotificationTypes
	isSet bool
}

func (v NullableFileNotificationTypes) Get() *FileNotificationTypes {
	return v.value
}

func (v *NullableFileNotificationTypes) Set(val *FileNotificationTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableFileNotificationTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableFileNotificationTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileNotificationTypes(val *FileNotificationTypes) *NullableFileNotificationTypes {
	return &NullableFileNotificationTypes{value: val, isSet: true}
}

func (v NullableFileNotificationTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileNotificationTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

