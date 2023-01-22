/*
MSGS_MSGDelivery

API for MSGG MSGin5G Server Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MSGS_MSGDelivery

import (
	"encoding/json"
	"fmt"
)

// DeliveryStatusAnyOf the model 'DeliveryStatusAnyOf'
type DeliveryStatusAnyOf string

// List of DeliveryStatus_anyOf
const (
	FAILED DeliveryStatusAnyOf = "DELY_FAILED"
	STORED DeliveryStatusAnyOf = "DELY_STORED"
)

// All allowed values of DeliveryStatusAnyOf enum
var AllowedDeliveryStatusAnyOfEnumValues = []DeliveryStatusAnyOf{
	"DELY_FAILED",
	"DELY_STORED",
}

func (v *DeliveryStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeliveryStatusAnyOf(value)
	for _, existing := range AllowedDeliveryStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeliveryStatusAnyOf", value)
}

// NewDeliveryStatusAnyOfFromValue returns a pointer to a valid DeliveryStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeliveryStatusAnyOfFromValue(v string) (*DeliveryStatusAnyOf, error) {
	ev := DeliveryStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeliveryStatusAnyOf: valid values are %v", v, AllowedDeliveryStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeliveryStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedDeliveryStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeliveryStatus_anyOf value
func (v DeliveryStatusAnyOf) Ptr() *DeliveryStatusAnyOf {
	return &v
}

type NullableDeliveryStatusAnyOf struct {
	value *DeliveryStatusAnyOf
	isSet bool
}

func (v NullableDeliveryStatusAnyOf) Get() *DeliveryStatusAnyOf {
	return v.value
}

func (v *NullableDeliveryStatusAnyOf) Set(val *DeliveryStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryStatusAnyOf(val *DeliveryStatusAnyOf) *NullableDeliveryStatusAnyOf {
	return &NullableDeliveryStatusAnyOf{value: val, isSet: true}
}

func (v NullableDeliveryStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

