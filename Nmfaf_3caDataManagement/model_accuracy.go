/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmfaf_3caDataManagement

import (
	"encoding/json"
	"fmt"
)

// Accuracy Possible values are: - LOW: Low accuracy.   - HIGH: High accuracy. 
type Accuracy struct {
	AccuracyAnyOf *AccuracyAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Accuracy) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AccuracyAnyOf
	err = json.Unmarshal(data, &dst.AccuracyAnyOf);
	if err == nil {
		jsonAccuracyAnyOf, _ := json.Marshal(dst.AccuracyAnyOf)
		if string(jsonAccuracyAnyOf) == "{}" { // empty struct
			dst.AccuracyAnyOf = nil
		} else {
			return nil // data stored in dst.AccuracyAnyOf, return on the first match
		}
	} else {
		dst.AccuracyAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(Accuracy)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Accuracy) MarshalJSON() ([]byte, error) {
	if src.AccuracyAnyOf != nil {
		return json.Marshal(&src.AccuracyAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAccuracy struct {
	value *Accuracy
	isSet bool
}

func (v NullableAccuracy) Get() *Accuracy {
	return v.value
}

func (v *NullableAccuracy) Set(val *Accuracy) {
	v.value = val
	v.isSet = true
}

func (v NullableAccuracy) IsSet() bool {
	return v.isSet
}

func (v *NullableAccuracy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccuracy(val *Accuracy) *NullableAccuracy {
	return &NullableAccuracy{value: val, isSet: true}
}

func (v NullableAccuracy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccuracy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


