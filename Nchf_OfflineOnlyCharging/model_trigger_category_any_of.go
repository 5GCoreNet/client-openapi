/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_OfflineOnlyCharging

import (
	"encoding/json"
	"fmt"
)

// TriggerCategoryAnyOf the model 'TriggerCategoryAnyOf'
type TriggerCategoryAnyOf string

// List of TriggerCategory_anyOf
const (
	IMMEDIATE_REPORT TriggerCategoryAnyOf = "IMMEDIATE_REPORT"
	DEFERRED_REPORT TriggerCategoryAnyOf = "DEFERRED_REPORT"
)

// All allowed values of TriggerCategoryAnyOf enum
var AllowedTriggerCategoryAnyOfEnumValues = []TriggerCategoryAnyOf{
	"IMMEDIATE_REPORT",
	"DEFERRED_REPORT",
}

func (v *TriggerCategoryAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TriggerCategoryAnyOf(value)
	for _, existing := range AllowedTriggerCategoryAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TriggerCategoryAnyOf", value)
}

// NewTriggerCategoryAnyOfFromValue returns a pointer to a valid TriggerCategoryAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTriggerCategoryAnyOfFromValue(v string) (*TriggerCategoryAnyOf, error) {
	ev := TriggerCategoryAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TriggerCategoryAnyOf: valid values are %v", v, AllowedTriggerCategoryAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TriggerCategoryAnyOf) IsValid() bool {
	for _, existing := range AllowedTriggerCategoryAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TriggerCategory_anyOf value
func (v TriggerCategoryAnyOf) Ptr() *TriggerCategoryAnyOf {
	return &v
}

type NullableTriggerCategoryAnyOf struct {
	value *TriggerCategoryAnyOf
	isSet bool
}

func (v NullableTriggerCategoryAnyOf) Get() *TriggerCategoryAnyOf {
	return v.value
}

func (v *NullableTriggerCategoryAnyOf) Set(val *TriggerCategoryAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerCategoryAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerCategoryAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerCategoryAnyOf(val *TriggerCategoryAnyOf) *NullableTriggerCategoryAnyOf {
	return &NullableTriggerCategoryAnyOf{value: val, isSet: true}
}

func (v NullableTriggerCategoryAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerCategoryAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

