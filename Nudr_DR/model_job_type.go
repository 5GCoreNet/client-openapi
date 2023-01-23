/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// JobType The enumeration JobType defines Job Type in the trace. See 3GPP TS 32.422 for further  description of the values. It shall comply with the provisions defined in table 5.6.3.3-1. 
type JobType struct {
	JobTypeAnyOf *JobTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *JobType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into JobTypeAnyOf
	err = json.Unmarshal(data, &dst.JobTypeAnyOf);
	if err == nil {
		jsonJobTypeAnyOf, _ := json.Marshal(dst.JobTypeAnyOf)
		if string(jsonJobTypeAnyOf) == "{}" { // empty struct
			dst.JobTypeAnyOf = nil
		} else {
			return nil // data stored in dst.JobTypeAnyOf, return on the first match
		}
	} else {
		dst.JobTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(JobType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *JobType) MarshalJSON() ([]byte, error) {
	if src.JobTypeAnyOf != nil {
		return json.Marshal(&src.JobTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableJobType struct {
	value *JobType
	isSet bool
}

func (v NullableJobType) Get() *JobType {
	return v.value
}

func (v *NullableJobType) Set(val *JobType) {
	v.value = val
	v.isSet = true
}

func (v NullableJobType) IsSet() bool {
	return v.isSet
}

func (v *NullableJobType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobType(val *JobType) *NullableJobType {
	return &NullableJobType{value: val, isSet: true}
}

func (v NullableJobType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


