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

// APIDirectionAnyOf the model 'APIDirectionAnyOf'
type APIDirectionAnyOf string

// List of APIDirection_anyOf
const (
	INVOCATION APIDirectionAnyOf = "INVOCATION"
	NOTIFICATION APIDirectionAnyOf = "NOTIFICATION"
)

// All allowed values of APIDirectionAnyOf enum
var AllowedAPIDirectionAnyOfEnumValues = []APIDirectionAnyOf{
	"INVOCATION",
	"NOTIFICATION",
}

func (v *APIDirectionAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := APIDirectionAnyOf(value)
	for _, existing := range AllowedAPIDirectionAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid APIDirectionAnyOf", value)
}

// NewAPIDirectionAnyOfFromValue returns a pointer to a valid APIDirectionAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAPIDirectionAnyOfFromValue(v string) (*APIDirectionAnyOf, error) {
	ev := APIDirectionAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for APIDirectionAnyOf: valid values are %v", v, AllowedAPIDirectionAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v APIDirectionAnyOf) IsValid() bool {
	for _, existing := range AllowedAPIDirectionAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to APIDirection_anyOf value
func (v APIDirectionAnyOf) Ptr() *APIDirectionAnyOf {
	return &v
}

type NullableAPIDirectionAnyOf struct {
	value *APIDirectionAnyOf
	isSet bool
}

func (v NullableAPIDirectionAnyOf) Get() *APIDirectionAnyOf {
	return v.value
}

func (v *NullableAPIDirectionAnyOf) Set(val *APIDirectionAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIDirectionAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIDirectionAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIDirectionAnyOf(val *APIDirectionAnyOf) *NullableAPIDirectionAnyOf {
	return &NullableAPIDirectionAnyOf{value: val, isSet: true}
}

func (v NullableAPIDirectionAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIDirectionAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

