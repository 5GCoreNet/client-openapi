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

// HssAuthTypeInUriAnyOf the model 'HssAuthTypeInUriAnyOf'
type HssAuthTypeInUriAnyOf string

// List of HssAuthTypeInUri_anyOf
const (
	EPS_AKA HssAuthTypeInUriAnyOf = "eps-aka"
	EAP_AKA HssAuthTypeInUriAnyOf = "eap-aka"
	EAP_AKA_PRIME HssAuthTypeInUriAnyOf = "eap-aka-prime"
	IMS_AKA HssAuthTypeInUriAnyOf = "ims-aka"
	GBA_AKA HssAuthTypeInUriAnyOf = "gba-aka"
)

// All allowed values of HssAuthTypeInUriAnyOf enum
var AllowedHssAuthTypeInUriAnyOfEnumValues = []HssAuthTypeInUriAnyOf{
	"eps-aka",
	"eap-aka",
	"eap-aka-prime",
	"ims-aka",
	"gba-aka",
}

func (v *HssAuthTypeInUriAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HssAuthTypeInUriAnyOf(value)
	for _, existing := range AllowedHssAuthTypeInUriAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HssAuthTypeInUriAnyOf", value)
}

// NewHssAuthTypeInUriAnyOfFromValue returns a pointer to a valid HssAuthTypeInUriAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHssAuthTypeInUriAnyOfFromValue(v string) (*HssAuthTypeInUriAnyOf, error) {
	ev := HssAuthTypeInUriAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HssAuthTypeInUriAnyOf: valid values are %v", v, AllowedHssAuthTypeInUriAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HssAuthTypeInUriAnyOf) IsValid() bool {
	for _, existing := range AllowedHssAuthTypeInUriAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HssAuthTypeInUri_anyOf value
func (v HssAuthTypeInUriAnyOf) Ptr() *HssAuthTypeInUriAnyOf {
	return &v
}

type NullableHssAuthTypeInUriAnyOf struct {
	value *HssAuthTypeInUriAnyOf
	isSet bool
}

func (v NullableHssAuthTypeInUriAnyOf) Get() *HssAuthTypeInUriAnyOf {
	return v.value
}

func (v *NullableHssAuthTypeInUriAnyOf) Set(val *HssAuthTypeInUriAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHssAuthTypeInUriAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHssAuthTypeInUriAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHssAuthTypeInUriAnyOf(val *HssAuthTypeInUriAnyOf) *NullableHssAuthTypeInUriAnyOf {
	return &NullableHssAuthTypeInUriAnyOf{value: val, isSet: true}
}

func (v NullableHssAuthTypeInUriAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHssAuthTypeInUriAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

