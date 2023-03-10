/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// OneTimeEventTypeAnyOf the model 'OneTimeEventTypeAnyOf'
type OneTimeEventTypeAnyOf string

// List of oneTimeEventType_anyOf
const (
	IEC OneTimeEventTypeAnyOf = "IEC"
	PEC OneTimeEventTypeAnyOf = "PEC"
)

// All allowed values of OneTimeEventTypeAnyOf enum
var AllowedOneTimeEventTypeAnyOfEnumValues = []OneTimeEventTypeAnyOf{
	"IEC",
	"PEC",
}

func (v *OneTimeEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OneTimeEventTypeAnyOf(value)
	for _, existing := range AllowedOneTimeEventTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OneTimeEventTypeAnyOf", value)
}

// NewOneTimeEventTypeAnyOfFromValue returns a pointer to a valid OneTimeEventTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOneTimeEventTypeAnyOfFromValue(v string) (*OneTimeEventTypeAnyOf, error) {
	ev := OneTimeEventTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OneTimeEventTypeAnyOf: valid values are %v", v, AllowedOneTimeEventTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OneTimeEventTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedOneTimeEventTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to oneTimeEventType_anyOf value
func (v OneTimeEventTypeAnyOf) Ptr() *OneTimeEventTypeAnyOf {
	return &v
}

type NullableOneTimeEventTypeAnyOf struct {
	value *OneTimeEventTypeAnyOf
	isSet bool
}

func (v NullableOneTimeEventTypeAnyOf) Get() *OneTimeEventTypeAnyOf {
	return v.value
}

func (v *NullableOneTimeEventTypeAnyOf) Set(val *OneTimeEventTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOneTimeEventTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOneTimeEventTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneTimeEventTypeAnyOf(val *OneTimeEventTypeAnyOf) *NullableOneTimeEventTypeAnyOf {
	return &NullableOneTimeEventTypeAnyOf{value: val, isSet: true}
}

func (v NullableOneTimeEventTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOneTimeEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

