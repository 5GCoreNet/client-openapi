/*
M5_ConsumptionReporting

5GMS AF M5 Consumption Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package M5_ConsumptionReporting

import (
	"encoding/json"
	"fmt"
)

// CellIdentifierTypeAnyOf the model 'CellIdentifierTypeAnyOf'
type CellIdentifierTypeAnyOf string

// List of CellIdentifierType_anyOf
const (
	CGI CellIdentifierTypeAnyOf = "CGI"
	ECGI CellIdentifierTypeAnyOf = "ECGI"
	NCGI CellIdentifierTypeAnyOf = "NCGI"
)

// All allowed values of CellIdentifierTypeAnyOf enum
var AllowedCellIdentifierTypeAnyOfEnumValues = []CellIdentifierTypeAnyOf{
	"CGI",
	"ECGI",
	"NCGI",
}

func (v *CellIdentifierTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CellIdentifierTypeAnyOf(value)
	for _, existing := range AllowedCellIdentifierTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CellIdentifierTypeAnyOf", value)
}

// NewCellIdentifierTypeAnyOfFromValue returns a pointer to a valid CellIdentifierTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCellIdentifierTypeAnyOfFromValue(v string) (*CellIdentifierTypeAnyOf, error) {
	ev := CellIdentifierTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CellIdentifierTypeAnyOf: valid values are %v", v, AllowedCellIdentifierTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CellIdentifierTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedCellIdentifierTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CellIdentifierType_anyOf value
func (v CellIdentifierTypeAnyOf) Ptr() *CellIdentifierTypeAnyOf {
	return &v
}

type NullableCellIdentifierTypeAnyOf struct {
	value *CellIdentifierTypeAnyOf
	isSet bool
}

func (v NullableCellIdentifierTypeAnyOf) Get() *CellIdentifierTypeAnyOf {
	return v.value
}

func (v *NullableCellIdentifierTypeAnyOf) Set(val *CellIdentifierTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCellIdentifierTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCellIdentifierTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCellIdentifierTypeAnyOf(val *CellIdentifierTypeAnyOf) *NullableCellIdentifierTypeAnyOf {
	return &NullableCellIdentifierTypeAnyOf{value: val, isSet: true}
}

func (v NullableCellIdentifierTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCellIdentifierTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

