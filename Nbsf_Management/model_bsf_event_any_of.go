/*
Nbsf_Management

Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsf_Management

import (
	"encoding/json"
	"fmt"
)

// BsfEventAnyOf the model 'BsfEventAnyOf'
type BsfEventAnyOf string

// List of BsfEvent_anyOf
const (
	PCF_PDU_SESSION_BINDING_REGISTRATION BsfEventAnyOf = "PCF_PDU_SESSION_BINDING_REGISTRATION"
	PCF_PDU_SESSION_BINDING_DEREGISTRATION BsfEventAnyOf = "PCF_PDU_SESSION_BINDING_DEREGISTRATION"
	PCF_UE_BINDING_REGISTRATION BsfEventAnyOf = "PCF_UE_BINDING_REGISTRATION"
	PCF_UE_BINDING_DEREGISTRATION BsfEventAnyOf = "PCF_UE_BINDING_DEREGISTRATION"
	SNSSAI_DNN_BINDING_REGISTRATION BsfEventAnyOf = "SNSSAI_DNN_BINDING_REGISTRATION"
	SNSSAI_DNN_BINDING_DEREGISTRATION BsfEventAnyOf = "SNSSAI_DNN_BINDING_DEREGISTRATION"
)

// All allowed values of BsfEventAnyOf enum
var AllowedBsfEventAnyOfEnumValues = []BsfEventAnyOf{
	"PCF_PDU_SESSION_BINDING_REGISTRATION",
	"PCF_PDU_SESSION_BINDING_DEREGISTRATION",
	"PCF_UE_BINDING_REGISTRATION",
	"PCF_UE_BINDING_DEREGISTRATION",
	"SNSSAI_DNN_BINDING_REGISTRATION",
	"SNSSAI_DNN_BINDING_DEREGISTRATION",
}

func (v *BsfEventAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BsfEventAnyOf(value)
	for _, existing := range AllowedBsfEventAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BsfEventAnyOf", value)
}

// NewBsfEventAnyOfFromValue returns a pointer to a valid BsfEventAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBsfEventAnyOfFromValue(v string) (*BsfEventAnyOf, error) {
	ev := BsfEventAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BsfEventAnyOf: valid values are %v", v, AllowedBsfEventAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BsfEventAnyOf) IsValid() bool {
	for _, existing := range AllowedBsfEventAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BsfEvent_anyOf value
func (v BsfEventAnyOf) Ptr() *BsfEventAnyOf {
	return &v
}

type NullableBsfEventAnyOf struct {
	value *BsfEventAnyOf
	isSet bool
}

func (v NullableBsfEventAnyOf) Get() *BsfEventAnyOf {
	return v.value
}

func (v *NullableBsfEventAnyOf) Set(val *BsfEventAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBsfEventAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBsfEventAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBsfEventAnyOf(val *BsfEventAnyOf) *NullableBsfEventAnyOf {
	return &NullableBsfEventAnyOf{value: val, isSet: true}
}

func (v NullableBsfEventAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBsfEventAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

