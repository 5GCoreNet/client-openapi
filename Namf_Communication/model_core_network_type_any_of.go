/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// CoreNetworkTypeAnyOf the model 'CoreNetworkTypeAnyOf'
type CoreNetworkTypeAnyOf string

// List of CoreNetworkType_anyOf
const (
	_5_GC CoreNetworkTypeAnyOf = "5GC"
	EPC CoreNetworkTypeAnyOf = "EPC"
)

// All allowed values of CoreNetworkTypeAnyOf enum
var AllowedCoreNetworkTypeAnyOfEnumValues = []CoreNetworkTypeAnyOf{
	"5GC",
	"EPC",
}

func (v *CoreNetworkTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CoreNetworkTypeAnyOf(value)
	for _, existing := range AllowedCoreNetworkTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CoreNetworkTypeAnyOf", value)
}

// NewCoreNetworkTypeAnyOfFromValue returns a pointer to a valid CoreNetworkTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCoreNetworkTypeAnyOfFromValue(v string) (*CoreNetworkTypeAnyOf, error) {
	ev := CoreNetworkTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CoreNetworkTypeAnyOf: valid values are %v", v, AllowedCoreNetworkTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CoreNetworkTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedCoreNetworkTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CoreNetworkType_anyOf value
func (v CoreNetworkTypeAnyOf) Ptr() *CoreNetworkTypeAnyOf {
	return &v
}

type NullableCoreNetworkTypeAnyOf struct {
	value *CoreNetworkTypeAnyOf
	isSet bool
}

func (v NullableCoreNetworkTypeAnyOf) Get() *CoreNetworkTypeAnyOf {
	return v.value
}

func (v *NullableCoreNetworkTypeAnyOf) Set(val *CoreNetworkTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreNetworkTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreNetworkTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreNetworkTypeAnyOf(val *CoreNetworkTypeAnyOf) *NullableCoreNetworkTypeAnyOf {
	return &NullableCoreNetworkTypeAnyOf{value: val, isSet: true}
}

func (v NullableCoreNetworkTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreNetworkTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

