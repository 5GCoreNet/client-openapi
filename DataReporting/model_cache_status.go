/*
3gpp-data-reporting

API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_DataReporting

import (
	"encoding/json"
	"fmt"
)

// CacheStatus struct for CacheStatus
type CacheStatus struct {
	CacheStatusAnyOf *CacheStatusAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CacheStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CacheStatusAnyOf
	err = json.Unmarshal(data, &dst.CacheStatusAnyOf);
	if err == nil {
		jsonCacheStatusAnyOf, _ := json.Marshal(dst.CacheStatusAnyOf)
		if string(jsonCacheStatusAnyOf) == "{}" { // empty struct
			dst.CacheStatusAnyOf = nil
		} else {
			return nil // data stored in dst.CacheStatusAnyOf, return on the first match
		}
	} else {
		dst.CacheStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(CacheStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CacheStatus) MarshalJSON() ([]byte, error) {
	if src.CacheStatusAnyOf != nil {
		return json.Marshal(&src.CacheStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCacheStatus struct {
	value *CacheStatus
	isSet bool
}

func (v NullableCacheStatus) Get() *CacheStatus {
	return v.value
}

func (v *NullableCacheStatus) Set(val *CacheStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCacheStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCacheStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCacheStatus(val *CacheStatus) *NullableCacheStatus {
	return &NullableCacheStatus{value: val, isSet: true}
}

func (v NullableCacheStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCacheStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


