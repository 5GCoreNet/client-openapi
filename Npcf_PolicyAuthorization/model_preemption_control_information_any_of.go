/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// PreemptionControlInformationAnyOf the model 'PreemptionControlInformationAnyOf'
type PreemptionControlInformationAnyOf string

// List of PreemptionControlInformation_anyOf
const (
	MOST_RECENT PreemptionControlInformationAnyOf = "MOST_RECENT"
	LEAST_RECENT PreemptionControlInformationAnyOf = "LEAST_RECENT"
	HIGHEST_BW PreemptionControlInformationAnyOf = "HIGHEST_BW"
)

// All allowed values of PreemptionControlInformationAnyOf enum
var AllowedPreemptionControlInformationAnyOfEnumValues = []PreemptionControlInformationAnyOf{
	"MOST_RECENT",
	"LEAST_RECENT",
	"HIGHEST_BW",
}

func (v *PreemptionControlInformationAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PreemptionControlInformationAnyOf(value)
	for _, existing := range AllowedPreemptionControlInformationAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PreemptionControlInformationAnyOf", value)
}

// NewPreemptionControlInformationAnyOfFromValue returns a pointer to a valid PreemptionControlInformationAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPreemptionControlInformationAnyOfFromValue(v string) (*PreemptionControlInformationAnyOf, error) {
	ev := PreemptionControlInformationAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PreemptionControlInformationAnyOf: valid values are %v", v, AllowedPreemptionControlInformationAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PreemptionControlInformationAnyOf) IsValid() bool {
	for _, existing := range AllowedPreemptionControlInformationAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PreemptionControlInformation_anyOf value
func (v PreemptionControlInformationAnyOf) Ptr() *PreemptionControlInformationAnyOf {
	return &v
}

type NullablePreemptionControlInformationAnyOf struct {
	value *PreemptionControlInformationAnyOf
	isSet bool
}

func (v NullablePreemptionControlInformationAnyOf) Get() *PreemptionControlInformationAnyOf {
	return v.value
}

func (v *NullablePreemptionControlInformationAnyOf) Set(val *PreemptionControlInformationAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePreemptionControlInformationAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePreemptionControlInformationAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreemptionControlInformationAnyOf(val *PreemptionControlInformationAnyOf) *NullablePreemptionControlInformationAnyOf {
	return &NullablePreemptionControlInformationAnyOf{value: val, isSet: true}
}

func (v NullablePreemptionControlInformationAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreemptionControlInformationAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

