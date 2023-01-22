/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// NgapIeTypeAnyOf the model 'NgapIeTypeAnyOf'
type NgapIeTypeAnyOf string

// List of NgapIeType_anyOf
const (
	PDU_RES_SETUP_REQ NgapIeTypeAnyOf = "PDU_RES_SETUP_REQ"
	PDU_RES_REL_CMD NgapIeTypeAnyOf = "PDU_RES_REL_CMD"
	PDU_RES_MOD_REQ NgapIeTypeAnyOf = "PDU_RES_MOD_REQ"
	HANDOVER_CMD NgapIeTypeAnyOf = "HANDOVER_CMD"
	HANDOVER_REQUIRED NgapIeTypeAnyOf = "HANDOVER_REQUIRED"
	HANDOVER_PREP_FAIL NgapIeTypeAnyOf = "HANDOVER_PREP_FAIL"
	SRC_TO_TAR_CONTAINER NgapIeTypeAnyOf = "SRC_TO_TAR_CONTAINER"
	TAR_TO_SRC_CONTAINER NgapIeTypeAnyOf = "TAR_TO_SRC_CONTAINER"
	TAR_TO_SRC_FAIL_CONTAINER NgapIeTypeAnyOf = "TAR_TO_SRC_FAIL_CONTAINER"
	RAN_STATUS_TRANS_CONTAINER NgapIeTypeAnyOf = "RAN_STATUS_TRANS_CONTAINER"
	SON_CONFIG_TRANSFER NgapIeTypeAnyOf = "SON_CONFIG_TRANSFER"
	NRPPA_PDU NgapIeTypeAnyOf = "NRPPA_PDU"
	UE_RADIO_CAPABILITY NgapIeTypeAnyOf = "UE_RADIO_CAPABILITY"
	RIM_INFO_TRANSFER NgapIeTypeAnyOf = "RIM_INFO_TRANSFER"
	SECONDARY_RAT_USAGE NgapIeTypeAnyOf = "SECONDARY_RAT_USAGE"
	PC5_QOS_PARA NgapIeTypeAnyOf = "PC5_QOS_PARA"
	EARLY_STATUS_TRANS_CONTAINER NgapIeTypeAnyOf = "EARLY_STATUS_TRANS_CONTAINER"
	UE_RADIO_CAPABILITY_FOR_PAGING NgapIeTypeAnyOf = "UE_RADIO_CAPABILITY_FOR_PAGING"
)

// All allowed values of NgapIeTypeAnyOf enum
var AllowedNgapIeTypeAnyOfEnumValues = []NgapIeTypeAnyOf{
	"PDU_RES_SETUP_REQ",
	"PDU_RES_REL_CMD",
	"PDU_RES_MOD_REQ",
	"HANDOVER_CMD",
	"HANDOVER_REQUIRED",
	"HANDOVER_PREP_FAIL",
	"SRC_TO_TAR_CONTAINER",
	"TAR_TO_SRC_CONTAINER",
	"TAR_TO_SRC_FAIL_CONTAINER",
	"RAN_STATUS_TRANS_CONTAINER",
	"SON_CONFIG_TRANSFER",
	"NRPPA_PDU",
	"UE_RADIO_CAPABILITY",
	"RIM_INFO_TRANSFER",
	"SECONDARY_RAT_USAGE",
	"PC5_QOS_PARA",
	"EARLY_STATUS_TRANS_CONTAINER",
	"UE_RADIO_CAPABILITY_FOR_PAGING",
}

func (v *NgapIeTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NgapIeTypeAnyOf(value)
	for _, existing := range AllowedNgapIeTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NgapIeTypeAnyOf", value)
}

// NewNgapIeTypeAnyOfFromValue returns a pointer to a valid NgapIeTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNgapIeTypeAnyOfFromValue(v string) (*NgapIeTypeAnyOf, error) {
	ev := NgapIeTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NgapIeTypeAnyOf: valid values are %v", v, AllowedNgapIeTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NgapIeTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedNgapIeTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NgapIeType_anyOf value
func (v NgapIeTypeAnyOf) Ptr() *NgapIeTypeAnyOf {
	return &v
}

type NullableNgapIeTypeAnyOf struct {
	value *NgapIeTypeAnyOf
	isSet bool
}

func (v NullableNgapIeTypeAnyOf) Get() *NgapIeTypeAnyOf {
	return v.value
}

func (v *NullableNgapIeTypeAnyOf) Set(val *NgapIeTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNgapIeTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNgapIeTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNgapIeTypeAnyOf(val *NgapIeTypeAnyOf) *NullableNgapIeTypeAnyOf {
	return &NullableNgapIeTypeAnyOf{value: val, isSet: true}
}

func (v NullableNgapIeTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNgapIeTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

