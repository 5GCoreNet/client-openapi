/*
Nudm_UECM

Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_UECM

import (
	"encoding/json"
	"fmt"
)

// DeregistrationReasonAnyOf the model 'DeregistrationReasonAnyOf'
type DeregistrationReasonAnyOf string

// List of DeregistrationReason_anyOf
const (
	UE_INITIAL_REGISTRATION DeregistrationReasonAnyOf = "UE_INITIAL_REGISTRATION"
	UE_REGISTRATION_AREA_CHANGE DeregistrationReasonAnyOf = "UE_REGISTRATION_AREA_CHANGE"
	SUBSCRIPTION_WITHDRAWN DeregistrationReasonAnyOf = "SUBSCRIPTION_WITHDRAWN"
	_5_GS_TO_EPS_MOBILITY DeregistrationReasonAnyOf = "5GS_TO_EPS_MOBILITY"
	_5_GS_TO_EPS_MOBILITY_UE_INITIAL_REGISTRATION DeregistrationReasonAnyOf = "5GS_TO_EPS_MOBILITY_UE_INITIAL_REGISTRATION"
	REREGISTRATION_REQUIRED DeregistrationReasonAnyOf = "REREGISTRATION_REQUIRED"
	SMF_CONTEXT_TRANSFERRED DeregistrationReasonAnyOf = "SMF_CONTEXT_TRANSFERRED"
	DUPLICATE_PDU_SESSION DeregistrationReasonAnyOf = "DUPLICATE_PDU_SESSION"
	_5_G_SRVCC_TO_UTRAN_MOBILITY DeregistrationReasonAnyOf = "5G_SRVCC_TO_UTRAN_MOBILITY"
)

// All allowed values of DeregistrationReasonAnyOf enum
var AllowedDeregistrationReasonAnyOfEnumValues = []DeregistrationReasonAnyOf{
	"UE_INITIAL_REGISTRATION",
	"UE_REGISTRATION_AREA_CHANGE",
	"SUBSCRIPTION_WITHDRAWN",
	"5GS_TO_EPS_MOBILITY",
	"5GS_TO_EPS_MOBILITY_UE_INITIAL_REGISTRATION",
	"REREGISTRATION_REQUIRED",
	"SMF_CONTEXT_TRANSFERRED",
	"DUPLICATE_PDU_SESSION",
	"5G_SRVCC_TO_UTRAN_MOBILITY",
}

func (v *DeregistrationReasonAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeregistrationReasonAnyOf(value)
	for _, existing := range AllowedDeregistrationReasonAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeregistrationReasonAnyOf", value)
}

// NewDeregistrationReasonAnyOfFromValue returns a pointer to a valid DeregistrationReasonAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeregistrationReasonAnyOfFromValue(v string) (*DeregistrationReasonAnyOf, error) {
	ev := DeregistrationReasonAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeregistrationReasonAnyOf: valid values are %v", v, AllowedDeregistrationReasonAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeregistrationReasonAnyOf) IsValid() bool {
	for _, existing := range AllowedDeregistrationReasonAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeregistrationReason_anyOf value
func (v DeregistrationReasonAnyOf) Ptr() *DeregistrationReasonAnyOf {
	return &v
}

type NullableDeregistrationReasonAnyOf struct {
	value *DeregistrationReasonAnyOf
	isSet bool
}

func (v NullableDeregistrationReasonAnyOf) Get() *DeregistrationReasonAnyOf {
	return v.value
}

func (v *NullableDeregistrationReasonAnyOf) Set(val *DeregistrationReasonAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeregistrationReasonAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeregistrationReasonAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeregistrationReasonAnyOf(val *DeregistrationReasonAnyOf) *NullableDeregistrationReasonAnyOf {
	return &NullableDeregistrationReasonAnyOf{value: val, isSet: true}
}

func (v NullableDeregistrationReasonAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeregistrationReasonAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

