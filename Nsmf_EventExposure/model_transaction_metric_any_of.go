/*
Nsmf_EventExposure

Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_EventExposure

import (
	"encoding/json"
	"fmt"
)

// TransactionMetricAnyOf the model 'TransactionMetricAnyOf'
type TransactionMetricAnyOf string

// List of TransactionMetric_anyOf
const (
	EST TransactionMetricAnyOf = "PDU_SES_EST"
	AUTH TransactionMetricAnyOf = "PDU_SES_AUTH"
	MODIF TransactionMetricAnyOf = "PDU_SES_MODIF"
	REL TransactionMetricAnyOf = "PDU_SES_REL"
)

// All allowed values of TransactionMetricAnyOf enum
var AllowedTransactionMetricAnyOfEnumValues = []TransactionMetricAnyOf{
	"PDU_SES_EST",
	"PDU_SES_AUTH",
	"PDU_SES_MODIF",
	"PDU_SES_REL",
}

func (v *TransactionMetricAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionMetricAnyOf(value)
	for _, existing := range AllowedTransactionMetricAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransactionMetricAnyOf", value)
}

// NewTransactionMetricAnyOfFromValue returns a pointer to a valid TransactionMetricAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionMetricAnyOfFromValue(v string) (*TransactionMetricAnyOf, error) {
	ev := TransactionMetricAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionMetricAnyOf: valid values are %v", v, AllowedTransactionMetricAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionMetricAnyOf) IsValid() bool {
	for _, existing := range AllowedTransactionMetricAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionMetric_anyOf value
func (v TransactionMetricAnyOf) Ptr() *TransactionMetricAnyOf {
	return &v
}

type NullableTransactionMetricAnyOf struct {
	value *TransactionMetricAnyOf
	isSet bool
}

func (v NullableTransactionMetricAnyOf) Get() *TransactionMetricAnyOf {
	return v.value
}

func (v *NullableTransactionMetricAnyOf) Set(val *TransactionMetricAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionMetricAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionMetricAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionMetricAnyOf(val *TransactionMetricAnyOf) *NullableTransactionMetricAnyOf {
	return &NullableTransactionMetricAnyOf{value: val, isSet: true}
}

func (v NullableTransactionMetricAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionMetricAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

