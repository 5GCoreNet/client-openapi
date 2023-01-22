/*
M1_ProvisioningSessions

5GMS AF M1 Provisioning Sessions API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package M1_ProvisioningSessions

import (
	"encoding/json"
	"fmt"
)

// ProvisioningSessionTypeAnyOf the model 'ProvisioningSessionTypeAnyOf'
type ProvisioningSessionTypeAnyOf string

// List of ProvisioningSessionType_anyOf
const (
	DOWNLINK ProvisioningSessionTypeAnyOf = "DOWNLINK"
	UPLINK ProvisioningSessionTypeAnyOf = "UPLINK"
)

// All allowed values of ProvisioningSessionTypeAnyOf enum
var AllowedProvisioningSessionTypeAnyOfEnumValues = []ProvisioningSessionTypeAnyOf{
	"DOWNLINK",
	"UPLINK",
}

func (v *ProvisioningSessionTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProvisioningSessionTypeAnyOf(value)
	for _, existing := range AllowedProvisioningSessionTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProvisioningSessionTypeAnyOf", value)
}

// NewProvisioningSessionTypeAnyOfFromValue returns a pointer to a valid ProvisioningSessionTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProvisioningSessionTypeAnyOfFromValue(v string) (*ProvisioningSessionTypeAnyOf, error) {
	ev := ProvisioningSessionTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProvisioningSessionTypeAnyOf: valid values are %v", v, AllowedProvisioningSessionTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProvisioningSessionTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedProvisioningSessionTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProvisioningSessionType_anyOf value
func (v ProvisioningSessionTypeAnyOf) Ptr() *ProvisioningSessionTypeAnyOf {
	return &v
}

type NullableProvisioningSessionTypeAnyOf struct {
	value *ProvisioningSessionTypeAnyOf
	isSet bool
}

func (v NullableProvisioningSessionTypeAnyOf) Get() *ProvisioningSessionTypeAnyOf {
	return v.value
}

func (v *NullableProvisioningSessionTypeAnyOf) Set(val *ProvisioningSessionTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningSessionTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningSessionTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningSessionTypeAnyOf(val *ProvisioningSessionTypeAnyOf) *NullableProvisioningSessionTypeAnyOf {
	return &NullableProvisioningSessionTypeAnyOf{value: val, isSet: true}
}

func (v NullableProvisioningSessionTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningSessionTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

