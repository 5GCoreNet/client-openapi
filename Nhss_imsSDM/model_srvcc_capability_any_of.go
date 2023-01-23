/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// SrvccCapabilityAnyOf the model 'SrvccCapabilityAnyOf'
type SrvccCapabilityAnyOf string

// List of SrvccCapability_anyOf
const (
	_4_G_SRVCC_CAPABLE SrvccCapabilityAnyOf = "UE_4G_SRVCC_CAPABLE"
	_5_G_SRVCC_CAPABLE SrvccCapabilityAnyOf = "UE_5G_SRVCC_CAPABLE"
)

// All allowed values of SrvccCapabilityAnyOf enum
var AllowedSrvccCapabilityAnyOfEnumValues = []SrvccCapabilityAnyOf{
	"UE_4G_SRVCC_CAPABLE",
	"UE_5G_SRVCC_CAPABLE",
}

func (v *SrvccCapabilityAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SrvccCapabilityAnyOf(value)
	for _, existing := range AllowedSrvccCapabilityAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SrvccCapabilityAnyOf", value)
}

// NewSrvccCapabilityAnyOfFromValue returns a pointer to a valid SrvccCapabilityAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSrvccCapabilityAnyOfFromValue(v string) (*SrvccCapabilityAnyOf, error) {
	ev := SrvccCapabilityAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SrvccCapabilityAnyOf: valid values are %v", v, AllowedSrvccCapabilityAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SrvccCapabilityAnyOf) IsValid() bool {
	for _, existing := range AllowedSrvccCapabilityAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SrvccCapability_anyOf value
func (v SrvccCapabilityAnyOf) Ptr() *SrvccCapabilityAnyOf {
	return &v
}

type NullableSrvccCapabilityAnyOf struct {
	value *SrvccCapabilityAnyOf
	isSet bool
}

func (v NullableSrvccCapabilityAnyOf) Get() *SrvccCapabilityAnyOf {
	return v.value
}

func (v *NullableSrvccCapabilityAnyOf) Set(val *SrvccCapabilityAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSrvccCapabilityAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSrvccCapabilityAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSrvccCapabilityAnyOf(val *SrvccCapabilityAnyOf) *NullableSrvccCapabilityAnyOf {
	return &NullableSrvccCapabilityAnyOf{value: val, isSet: true}
}

func (v NullableSrvccCapabilityAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSrvccCapabilityAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

