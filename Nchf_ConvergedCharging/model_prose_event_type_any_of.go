/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// ProseEventTypeAnyOf the model 'ProseEventTypeAnyOf'
type ProseEventTypeAnyOf string

// List of ProseEventType_anyOf
const (
	ANNOUNCING ProseEventTypeAnyOf = "ANNOUNCING"
	MONITORING ProseEventTypeAnyOf = "MONITORING"
	MATCH_REPORT ProseEventTypeAnyOf = "MATCH_REPORT"
)

// All allowed values of ProseEventTypeAnyOf enum
var AllowedProseEventTypeAnyOfEnumValues = []ProseEventTypeAnyOf{
	"ANNOUNCING",
	"MONITORING",
	"MATCH_REPORT",
}

func (v *ProseEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProseEventTypeAnyOf(value)
	for _, existing := range AllowedProseEventTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProseEventTypeAnyOf", value)
}

// NewProseEventTypeAnyOfFromValue returns a pointer to a valid ProseEventTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProseEventTypeAnyOfFromValue(v string) (*ProseEventTypeAnyOf, error) {
	ev := ProseEventTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProseEventTypeAnyOf: valid values are %v", v, AllowedProseEventTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProseEventTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedProseEventTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProseEventType_anyOf value
func (v ProseEventTypeAnyOf) Ptr() *ProseEventTypeAnyOf {
	return &v
}

type NullableProseEventTypeAnyOf struct {
	value *ProseEventTypeAnyOf
	isSet bool
}

func (v NullableProseEventTypeAnyOf) Get() *ProseEventTypeAnyOf {
	return v.value
}

func (v *NullableProseEventTypeAnyOf) Set(val *ProseEventTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProseEventTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProseEventTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProseEventTypeAnyOf(val *ProseEventTypeAnyOf) *NullableProseEventTypeAnyOf {
	return &NullableProseEventTypeAnyOf{value: val, isSet: true}
}

func (v NullableProseEventTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProseEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

