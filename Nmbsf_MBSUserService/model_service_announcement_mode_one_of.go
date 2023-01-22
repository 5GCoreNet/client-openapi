/*
nmbsf-mbs-us

API for MBS User Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbsf_MBSUserService

import (
	"encoding/json"
	"fmt"
)

// ServiceAnnouncementModeOneOf the model 'ServiceAnnouncementModeOneOf'
type ServiceAnnouncementModeOneOf string

// List of ServiceAnnouncementMode_oneOf
const (
	VIA_MBS_5 ServiceAnnouncementModeOneOf = "VIA_MBS_5"
	VIA_MBS_DISTRIBUTION_SESSION ServiceAnnouncementModeOneOf = "VIA_MBS_DISTRIBUTION_SESSION"
	PASSED_BACK ServiceAnnouncementModeOneOf = "PASSED_BACK"
)

// All allowed values of ServiceAnnouncementModeOneOf enum
var AllowedServiceAnnouncementModeOneOfEnumValues = []ServiceAnnouncementModeOneOf{
	"VIA_MBS_5",
	"VIA_MBS_DISTRIBUTION_SESSION",
	"PASSED_BACK",
}

func (v *ServiceAnnouncementModeOneOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceAnnouncementModeOneOf(value)
	for _, existing := range AllowedServiceAnnouncementModeOneOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceAnnouncementModeOneOf", value)
}

// NewServiceAnnouncementModeOneOfFromValue returns a pointer to a valid ServiceAnnouncementModeOneOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceAnnouncementModeOneOfFromValue(v string) (*ServiceAnnouncementModeOneOf, error) {
	ev := ServiceAnnouncementModeOneOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceAnnouncementModeOneOf: valid values are %v", v, AllowedServiceAnnouncementModeOneOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceAnnouncementModeOneOf) IsValid() bool {
	for _, existing := range AllowedServiceAnnouncementModeOneOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceAnnouncementMode_oneOf value
func (v ServiceAnnouncementModeOneOf) Ptr() *ServiceAnnouncementModeOneOf {
	return &v
}

type NullableServiceAnnouncementModeOneOf struct {
	value *ServiceAnnouncementModeOneOf
	isSet bool
}

func (v NullableServiceAnnouncementModeOneOf) Get() *ServiceAnnouncementModeOneOf {
	return v.value
}

func (v *NullableServiceAnnouncementModeOneOf) Set(val *ServiceAnnouncementModeOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAnnouncementModeOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAnnouncementModeOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAnnouncementModeOneOf(val *ServiceAnnouncementModeOneOf) *NullableServiceAnnouncementModeOneOf {
	return &NullableServiceAnnouncementModeOneOf{value: val, isSet: true}
}

func (v NullableServiceAnnouncementModeOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAnnouncementModeOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

