/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoslaNrm

import (
	"encoding/json"
	"fmt"
)

// AssuranceGoalStatusPredicted the model 'AssuranceGoalStatusPredicted'
type AssuranceGoalStatusPredicted string

// List of AssuranceGoalStatusPredicted
const (
	FULFILLED AssuranceGoalStatusPredicted = "FULFILLED"
	NOT_FULFILLED AssuranceGoalStatusPredicted = "NOT_FULFILLED"
)

// All allowed values of AssuranceGoalStatusPredicted enum
var AllowedAssuranceGoalStatusPredictedEnumValues = []AssuranceGoalStatusPredicted{
	"FULFILLED",
	"NOT_FULFILLED",
}

func (v *AssuranceGoalStatusPredicted) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AssuranceGoalStatusPredicted(value)
	for _, existing := range AllowedAssuranceGoalStatusPredictedEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AssuranceGoalStatusPredicted", value)
}

// NewAssuranceGoalStatusPredictedFromValue returns a pointer to a valid AssuranceGoalStatusPredicted
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssuranceGoalStatusPredictedFromValue(v string) (*AssuranceGoalStatusPredicted, error) {
	ev := AssuranceGoalStatusPredicted(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssuranceGoalStatusPredicted: valid values are %v", v, AllowedAssuranceGoalStatusPredictedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssuranceGoalStatusPredicted) IsValid() bool {
	for _, existing := range AllowedAssuranceGoalStatusPredictedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssuranceGoalStatusPredicted value
func (v AssuranceGoalStatusPredicted) Ptr() *AssuranceGoalStatusPredicted {
	return &v
}

type NullableAssuranceGoalStatusPredicted struct {
	value *AssuranceGoalStatusPredicted
	isSet bool
}

func (v NullableAssuranceGoalStatusPredicted) Get() *AssuranceGoalStatusPredicted {
	return v.value
}

func (v *NullableAssuranceGoalStatusPredicted) Set(val *AssuranceGoalStatusPredicted) {
	v.value = val
	v.isSet = true
}

func (v NullableAssuranceGoalStatusPredicted) IsSet() bool {
	return v.isSet
}

func (v *NullableAssuranceGoalStatusPredicted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssuranceGoalStatusPredicted(val *AssuranceGoalStatusPredicted) *NullableAssuranceGoalStatusPredicted {
	return &NullableAssuranceGoalStatusPredicted{value: val, isSet: true}
}

func (v NullableAssuranceGoalStatusPredicted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssuranceGoalStatusPredicted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

