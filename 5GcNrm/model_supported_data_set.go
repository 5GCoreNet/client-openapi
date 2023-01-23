/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
	"fmt"
)

// SupportedDataSet any of enumrated value
type SupportedDataSet string

// List of SupportedDataSet
const (
	SUBSCRIPTION SupportedDataSet = "SUBSCRIPTION"
	POLICY SupportedDataSet = "POLICY"
	EXPOSURE SupportedDataSet = "EXPOSURE"
	APPLICATION SupportedDataSet = "APPLICATION"
	A_PFD SupportedDataSet = "A_PFD"
	A_AFTI SupportedDataSet = "A_AFTI"
	A_IPTV SupportedDataSet = "A_IPTV"
	A_BDT SupportedDataSet = "A_BDT"
	A_SPD SupportedDataSet = "A_SPD"
	A_EASD SupportedDataSet = "A_EASD"
	A_AMI SupportedDataSet = "A_AMI"
	P_UE SupportedDataSet = "P_UE"
	P_SCD SupportedDataSet = "P_SCD"
	P_BDT SupportedDataSet = "P_BDT"
	P_PLMNUE SupportedDataSet = "P_PLMNUE"
	P_NSSCD SupportedDataSet = "P_NSSCD"
)

// All allowed values of SupportedDataSet enum
var AllowedSupportedDataSetEnumValues = []SupportedDataSet{
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

func (v *SupportedDataSet) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SupportedDataSet(value)
	for _, existing := range AllowedSupportedDataSetEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SupportedDataSet", value)
}

// NewSupportedDataSetFromValue returns a pointer to a valid SupportedDataSet
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSupportedDataSetFromValue(v string) (*SupportedDataSet, error) {
	ev := SupportedDataSet(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SupportedDataSet: valid values are %v", v, AllowedSupportedDataSetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SupportedDataSet) IsValid() bool {
	for _, existing := range AllowedSupportedDataSetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SupportedDataSet value
func (v SupportedDataSet) Ptr() *SupportedDataSet {
	return &v
}

type NullableSupportedDataSet struct {
	value *SupportedDataSet
	isSet bool
}

func (v NullableSupportedDataSet) Get() *SupportedDataSet {
	return v.value
}

func (v *NullableSupportedDataSet) Set(val *SupportedDataSet) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedDataSet) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedDataSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedDataSet(val *SupportedDataSet) *NullableSupportedDataSet {
	return &NullableSupportedDataSet{value: val, isSet: true}
}

func (v NullableSupportedDataSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedDataSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

