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

// RadioResourcesIdAnyOf the model 'RadioResourcesIdAnyOf'
type RadioResourcesIdAnyOf string

// List of RadioResourcesId_anyOf
const (
	OPERATOR_PROVIDED RadioResourcesIdAnyOf = "OPERATOR_PROVIDED"
	CONFIGURED RadioResourcesIdAnyOf = "CONFIGURED"
)

// All allowed values of RadioResourcesIdAnyOf enum
var AllowedRadioResourcesIdAnyOfEnumValues = []RadioResourcesIdAnyOf{
	"OPERATOR_PROVIDED",
	"CONFIGURED",
}

func (v *RadioResourcesIdAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RadioResourcesIdAnyOf(value)
	for _, existing := range AllowedRadioResourcesIdAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RadioResourcesIdAnyOf", value)
}

// NewRadioResourcesIdAnyOfFromValue returns a pointer to a valid RadioResourcesIdAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRadioResourcesIdAnyOfFromValue(v string) (*RadioResourcesIdAnyOf, error) {
	ev := RadioResourcesIdAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RadioResourcesIdAnyOf: valid values are %v", v, AllowedRadioResourcesIdAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RadioResourcesIdAnyOf) IsValid() bool {
	for _, existing := range AllowedRadioResourcesIdAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RadioResourcesId_anyOf value
func (v RadioResourcesIdAnyOf) Ptr() *RadioResourcesIdAnyOf {
	return &v
}

type NullableRadioResourcesIdAnyOf struct {
	value *RadioResourcesIdAnyOf
	isSet bool
}

func (v NullableRadioResourcesIdAnyOf) Get() *RadioResourcesIdAnyOf {
	return v.value
}

func (v *NullableRadioResourcesIdAnyOf) Set(val *RadioResourcesIdAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRadioResourcesIdAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRadioResourcesIdAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadioResourcesIdAnyOf(val *RadioResourcesIdAnyOf) *NullableRadioResourcesIdAnyOf {
	return &NullableRadioResourcesIdAnyOf{value: val, isSet: true}
}

func (v NullableRadioResourcesIdAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadioResourcesIdAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

