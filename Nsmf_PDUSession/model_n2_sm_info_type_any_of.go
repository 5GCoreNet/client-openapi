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

// N2SmInfoTypeAnyOf the model 'N2SmInfoTypeAnyOf'
type N2SmInfoTypeAnyOf string

// List of N2SmInfoType_anyOf
const (
	PDU_RES_SETUP_REQ N2SmInfoTypeAnyOf = "PDU_RES_SETUP_REQ"
	PDU_RES_SETUP_RSP N2SmInfoTypeAnyOf = "PDU_RES_SETUP_RSP"
	PDU_RES_SETUP_FAIL N2SmInfoTypeAnyOf = "PDU_RES_SETUP_FAIL"
	PDU_RES_REL_CMD N2SmInfoTypeAnyOf = "PDU_RES_REL_CMD"
	PDU_RES_REL_RSP N2SmInfoTypeAnyOf = "PDU_RES_REL_RSP"
	PDU_RES_MOD_REQ N2SmInfoTypeAnyOf = "PDU_RES_MOD_REQ"
	PDU_RES_MOD_RSP N2SmInfoTypeAnyOf = "PDU_RES_MOD_RSP"
	PDU_RES_MOD_FAIL N2SmInfoTypeAnyOf = "PDU_RES_MOD_FAIL"
	PDU_RES_NTY N2SmInfoTypeAnyOf = "PDU_RES_NTY"
	PDU_RES_NTY_REL N2SmInfoTypeAnyOf = "PDU_RES_NTY_REL"
	PDU_RES_MOD_IND N2SmInfoTypeAnyOf = "PDU_RES_MOD_IND"
	PDU_RES_MOD_CFM N2SmInfoTypeAnyOf = "PDU_RES_MOD_CFM"
	PATH_SWITCH_REQ N2SmInfoTypeAnyOf = "PATH_SWITCH_REQ"
	PATH_SWITCH_SETUP_FAIL N2SmInfoTypeAnyOf = "PATH_SWITCH_SETUP_FAIL"
	PATH_SWITCH_REQ_ACK N2SmInfoTypeAnyOf = "PATH_SWITCH_REQ_ACK"
	PATH_SWITCH_REQ_FAIL N2SmInfoTypeAnyOf = "PATH_SWITCH_REQ_FAIL"
	HANDOVER_REQUIRED N2SmInfoTypeAnyOf = "HANDOVER_REQUIRED"
	HANDOVER_CMD N2SmInfoTypeAnyOf = "HANDOVER_CMD"
	HANDOVER_PREP_FAIL N2SmInfoTypeAnyOf = "HANDOVER_PREP_FAIL"
	HANDOVER_REQ_ACK N2SmInfoTypeAnyOf = "HANDOVER_REQ_ACK"
	HANDOVER_RES_ALLOC_FAIL N2SmInfoTypeAnyOf = "HANDOVER_RES_ALLOC_FAIL"
	SECONDARY_RAT_USAGE N2SmInfoTypeAnyOf = "SECONDARY_RAT_USAGE"
	PDU_RES_MOD_IND_FAIL N2SmInfoTypeAnyOf = "PDU_RES_MOD_IND_FAIL"
	UE_CONTEXT_RESUME_REQ N2SmInfoTypeAnyOf = "UE_CONTEXT_RESUME_REQ"
	UE_CONTEXT_RESUME_RSP N2SmInfoTypeAnyOf = "UE_CONTEXT_RESUME_RSP"
	UE_CONTEXT_SUSPEND_REQ N2SmInfoTypeAnyOf = "UE_CONTEXT_SUSPEND_REQ"
)

// All allowed values of N2SmInfoTypeAnyOf enum
var AllowedN2SmInfoTypeAnyOfEnumValues = []N2SmInfoTypeAnyOf{
	"PDU_RES_SETUP_REQ",
	"PDU_RES_SETUP_RSP",
	"PDU_RES_SETUP_FAIL",
	"PDU_RES_REL_CMD",
	"PDU_RES_REL_RSP",
	"PDU_RES_MOD_REQ",
	"PDU_RES_MOD_RSP",
	"PDU_RES_MOD_FAIL",
	"PDU_RES_NTY",
	"PDU_RES_NTY_REL",
	"PDU_RES_MOD_IND",
	"PDU_RES_MOD_CFM",
	"PATH_SWITCH_REQ",
	"PATH_SWITCH_SETUP_FAIL",
	"PATH_SWITCH_REQ_ACK",
	"PATH_SWITCH_REQ_FAIL",
	"HANDOVER_REQUIRED",
	"HANDOVER_CMD",
	"HANDOVER_PREP_FAIL",
	"HANDOVER_REQ_ACK",
	"HANDOVER_RES_ALLOC_FAIL",
	"SECONDARY_RAT_USAGE",
	"PDU_RES_MOD_IND_FAIL",
	"UE_CONTEXT_RESUME_REQ",
	"UE_CONTEXT_RESUME_RSP",
	"UE_CONTEXT_SUSPEND_REQ",
}

func (v *N2SmInfoTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := N2SmInfoTypeAnyOf(value)
	for _, existing := range AllowedN2SmInfoTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid N2SmInfoTypeAnyOf", value)
}

// NewN2SmInfoTypeAnyOfFromValue returns a pointer to a valid N2SmInfoTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewN2SmInfoTypeAnyOfFromValue(v string) (*N2SmInfoTypeAnyOf, error) {
	ev := N2SmInfoTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for N2SmInfoTypeAnyOf: valid values are %v", v, AllowedN2SmInfoTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v N2SmInfoTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedN2SmInfoTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to N2SmInfoType_anyOf value
func (v N2SmInfoTypeAnyOf) Ptr() *N2SmInfoTypeAnyOf {
	return &v
}

type NullableN2SmInfoTypeAnyOf struct {
	value *N2SmInfoTypeAnyOf
	isSet bool
}

func (v NullableN2SmInfoTypeAnyOf) Get() *N2SmInfoTypeAnyOf {
	return v.value
}

func (v *NullableN2SmInfoTypeAnyOf) Set(val *N2SmInfoTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableN2SmInfoTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableN2SmInfoTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2SmInfoTypeAnyOf(val *N2SmInfoTypeAnyOf) *NullableN2SmInfoTypeAnyOf {
	return &NullableN2SmInfoTypeAnyOf{value: val, isSet: true}
}

func (v NullableN2SmInfoTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2SmInfoTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

