/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_EventsSubscription

import (
	"encoding/json"
	"fmt"
)

// TransferRequestTypeAnyOf the model 'TransferRequestTypeAnyOf'
type TransferRequestTypeAnyOf string

// List of TransferRequestType_anyOf
const (
	PREPARE TransferRequestTypeAnyOf = "PREPARE"
	TRANSFER TransferRequestTypeAnyOf = "TRANSFER"
)

// All allowed values of TransferRequestTypeAnyOf enum
var AllowedTransferRequestTypeAnyOfEnumValues = []TransferRequestTypeAnyOf{
	"PREPARE",
	"TRANSFER",
}

func (v *TransferRequestTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferRequestTypeAnyOf(value)
	for _, existing := range AllowedTransferRequestTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferRequestTypeAnyOf", value)
}

// NewTransferRequestTypeAnyOfFromValue returns a pointer to a valid TransferRequestTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferRequestTypeAnyOfFromValue(v string) (*TransferRequestTypeAnyOf, error) {
	ev := TransferRequestTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferRequestTypeAnyOf: valid values are %v", v, AllowedTransferRequestTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferRequestTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedTransferRequestTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferRequestType_anyOf value
func (v TransferRequestTypeAnyOf) Ptr() *TransferRequestTypeAnyOf {
	return &v
}

type NullableTransferRequestTypeAnyOf struct {
	value *TransferRequestTypeAnyOf
	isSet bool
}

func (v NullableTransferRequestTypeAnyOf) Get() *TransferRequestTypeAnyOf {
	return v.value
}

func (v *NullableTransferRequestTypeAnyOf) Set(val *TransferRequestTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRequestTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRequestTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRequestTypeAnyOf(val *TransferRequestTypeAnyOf) *NullableTransferRequestTypeAnyOf {
	return &NullableTransferRequestTypeAnyOf{value: val, isSet: true}
}

func (v NullableTransferRequestTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRequestTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

