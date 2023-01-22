/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_SDM

import (
	"encoding/json"
	"fmt"
)

// LcsClientClassAnyOf the model 'LcsClientClassAnyOf'
type LcsClientClassAnyOf string

// List of LcsClientClass_anyOf
const (
	BROADCAST_SERVICE LcsClientClassAnyOf = "BROADCAST_SERVICE"
	OM_IN_HPLMN LcsClientClassAnyOf = "OM_IN_HPLMN"
	OM_IN_VPLMN LcsClientClassAnyOf = "OM_IN_VPLMN"
	ANONYMOUS_LOCATION_SERVICE LcsClientClassAnyOf = "ANONYMOUS_LOCATION_SERVICE"
	SPECIFIC_SERVICE LcsClientClassAnyOf = "SPECIFIC_SERVICE"
)

// All allowed values of LcsClientClassAnyOf enum
var AllowedLcsClientClassAnyOfEnumValues = []LcsClientClassAnyOf{
	"BROADCAST_SERVICE",
	"OM_IN_HPLMN",
	"OM_IN_VPLMN",
	"ANONYMOUS_LOCATION_SERVICE",
	"SPECIFIC_SERVICE",
}

func (v *LcsClientClassAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LcsClientClassAnyOf(value)
	for _, existing := range AllowedLcsClientClassAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LcsClientClassAnyOf", value)
}

// NewLcsClientClassAnyOfFromValue returns a pointer to a valid LcsClientClassAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLcsClientClassAnyOfFromValue(v string) (*LcsClientClassAnyOf, error) {
	ev := LcsClientClassAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LcsClientClassAnyOf: valid values are %v", v, AllowedLcsClientClassAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LcsClientClassAnyOf) IsValid() bool {
	for _, existing := range AllowedLcsClientClassAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LcsClientClass_anyOf value
func (v LcsClientClassAnyOf) Ptr() *LcsClientClassAnyOf {
	return &v
}

type NullableLcsClientClassAnyOf struct {
	value *LcsClientClassAnyOf
	isSet bool
}

func (v NullableLcsClientClassAnyOf) Get() *LcsClientClassAnyOf {
	return v.value
}

func (v *NullableLcsClientClassAnyOf) Set(val *LcsClientClassAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsClientClassAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsClientClassAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsClientClassAnyOf(val *LcsClientClassAnyOf) *NullableLcsClientClassAnyOf {
	return &NullableLcsClientClassAnyOf{value: val, isSet: true}
}

func (v NullableLcsClientClassAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsClientClassAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

