/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"fmt"
)

// Operation the model 'Operation'
type Operation string

// List of Operation
const (
	ADD Operation = "add"
	REMOVE Operation = "remove"
	REPLACE Operation = "replace"
)

// All allowed values of Operation enum
var AllowedOperationEnumValues = []Operation{
	"add",
	"remove",
	"replace",
}

func (v *Operation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Operation(value)
	for _, existing := range AllowedOperationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Operation", value)
}

// NewOperationFromValue returns a pointer to a valid Operation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperationFromValue(v string) (*Operation, error) {
	ev := Operation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Operation: valid values are %v", v, AllowedOperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Operation) IsValid() bool {
	for _, existing := range AllowedOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Operation value
func (v Operation) Ptr() *Operation {
	return &v
}

type NullableOperation struct {
	value *Operation
	isSet bool
}

func (v NullableOperation) Get() *Operation {
	return v.value
}

func (v *NullableOperation) Set(val *Operation) {
	v.value = val
	v.isSet = true
}

func (v NullableOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperation(val *Operation) *NullableOperation {
	return &NullableOperation{value: val, isSet: true}
}

func (v NullableOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

