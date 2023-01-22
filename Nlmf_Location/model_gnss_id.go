/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nlmf_Location

import (
	"encoding/json"
	"fmt"
)

// GnssId Global Navigation Satellite System (GNSS) ID.
type GnssId struct {
	GnssIdAnyOf *GnssIdAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GnssId) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into GnssIdAnyOf
	err = json.Unmarshal(data, &dst.GnssIdAnyOf);
	if err == nil {
		jsonGnssIdAnyOf, _ := json.Marshal(dst.GnssIdAnyOf)
		if string(jsonGnssIdAnyOf) == "{}" { // empty struct
			dst.GnssIdAnyOf = nil
		} else {
			return nil // data stored in dst.GnssIdAnyOf, return on the first match
		}
	} else {
		dst.GnssIdAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(GnssId)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GnssId) MarshalJSON() ([]byte, error) {
	if src.GnssIdAnyOf != nil {
		return json.Marshal(&src.GnssIdAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGnssId struct {
	value *GnssId
	isSet bool
}

func (v NullableGnssId) Get() *GnssId {
	return v.value
}

func (v *NullableGnssId) Set(val *GnssId) {
	v.value = val
	v.isSet = true
}

func (v NullableGnssId) IsSet() bool {
	return v.isSet
}

func (v *NullableGnssId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGnssId(val *GnssId) *NullableGnssId {
	return &NullableGnssId{value: val, isSet: true}
}

func (v NullableGnssId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGnssId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


