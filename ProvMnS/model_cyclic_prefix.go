/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"fmt"
)

// CyclicPrefix the model 'CyclicPrefix'
type CyclicPrefix string

// List of CyclicPrefix
const (
	_15 CyclicPrefix = "15"
	_30 CyclicPrefix = "30"
	_60 CyclicPrefix = "60"
	_120 CyclicPrefix = "120"
)

// All allowed values of CyclicPrefix enum
var AllowedCyclicPrefixEnumValues = []CyclicPrefix{
	"15",
	"30",
	"60",
	"120",
}

func (v *CyclicPrefix) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CyclicPrefix(value)
	for _, existing := range AllowedCyclicPrefixEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CyclicPrefix", value)
}

// NewCyclicPrefixFromValue returns a pointer to a valid CyclicPrefix
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCyclicPrefixFromValue(v string) (*CyclicPrefix, error) {
	ev := CyclicPrefix(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CyclicPrefix: valid values are %v", v, AllowedCyclicPrefixEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CyclicPrefix) IsValid() bool {
	for _, existing := range AllowedCyclicPrefixEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CyclicPrefix value
func (v CyclicPrefix) Ptr() *CyclicPrefix {
	return &v
}

type NullableCyclicPrefix struct {
	value *CyclicPrefix
	isSet bool
}

func (v NullableCyclicPrefix) Get() *CyclicPrefix {
	return v.value
}

func (v *NullableCyclicPrefix) Set(val *CyclicPrefix) {
	v.value = val
	v.isSet = true
}

func (v NullableCyclicPrefix) IsSet() bool {
	return v.isSet
}

func (v *NullableCyclicPrefix) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCyclicPrefix(val *CyclicPrefix) *NullableCyclicPrefix {
	return &NullableCyclicPrefix{value: val, isSet: true}
}

func (v NullableCyclicPrefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCyclicPrefix) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

