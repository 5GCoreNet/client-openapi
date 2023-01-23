/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
	"fmt"
)

// DataSetIdAnyOf the model 'DataSetIdAnyOf'
type DataSetIdAnyOf string

// List of DataSetId_anyOf
const (
	SUBSCRIPTION DataSetIdAnyOf = "SUBSCRIPTION"
	POLICY DataSetIdAnyOf = "POLICY"
	EXPOSURE DataSetIdAnyOf = "EXPOSURE"
	APPLICATION DataSetIdAnyOf = "APPLICATION"
	A_PFD DataSetIdAnyOf = "A_PFD"
	A_AFTI DataSetIdAnyOf = "A_AFTI"
	A_IPTV DataSetIdAnyOf = "A_IPTV"
	A_BDT DataSetIdAnyOf = "A_BDT"
	A_SPD DataSetIdAnyOf = "A_SPD"
	A_EASD DataSetIdAnyOf = "A_EASD"
	A_AMI DataSetIdAnyOf = "A_AMI"
	P_UE DataSetIdAnyOf = "P_UE"
	P_SCD DataSetIdAnyOf = "P_SCD"
	P_BDT DataSetIdAnyOf = "P_BDT"
	P_PLMNUE DataSetIdAnyOf = "P_PLMNUE"
	P_NSSCD DataSetIdAnyOf = "P_NSSCD"
)

// All allowed values of DataSetIdAnyOf enum
var AllowedDataSetIdAnyOfEnumValues = []DataSetIdAnyOf{
	"SUBSCRIPTION",
	"POLICY",
	"EXPOSURE",
	"APPLICATION",
	"A_PFD",
	"A_AFTI",
	"A_IPTV",
	"A_BDT",
	"A_SPD",
	"A_EASD",
	"A_AMI",
	"P_UE",
	"P_SCD",
	"P_BDT",
	"P_PLMNUE",
	"P_NSSCD",
}

func (v *DataSetIdAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataSetIdAnyOf(value)
	for _, existing := range AllowedDataSetIdAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataSetIdAnyOf", value)
}

// NewDataSetIdAnyOfFromValue returns a pointer to a valid DataSetIdAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataSetIdAnyOfFromValue(v string) (*DataSetIdAnyOf, error) {
	ev := DataSetIdAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataSetIdAnyOf: valid values are %v", v, AllowedDataSetIdAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataSetIdAnyOf) IsValid() bool {
	for _, existing := range AllowedDataSetIdAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataSetId_anyOf value
func (v DataSetIdAnyOf) Ptr() *DataSetIdAnyOf {
	return &v
}

type NullableDataSetIdAnyOf struct {
	value *DataSetIdAnyOf
	isSet bool
}

func (v NullableDataSetIdAnyOf) Get() *DataSetIdAnyOf {
	return v.value
}

func (v *NullableDataSetIdAnyOf) Set(val *DataSetIdAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSetIdAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSetIdAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSetIdAnyOf(val *DataSetIdAnyOf) *NullableDataSetIdAnyOf {
	return &NullableDataSetIdAnyOf{value: val, isSet: true}
}

func (v NullableDataSetIdAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSetIdAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

