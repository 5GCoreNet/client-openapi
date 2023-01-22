/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// SmfEventAnyOf the model 'SmfEventAnyOf'
type SmfEventAnyOf string

// List of SmfEvent_anyOf
const (
	AC_TY_CH SmfEventAnyOf = "AC_TY_CH"
	UP_PATH_CH SmfEventAnyOf = "UP_PATH_CH"
	PDU_SES_REL SmfEventAnyOf = "PDU_SES_REL"
	PLMN_CH SmfEventAnyOf = "PLMN_CH"
	UE_IP_CH SmfEventAnyOf = "UE_IP_CH"
	RAT_TY_CH SmfEventAnyOf = "RAT_TY_CH"
	DDDS SmfEventAnyOf = "DDDS"
	COMM_FAIL SmfEventAnyOf = "COMM_FAIL"
	PDU_SES_EST SmfEventAnyOf = "PDU_SES_EST"
	QFI_ALLOC SmfEventAnyOf = "QFI_ALLOC"
	QOS_MON SmfEventAnyOf = "QOS_MON"
	SMCC_EXP SmfEventAnyOf = "SMCC_EXP"
	DISPERSION SmfEventAnyOf = "DISPERSION"
	RED_TRANS_EXP SmfEventAnyOf = "RED_TRANS_EXP"
	WLAN_INFO SmfEventAnyOf = "WLAN_INFO"
	UPF_INFO SmfEventAnyOf = "UPF_INFO"
)

// All allowed values of SmfEventAnyOf enum
var AllowedSmfEventAnyOfEnumValues = []SmfEventAnyOf{
	"AC_TY_CH",
	"UP_PATH_CH",
	"PDU_SES_REL",
	"PLMN_CH",
	"UE_IP_CH",
	"RAT_TY_CH",
	"DDDS",
	"COMM_FAIL",
	"PDU_SES_EST",
	"QFI_ALLOC",
	"QOS_MON",
	"SMCC_EXP",
	"DISPERSION",
	"RED_TRANS_EXP",
	"WLAN_INFO",
	"UPF_INFO",
}

func (v *SmfEventAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SmfEventAnyOf(value)
	for _, existing := range AllowedSmfEventAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SmfEventAnyOf", value)
}

// NewSmfEventAnyOfFromValue returns a pointer to a valid SmfEventAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSmfEventAnyOfFromValue(v string) (*SmfEventAnyOf, error) {
	ev := SmfEventAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SmfEventAnyOf: valid values are %v", v, AllowedSmfEventAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SmfEventAnyOf) IsValid() bool {
	for _, existing := range AllowedSmfEventAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SmfEvent_anyOf value
func (v SmfEventAnyOf) Ptr() *SmfEventAnyOf {
	return &v
}

type NullableSmfEventAnyOf struct {
	value *SmfEventAnyOf
	isSet bool
}

func (v NullableSmfEventAnyOf) Get() *SmfEventAnyOf {
	return v.value
}

func (v *NullableSmfEventAnyOf) Set(val *SmfEventAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSmfEventAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSmfEventAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmfEventAnyOf(val *SmfEventAnyOf) *NullableSmfEventAnyOf {
	return &NullableSmfEventAnyOf{value: val, isSet: true}
}

func (v NullableSmfEventAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmfEventAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

