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

// TADIdentifierAnyOf the model 'TADIdentifierAnyOf'
type TADIdentifierAnyOf string

// List of TADIdentifier_anyOf
const (
	CS TADIdentifierAnyOf = "CS"
	PS TADIdentifierAnyOf = "PS"
)

// All allowed values of TADIdentifierAnyOf enum
var AllowedTADIdentifierAnyOfEnumValues = []TADIdentifierAnyOf{
	"CS",
	"PS",
}

func (v *TADIdentifierAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TADIdentifierAnyOf(value)
	for _, existing := range AllowedTADIdentifierAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TADIdentifierAnyOf", value)
}

// NewTADIdentifierAnyOfFromValue returns a pointer to a valid TADIdentifierAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTADIdentifierAnyOfFromValue(v string) (*TADIdentifierAnyOf, error) {
	ev := TADIdentifierAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TADIdentifierAnyOf: valid values are %v", v, AllowedTADIdentifierAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TADIdentifierAnyOf) IsValid() bool {
	for _, existing := range AllowedTADIdentifierAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TADIdentifier_anyOf value
func (v TADIdentifierAnyOf) Ptr() *TADIdentifierAnyOf {
	return &v
}

type NullableTADIdentifierAnyOf struct {
	value *TADIdentifierAnyOf
	isSet bool
}

func (v NullableTADIdentifierAnyOf) Get() *TADIdentifierAnyOf {
	return v.value
}

func (v *NullableTADIdentifierAnyOf) Set(val *TADIdentifierAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTADIdentifierAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTADIdentifierAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTADIdentifierAnyOf(val *TADIdentifierAnyOf) *NullableTADIdentifierAnyOf {
	return &NullableTADIdentifierAnyOf{value: val, isSet: true}
}

func (v NullableTADIdentifierAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTADIdentifierAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

