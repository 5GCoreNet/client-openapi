/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// InterfaceTypeAnyOf the model 'InterfaceTypeAnyOf'
type InterfaceTypeAnyOf string

// List of InterfaceType_anyOf
const (
	UNKNOWN InterfaceTypeAnyOf = "UNKNOWN"
	MOBILE_ORIGINATING InterfaceTypeAnyOf = "MOBILE_ORIGINATING"
	MOBILE_TERMINATING InterfaceTypeAnyOf = "MOBILE_TERMINATING"
	APPLICATION_ORIGINATING InterfaceTypeAnyOf = "APPLICATION_ORIGINATING"
	APPLICATION_TERMINATING InterfaceTypeAnyOf = "APPLICATION_TERMINATING"
)

// All allowed values of InterfaceTypeAnyOf enum
var AllowedInterfaceTypeAnyOfEnumValues = []InterfaceTypeAnyOf{
	"UNKNOWN",
	"MOBILE_ORIGINATING",
	"MOBILE_TERMINATING",
	"APPLICATION_ORIGINATING",
	"APPLICATION_TERMINATING",
}

func (v *InterfaceTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceTypeAnyOf(value)
	for _, existing := range AllowedInterfaceTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceTypeAnyOf", value)
}

// NewInterfaceTypeAnyOfFromValue returns a pointer to a valid InterfaceTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceTypeAnyOfFromValue(v string) (*InterfaceTypeAnyOf, error) {
	ev := InterfaceTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceTypeAnyOf: valid values are %v", v, AllowedInterfaceTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedInterfaceTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InterfaceType_anyOf value
func (v InterfaceTypeAnyOf) Ptr() *InterfaceTypeAnyOf {
	return &v
}

type NullableInterfaceTypeAnyOf struct {
	value *InterfaceTypeAnyOf
	isSet bool
}

func (v NullableInterfaceTypeAnyOf) Get() *InterfaceTypeAnyOf {
	return v.value
}

func (v *NullableInterfaceTypeAnyOf) Set(val *InterfaceTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceTypeAnyOf(val *InterfaceTypeAnyOf) *NullableInterfaceTypeAnyOf {
	return &NullableInterfaceTypeAnyOf{value: val, isSet: true}
}

func (v NullableInterfaceTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
