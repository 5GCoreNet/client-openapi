/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_UEAU

import (
	"encoding/json"
	"fmt"
)

// HssAvTypeAnyOf the model 'HssAvTypeAnyOf'
type HssAvTypeAnyOf string

// List of HssAvType_anyOf
const (
	EPS_AKA HssAvTypeAnyOf = "EPS_AKA"
	EAP_AKA HssAvTypeAnyOf = "EAP_AKA"
	IMS_AKA HssAvTypeAnyOf = "IMS_AKA"
	GBA_AKA HssAvTypeAnyOf = "GBA_AKA"
	UMTS_AKA HssAvTypeAnyOf = "UMTS_AKA"
)

// All allowed values of HssAvTypeAnyOf enum
var AllowedHssAvTypeAnyOfEnumValues = []HssAvTypeAnyOf{
	"EPS_AKA",
	"EAP_AKA",
	"IMS_AKA",
	"GBA_AKA",
	"UMTS_AKA",
}

func (v *HssAvTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HssAvTypeAnyOf(value)
	for _, existing := range AllowedHssAvTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HssAvTypeAnyOf", value)
}

// NewHssAvTypeAnyOfFromValue returns a pointer to a valid HssAvTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHssAvTypeAnyOfFromValue(v string) (*HssAvTypeAnyOf, error) {
	ev := HssAvTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HssAvTypeAnyOf: valid values are %v", v, AllowedHssAvTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HssAvTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedHssAvTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HssAvType_anyOf value
func (v HssAvTypeAnyOf) Ptr() *HssAvTypeAnyOf {
	return &v
}

type NullableHssAvTypeAnyOf struct {
	value *HssAvTypeAnyOf
	isSet bool
}

func (v NullableHssAvTypeAnyOf) Get() *HssAvTypeAnyOf {
	return v.value
}

func (v *NullableHssAvTypeAnyOf) Set(val *HssAvTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHssAvTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHssAvTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHssAvTypeAnyOf(val *HssAvTypeAnyOf) *NullableHssAvTypeAnyOf {
	return &NullableHssAvTypeAnyOf{value: val, isSet: true}
}

func (v NullableHssAvTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHssAvTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

