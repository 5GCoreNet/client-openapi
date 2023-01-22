/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// AccessTransferTypeAnyOf the model 'AccessTransferTypeAnyOf'
type AccessTransferTypeAnyOf string

// List of AccessTransferType_anyOf
const (
	PS_TO_CS AccessTransferTypeAnyOf = "PS_TO_CS"
	CS_TO_PS AccessTransferTypeAnyOf = "CS_TO_PS"
	PS_TO_PS AccessTransferTypeAnyOf = "PS_TO_PS"
	CS_TO_CS AccessTransferTypeAnyOf = "CS_TO_CS"
)

// All allowed values of AccessTransferTypeAnyOf enum
var AllowedAccessTransferTypeAnyOfEnumValues = []AccessTransferTypeAnyOf{
	"PS_TO_CS",
	"CS_TO_PS",
	"PS_TO_PS",
	"CS_TO_CS",
}

func (v *AccessTransferTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccessTransferTypeAnyOf(value)
	for _, existing := range AllowedAccessTransferTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccessTransferTypeAnyOf", value)
}

// NewAccessTransferTypeAnyOfFromValue returns a pointer to a valid AccessTransferTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessTransferTypeAnyOfFromValue(v string) (*AccessTransferTypeAnyOf, error) {
	ev := AccessTransferTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccessTransferTypeAnyOf: valid values are %v", v, AllowedAccessTransferTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessTransferTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedAccessTransferTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessTransferType_anyOf value
func (v AccessTransferTypeAnyOf) Ptr() *AccessTransferTypeAnyOf {
	return &v
}

type NullableAccessTransferTypeAnyOf struct {
	value *AccessTransferTypeAnyOf
	isSet bool
}

func (v NullableAccessTransferTypeAnyOf) Get() *AccessTransferTypeAnyOf {
	return v.value
}

func (v *NullableAccessTransferTypeAnyOf) Set(val *AccessTransferTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTransferTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTransferTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTransferTypeAnyOf(val *AccessTransferTypeAnyOf) *NullableAccessTransferTypeAnyOf {
	return &NullableAccessTransferTypeAnyOf{value: val, isSet: true}
}

func (v NullableAccessTransferTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTransferTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

