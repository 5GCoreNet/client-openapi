/*
5GMS Data Reporting data types

5GMS Data Reporting data types © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_R4_DataReporting

import (
	"encoding/json"
	"fmt"
)

// CacheStatusAnyOf the model 'CacheStatusAnyOf'
type CacheStatusAnyOf string

// List of CacheStatus_anyOf
const (
	HIT CacheStatusAnyOf = "HIT"
	MISS CacheStatusAnyOf = "MISS"
	EXPIRED CacheStatusAnyOf = "EXPIRED"
)

// All allowed values of CacheStatusAnyOf enum
var AllowedCacheStatusAnyOfEnumValues = []CacheStatusAnyOf{
	"HIT",
	"MISS",
	"EXPIRED",
}

func (v *CacheStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CacheStatusAnyOf(value)
	for _, existing := range AllowedCacheStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CacheStatusAnyOf", value)
}

// NewCacheStatusAnyOfFromValue returns a pointer to a valid CacheStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCacheStatusAnyOfFromValue(v string) (*CacheStatusAnyOf, error) {
	ev := CacheStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CacheStatusAnyOf: valid values are %v", v, AllowedCacheStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CacheStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedCacheStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CacheStatus_anyOf value
func (v CacheStatusAnyOf) Ptr() *CacheStatusAnyOf {
	return &v
}

type NullableCacheStatusAnyOf struct {
	value *CacheStatusAnyOf
	isSet bool
}

func (v NullableCacheStatusAnyOf) Get() *CacheStatusAnyOf {
	return v.value
}

func (v *NullableCacheStatusAnyOf) Set(val *CacheStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCacheStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCacheStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCacheStatusAnyOf(val *CacheStatusAnyOf) *NullableCacheStatusAnyOf {
	return &NullableCacheStatusAnyOf{value: val, isSet: true}
}

func (v NullableCacheStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCacheStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

