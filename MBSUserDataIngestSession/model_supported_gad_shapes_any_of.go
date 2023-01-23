/*
3gpp-mbs-ud-ingest

API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSUserDataIngestSession

import (
	"encoding/json"
	"fmt"
)

// SupportedGADShapesAnyOf the model 'SupportedGADShapesAnyOf'
type SupportedGADShapesAnyOf string

// List of SupportedGADShapes_anyOf
const (
	POINT SupportedGADShapesAnyOf = "POINT"
	POINT_UNCERTAINTY_CIRCLE SupportedGADShapesAnyOf = "POINT_UNCERTAINTY_CIRCLE"
	POINT_UNCERTAINTY_ELLIPSE SupportedGADShapesAnyOf = "POINT_UNCERTAINTY_ELLIPSE"
	POLYGON SupportedGADShapesAnyOf = "POLYGON"
	POINT_ALTITUDE SupportedGADShapesAnyOf = "POINT_ALTITUDE"
	POINT_ALTITUDE_UNCERTAINTY SupportedGADShapesAnyOf = "POINT_ALTITUDE_UNCERTAINTY"
	ELLIPSOID_ARC SupportedGADShapesAnyOf = "ELLIPSOID_ARC"
	LOCAL_2_D_POINT_UNCERTAINTY_ELLIPSE SupportedGADShapesAnyOf = "LOCAL_2D_POINT_UNCERTAINTY_ELLIPSE"
	LOCAL_3_D_POINT_UNCERTAINTY_ELLIPSOID SupportedGADShapesAnyOf = "LOCAL_3D_POINT_UNCERTAINTY_ELLIPSOID"
)

// All allowed values of SupportedGADShapesAnyOf enum
var AllowedSupportedGADShapesAnyOfEnumValues = []SupportedGADShapesAnyOf{
	"POINT",
	"POINT_UNCERTAINTY_CIRCLE",
	"POINT_UNCERTAINTY_ELLIPSE",
	"POLYGON",
	"POINT_ALTITUDE",
	"POINT_ALTITUDE_UNCERTAINTY",
	"ELLIPSOID_ARC",
	"LOCAL_2D_POINT_UNCERTAINTY_ELLIPSE",
	"LOCAL_3D_POINT_UNCERTAINTY_ELLIPSOID",
}

func (v *SupportedGADShapesAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SupportedGADShapesAnyOf(value)
	for _, existing := range AllowedSupportedGADShapesAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SupportedGADShapesAnyOf", value)
}

// NewSupportedGADShapesAnyOfFromValue returns a pointer to a valid SupportedGADShapesAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSupportedGADShapesAnyOfFromValue(v string) (*SupportedGADShapesAnyOf, error) {
	ev := SupportedGADShapesAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SupportedGADShapesAnyOf: valid values are %v", v, AllowedSupportedGADShapesAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SupportedGADShapesAnyOf) IsValid() bool {
	for _, existing := range AllowedSupportedGADShapesAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SupportedGADShapes_anyOf value
func (v SupportedGADShapesAnyOf) Ptr() *SupportedGADShapesAnyOf {
	return &v
}

type NullableSupportedGADShapesAnyOf struct {
	value *SupportedGADShapesAnyOf
	isSet bool
}

func (v NullableSupportedGADShapesAnyOf) Get() *SupportedGADShapesAnyOf {
	return v.value
}

func (v *NullableSupportedGADShapesAnyOf) Set(val *SupportedGADShapesAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedGADShapesAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedGADShapesAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedGADShapesAnyOf(val *SupportedGADShapesAnyOf) *NullableSupportedGADShapesAnyOf {
	return &NullableSupportedGADShapesAnyOf{value: val, isSet: true}
}

func (v NullableSupportedGADShapesAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedGADShapesAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

