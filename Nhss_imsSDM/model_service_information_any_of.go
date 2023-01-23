/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// ServiceInformationAnyOf the model 'ServiceInformationAnyOf'
type ServiceInformationAnyOf string

// List of ServiceInformation_anyOf
const (
	REQUEST ServiceInformationAnyOf = "INCLUDE_REGISTER_REQUEST"
	RESPONSE ServiceInformationAnyOf = "INCLUDE_REGISTER_RESPONSE"
)

// All allowed values of ServiceInformationAnyOf enum
var AllowedServiceInformationAnyOfEnumValues = []ServiceInformationAnyOf{
	"INCLUDE_REGISTER_REQUEST",
	"INCLUDE_REGISTER_RESPONSE",
}

func (v *ServiceInformationAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceInformationAnyOf(value)
	for _, existing := range AllowedServiceInformationAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceInformationAnyOf", value)
}

// NewServiceInformationAnyOfFromValue returns a pointer to a valid ServiceInformationAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceInformationAnyOfFromValue(v string) (*ServiceInformationAnyOf, error) {
	ev := ServiceInformationAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceInformationAnyOf: valid values are %v", v, AllowedServiceInformationAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceInformationAnyOf) IsValid() bool {
	for _, existing := range AllowedServiceInformationAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceInformation_anyOf value
func (v ServiceInformationAnyOf) Ptr() *ServiceInformationAnyOf {
	return &v
}

type NullableServiceInformationAnyOf struct {
	value *ServiceInformationAnyOf
	isSet bool
}

func (v NullableServiceInformationAnyOf) Get() *ServiceInformationAnyOf {
	return v.value
}

func (v *NullableServiceInformationAnyOf) Set(val *ServiceInformationAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceInformationAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceInformationAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceInformationAnyOf(val *ServiceInformationAnyOf) *NullableServiceInformationAnyOf {
	return &NullableServiceInformationAnyOf{value: val, isSet: true}
}

func (v NullableServiceInformationAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceInformationAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

