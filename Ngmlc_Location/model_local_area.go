/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
	"fmt"
)

// LocalArea - Local area specified by different shape
type LocalArea struct {
	Local2dPointUncertaintyEllipse *Local2dPointUncertaintyEllipse
	Local3dPointUncertaintyEllipsoid *Local3dPointUncertaintyEllipsoid
}

// Local2dPointUncertaintyEllipseAsLocalArea is a convenience function that returns Local2dPointUncertaintyEllipse wrapped in LocalArea
func Local2dPointUncertaintyEllipseAsLocalArea(v *Local2dPointUncertaintyEllipse) LocalArea {
	return LocalArea{
		Local2dPointUncertaintyEllipse: v,
	}
}

// Local3dPointUncertaintyEllipsoidAsLocalArea is a convenience function that returns Local3dPointUncertaintyEllipsoid wrapped in LocalArea
func Local3dPointUncertaintyEllipsoidAsLocalArea(v *Local3dPointUncertaintyEllipsoid) LocalArea {
	return LocalArea{
		Local3dPointUncertaintyEllipsoid: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *LocalArea) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Local2dPointUncertaintyEllipse
	err = newStrictDecoder(data).Decode(&dst.Local2dPointUncertaintyEllipse)
	if err == nil {
		jsonLocal2dPointUncertaintyEllipse, _ := json.Marshal(dst.Local2dPointUncertaintyEllipse)
		if string(jsonLocal2dPointUncertaintyEllipse) == "{}" { // empty struct
			dst.Local2dPointUncertaintyEllipse = nil
		} else {
			match++
		}
	} else {
		dst.Local2dPointUncertaintyEllipse = nil
	}

	// try to unmarshal data into Local3dPointUncertaintyEllipsoid
	err = newStrictDecoder(data).Decode(&dst.Local3dPointUncertaintyEllipsoid)
	if err == nil {
		jsonLocal3dPointUncertaintyEllipsoid, _ := json.Marshal(dst.Local3dPointUncertaintyEllipsoid)
		if string(jsonLocal3dPointUncertaintyEllipsoid) == "{}" { // empty struct
			dst.Local3dPointUncertaintyEllipsoid = nil
		} else {
			match++
		}
	} else {
		dst.Local3dPointUncertaintyEllipsoid = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Local2dPointUncertaintyEllipse = nil
		dst.Local3dPointUncertaintyEllipsoid = nil

		return fmt.Errorf("data matches more than one schema in oneOf(LocalArea)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LocalArea)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LocalArea) MarshalJSON() ([]byte, error) {
	if src.Local2dPointUncertaintyEllipse != nil {
		return json.Marshal(&src.Local2dPointUncertaintyEllipse)
	}

	if src.Local3dPointUncertaintyEllipsoid != nil {
		return json.Marshal(&src.Local3dPointUncertaintyEllipsoid)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LocalArea) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Local2dPointUncertaintyEllipse != nil {
		return obj.Local2dPointUncertaintyEllipse
	}

	if obj.Local3dPointUncertaintyEllipsoid != nil {
		return obj.Local3dPointUncertaintyEllipsoid
	}

	// all schemas are nil
	return nil
}

type NullableLocalArea struct {
	value *LocalArea
	isSet bool
}

func (v NullableLocalArea) Get() *LocalArea {
	return v.value
}

func (v *NullableLocalArea) Set(val *LocalArea) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalArea) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalArea(val *LocalArea) *NullableLocalArea {
	return &NullableLocalArea{value: val, isSet: true}
}

func (v NullableLocalArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


