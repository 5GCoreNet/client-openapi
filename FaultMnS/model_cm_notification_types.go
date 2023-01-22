/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package FaultMnS

import (
	"encoding/json"
	"fmt"
)

// CmNotificationTypes the model 'CmNotificationTypes'
type CmNotificationTypes string

// List of CmNotificationTypes
const (
	NOTIFY_MOI_CREATION CmNotificationTypes = "notifyMOICreation"
	NOTIFY_MOI_DELETION CmNotificationTypes = "notifyMOIDeletion"
	NOTIFY_MOI_ATTRIBUTE_VALUE_CHANGES CmNotificationTypes = "notifyMOIAttributeValueChanges"
	NOTIFY_MOI_CHANGES CmNotificationTypes = "notifyMOIChanges"
)

// All allowed values of CmNotificationTypes enum
var AllowedCmNotificationTypesEnumValues = []CmNotificationTypes{
	"notifyMOICreation",
	"notifyMOIDeletion",
	"notifyMOIAttributeValueChanges",
	"notifyMOIChanges",
}

func (v *CmNotificationTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CmNotificationTypes(value)
	for _, existing := range AllowedCmNotificationTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CmNotificationTypes", value)
}

// NewCmNotificationTypesFromValue returns a pointer to a valid CmNotificationTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCmNotificationTypesFromValue(v string) (*CmNotificationTypes, error) {
	ev := CmNotificationTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CmNotificationTypes: valid values are %v", v, AllowedCmNotificationTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CmNotificationTypes) IsValid() bool {
	for _, existing := range AllowedCmNotificationTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CmNotificationTypes value
func (v CmNotificationTypes) Ptr() *CmNotificationTypes {
	return &v
}

type NullableCmNotificationTypes struct {
	value *CmNotificationTypes
	isSet bool
}

func (v NullableCmNotificationTypes) Get() *CmNotificationTypes {
	return v.value
}

func (v *NullableCmNotificationTypes) Set(val *CmNotificationTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableCmNotificationTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableCmNotificationTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmNotificationTypes(val *CmNotificationTypes) *NullableCmNotificationTypes {
	return &NullableCmNotificationTypes{value: val, isSet: true}
}

func (v NullableCmNotificationTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmNotificationTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

