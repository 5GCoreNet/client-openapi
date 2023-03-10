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

// CollectionPeriodM6NRType See details in 3GPP TS 32.422 clause 5.10.34.
type CollectionPeriodM6NRType string

// List of collectionPeriodM6NR-Type
const (
	_120MS CollectionPeriodM6NRType = "120ms"
	_240MS CollectionPeriodM6NRType = "240ms"
	_480MS CollectionPeriodM6NRType = "480ms"
	_640MS CollectionPeriodM6NRType = "640ms"
	_1024MS CollectionPeriodM6NRType = "1024ms"
	_2048MS CollectionPeriodM6NRType = "2048ms"
	_5120MS CollectionPeriodM6NRType = "5120ms"
	_10240MS CollectionPeriodM6NRType = "10240ms"
	_20480MS CollectionPeriodM6NRType = "20480ms"
	_40960MS CollectionPeriodM6NRType = "40960ms"
	_1MIN CollectionPeriodM6NRType = "1min"
	_6MIN CollectionPeriodM6NRType = "6min"
	_12MIN CollectionPeriodM6NRType = "12min"
	_30MIN CollectionPeriodM6NRType = "30min"
)

// All allowed values of CollectionPeriodM6NRType enum
var AllowedCollectionPeriodM6NRTypeEnumValues = []CollectionPeriodM6NRType{
	"120ms",
	"240ms",
	"480ms",
	"640ms",
	"1024ms",
	"2048ms",
	"5120ms",
	"10240ms",
	"20480ms",
	"40960ms",
	"1min",
	"6min",
	"12min",
	"30min",
}

func (v *CollectionPeriodM6NRType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionPeriodM6NRType(value)
	for _, existing := range AllowedCollectionPeriodM6NRTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionPeriodM6NRType", value)
}

// NewCollectionPeriodM6NRTypeFromValue returns a pointer to a valid CollectionPeriodM6NRType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectionPeriodM6NRTypeFromValue(v string) (*CollectionPeriodM6NRType, error) {
	ev := CollectionPeriodM6NRType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectionPeriodM6NRType: valid values are %v", v, AllowedCollectionPeriodM6NRTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectionPeriodM6NRType) IsValid() bool {
	for _, existing := range AllowedCollectionPeriodM6NRTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to collectionPeriodM6NR-Type value
func (v CollectionPeriodM6NRType) Ptr() *CollectionPeriodM6NRType {
	return &v
}

type NullableCollectionPeriodM6NRType struct {
	value *CollectionPeriodM6NRType
	isSet bool
}

func (v NullableCollectionPeriodM6NRType) Get() *CollectionPeriodM6NRType {
	return v.value
}

func (v *NullableCollectionPeriodM6NRType) Set(val *CollectionPeriodM6NRType) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPeriodM6NRType) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPeriodM6NRType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPeriodM6NRType(val *CollectionPeriodM6NRType) *NullableCollectionPeriodM6NRType {
	return &NullableCollectionPeriodM6NRType{value: val, isSet: true}
}

func (v NullableCollectionPeriodM6NRType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPeriodM6NRType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

