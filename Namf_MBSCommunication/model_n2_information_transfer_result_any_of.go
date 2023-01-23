/*
Namf_MBSCommunication

AMF Communication Service for MBS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MBSCommunication

import (
	"encoding/json"
	"fmt"
)

// N2InformationTransferResultAnyOf the model 'N2InformationTransferResultAnyOf'
type N2InformationTransferResultAnyOf string

// List of N2InformationTransferResult_anyOf
const (
	N2_INFO_TRANSFER_INITIATED N2InformationTransferResultAnyOf = "N2_INFO_TRANSFER_INITIATED"
)

// All allowed values of N2InformationTransferResultAnyOf enum
var AllowedN2InformationTransferResultAnyOfEnumValues = []N2InformationTransferResultAnyOf{
	"N2_INFO_TRANSFER_INITIATED",
}

func (v *N2InformationTransferResultAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := N2InformationTransferResultAnyOf(value)
	for _, existing := range AllowedN2InformationTransferResultAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid N2InformationTransferResultAnyOf", value)
}

// NewN2InformationTransferResultAnyOfFromValue returns a pointer to a valid N2InformationTransferResultAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewN2InformationTransferResultAnyOfFromValue(v string) (*N2InformationTransferResultAnyOf, error) {
	ev := N2InformationTransferResultAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for N2InformationTransferResultAnyOf: valid values are %v", v, AllowedN2InformationTransferResultAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v N2InformationTransferResultAnyOf) IsValid() bool {
	for _, existing := range AllowedN2InformationTransferResultAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to N2InformationTransferResult_anyOf value
func (v N2InformationTransferResultAnyOf) Ptr() *N2InformationTransferResultAnyOf {
	return &v
}

type NullableN2InformationTransferResultAnyOf struct {
	value *N2InformationTransferResultAnyOf
	isSet bool
}

func (v NullableN2InformationTransferResultAnyOf) Get() *N2InformationTransferResultAnyOf {
	return v.value
}

func (v *NullableN2InformationTransferResultAnyOf) Set(val *N2InformationTransferResultAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableN2InformationTransferResultAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableN2InformationTransferResultAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2InformationTransferResultAnyOf(val *N2InformationTransferResultAnyOf) *NullableN2InformationTransferResultAnyOf {
	return &NullableN2InformationTransferResultAnyOf{value: val, isSet: true}
}

func (v NullableN2InformationTransferResultAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2InformationTransferResultAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

