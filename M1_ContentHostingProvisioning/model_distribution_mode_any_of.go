/*
M1_ContentHostingProvisioning

5GMS AF M1 Content Hosting Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package M1_ContentHostingProvisioning

import (
	"encoding/json"
	"fmt"
)

// DistributionModeAnyOf the model 'DistributionModeAnyOf'
type DistributionModeAnyOf string

// List of DistributionMode_anyOf
const (
	EXCLUSIVE DistributionModeAnyOf = "MODE_EXCLUSIVE"
	HYBRID DistributionModeAnyOf = "MODE_HYBRID"
	DYNAMIC DistributionModeAnyOf = "MODE_DYNAMIC"
)

// All allowed values of DistributionModeAnyOf enum
var AllowedDistributionModeAnyOfEnumValues = []DistributionModeAnyOf{
	"MODE_EXCLUSIVE",
	"MODE_HYBRID",
	"MODE_DYNAMIC",
}

func (v *DistributionModeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DistributionModeAnyOf(value)
	for _, existing := range AllowedDistributionModeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DistributionModeAnyOf", value)
}

// NewDistributionModeAnyOfFromValue returns a pointer to a valid DistributionModeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDistributionModeAnyOfFromValue(v string) (*DistributionModeAnyOf, error) {
	ev := DistributionModeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DistributionModeAnyOf: valid values are %v", v, AllowedDistributionModeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DistributionModeAnyOf) IsValid() bool {
	for _, existing := range AllowedDistributionModeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DistributionMode_anyOf value
func (v DistributionModeAnyOf) Ptr() *DistributionModeAnyOf {
	return &v
}

type NullableDistributionModeAnyOf struct {
	value *DistributionModeAnyOf
	isSet bool
}

func (v NullableDistributionModeAnyOf) Get() *DistributionModeAnyOf {
	return v.value
}

func (v *NullableDistributionModeAnyOf) Set(val *DistributionModeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionModeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionModeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionModeAnyOf(val *DistributionModeAnyOf) *NullableDistributionModeAnyOf {
	return &NullableDistributionModeAnyOf{value: val, isSet: true}
}

func (v NullableDistributionModeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionModeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

