/*
3gpp-mbs-session

API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSSession

import (
	"encoding/json"
	"fmt"
)

// BroadcastDeliveryStatusAnyOf the model 'BroadcastDeliveryStatusAnyOf'
type BroadcastDeliveryStatusAnyOf string

// List of BroadcastDeliveryStatus_anyOf
const (
	ACTIVATED BroadcastDeliveryStatusAnyOf = "ACTIVATED"
	TERMINATED BroadcastDeliveryStatusAnyOf = "TERMINATED"
)

// All allowed values of BroadcastDeliveryStatusAnyOf enum
var AllowedBroadcastDeliveryStatusAnyOfEnumValues = []BroadcastDeliveryStatusAnyOf{
	"ACTIVATED",
	"TERMINATED",
}

func (v *BroadcastDeliveryStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BroadcastDeliveryStatusAnyOf(value)
	for _, existing := range AllowedBroadcastDeliveryStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BroadcastDeliveryStatusAnyOf", value)
}

// NewBroadcastDeliveryStatusAnyOfFromValue returns a pointer to a valid BroadcastDeliveryStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBroadcastDeliveryStatusAnyOfFromValue(v string) (*BroadcastDeliveryStatusAnyOf, error) {
	ev := BroadcastDeliveryStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BroadcastDeliveryStatusAnyOf: valid values are %v", v, AllowedBroadcastDeliveryStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BroadcastDeliveryStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedBroadcastDeliveryStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BroadcastDeliveryStatus_anyOf value
func (v BroadcastDeliveryStatusAnyOf) Ptr() *BroadcastDeliveryStatusAnyOf {
	return &v
}

type NullableBroadcastDeliveryStatusAnyOf struct {
	value *BroadcastDeliveryStatusAnyOf
	isSet bool
}

func (v NullableBroadcastDeliveryStatusAnyOf) Get() *BroadcastDeliveryStatusAnyOf {
	return v.value
}

func (v *NullableBroadcastDeliveryStatusAnyOf) Set(val *BroadcastDeliveryStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastDeliveryStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastDeliveryStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastDeliveryStatusAnyOf(val *BroadcastDeliveryStatusAnyOf) *NullableBroadcastDeliveryStatusAnyOf {
	return &NullableBroadcastDeliveryStatusAnyOf{value: val, isSet: true}
}

func (v NullableBroadcastDeliveryStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastDeliveryStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

