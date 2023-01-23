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

// MeasurementPeriodLTEType See details in 3GPP TS 32.422 clause 5.10.23.
type MeasurementPeriodLTEType string

// List of measurementPeriodLTE-Type
const (
	_1024MS MeasurementPeriodLTEType = "1024ms"
	_2048MS MeasurementPeriodLTEType = "2048ms"
	_5120MS MeasurementPeriodLTEType = "5120ms"
	_10240MS MeasurementPeriodLTEType = "10240ms"
	_1MIN MeasurementPeriodLTEType = "1min"
)

// All allowed values of MeasurementPeriodLTEType enum
var AllowedMeasurementPeriodLTETypeEnumValues = []MeasurementPeriodLTEType{
	"1024ms",
	"2048ms",
	"5120ms",
	"10240ms",
	"1min",
}

func (v *MeasurementPeriodLTEType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MeasurementPeriodLTEType(value)
	for _, existing := range AllowedMeasurementPeriodLTETypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MeasurementPeriodLTEType", value)
}

// NewMeasurementPeriodLTETypeFromValue returns a pointer to a valid MeasurementPeriodLTEType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMeasurementPeriodLTETypeFromValue(v string) (*MeasurementPeriodLTEType, error) {
	ev := MeasurementPeriodLTEType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MeasurementPeriodLTEType: valid values are %v", v, AllowedMeasurementPeriodLTETypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MeasurementPeriodLTEType) IsValid() bool {
	for _, existing := range AllowedMeasurementPeriodLTETypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to measurementPeriodLTE-Type value
func (v MeasurementPeriodLTEType) Ptr() *MeasurementPeriodLTEType {
	return &v
}

type NullableMeasurementPeriodLTEType struct {
	value *MeasurementPeriodLTEType
	isSet bool
}

func (v NullableMeasurementPeriodLTEType) Get() *MeasurementPeriodLTEType {
	return v.value
}

func (v *NullableMeasurementPeriodLTEType) Set(val *MeasurementPeriodLTEType) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementPeriodLTEType) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementPeriodLTEType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementPeriodLTEType(val *MeasurementPeriodLTEType) *NullableMeasurementPeriodLTEType {
	return &NullableMeasurementPeriodLTEType{value: val, isSet: true}
}

func (v NullableMeasurementPeriodLTEType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementPeriodLTEType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

