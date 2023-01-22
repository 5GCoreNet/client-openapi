/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// N4MessageTypeAnyOf the model 'N4MessageTypeAnyOf'
type N4MessageTypeAnyOf string

// List of N4MessageType_anyOf
const (
	EST_REQ N4MessageTypeAnyOf = "PFCP_SES_EST_REQ"
	EST_RSP N4MessageTypeAnyOf = "PFCP_SES_EST_RSP"
	MOD_REQ N4MessageTypeAnyOf = "PFCP_SES_MOD_REQ"
	MOD_RSP N4MessageTypeAnyOf = "PFCP_SES_MOD_RSP"
	DEL_REQ N4MessageTypeAnyOf = "PFCP_SES_DEL_REQ"
	DEL_RSP N4MessageTypeAnyOf = "PFCP_SES_DEL_RSP"
	REP_REQ N4MessageTypeAnyOf = "PFCP_SES_REP_REQ"
	REP_RSP N4MessageTypeAnyOf = "PFCP_SES_REP_RSP"
)

// All allowed values of N4MessageTypeAnyOf enum
var AllowedN4MessageTypeAnyOfEnumValues = []N4MessageTypeAnyOf{
	"PFCP_SES_EST_REQ",
	"PFCP_SES_EST_RSP",
	"PFCP_SES_MOD_REQ",
	"PFCP_SES_MOD_RSP",
	"PFCP_SES_DEL_REQ",
	"PFCP_SES_DEL_RSP",
	"PFCP_SES_REP_REQ",
	"PFCP_SES_REP_RSP",
}

func (v *N4MessageTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := N4MessageTypeAnyOf(value)
	for _, existing := range AllowedN4MessageTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid N4MessageTypeAnyOf", value)
}

// NewN4MessageTypeAnyOfFromValue returns a pointer to a valid N4MessageTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewN4MessageTypeAnyOfFromValue(v string) (*N4MessageTypeAnyOf, error) {
	ev := N4MessageTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for N4MessageTypeAnyOf: valid values are %v", v, AllowedN4MessageTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v N4MessageTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedN4MessageTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to N4MessageType_anyOf value
func (v N4MessageTypeAnyOf) Ptr() *N4MessageTypeAnyOf {
	return &v
}

type NullableN4MessageTypeAnyOf struct {
	value *N4MessageTypeAnyOf
	isSet bool
}

func (v NullableN4MessageTypeAnyOf) Get() *N4MessageTypeAnyOf {
	return v.value
}

func (v *NullableN4MessageTypeAnyOf) Set(val *N4MessageTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableN4MessageTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableN4MessageTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN4MessageTypeAnyOf(val *N4MessageTypeAnyOf) *NullableN4MessageTypeAnyOf {
	return &NullableN4MessageTypeAnyOf{value: val, isSet: true}
}

func (v NullableN4MessageTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN4MessageTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

