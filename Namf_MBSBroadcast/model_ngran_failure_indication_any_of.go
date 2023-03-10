/*
Namf_MBSBroadcast

AMF MBSBroadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MBSBroadcast

import (
	"encoding/json"
	"fmt"
)

// NgranFailureIndicationAnyOf the model 'NgranFailureIndicationAnyOf'
type NgranFailureIndicationAnyOf string

// List of NgranFailureIndication_anyOf
const (
	RESTART_OR_START NgranFailureIndicationAnyOf = "NG_RAN_RESTART_OR_START"
	FAILURE_WITHOUT_RESTART NgranFailureIndicationAnyOf = "NG_RAN_FAILURE_WITHOUT_RESTART"
	NOT_REACHABLE NgranFailureIndicationAnyOf = "NG_RAN_NOT_REACHABLE"
)

// All allowed values of NgranFailureIndicationAnyOf enum
var AllowedNgranFailureIndicationAnyOfEnumValues = []NgranFailureIndicationAnyOf{
	"NG_RAN_RESTART_OR_START",
	"NG_RAN_FAILURE_WITHOUT_RESTART",
	"NG_RAN_NOT_REACHABLE",
}

func (v *NgranFailureIndicationAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NgranFailureIndicationAnyOf(value)
	for _, existing := range AllowedNgranFailureIndicationAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NgranFailureIndicationAnyOf", value)
}

// NewNgranFailureIndicationAnyOfFromValue returns a pointer to a valid NgranFailureIndicationAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNgranFailureIndicationAnyOfFromValue(v string) (*NgranFailureIndicationAnyOf, error) {
	ev := NgranFailureIndicationAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NgranFailureIndicationAnyOf: valid values are %v", v, AllowedNgranFailureIndicationAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NgranFailureIndicationAnyOf) IsValid() bool {
	for _, existing := range AllowedNgranFailureIndicationAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NgranFailureIndication_anyOf value
func (v NgranFailureIndicationAnyOf) Ptr() *NgranFailureIndicationAnyOf {
	return &v
}

type NullableNgranFailureIndicationAnyOf struct {
	value *NgranFailureIndicationAnyOf
	isSet bool
}

func (v NullableNgranFailureIndicationAnyOf) Get() *NgranFailureIndicationAnyOf {
	return v.value
}

func (v *NullableNgranFailureIndicationAnyOf) Set(val *NgranFailureIndicationAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNgranFailureIndicationAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNgranFailureIndicationAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNgranFailureIndicationAnyOf(val *NgranFailureIndicationAnyOf) *NullableNgranFailureIndicationAnyOf {
	return &NullableNgranFailureIndicationAnyOf{value: val, isSet: true}
}

func (v NullableNgranFailureIndicationAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNgranFailureIndicationAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

