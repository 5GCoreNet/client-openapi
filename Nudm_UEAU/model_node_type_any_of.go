/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_UEAU

import (
	"encoding/json"
	"fmt"
)

// NodeTypeAnyOf the model 'NodeTypeAnyOf'
type NodeTypeAnyOf string

// List of NodeType_anyOf
const (
	AUSF NodeTypeAnyOf = "AUSF"
	VLR NodeTypeAnyOf = "VLR"
	SGSN NodeTypeAnyOf = "SGSN"
	S_CSCF NodeTypeAnyOf = "S_CSCF"
	BSF NodeTypeAnyOf = "BSF"
	GAN_AAA_SERVER NodeTypeAnyOf = "GAN_AAA_SERVER"
	WLAN_AAA_SERVER NodeTypeAnyOf = "WLAN_AAA_SERVER"
	MME NodeTypeAnyOf = "MME"
)

// All allowed values of NodeTypeAnyOf enum
var AllowedNodeTypeAnyOfEnumValues = []NodeTypeAnyOf{
	"AUSF",
	"VLR",
	"SGSN",
	"S_CSCF",
	"BSF",
	"GAN_AAA_SERVER",
	"WLAN_AAA_SERVER",
	"MME",
}

func (v *NodeTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NodeTypeAnyOf(value)
	for _, existing := range AllowedNodeTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NodeTypeAnyOf", value)
}

// NewNodeTypeAnyOfFromValue returns a pointer to a valid NodeTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNodeTypeAnyOfFromValue(v string) (*NodeTypeAnyOf, error) {
	ev := NodeTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NodeTypeAnyOf: valid values are %v", v, AllowedNodeTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NodeTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedNodeTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NodeType_anyOf value
func (v NodeTypeAnyOf) Ptr() *NodeTypeAnyOf {
	return &v
}

type NullableNodeTypeAnyOf struct {
	value *NodeTypeAnyOf
	isSet bool
}

func (v NullableNodeTypeAnyOf) Get() *NodeTypeAnyOf {
	return v.value
}

func (v *NullableNodeTypeAnyOf) Set(val *NodeTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeTypeAnyOf(val *NodeTypeAnyOf) *NullableNodeTypeAnyOf {
	return &NullableNodeTypeAnyOf{value: val, isSet: true}
}

func (v NullableNodeTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

