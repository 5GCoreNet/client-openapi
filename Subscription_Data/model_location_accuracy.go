/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// LocationAccuracy struct for LocationAccuracy
type LocationAccuracy struct {
	LocationAccuracyAnyOf *LocationAccuracyAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LocationAccuracy) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LocationAccuracyAnyOf
	err = json.Unmarshal(data, &dst.LocationAccuracyAnyOf);
	if err == nil {
		jsonLocationAccuracyAnyOf, _ := json.Marshal(dst.LocationAccuracyAnyOf)
		if string(jsonLocationAccuracyAnyOf) == "{}" { // empty struct
			dst.LocationAccuracyAnyOf = nil
		} else {
			return nil // data stored in dst.LocationAccuracyAnyOf, return on the first match
		}
	} else {
		dst.LocationAccuracyAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(LocationAccuracy)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LocationAccuracy) MarshalJSON() ([]byte, error) {
	if src.LocationAccuracyAnyOf != nil {
		return json.Marshal(&src.LocationAccuracyAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLocationAccuracy struct {
	value *LocationAccuracy
	isSet bool
}

func (v NullableLocationAccuracy) Get() *LocationAccuracy {
	return v.value
}

func (v *NullableLocationAccuracy) Set(val *LocationAccuracy) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationAccuracy) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationAccuracy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationAccuracy(val *LocationAccuracy) *NullableLocationAccuracy {
	return &NullableLocationAccuracy{value: val, isSet: true}
}

func (v NullableLocationAccuracy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationAccuracy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


