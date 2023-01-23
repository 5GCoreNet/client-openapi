/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// ServiceExperienceTypeAnyOf the model 'ServiceExperienceTypeAnyOf'
type ServiceExperienceTypeAnyOf string

// List of ServiceExperienceType_anyOf
const (
	VOICE ServiceExperienceTypeAnyOf = "VOICE"
	VIDEO ServiceExperienceTypeAnyOf = "VIDEO"
	OTHER ServiceExperienceTypeAnyOf = "OTHER"
)

// All allowed values of ServiceExperienceTypeAnyOf enum
var AllowedServiceExperienceTypeAnyOfEnumValues = []ServiceExperienceTypeAnyOf{
	"VOICE",
	"VIDEO",
	"OTHER",
}

func (v *ServiceExperienceTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceExperienceTypeAnyOf(value)
	for _, existing := range AllowedServiceExperienceTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceExperienceTypeAnyOf", value)
}

// NewServiceExperienceTypeAnyOfFromValue returns a pointer to a valid ServiceExperienceTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceExperienceTypeAnyOfFromValue(v string) (*ServiceExperienceTypeAnyOf, error) {
	ev := ServiceExperienceTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceExperienceTypeAnyOf: valid values are %v", v, AllowedServiceExperienceTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceExperienceTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedServiceExperienceTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceExperienceType_anyOf value
func (v ServiceExperienceTypeAnyOf) Ptr() *ServiceExperienceTypeAnyOf {
	return &v
}

type NullableServiceExperienceTypeAnyOf struct {
	value *ServiceExperienceTypeAnyOf
	isSet bool
}

func (v NullableServiceExperienceTypeAnyOf) Get() *ServiceExperienceTypeAnyOf {
	return v.value
}

func (v *NullableServiceExperienceTypeAnyOf) Set(val *ServiceExperienceTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceExperienceTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceExperienceTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceExperienceTypeAnyOf(val *ServiceExperienceTypeAnyOf) *NullableServiceExperienceTypeAnyOf {
	return &NullableServiceExperienceTypeAnyOf{value: val, isSet: true}
}

func (v NullableServiceExperienceTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceExperienceTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

