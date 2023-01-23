/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
	"fmt"
)

// CommModelType any of enumrated value
type CommModelType string

// List of CommModelType
const (
	DIRECT_COMMUNICATION_WO_NRF CommModelType = "DIRECT_COMMUNICATION_WO_NRF"
	DIRECT_COMMUNICATION_WITH_NRF CommModelType = "DIRECT_COMMUNICATION_WITH_NRF"
	INDIRECT_COMMUNICATION_WO_DEDICATED_DISCOVERY CommModelType = "INDIRECT_COMMUNICATION_WO_DEDICATED_DISCOVERY"
	INDIRECT_COMMUNICATION_WITH_DEDICATED_DISCOVERY CommModelType = "INDIRECT_COMMUNICATION_WITH_DEDICATED_DISCOVERY"
)

// All allowed values of CommModelType enum
var AllowedCommModelTypeEnumValues = []CommModelType{
	"DIRECT_COMMUNICATION_WO_NRF",
	"DIRECT_COMMUNICATION_WITH_NRF",
	"INDIRECT_COMMUNICATION_WO_DEDICATED_DISCOVERY",
	"INDIRECT_COMMUNICATION_WITH_DEDICATED_DISCOVERY",
}

func (v *CommModelType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CommModelType(value)
	for _, existing := range AllowedCommModelTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CommModelType", value)
}

// NewCommModelTypeFromValue returns a pointer to a valid CommModelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCommModelTypeFromValue(v string) (*CommModelType, error) {
	ev := CommModelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CommModelType: valid values are %v", v, AllowedCommModelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CommModelType) IsValid() bool {
	for _, existing := range AllowedCommModelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommModelType value
func (v CommModelType) Ptr() *CommModelType {
	return &v
}

type NullableCommModelType struct {
	value *CommModelType
	isSet bool
}

func (v NullableCommModelType) Get() *CommModelType {
	return v.value
}

func (v *NullableCommModelType) Set(val *CommModelType) {
	v.value = val
	v.isSet = true
}

func (v NullableCommModelType) IsSet() bool {
	return v.isSet
}

func (v *NullableCommModelType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommModelType(val *CommModelType) *NullableCommModelType {
	return &NullableCommModelType{value: val, isSet: true}
}

func (v NullableCommModelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommModelType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

