/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// QosMonitoringReqAnyOf the model 'QosMonitoringReqAnyOf'
type QosMonitoringReqAnyOf string

// List of QosMonitoringReq_anyOf
const (
	UL QosMonitoringReqAnyOf = "UL"
	DL QosMonitoringReqAnyOf = "DL"
	BOTH QosMonitoringReqAnyOf = "BOTH"
	NONE QosMonitoringReqAnyOf = "NONE"
)

// All allowed values of QosMonitoringReqAnyOf enum
var AllowedQosMonitoringReqAnyOfEnumValues = []QosMonitoringReqAnyOf{
	"UL",
	"DL",
	"BOTH",
	"NONE",
}

func (v *QosMonitoringReqAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QosMonitoringReqAnyOf(value)
	for _, existing := range AllowedQosMonitoringReqAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QosMonitoringReqAnyOf", value)
}

// NewQosMonitoringReqAnyOfFromValue returns a pointer to a valid QosMonitoringReqAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQosMonitoringReqAnyOfFromValue(v string) (*QosMonitoringReqAnyOf, error) {
	ev := QosMonitoringReqAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QosMonitoringReqAnyOf: valid values are %v", v, AllowedQosMonitoringReqAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QosMonitoringReqAnyOf) IsValid() bool {
	for _, existing := range AllowedQosMonitoringReqAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QosMonitoringReq_anyOf value
func (v QosMonitoringReqAnyOf) Ptr() *QosMonitoringReqAnyOf {
	return &v
}

type NullableQosMonitoringReqAnyOf struct {
	value *QosMonitoringReqAnyOf
	isSet bool
}

func (v NullableQosMonitoringReqAnyOf) Get() *QosMonitoringReqAnyOf {
	return v.value
}

func (v *NullableQosMonitoringReqAnyOf) Set(val *QosMonitoringReqAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableQosMonitoringReqAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableQosMonitoringReqAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosMonitoringReqAnyOf(val *QosMonitoringReqAnyOf) *NullableQosMonitoringReqAnyOf {
	return &NullableQosMonitoringReqAnyOf{value: val, isSet: true}
}

func (v NullableQosMonitoringReqAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosMonitoringReqAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

