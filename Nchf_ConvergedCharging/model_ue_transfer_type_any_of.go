/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// UETransferTypeAnyOf the model 'UETransferTypeAnyOf'
type UETransferTypeAnyOf string

// List of UETransferType_anyOf
const (
	INTRA_UE UETransferTypeAnyOf = "INTRA_UE"
	INTER_UE UETransferTypeAnyOf = "INTER_UE"
)

// All allowed values of UETransferTypeAnyOf enum
var AllowedUETransferTypeAnyOfEnumValues = []UETransferTypeAnyOf{
	"INTRA_UE",
	"INTER_UE",
}

func (v *UETransferTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UETransferTypeAnyOf(value)
	for _, existing := range AllowedUETransferTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UETransferTypeAnyOf", value)
}

// NewUETransferTypeAnyOfFromValue returns a pointer to a valid UETransferTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUETransferTypeAnyOfFromValue(v string) (*UETransferTypeAnyOf, error) {
	ev := UETransferTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UETransferTypeAnyOf: valid values are %v", v, AllowedUETransferTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UETransferTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedUETransferTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UETransferType_anyOf value
func (v UETransferTypeAnyOf) Ptr() *UETransferTypeAnyOf {
	return &v
}

type NullableUETransferTypeAnyOf struct {
	value *UETransferTypeAnyOf
	isSet bool
}

func (v NullableUETransferTypeAnyOf) Get() *UETransferTypeAnyOf {
	return v.value
}

func (v *NullableUETransferTypeAnyOf) Set(val *UETransferTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUETransferTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUETransferTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUETransferTypeAnyOf(val *UETransferTypeAnyOf) *NullableUETransferTypeAnyOf {
	return &NullableUETransferTypeAnyOf{value: val, isSet: true}
}

func (v NullableUETransferTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUETransferTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

