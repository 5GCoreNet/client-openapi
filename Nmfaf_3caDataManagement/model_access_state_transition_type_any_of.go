/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
	"fmt"
)

// AccessStateTransitionTypeAnyOf the model 'AccessStateTransitionTypeAnyOf'
type AccessStateTransitionTypeAnyOf string

// List of AccessStateTransitionType_anyOf
const (
	ACCESS_TYPE_CHANGE_3_GPP AccessStateTransitionTypeAnyOf = "ACCESS_TYPE_CHANGE_3GPP"
	ACCESS_TYPE_CHANGE_N3_GPP AccessStateTransitionTypeAnyOf = "ACCESS_TYPE_CHANGE_N3GPP"
	RM_STATE_CHANGE_DEREGISTERED AccessStateTransitionTypeAnyOf = "RM_STATE_CHANGE_DEREGISTERED"
	RM_STATE_CHANGE_REGISTERED AccessStateTransitionTypeAnyOf = "RM_STATE_CHANGE_REGISTERED"
	CM_STATE_CHANGE_IDLE AccessStateTransitionTypeAnyOf = "CM_STATE_CHANGE_IDLE"
	CM_STATE_CHANGE_CONNECTED AccessStateTransitionTypeAnyOf = "CM_STATE_CHANGE_CONNECTED"
	HANDOVER AccessStateTransitionTypeAnyOf = "HANDOVER"
	MOBILITY_REGISTRATION_UPDATE AccessStateTransitionTypeAnyOf = "MOBILITY_REGISTRATION_UPDATE"
)

// All allowed values of AccessStateTransitionTypeAnyOf enum
var AllowedAccessStateTransitionTypeAnyOfEnumValues = []AccessStateTransitionTypeAnyOf{
	"ACCESS_TYPE_CHANGE_3GPP",
	"ACCESS_TYPE_CHANGE_N3GPP",
	"RM_STATE_CHANGE_DEREGISTERED",
	"RM_STATE_CHANGE_REGISTERED",
	"CM_STATE_CHANGE_IDLE",
	"CM_STATE_CHANGE_CONNECTED",
	"HANDOVER",
	"MOBILITY_REGISTRATION_UPDATE",
}

func (v *AccessStateTransitionTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccessStateTransitionTypeAnyOf(value)
	for _, existing := range AllowedAccessStateTransitionTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccessStateTransitionTypeAnyOf", value)
}

// NewAccessStateTransitionTypeAnyOfFromValue returns a pointer to a valid AccessStateTransitionTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessStateTransitionTypeAnyOfFromValue(v string) (*AccessStateTransitionTypeAnyOf, error) {
	ev := AccessStateTransitionTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccessStateTransitionTypeAnyOf: valid values are %v", v, AllowedAccessStateTransitionTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessStateTransitionTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedAccessStateTransitionTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessStateTransitionType_anyOf value
func (v AccessStateTransitionTypeAnyOf) Ptr() *AccessStateTransitionTypeAnyOf {
	return &v
}

type NullableAccessStateTransitionTypeAnyOf struct {
	value *AccessStateTransitionTypeAnyOf
	isSet bool
}

func (v NullableAccessStateTransitionTypeAnyOf) Get() *AccessStateTransitionTypeAnyOf {
	return v.value
}

func (v *NullableAccessStateTransitionTypeAnyOf) Set(val *AccessStateTransitionTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessStateTransitionTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessStateTransitionTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessStateTransitionTypeAnyOf(val *AccessStateTransitionTypeAnyOf) *NullableAccessStateTransitionTypeAnyOf {
	return &NullableAccessStateTransitionTypeAnyOf{value: val, isSet: true}
}

func (v NullableAccessStateTransitionTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessStateTransitionTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

