/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// AfEventAnyOf the model 'AfEventAnyOf'
type AfEventAnyOf string

// List of AfEvent_anyOf
const (
	ACCESS_TYPE_CHANGE AfEventAnyOf = "ACCESS_TYPE_CHANGE"
	ANI_REPORT AfEventAnyOf = "ANI_REPORT"
	APP_DETECTION AfEventAnyOf = "APP_DETECTION"
	CHARGING_CORRELATION AfEventAnyOf = "CHARGING_CORRELATION"
	EPS_FALLBACK AfEventAnyOf = "EPS_FALLBACK"
	FAILED_QOS_UPDATE AfEventAnyOf = "FAILED_QOS_UPDATE"
	FAILED_RESOURCES_ALLOCATION AfEventAnyOf = "FAILED_RESOURCES_ALLOCATION"
	OUT_OF_CREDIT AfEventAnyOf = "OUT_OF_CREDIT"
	PDU_SESSION_STATUS AfEventAnyOf = "PDU_SESSION_STATUS"
	PLMN_CHG AfEventAnyOf = "PLMN_CHG"
	QOS_MONITORING AfEventAnyOf = "QOS_MONITORING"
	QOS_NOTIF AfEventAnyOf = "QOS_NOTIF"
	RAN_NAS_CAUSE AfEventAnyOf = "RAN_NAS_CAUSE"
	REALLOCATION_OF_CREDIT AfEventAnyOf = "REALLOCATION_OF_CREDIT"
	SAT_CATEGORY_CHG AfEventAnyOf = "SAT_CATEGORY_CHG"
	SUCCESSFUL_QOS_UPDATE AfEventAnyOf = "SUCCESSFUL_QOS_UPDATE"
	SUCCESSFUL_RESOURCES_ALLOCATION AfEventAnyOf = "SUCCESSFUL_RESOURCES_ALLOCATION"
	TSN_BRIDGE_INFO AfEventAnyOf = "TSN_BRIDGE_INFO"
	UP_PATH_CHG_FAILURE AfEventAnyOf = "UP_PATH_CHG_FAILURE"
	USAGE_REPORT AfEventAnyOf = "USAGE_REPORT"
)

// All allowed values of AfEventAnyOf enum
var AllowedAfEventAnyOfEnumValues = []AfEventAnyOf{
	"ACCESS_TYPE_CHANGE",
	"ANI_REPORT",
	"APP_DETECTION",
	"CHARGING_CORRELATION",
	"EPS_FALLBACK",
	"FAILED_QOS_UPDATE",
	"FAILED_RESOURCES_ALLOCATION",
	"OUT_OF_CREDIT",
	"PDU_SESSION_STATUS",
	"PLMN_CHG",
	"QOS_MONITORING",
	"QOS_NOTIF",
	"RAN_NAS_CAUSE",
	"REALLOCATION_OF_CREDIT",
	"SAT_CATEGORY_CHG",
	"SUCCESSFUL_QOS_UPDATE",
	"SUCCESSFUL_RESOURCES_ALLOCATION",
	"TSN_BRIDGE_INFO",
	"UP_PATH_CHG_FAILURE",
	"USAGE_REPORT",
}

func (v *AfEventAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AfEventAnyOf(value)
	for _, existing := range AllowedAfEventAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AfEventAnyOf", value)
}

// NewAfEventAnyOfFromValue returns a pointer to a valid AfEventAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAfEventAnyOfFromValue(v string) (*AfEventAnyOf, error) {
	ev := AfEventAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AfEventAnyOf: valid values are %v", v, AllowedAfEventAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AfEventAnyOf) IsValid() bool {
	for _, existing := range AllowedAfEventAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AfEvent_anyOf value
func (v AfEventAnyOf) Ptr() *AfEventAnyOf {
	return &v
}

type NullableAfEventAnyOf struct {
	value *AfEventAnyOf
	isSet bool
}

func (v NullableAfEventAnyOf) Get() *AfEventAnyOf {
	return v.value
}

func (v *NullableAfEventAnyOf) Set(val *AfEventAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAfEventAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAfEventAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfEventAnyOf(val *AfEventAnyOf) *NullableAfEventAnyOf {
	return &NullableAfEventAnyOf{value: val, isSet: true}
}

func (v NullableAfEventAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfEventAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

