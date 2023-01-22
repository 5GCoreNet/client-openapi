/*
3gpp-data-reporting

API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package DataReporting

import (
	"encoding/json"
	"fmt"
)

// Usage Indicates usage made of the location measurement.
type Usage struct {
	UsageAnyOf *UsageAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Usage) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UsageAnyOf
	err = json.Unmarshal(data, &dst.UsageAnyOf);
	if err == nil {
		jsonUsageAnyOf, _ := json.Marshal(dst.UsageAnyOf)
		if string(jsonUsageAnyOf) == "{}" { // empty struct
			dst.UsageAnyOf = nil
		} else {
			return nil // data stored in dst.UsageAnyOf, return on the first match
		}
	} else {
		dst.UsageAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(Usage)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Usage) MarshalJSON() ([]byte, error) {
	if src.UsageAnyOf != nil {
		return json.Marshal(&src.UsageAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUsage struct {
	value *Usage
	isSet bool
}

func (v NullableUsage) Get() *Usage {
	return v.value
}

func (v *NullableUsage) Set(val *Usage) {
	v.value = val
	v.isSet = true
}

func (v NullableUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsage(val *Usage) *NullableUsage {
	return &NullableUsage{value: val, isSet: true}
}

func (v NullableUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


