/*
M1_EdgeResourcesProvisioning

5GMS AF M1 Edge Resources Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package M1_EdgeResourcesProvisioning

import (
	"encoding/json"
	"fmt"
)

// ACRScenarioAnyOf the model 'ACRScenarioAnyOf'
type ACRScenarioAnyOf string

// List of ACRScenario_anyOf
const (
	EEC_INITIATED ACRScenarioAnyOf = "EEC_INITIATED"
	EEC_EXECUTED_VIA_SOURCE_EES ACRScenarioAnyOf = "EEC_EXECUTED_VIA_SOURCE_EES"
	EEC_EXECUTED_VIA_TARGET_EES ACRScenarioAnyOf = "EEC_EXECUTED_VIA_TARGET_EES"
	SOURCE_EAS_DECIDED ACRScenarioAnyOf = "SOURCE_EAS_DECIDED"
	SOURCE_EES_EXECUTED ACRScenarioAnyOf = "SOURCE_EES_EXECUTED"
	EEL_MANAGED_ACR ACRScenarioAnyOf = "EEL_MANAGED_ACR"
)

// All allowed values of ACRScenarioAnyOf enum
var AllowedACRScenarioAnyOfEnumValues = []ACRScenarioAnyOf{
	"EEC_INITIATED",
	"EEC_EXECUTED_VIA_SOURCE_EES",
	"EEC_EXECUTED_VIA_TARGET_EES",
	"SOURCE_EAS_DECIDED",
	"SOURCE_EES_EXECUTED",
	"EEL_MANAGED_ACR",
}

func (v *ACRScenarioAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ACRScenarioAnyOf(value)
	for _, existing := range AllowedACRScenarioAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ACRScenarioAnyOf", value)
}

// NewACRScenarioAnyOfFromValue returns a pointer to a valid ACRScenarioAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewACRScenarioAnyOfFromValue(v string) (*ACRScenarioAnyOf, error) {
	ev := ACRScenarioAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ACRScenarioAnyOf: valid values are %v", v, AllowedACRScenarioAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ACRScenarioAnyOf) IsValid() bool {
	for _, existing := range AllowedACRScenarioAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ACRScenario_anyOf value
func (v ACRScenarioAnyOf) Ptr() *ACRScenarioAnyOf {
	return &v
}

type NullableACRScenarioAnyOf struct {
	value *ACRScenarioAnyOf
	isSet bool
}

func (v NullableACRScenarioAnyOf) Get() *ACRScenarioAnyOf {
	return v.value
}

func (v *NullableACRScenarioAnyOf) Set(val *ACRScenarioAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableACRScenarioAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableACRScenarioAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACRScenarioAnyOf(val *ACRScenarioAnyOf) *NullableACRScenarioAnyOf {
	return &NullableACRScenarioAnyOf{value: val, isSet: true}
}

func (v NullableACRScenarioAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACRScenarioAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

