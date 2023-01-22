/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_SDM

import (
	"encoding/json"
	"fmt"
)

// DataSetNameAnyOf the model 'DataSetNameAnyOf'
type DataSetNameAnyOf string

// List of DataSetName_anyOf
const (
	AM DataSetNameAnyOf = "AM"
	SMF_SEL DataSetNameAnyOf = "SMF_SEL"
	UEC_SMF DataSetNameAnyOf = "UEC_SMF"
	UEC_SMSF DataSetNameAnyOf = "UEC_SMSF"
	SMS_SUB DataSetNameAnyOf = "SMS_SUB"
	SM DataSetNameAnyOf = "SM"
	TRACE DataSetNameAnyOf = "TRACE"
	SMS_MNG DataSetNameAnyOf = "SMS_MNG"
	LCS_PRIVACY DataSetNameAnyOf = "LCS_PRIVACY"
	LCS_MO DataSetNameAnyOf = "LCS_MO"
	UEC_AMF DataSetNameAnyOf = "UEC_AMF"
	V2_X DataSetNameAnyOf = "V2X"
	LCS_BCA DataSetNameAnyOf = "LCS_BCA"
	PROSE DataSetNameAnyOf = "PROSE"
	UC DataSetNameAnyOf = "UC"
	MBS DataSetNameAnyOf = "MBS"
)

// All allowed values of DataSetNameAnyOf enum
var AllowedDataSetNameAnyOfEnumValues = []DataSetNameAnyOf{
	"AM",
	"SMF_SEL",
	"UEC_SMF",
	"UEC_SMSF",
	"SMS_SUB",
	"SM",
	"TRACE",
	"SMS_MNG",
	"LCS_PRIVACY",
	"LCS_MO",
	"UEC_AMF",
	"V2X",
	"LCS_BCA",
	"PROSE",
	"UC",
	"MBS",
}

func (v *DataSetNameAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataSetNameAnyOf(value)
	for _, existing := range AllowedDataSetNameAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataSetNameAnyOf", value)
}

// NewDataSetNameAnyOfFromValue returns a pointer to a valid DataSetNameAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataSetNameAnyOfFromValue(v string) (*DataSetNameAnyOf, error) {
	ev := DataSetNameAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataSetNameAnyOf: valid values are %v", v, AllowedDataSetNameAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataSetNameAnyOf) IsValid() bool {
	for _, existing := range AllowedDataSetNameAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataSetName_anyOf value
func (v DataSetNameAnyOf) Ptr() *DataSetNameAnyOf {
	return &v
}

type NullableDataSetNameAnyOf struct {
	value *DataSetNameAnyOf
	isSet bool
}

func (v NullableDataSetNameAnyOf) Get() *DataSetNameAnyOf {
	return v.value
}

func (v *NullableDataSetNameAnyOf) Set(val *DataSetNameAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSetNameAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSetNameAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSetNameAnyOf(val *DataSetNameAnyOf) *NullableDataSetNameAnyOf {
	return &NullableDataSetNameAnyOf{value: val, isSet: true}
}

func (v NullableDataSetNameAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSetNameAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

