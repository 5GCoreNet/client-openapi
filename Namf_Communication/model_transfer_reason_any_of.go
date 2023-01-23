/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// TransferReasonAnyOf the model 'TransferReasonAnyOf'
type TransferReasonAnyOf string

// List of TransferReason_anyOf
const (
	INIT_REG TransferReasonAnyOf = "INIT_REG"
	MOBI_REG TransferReasonAnyOf = "MOBI_REG"
	MOBI_REG_UE_VALIDATED TransferReasonAnyOf = "MOBI_REG_UE_VALIDATED"
)

// All allowed values of TransferReasonAnyOf enum
var AllowedTransferReasonAnyOfEnumValues = []TransferReasonAnyOf{
	"INIT_REG",
	"MOBI_REG",
	"MOBI_REG_UE_VALIDATED",
}

func (v *TransferReasonAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferReasonAnyOf(value)
	for _, existing := range AllowedTransferReasonAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferReasonAnyOf", value)
}

// NewTransferReasonAnyOfFromValue returns a pointer to a valid TransferReasonAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferReasonAnyOfFromValue(v string) (*TransferReasonAnyOf, error) {
	ev := TransferReasonAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferReasonAnyOf: valid values are %v", v, AllowedTransferReasonAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferReasonAnyOf) IsValid() bool {
	for _, existing := range AllowedTransferReasonAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferReason_anyOf value
func (v TransferReasonAnyOf) Ptr() *TransferReasonAnyOf {
	return &v
}

type NullableTransferReasonAnyOf struct {
	value *TransferReasonAnyOf
	isSet bool
}

func (v NullableTransferReasonAnyOf) Get() *TransferReasonAnyOf {
	return v.value
}

func (v *NullableTransferReasonAnyOf) Set(val *TransferReasonAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferReasonAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferReasonAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferReasonAnyOf(val *TransferReasonAnyOf) *NullableTransferReasonAnyOf {
	return &NullableTransferReasonAnyOf{value: val, isSet: true}
}

func (v NullableTransferReasonAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferReasonAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

