/*
File Data Reporting MnS

OAS 3.0.1 definition of the File Data Reporting MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package FileDataReportingMnS

import (
	"encoding/json"
	"fmt"
)

// FileDataType the model 'FileDataType'
type FileDataType string

// List of FileDataType
const (
	PERFORMANCE FileDataType = "Performance"
	TRACE FileDataType = "Trace"
	ANATYTICS FileDataType = "Anatytics"
	PROPRIETARY FileDataType = "Proprietary"
)

// All allowed values of FileDataType enum
var AllowedFileDataTypeEnumValues = []FileDataType{
	"Performance",
	"Trace",
	"Anatytics",
	"Proprietary",
}

func (v *FileDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileDataType(value)
	for _, existing := range AllowedFileDataTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileDataType", value)
}

// NewFileDataTypeFromValue returns a pointer to a valid FileDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileDataTypeFromValue(v string) (*FileDataType, error) {
	ev := FileDataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileDataType: valid values are %v", v, AllowedFileDataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileDataType) IsValid() bool {
	for _, existing := range AllowedFileDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileDataType value
func (v FileDataType) Ptr() *FileDataType {
	return &v
}

type NullableFileDataType struct {
	value *FileDataType
	isSet bool
}

func (v NullableFileDataType) Get() *FileDataType {
	return v.value
}

func (v *NullableFileDataType) Set(val *FileDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDataType(val *FileDataType) *NullableFileDataType {
	return &NullableFileDataType{value: val, isSet: true}
}

func (v NullableFileDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

