/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
	"fmt"
)

// SatelliteBackhaulCategoryRm Provides information about the satellite backhaul but with the OpenAPI 'nullable: true' property.  
type SatelliteBackhaulCategoryRm struct {
	NullValue *NullValue
	SatelliteBackhaulCategory *SatelliteBackhaulCategory
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SatelliteBackhaulCategoryRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue);
	if err == nil {
		jsonNullValue, _ := json.Marshal(dst.NullValue)
		if string(jsonNullValue) == "{}" { // empty struct
			dst.NullValue = nil
		} else {
			return nil // data stored in dst.NullValue, return on the first match
		}
	} else {
		dst.NullValue = nil
	}

	// try to unmarshal JSON data into SatelliteBackhaulCategory
	err = json.Unmarshal(data, &dst.SatelliteBackhaulCategory);
	if err == nil {
		jsonSatelliteBackhaulCategory, _ := json.Marshal(dst.SatelliteBackhaulCategory)
		if string(jsonSatelliteBackhaulCategory) == "{}" { // empty struct
			dst.SatelliteBackhaulCategory = nil
		} else {
			return nil // data stored in dst.SatelliteBackhaulCategory, return on the first match
		}
	} else {
		dst.SatelliteBackhaulCategory = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SatelliteBackhaulCategoryRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SatelliteBackhaulCategoryRm) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.SatelliteBackhaulCategory != nil {
		return json.Marshal(&src.SatelliteBackhaulCategory)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSatelliteBackhaulCategoryRm struct {
	value *SatelliteBackhaulCategoryRm
	isSet bool
}

func (v NullableSatelliteBackhaulCategoryRm) Get() *SatelliteBackhaulCategoryRm {
	return v.value
}

func (v *NullableSatelliteBackhaulCategoryRm) Set(val *SatelliteBackhaulCategoryRm) {
	v.value = val
	v.isSet = true
}

func (v NullableSatelliteBackhaulCategoryRm) IsSet() bool {
	return v.isSet
}

func (v *NullableSatelliteBackhaulCategoryRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSatelliteBackhaulCategoryRm(val *SatelliteBackhaulCategoryRm) *NullableSatelliteBackhaulCategoryRm {
	return &NullableSatelliteBackhaulCategoryRm{value: val, isSet: true}
}

func (v NullableSatelliteBackhaulCategoryRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSatelliteBackhaulCategoryRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


