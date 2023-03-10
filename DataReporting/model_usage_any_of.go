/*
3gpp-data-reporting

API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_DataReporting

import (
	"encoding/json"
	"fmt"
)

// UsageAnyOf the model 'UsageAnyOf'
type UsageAnyOf string

// List of Usage_anyOf
const (
	UNSUCCESS UsageAnyOf = "UNSUCCESS"
	SUCCESS_RESULTS_NOT_USED UsageAnyOf = "SUCCESS_RESULTS_NOT_USED"
	SUCCESS_RESULTS_USED_TO_VERIFY_LOCATION UsageAnyOf = "SUCCESS_RESULTS_USED_TO_VERIFY_LOCATION"
	SUCCESS_RESULTS_USED_TO_GENERATE_LOCATION UsageAnyOf = "SUCCESS_RESULTS_USED_TO_GENERATE_LOCATION"
	SUCCESS_METHOD_NOT_DETERMINED UsageAnyOf = "SUCCESS_METHOD_NOT_DETERMINED"
)

// All allowed values of UsageAnyOf enum
var AllowedUsageAnyOfEnumValues = []UsageAnyOf{
	"UNSUCCESS",
	"SUCCESS_RESULTS_NOT_USED",
	"SUCCESS_RESULTS_USED_TO_VERIFY_LOCATION",
	"SUCCESS_RESULTS_USED_TO_GENERATE_LOCATION",
	"SUCCESS_METHOD_NOT_DETERMINED",
}

func (v *UsageAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UsageAnyOf(value)
	for _, existing := range AllowedUsageAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UsageAnyOf", value)
}

// NewUsageAnyOfFromValue returns a pointer to a valid UsageAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUsageAnyOfFromValue(v string) (*UsageAnyOf, error) {
	ev := UsageAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UsageAnyOf: valid values are %v", v, AllowedUsageAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UsageAnyOf) IsValid() bool {
	for _, existing := range AllowedUsageAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Usage_anyOf value
func (v UsageAnyOf) Ptr() *UsageAnyOf {
	return &v
}

type NullableUsageAnyOf struct {
	value *UsageAnyOf
	isSet bool
}

func (v NullableUsageAnyOf) Get() *UsageAnyOf {
	return v.value
}

func (v *NullableUsageAnyOf) Set(val *UsageAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageAnyOf(val *UsageAnyOf) *NullableUsageAnyOf {
	return &NullableUsageAnyOf{value: val, isSet: true}
}

func (v NullableUsageAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

