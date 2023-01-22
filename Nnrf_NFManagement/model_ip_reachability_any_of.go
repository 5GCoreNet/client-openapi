/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnrf_NFManagement

import (
	"encoding/json"
	"fmt"
)

// IpReachabilityAnyOf the model 'IpReachabilityAnyOf'
type IpReachabilityAnyOf string

// List of IpReachability_anyOf
const (
	IPV4 IpReachabilityAnyOf = "IPV4"
	IPV6 IpReachabilityAnyOf = "IPV6"
	IPV4_V6 IpReachabilityAnyOf = "IPV4V6"
)

// All allowed values of IpReachabilityAnyOf enum
var AllowedIpReachabilityAnyOfEnumValues = []IpReachabilityAnyOf{
	"IPV4",
	"IPV6",
	"IPV4V6",
}

func (v *IpReachabilityAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IpReachabilityAnyOf(value)
	for _, existing := range AllowedIpReachabilityAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IpReachabilityAnyOf", value)
}

// NewIpReachabilityAnyOfFromValue returns a pointer to a valid IpReachabilityAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIpReachabilityAnyOfFromValue(v string) (*IpReachabilityAnyOf, error) {
	ev := IpReachabilityAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IpReachabilityAnyOf: valid values are %v", v, AllowedIpReachabilityAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IpReachabilityAnyOf) IsValid() bool {
	for _, existing := range AllowedIpReachabilityAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IpReachability_anyOf value
func (v IpReachabilityAnyOf) Ptr() *IpReachabilityAnyOf {
	return &v
}

type NullableIpReachabilityAnyOf struct {
	value *IpReachabilityAnyOf
	isSet bool
}

func (v NullableIpReachabilityAnyOf) Get() *IpReachabilityAnyOf {
	return v.value
}

func (v *NullableIpReachabilityAnyOf) Set(val *IpReachabilityAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIpReachabilityAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIpReachabilityAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpReachabilityAnyOf(val *IpReachabilityAnyOf) *NullableIpReachabilityAnyOf {
	return &NullableIpReachabilityAnyOf{value: val, isSet: true}
}

func (v NullableIpReachabilityAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpReachabilityAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

