/*
3gpp-mbs-ud-ingest

API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSUserDataIngestSession

import (
	"encoding/json"
	"fmt"
)

// DistributionMethodOneOf the model 'DistributionMethodOneOf'
type DistributionMethodOneOf string

// List of DistributionMethod_oneOf
const (
	OBJECT DistributionMethodOneOf = "OBJECT"
	PACKET DistributionMethodOneOf = "PACKET"
)

// All allowed values of DistributionMethodOneOf enum
var AllowedDistributionMethodOneOfEnumValues = []DistributionMethodOneOf{
	"OBJECT",
	"PACKET",
}

func (v *DistributionMethodOneOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DistributionMethodOneOf(value)
	for _, existing := range AllowedDistributionMethodOneOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DistributionMethodOneOf", value)
}

// NewDistributionMethodOneOfFromValue returns a pointer to a valid DistributionMethodOneOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDistributionMethodOneOfFromValue(v string) (*DistributionMethodOneOf, error) {
	ev := DistributionMethodOneOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DistributionMethodOneOf: valid values are %v", v, AllowedDistributionMethodOneOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DistributionMethodOneOf) IsValid() bool {
	for _, existing := range AllowedDistributionMethodOneOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DistributionMethod_oneOf value
func (v DistributionMethodOneOf) Ptr() *DistributionMethodOneOf {
	return &v
}

type NullableDistributionMethodOneOf struct {
	value *DistributionMethodOneOf
	isSet bool
}

func (v NullableDistributionMethodOneOf) Get() *DistributionMethodOneOf {
	return v.value
}

func (v *NullableDistributionMethodOneOf) Set(val *DistributionMethodOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionMethodOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionMethodOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionMethodOneOf(val *DistributionMethodOneOf) *NullableDistributionMethodOneOf {
	return &NullableDistributionMethodOneOf{value: val, isSet: true}
}

func (v NullableDistributionMethodOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionMethodOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

