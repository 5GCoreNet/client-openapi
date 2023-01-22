/*
MSGG_N3GDelivery

API for MSGG N3G Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MSGG_N3GDelivery

import (
	"encoding/json"
	"fmt"
)

// ReportDeliveryStatusAnyOf the model 'ReportDeliveryStatusAnyOf'
type ReportDeliveryStatusAnyOf string

// List of ReportDeliveryStatus_anyOf
const (
	SUCCESS ReportDeliveryStatusAnyOf = "REPT_DELY_SUCCESS"
	FAILED ReportDeliveryStatusAnyOf = "REPT_DELY_FAILED"
)

// All allowed values of ReportDeliveryStatusAnyOf enum
var AllowedReportDeliveryStatusAnyOfEnumValues = []ReportDeliveryStatusAnyOf{
	"REPT_DELY_SUCCESS",
	"REPT_DELY_FAILED",
}

func (v *ReportDeliveryStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportDeliveryStatusAnyOf(value)
	for _, existing := range AllowedReportDeliveryStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportDeliveryStatusAnyOf", value)
}

// NewReportDeliveryStatusAnyOfFromValue returns a pointer to a valid ReportDeliveryStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportDeliveryStatusAnyOfFromValue(v string) (*ReportDeliveryStatusAnyOf, error) {
	ev := ReportDeliveryStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportDeliveryStatusAnyOf: valid values are %v", v, AllowedReportDeliveryStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportDeliveryStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedReportDeliveryStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportDeliveryStatus_anyOf value
func (v ReportDeliveryStatusAnyOf) Ptr() *ReportDeliveryStatusAnyOf {
	return &v
}

type NullableReportDeliveryStatusAnyOf struct {
	value *ReportDeliveryStatusAnyOf
	isSet bool
}

func (v NullableReportDeliveryStatusAnyOf) Get() *ReportDeliveryStatusAnyOf {
	return v.value
}

func (v *NullableReportDeliveryStatusAnyOf) Set(val *ReportDeliveryStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReportDeliveryStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReportDeliveryStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportDeliveryStatusAnyOf(val *ReportDeliveryStatusAnyOf) *NullableReportDeliveryStatusAnyOf {
	return &NullableReportDeliveryStatusAnyOf{value: val, isSet: true}
}

func (v NullableReportDeliveryStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportDeliveryStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

