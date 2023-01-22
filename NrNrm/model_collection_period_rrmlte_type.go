/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NrNrm

import (
	"encoding/json"
	"fmt"
)

// CollectionPeriodRRMLTEType See details in 3GPP TS 32.422 clause 5.10.20.
type CollectionPeriodRRMLTEType string

// List of collectionPeriodRRMLTE-Type
const (
	_100MS CollectionPeriodRRMLTEType = "100ms"
	_1000MS CollectionPeriodRRMLTEType = "1000ms"
	_1024MS CollectionPeriodRRMLTEType = "1024ms"
	_1280MS CollectionPeriodRRMLTEType = "1280ms"
	_2048MS CollectionPeriodRRMLTEType = "2048ms"
	_2560MS CollectionPeriodRRMLTEType = "2560ms"
	_5120MS CollectionPeriodRRMLTEType = "5120ms"
	_10000MS CollectionPeriodRRMLTEType = "10000ms"
	_10240MS CollectionPeriodRRMLTEType = "10240ms"
	_60000MS CollectionPeriodRRMLTEType = "60000ms"
)

// All allowed values of CollectionPeriodRRMLTEType enum
var AllowedCollectionPeriodRRMLTETypeEnumValues = []CollectionPeriodRRMLTEType{
	"100ms",
	"1000ms",
	"1024ms",
	"1280ms",
	"2048ms",
	"2560ms",
	"5120ms",
	"10000ms",
	"10240ms",
	"60000ms",
}

func (v *CollectionPeriodRRMLTEType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionPeriodRRMLTEType(value)
	for _, existing := range AllowedCollectionPeriodRRMLTETypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionPeriodRRMLTEType", value)
}

// NewCollectionPeriodRRMLTETypeFromValue returns a pointer to a valid CollectionPeriodRRMLTEType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectionPeriodRRMLTETypeFromValue(v string) (*CollectionPeriodRRMLTEType, error) {
	ev := CollectionPeriodRRMLTEType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectionPeriodRRMLTEType: valid values are %v", v, AllowedCollectionPeriodRRMLTETypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectionPeriodRRMLTEType) IsValid() bool {
	for _, existing := range AllowedCollectionPeriodRRMLTETypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to collectionPeriodRRMLTE-Type value
func (v CollectionPeriodRRMLTEType) Ptr() *CollectionPeriodRRMLTEType {
	return &v
}

type NullableCollectionPeriodRRMLTEType struct {
	value *CollectionPeriodRRMLTEType
	isSet bool
}

func (v NullableCollectionPeriodRRMLTEType) Get() *CollectionPeriodRRMLTEType {
	return v.value
}

func (v *NullableCollectionPeriodRRMLTEType) Set(val *CollectionPeriodRRMLTEType) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPeriodRRMLTEType) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPeriodRRMLTEType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPeriodRRMLTEType(val *CollectionPeriodRRMLTEType) *NullableCollectionPeriodRRMLTEType {
	return &NullableCollectionPeriodRRMLTEType{value: val, isSet: true}
}

func (v NullableCollectionPeriodRRMLTEType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPeriodRRMLTEType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

