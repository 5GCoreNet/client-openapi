/*
Generic NRM

OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GenericNrm

import (
	"encoding/json"
	"fmt"
)

// MeasurementPeriodUMTSType See details in 3GPP TS 32.422 clause 5.10.22.
type MeasurementPeriodUMTSType string

// List of measurementPeriodUMTS-Type
const (
	_1000MS MeasurementPeriodUMTSType = "1000ms"
	_2000MS MeasurementPeriodUMTSType = "2000ms"
	_3000MS MeasurementPeriodUMTSType = "3000ms"
	_4000MS MeasurementPeriodUMTSType = "4000ms"
	_6000MS MeasurementPeriodUMTSType = "6000ms"
	_8000MS MeasurementPeriodUMTSType = "8000ms"
	_12000MS MeasurementPeriodUMTSType = "12000ms"
	_16000MS MeasurementPeriodUMTSType = "16000ms"
	_20000MS MeasurementPeriodUMTSType = "20000ms"
	_24000MS MeasurementPeriodUMTSType = "24000ms"
	_28000MS MeasurementPeriodUMTSType = "28000ms"
	_32000MS MeasurementPeriodUMTSType = "32000ms"
	_64000MS MeasurementPeriodUMTSType = "64000ms"
)

// All allowed values of MeasurementPeriodUMTSType enum
var AllowedMeasurementPeriodUMTSTypeEnumValues = []MeasurementPeriodUMTSType{
	"1000ms",
	"2000ms",
	"3000ms",
	"4000ms",
	"6000ms",
	"8000ms",
	"12000ms",
	"16000ms",
	"20000ms",
	"24000ms",
	"28000ms",
	"32000ms",
	"64000ms",
}

func (v *MeasurementPeriodUMTSType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MeasurementPeriodUMTSType(value)
	for _, existing := range AllowedMeasurementPeriodUMTSTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MeasurementPeriodUMTSType", value)
}

// NewMeasurementPeriodUMTSTypeFromValue returns a pointer to a valid MeasurementPeriodUMTSType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMeasurementPeriodUMTSTypeFromValue(v string) (*MeasurementPeriodUMTSType, error) {
	ev := MeasurementPeriodUMTSType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MeasurementPeriodUMTSType: valid values are %v", v, AllowedMeasurementPeriodUMTSTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MeasurementPeriodUMTSType) IsValid() bool {
	for _, existing := range AllowedMeasurementPeriodUMTSTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to measurementPeriodUMTS-Type value
func (v MeasurementPeriodUMTSType) Ptr() *MeasurementPeriodUMTSType {
	return &v
}

type NullableMeasurementPeriodUMTSType struct {
	value *MeasurementPeriodUMTSType
	isSet bool
}

func (v NullableMeasurementPeriodUMTSType) Get() *MeasurementPeriodUMTSType {
	return v.value
}

func (v *NullableMeasurementPeriodUMTSType) Set(val *MeasurementPeriodUMTSType) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementPeriodUMTSType) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementPeriodUMTSType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementPeriodUMTSType(val *MeasurementPeriodUMTSType) *NullableMeasurementPeriodUMTSType {
	return &NullableMeasurementPeriodUMTSType{value: val, isSet: true}
}

func (v NullableMeasurementPeriodUMTSType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementPeriodUMTSType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

