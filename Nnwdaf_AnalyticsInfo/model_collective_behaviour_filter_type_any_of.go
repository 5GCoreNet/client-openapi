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

// CollectiveBehaviourFilterTypeAnyOf the model 'CollectiveBehaviourFilterTypeAnyOf'
type CollectiveBehaviourFilterTypeAnyOf string

// List of CollectiveBehaviourFilterType_anyOf
const (
	COLLECTIVE_ATTRIBUTE CollectiveBehaviourFilterTypeAnyOf = "COLLECTIVE_ATTRIBUTE"
	DATA_PROCESSING CollectiveBehaviourFilterTypeAnyOf = "DATA_PROCESSING"
)

// All allowed values of CollectiveBehaviourFilterTypeAnyOf enum
var AllowedCollectiveBehaviourFilterTypeAnyOfEnumValues = []CollectiveBehaviourFilterTypeAnyOf{
	"COLLECTIVE_ATTRIBUTE",
	"DATA_PROCESSING",
}

func (v *CollectiveBehaviourFilterTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectiveBehaviourFilterTypeAnyOf(value)
	for _, existing := range AllowedCollectiveBehaviourFilterTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectiveBehaviourFilterTypeAnyOf", value)
}

// NewCollectiveBehaviourFilterTypeAnyOfFromValue returns a pointer to a valid CollectiveBehaviourFilterTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectiveBehaviourFilterTypeAnyOfFromValue(v string) (*CollectiveBehaviourFilterTypeAnyOf, error) {
	ev := CollectiveBehaviourFilterTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectiveBehaviourFilterTypeAnyOf: valid values are %v", v, AllowedCollectiveBehaviourFilterTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectiveBehaviourFilterTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedCollectiveBehaviourFilterTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CollectiveBehaviourFilterType_anyOf value
func (v CollectiveBehaviourFilterTypeAnyOf) Ptr() *CollectiveBehaviourFilterTypeAnyOf {
	return &v
}

type NullableCollectiveBehaviourFilterTypeAnyOf struct {
	value *CollectiveBehaviourFilterTypeAnyOf
	isSet bool
}

func (v NullableCollectiveBehaviourFilterTypeAnyOf) Get() *CollectiveBehaviourFilterTypeAnyOf {
	return v.value
}

func (v *NullableCollectiveBehaviourFilterTypeAnyOf) Set(val *CollectiveBehaviourFilterTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectiveBehaviourFilterTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectiveBehaviourFilterTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectiveBehaviourFilterTypeAnyOf(val *CollectiveBehaviourFilterTypeAnyOf) *NullableCollectiveBehaviourFilterTypeAnyOf {
	return &NullableCollectiveBehaviourFilterTypeAnyOf{value: val, isSet: true}
}

func (v NullableCollectiveBehaviourFilterTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectiveBehaviourFilterTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

