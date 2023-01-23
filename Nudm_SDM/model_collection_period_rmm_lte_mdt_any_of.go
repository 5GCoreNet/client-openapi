/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
	"fmt"
)

// CollectionPeriodRmmLteMdtAnyOf the model 'CollectionPeriodRmmLteMdtAnyOf'
type CollectionPeriodRmmLteMdtAnyOf string

// List of CollectionPeriodRmmLteMdt_anyOf
const (
	_1024 CollectionPeriodRmmLteMdtAnyOf = "1024"
	_1280 CollectionPeriodRmmLteMdtAnyOf = "1280"
	_2048 CollectionPeriodRmmLteMdtAnyOf = "2048"
	_2560 CollectionPeriodRmmLteMdtAnyOf = "2560"
	_5120 CollectionPeriodRmmLteMdtAnyOf = "5120"
	_10240 CollectionPeriodRmmLteMdtAnyOf = "10240"
	_60000 CollectionPeriodRmmLteMdtAnyOf = "60000"
)

// All allowed values of CollectionPeriodRmmLteMdtAnyOf enum
var AllowedCollectionPeriodRmmLteMdtAnyOfEnumValues = []CollectionPeriodRmmLteMdtAnyOf{
	"1024",
	"1280",
	"2048",
	"2560",
	"5120",
	"10240",
	"60000",
}

func (v *CollectionPeriodRmmLteMdtAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionPeriodRmmLteMdtAnyOf(value)
	for _, existing := range AllowedCollectionPeriodRmmLteMdtAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionPeriodRmmLteMdtAnyOf", value)
}

// NewCollectionPeriodRmmLteMdtAnyOfFromValue returns a pointer to a valid CollectionPeriodRmmLteMdtAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectionPeriodRmmLteMdtAnyOfFromValue(v string) (*CollectionPeriodRmmLteMdtAnyOf, error) {
	ev := CollectionPeriodRmmLteMdtAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectionPeriodRmmLteMdtAnyOf: valid values are %v", v, AllowedCollectionPeriodRmmLteMdtAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectionPeriodRmmLteMdtAnyOf) IsValid() bool {
	for _, existing := range AllowedCollectionPeriodRmmLteMdtAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CollectionPeriodRmmLteMdt_anyOf value
func (v CollectionPeriodRmmLteMdtAnyOf) Ptr() *CollectionPeriodRmmLteMdtAnyOf {
	return &v
}

type NullableCollectionPeriodRmmLteMdtAnyOf struct {
	value *CollectionPeriodRmmLteMdtAnyOf
	isSet bool
}

func (v NullableCollectionPeriodRmmLteMdtAnyOf) Get() *CollectionPeriodRmmLteMdtAnyOf {
	return v.value
}

func (v *NullableCollectionPeriodRmmLteMdtAnyOf) Set(val *CollectionPeriodRmmLteMdtAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPeriodRmmLteMdtAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPeriodRmmLteMdtAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPeriodRmmLteMdtAnyOf(val *CollectionPeriodRmmLteMdtAnyOf) *NullableCollectionPeriodRmmLteMdtAnyOf {
	return &NullableCollectionPeriodRmmLteMdtAnyOf{value: val, isSet: true}
}

func (v NullableCollectionPeriodRmmLteMdtAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPeriodRmmLteMdtAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

