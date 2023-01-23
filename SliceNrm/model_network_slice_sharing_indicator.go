/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
	"fmt"
)

// NetworkSliceSharingIndicator the model 'NetworkSliceSharingIndicator'
type NetworkSliceSharingIndicator string

// List of NetworkSliceSharingIndicator
const (
	SHARED NetworkSliceSharingIndicator = "SHARED"
	NON_SHARED NetworkSliceSharingIndicator = "NON-SHARED"
)

// All allowed values of NetworkSliceSharingIndicator enum
var AllowedNetworkSliceSharingIndicatorEnumValues = []NetworkSliceSharingIndicator{
	"SHARED",
	"NON-SHARED",
}

func (v *NetworkSliceSharingIndicator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkSliceSharingIndicator(value)
	for _, existing := range AllowedNetworkSliceSharingIndicatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkSliceSharingIndicator", value)
}

// NewNetworkSliceSharingIndicatorFromValue returns a pointer to a valid NetworkSliceSharingIndicator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkSliceSharingIndicatorFromValue(v string) (*NetworkSliceSharingIndicator, error) {
	ev := NetworkSliceSharingIndicator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkSliceSharingIndicator: valid values are %v", v, AllowedNetworkSliceSharingIndicatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkSliceSharingIndicator) IsValid() bool {
	for _, existing := range AllowedNetworkSliceSharingIndicatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkSliceSharingIndicator value
func (v NetworkSliceSharingIndicator) Ptr() *NetworkSliceSharingIndicator {
	return &v
}

type NullableNetworkSliceSharingIndicator struct {
	value *NetworkSliceSharingIndicator
	isSet bool
}

func (v NullableNetworkSliceSharingIndicator) Get() *NetworkSliceSharingIndicator {
	return v.value
}

func (v *NullableNetworkSliceSharingIndicator) Set(val *NetworkSliceSharingIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSliceSharingIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSliceSharingIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSliceSharingIndicator(val *NetworkSliceSharingIndicator) *NullableNetworkSliceSharingIndicator {
	return &NullableNetworkSliceSharingIndicator{value: val, isSet: true}
}

func (v NullableNetworkSliceSharingIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSliceSharingIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

