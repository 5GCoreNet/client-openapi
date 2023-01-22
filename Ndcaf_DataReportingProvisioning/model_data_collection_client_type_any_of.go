/*
Ndcaf_DataReportingProvisioning

Data Collection AF: Provisioning Sessions API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndcaf_DataReportingProvisioning

import (
	"encoding/json"
	"fmt"
)

// DataCollectionClientTypeAnyOf the model 'DataCollectionClientTypeAnyOf'
type DataCollectionClientTypeAnyOf string

// List of DataCollectionClientType_anyOf
const (
	DIRECT DataCollectionClientTypeAnyOf = "DIRECT"
	INDIRECT DataCollectionClientTypeAnyOf = "INDIRECT"
	APPLICATION_SERVER DataCollectionClientTypeAnyOf = "APPLICATION_SERVER"
)

// All allowed values of DataCollectionClientTypeAnyOf enum
var AllowedDataCollectionClientTypeAnyOfEnumValues = []DataCollectionClientTypeAnyOf{
	"DIRECT",
	"INDIRECT",
	"APPLICATION_SERVER",
}

func (v *DataCollectionClientTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataCollectionClientTypeAnyOf(value)
	for _, existing := range AllowedDataCollectionClientTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataCollectionClientTypeAnyOf", value)
}

// NewDataCollectionClientTypeAnyOfFromValue returns a pointer to a valid DataCollectionClientTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataCollectionClientTypeAnyOfFromValue(v string) (*DataCollectionClientTypeAnyOf, error) {
	ev := DataCollectionClientTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataCollectionClientTypeAnyOf: valid values are %v", v, AllowedDataCollectionClientTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataCollectionClientTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedDataCollectionClientTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataCollectionClientType_anyOf value
func (v DataCollectionClientTypeAnyOf) Ptr() *DataCollectionClientTypeAnyOf {
	return &v
}

type NullableDataCollectionClientTypeAnyOf struct {
	value *DataCollectionClientTypeAnyOf
	isSet bool
}

func (v NullableDataCollectionClientTypeAnyOf) Get() *DataCollectionClientTypeAnyOf {
	return v.value
}

func (v *NullableDataCollectionClientTypeAnyOf) Set(val *DataCollectionClientTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDataCollectionClientTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDataCollectionClientTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataCollectionClientTypeAnyOf(val *DataCollectionClientTypeAnyOf) *NullableDataCollectionClientTypeAnyOf {
	return &NullableDataCollectionClientTypeAnyOf{value: val, isSet: true}
}

func (v NullableDataCollectionClientTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataCollectionClientTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

