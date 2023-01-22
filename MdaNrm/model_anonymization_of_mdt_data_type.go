/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MdaNrm

import (
	"encoding/json"
	"fmt"
)

// AnonymizationOfMDTDataType Specifies level of MDT anonymization. For additional details see 3GPP TS 32.422 clause 5.10.12.
type AnonymizationOfMDTDataType string

// List of anonymizationOfMDTData-Type
const (
	NO_IDENTITY AnonymizationOfMDTDataType = "NO_IDENTITY"
	TAC_OF_IMEI AnonymizationOfMDTDataType = "TAC_OF_IMEI"
)

// All allowed values of AnonymizationOfMDTDataType enum
var AllowedAnonymizationOfMDTDataTypeEnumValues = []AnonymizationOfMDTDataType{
	"NO_IDENTITY",
	"TAC_OF_IMEI",
}

func (v *AnonymizationOfMDTDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AnonymizationOfMDTDataType(value)
	for _, existing := range AllowedAnonymizationOfMDTDataTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AnonymizationOfMDTDataType", value)
}

// NewAnonymizationOfMDTDataTypeFromValue returns a pointer to a valid AnonymizationOfMDTDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAnonymizationOfMDTDataTypeFromValue(v string) (*AnonymizationOfMDTDataType, error) {
	ev := AnonymizationOfMDTDataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AnonymizationOfMDTDataType: valid values are %v", v, AllowedAnonymizationOfMDTDataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AnonymizationOfMDTDataType) IsValid() bool {
	for _, existing := range AllowedAnonymizationOfMDTDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to anonymizationOfMDTData-Type value
func (v AnonymizationOfMDTDataType) Ptr() *AnonymizationOfMDTDataType {
	return &v
}

type NullableAnonymizationOfMDTDataType struct {
	value *AnonymizationOfMDTDataType
	isSet bool
}

func (v NullableAnonymizationOfMDTDataType) Get() *AnonymizationOfMDTDataType {
	return v.value
}

func (v *NullableAnonymizationOfMDTDataType) Set(val *AnonymizationOfMDTDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableAnonymizationOfMDTDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableAnonymizationOfMDTDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnonymizationOfMDTDataType(val *AnonymizationOfMDTDataType) *NullableAnonymizationOfMDTDataType {
	return &NullableAnonymizationOfMDTDataType{value: val, isSet: true}
}

func (v NullableAnonymizationOfMDTDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnonymizationOfMDTDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
