/*
VAE_ApplicationRequirement

API for VAE Application Requirement Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_VAE_ApplicationRequirement

import (
	"encoding/json"
	"fmt"
)

// ServiceLevelAnyOf the model 'ServiceLevelAnyOf'
type ServiceLevelAnyOf string

// List of ServiceLevel_anyOf
const (
	HIGH ServiceLevelAnyOf = "HIGH"
	MEDIUM ServiceLevelAnyOf = "MEDIUM"
	LOW ServiceLevelAnyOf = "LOW"
)

// All allowed values of ServiceLevelAnyOf enum
var AllowedServiceLevelAnyOfEnumValues = []ServiceLevelAnyOf{
	"HIGH",
	"MEDIUM",
	"LOW",
}

func (v *ServiceLevelAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceLevelAnyOf(value)
	for _, existing := range AllowedServiceLevelAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceLevelAnyOf", value)
}

// NewServiceLevelAnyOfFromValue returns a pointer to a valid ServiceLevelAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceLevelAnyOfFromValue(v string) (*ServiceLevelAnyOf, error) {
	ev := ServiceLevelAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceLevelAnyOf: valid values are %v", v, AllowedServiceLevelAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceLevelAnyOf) IsValid() bool {
	for _, existing := range AllowedServiceLevelAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceLevel_anyOf value
func (v ServiceLevelAnyOf) Ptr() *ServiceLevelAnyOf {
	return &v
}

type NullableServiceLevelAnyOf struct {
	value *ServiceLevelAnyOf
	isSet bool
}

func (v NullableServiceLevelAnyOf) Get() *ServiceLevelAnyOf {
	return v.value
}

func (v *NullableServiceLevelAnyOf) Set(val *ServiceLevelAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLevelAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceLevelAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLevelAnyOf(val *ServiceLevelAnyOf) *NullableServiceLevelAnyOf {
	return &NullableServiceLevelAnyOf{value: val, isSet: true}
}

func (v NullableServiceLevelAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLevelAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

