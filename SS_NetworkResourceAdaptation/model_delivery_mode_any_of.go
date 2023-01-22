/*
SS_NetworkResourceAdaptation

SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SS_NetworkResourceAdaptation

import (
	"encoding/json"
	"fmt"
)

// DeliveryModeAnyOf the model 'DeliveryModeAnyOf'
type DeliveryModeAnyOf string

// List of DeliveryMode_anyOf
const (
	UNICAST DeliveryModeAnyOf = "UNICAST"
	MULTICAST DeliveryModeAnyOf = "MULTICAST"
)

// All allowed values of DeliveryModeAnyOf enum
var AllowedDeliveryModeAnyOfEnumValues = []DeliveryModeAnyOf{
	"UNICAST",
	"MULTICAST",
}

func (v *DeliveryModeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeliveryModeAnyOf(value)
	for _, existing := range AllowedDeliveryModeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeliveryModeAnyOf", value)
}

// NewDeliveryModeAnyOfFromValue returns a pointer to a valid DeliveryModeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeliveryModeAnyOfFromValue(v string) (*DeliveryModeAnyOf, error) {
	ev := DeliveryModeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeliveryModeAnyOf: valid values are %v", v, AllowedDeliveryModeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeliveryModeAnyOf) IsValid() bool {
	for _, existing := range AllowedDeliveryModeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeliveryMode_anyOf value
func (v DeliveryModeAnyOf) Ptr() *DeliveryModeAnyOf {
	return &v
}

type NullableDeliveryModeAnyOf struct {
	value *DeliveryModeAnyOf
	isSet bool
}

func (v NullableDeliveryModeAnyOf) Get() *DeliveryModeAnyOf {
	return v.value
}

func (v *NullableDeliveryModeAnyOf) Set(val *DeliveryModeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryModeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryModeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryModeAnyOf(val *DeliveryModeAnyOf) *NullableDeliveryModeAnyOf {
	return &NullableDeliveryModeAnyOf{value: val, isSet: true}
}

func (v NullableDeliveryModeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryModeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

