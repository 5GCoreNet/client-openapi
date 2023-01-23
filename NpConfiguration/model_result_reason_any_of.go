/*
3gpp-network-parameter-configuration

API for network parameter configuration.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NpConfiguration

import (
	"encoding/json"
	"fmt"
)

// ResultReasonAnyOf the model 'ResultReasonAnyOf'
type ResultReasonAnyOf string

// List of ResultReason_anyOf
const (
	ROAMING_NOT_ALLOWED ResultReasonAnyOf = "ROAMING_NOT_ALLOWED"
	OTHER_REASON ResultReasonAnyOf = "OTHER_REASON"
)

// All allowed values of ResultReasonAnyOf enum
var AllowedResultReasonAnyOfEnumValues = []ResultReasonAnyOf{
	"ROAMING_NOT_ALLOWED",
	"OTHER_REASON",
}

func (v *ResultReasonAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResultReasonAnyOf(value)
	for _, existing := range AllowedResultReasonAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResultReasonAnyOf", value)
}

// NewResultReasonAnyOfFromValue returns a pointer to a valid ResultReasonAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResultReasonAnyOfFromValue(v string) (*ResultReasonAnyOf, error) {
	ev := ResultReasonAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResultReasonAnyOf: valid values are %v", v, AllowedResultReasonAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResultReasonAnyOf) IsValid() bool {
	for _, existing := range AllowedResultReasonAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResultReason_anyOf value
func (v ResultReasonAnyOf) Ptr() *ResultReasonAnyOf {
	return &v
}

type NullableResultReasonAnyOf struct {
	value *ResultReasonAnyOf
	isSet bool
}

func (v NullableResultReasonAnyOf) Get() *ResultReasonAnyOf {
	return v.value
}

func (v *NullableResultReasonAnyOf) Set(val *ResultReasonAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResultReasonAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResultReasonAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultReasonAnyOf(val *ResultReasonAnyOf) *NullableResultReasonAnyOf {
	return &NullableResultReasonAnyOf{value: val, isSet: true}
}

func (v NullableResultReasonAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultReasonAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

