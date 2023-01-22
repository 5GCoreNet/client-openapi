/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package 5GcNrm

import (
	"encoding/json"
	"fmt"
)

// CollectionPeriodRRMNRType See details in 3GPP TS 32.422 clause 5.10.30.
type CollectionPeriodRRMNRType string

// List of collectionPeriodRRMNR-Type
const (
	_1024MS CollectionPeriodRRMNRType = "1024ms"
	_2048MS CollectionPeriodRRMNRType = "2048ms"
	_5120MS CollectionPeriodRRMNRType = "5120ms"
	_10240MS CollectionPeriodRRMNRType = "10240ms"
	_60000MS CollectionPeriodRRMNRType = "60000ms"
)

// All allowed values of CollectionPeriodRRMNRType enum
var AllowedCollectionPeriodRRMNRTypeEnumValues = []CollectionPeriodRRMNRType{
	"1024ms",
	"2048ms",
	"5120ms",
	"10240ms",
	"60000ms",
}

func (v *CollectionPeriodRRMNRType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionPeriodRRMNRType(value)
	for _, existing := range AllowedCollectionPeriodRRMNRTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionPeriodRRMNRType", value)
}

// NewCollectionPeriodRRMNRTypeFromValue returns a pointer to a valid CollectionPeriodRRMNRType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectionPeriodRRMNRTypeFromValue(v string) (*CollectionPeriodRRMNRType, error) {
	ev := CollectionPeriodRRMNRType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectionPeriodRRMNRType: valid values are %v", v, AllowedCollectionPeriodRRMNRTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectionPeriodRRMNRType) IsValid() bool {
	for _, existing := range AllowedCollectionPeriodRRMNRTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to collectionPeriodRRMNR-Type value
func (v CollectionPeriodRRMNRType) Ptr() *CollectionPeriodRRMNRType {
	return &v
}

type NullableCollectionPeriodRRMNRType struct {
	value *CollectionPeriodRRMNRType
	isSet bool
}

func (v NullableCollectionPeriodRRMNRType) Get() *CollectionPeriodRRMNRType {
	return v.value
}

func (v *NullableCollectionPeriodRRMNRType) Set(val *CollectionPeriodRRMNRType) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPeriodRRMNRType) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPeriodRRMNRType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPeriodRRMNRType(val *CollectionPeriodRRMNRType) *NullableCollectionPeriodRRMNRType {
	return &NullableCollectionPeriodRRMNRType{value: val, isSet: true}
}

func (v NullableCollectionPeriodRRMNRType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPeriodRRMNRType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

