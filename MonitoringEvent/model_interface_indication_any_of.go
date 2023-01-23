/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
	"fmt"
)

// InterfaceIndicationAnyOf the model 'InterfaceIndicationAnyOf'
type InterfaceIndicationAnyOf string

// List of InterfaceIndication_anyOf
const (
	EXPOSURE_FUNCTION InterfaceIndicationAnyOf = "EXPOSURE_FUNCTION"
	PDN_GATEWAY InterfaceIndicationAnyOf = "PDN_GATEWAY"
)

// All allowed values of InterfaceIndicationAnyOf enum
var AllowedInterfaceIndicationAnyOfEnumValues = []InterfaceIndicationAnyOf{
	"EXPOSURE_FUNCTION",
	"PDN_GATEWAY",
}

func (v *InterfaceIndicationAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceIndicationAnyOf(value)
	for _, existing := range AllowedInterfaceIndicationAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceIndicationAnyOf", value)
}

// NewInterfaceIndicationAnyOfFromValue returns a pointer to a valid InterfaceIndicationAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceIndicationAnyOfFromValue(v string) (*InterfaceIndicationAnyOf, error) {
	ev := InterfaceIndicationAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceIndicationAnyOf: valid values are %v", v, AllowedInterfaceIndicationAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceIndicationAnyOf) IsValid() bool {
	for _, existing := range AllowedInterfaceIndicationAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InterfaceIndication_anyOf value
func (v InterfaceIndicationAnyOf) Ptr() *InterfaceIndicationAnyOf {
	return &v
}

type NullableInterfaceIndicationAnyOf struct {
	value *InterfaceIndicationAnyOf
	isSet bool
}

func (v NullableInterfaceIndicationAnyOf) Get() *InterfaceIndicationAnyOf {
	return v.value
}

func (v *NullableInterfaceIndicationAnyOf) Set(val *InterfaceIndicationAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceIndicationAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceIndicationAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceIndicationAnyOf(val *InterfaceIndicationAnyOf) *NullableInterfaceIndicationAnyOf {
	return &NullableInterfaceIndicationAnyOf{value: val, isSet: true}
}

func (v NullableInterfaceIndicationAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceIndicationAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

