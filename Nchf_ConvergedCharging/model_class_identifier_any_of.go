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

// ClassIdentifierAnyOf the model 'ClassIdentifierAnyOf'
type ClassIdentifierAnyOf string

// List of ClassIdentifier_anyOf
const (
	PERSONAL ClassIdentifierAnyOf = "PERSONAL"
	ADVERTISEMENT ClassIdentifierAnyOf = "ADVERTISEMENT"
	INFORMATIONAL ClassIdentifierAnyOf = "INFORMATIONAL"
	AUTO ClassIdentifierAnyOf = "AUTO"
)

// All allowed values of ClassIdentifierAnyOf enum
var AllowedClassIdentifierAnyOfEnumValues = []ClassIdentifierAnyOf{
	"PERSONAL",
	"ADVERTISEMENT",
	"INFORMATIONAL",
	"AUTO",
}

func (v *ClassIdentifierAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClassIdentifierAnyOf(value)
	for _, existing := range AllowedClassIdentifierAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClassIdentifierAnyOf", value)
}

// NewClassIdentifierAnyOfFromValue returns a pointer to a valid ClassIdentifierAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClassIdentifierAnyOfFromValue(v string) (*ClassIdentifierAnyOf, error) {
	ev := ClassIdentifierAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClassIdentifierAnyOf: valid values are %v", v, AllowedClassIdentifierAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClassIdentifierAnyOf) IsValid() bool {
	for _, existing := range AllowedClassIdentifierAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClassIdentifier_anyOf value
func (v ClassIdentifierAnyOf) Ptr() *ClassIdentifierAnyOf {
	return &v
}

type NullableClassIdentifierAnyOf struct {
	value *ClassIdentifierAnyOf
	isSet bool
}

func (v NullableClassIdentifierAnyOf) Get() *ClassIdentifierAnyOf {
	return v.value
}

func (v *NullableClassIdentifierAnyOf) Set(val *ClassIdentifierAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableClassIdentifierAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableClassIdentifierAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassIdentifierAnyOf(val *ClassIdentifierAnyOf) *NullableClassIdentifierAnyOf {
	return &NullableClassIdentifierAnyOf{value: val, isSet: true}
}

func (v NullableClassIdentifierAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassIdentifierAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

