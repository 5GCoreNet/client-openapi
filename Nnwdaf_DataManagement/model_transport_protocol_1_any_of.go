/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// TransportProtocol1AnyOf the model 'TransportProtocol1AnyOf'
type TransportProtocol1AnyOf string

// List of TransportProtocol_1_anyOf
const (
	TCP TransportProtocol1AnyOf = "TCP"
)

// All allowed values of TransportProtocol1AnyOf enum
var AllowedTransportProtocol1AnyOfEnumValues = []TransportProtocol1AnyOf{
	"TCP",
}

func (v *TransportProtocol1AnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransportProtocol1AnyOf(value)
	for _, existing := range AllowedTransportProtocol1AnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransportProtocol1AnyOf", value)
}

// NewTransportProtocol1AnyOfFromValue returns a pointer to a valid TransportProtocol1AnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransportProtocol1AnyOfFromValue(v string) (*TransportProtocol1AnyOf, error) {
	ev := TransportProtocol1AnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransportProtocol1AnyOf: valid values are %v", v, AllowedTransportProtocol1AnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransportProtocol1AnyOf) IsValid() bool {
	for _, existing := range AllowedTransportProtocol1AnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransportProtocol_1_anyOf value
func (v TransportProtocol1AnyOf) Ptr() *TransportProtocol1AnyOf {
	return &v
}

type NullableTransportProtocol1AnyOf struct {
	value *TransportProtocol1AnyOf
	isSet bool
}

func (v NullableTransportProtocol1AnyOf) Get() *TransportProtocol1AnyOf {
	return v.value
}

func (v *NullableTransportProtocol1AnyOf) Set(val *TransportProtocol1AnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportProtocol1AnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportProtocol1AnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportProtocol1AnyOf(val *TransportProtocol1AnyOf) *NullableTransportProtocol1AnyOf {
	return &NullableTransportProtocol1AnyOf{value: val, isSet: true}
}

func (v NullableTransportProtocol1AnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportProtocol1AnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

