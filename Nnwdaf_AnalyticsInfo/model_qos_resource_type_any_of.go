/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// QosResourceTypeAnyOf the model 'QosResourceTypeAnyOf'
type QosResourceTypeAnyOf string

// List of QosResourceType_anyOf
const (
	NON_GBR QosResourceTypeAnyOf = "NON_GBR"
	NON_CRITICAL_GBR QosResourceTypeAnyOf = "NON_CRITICAL_GBR"
	CRITICAL_GBR QosResourceTypeAnyOf = "CRITICAL_GBR"
)

// All allowed values of QosResourceTypeAnyOf enum
var AllowedQosResourceTypeAnyOfEnumValues = []QosResourceTypeAnyOf{
	"NON_GBR",
	"NON_CRITICAL_GBR",
	"CRITICAL_GBR",
}

func (v *QosResourceTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QosResourceTypeAnyOf(value)
	for _, existing := range AllowedQosResourceTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QosResourceTypeAnyOf", value)
}

// NewQosResourceTypeAnyOfFromValue returns a pointer to a valid QosResourceTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQosResourceTypeAnyOfFromValue(v string) (*QosResourceTypeAnyOf, error) {
	ev := QosResourceTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QosResourceTypeAnyOf: valid values are %v", v, AllowedQosResourceTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QosResourceTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedQosResourceTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QosResourceType_anyOf value
func (v QosResourceTypeAnyOf) Ptr() *QosResourceTypeAnyOf {
	return &v
}

type NullableQosResourceTypeAnyOf struct {
	value *QosResourceTypeAnyOf
	isSet bool
}

func (v NullableQosResourceTypeAnyOf) Get() *QosResourceTypeAnyOf {
	return v.value
}

func (v *NullableQosResourceTypeAnyOf) Set(val *QosResourceTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableQosResourceTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableQosResourceTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosResourceTypeAnyOf(val *QosResourceTypeAnyOf) *NullableQosResourceTypeAnyOf {
	return &NullableQosResourceTypeAnyOf{value: val, isSet: true}
}

func (v NullableQosResourceTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosResourceTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

