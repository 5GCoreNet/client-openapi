/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// SharedDataIdRange A range of SharedDataIds based on regular-expression matching
type SharedDataIdRange struct {
	Pattern *string `json:"pattern,omitempty"`
}

// NewSharedDataIdRange instantiates a new SharedDataIdRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedDataIdRange() *SharedDataIdRange {
	this := SharedDataIdRange{}
	return &this
}

// NewSharedDataIdRangeWithDefaults instantiates a new SharedDataIdRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedDataIdRangeWithDefaults() *SharedDataIdRange {
	this := SharedDataIdRange{}
	return &this
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *SharedDataIdRange) GetPattern() string {
	if o == nil || isNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDataIdRange) GetPatternOk() (*string, bool) {
	if o == nil || isNil(o.Pattern) {
    return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *SharedDataIdRange) HasPattern() bool {
	if o != nil && !isNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *SharedDataIdRange) SetPattern(v string) {
	o.Pattern = &v
}

func (o SharedDataIdRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Pattern) {
		toSerialize["pattern"] = o.Pattern
	}
	return json.Marshal(toSerialize)
}

type NullableSharedDataIdRange struct {
	value *SharedDataIdRange
	isSet bool
}

func (v NullableSharedDataIdRange) Get() *SharedDataIdRange {
	return v.value
}

func (v *NullableSharedDataIdRange) Set(val *SharedDataIdRange) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedDataIdRange) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedDataIdRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedDataIdRange(val *SharedDataIdRange) *NullableSharedDataIdRange {
	return &NullableSharedDataIdRange{value: val, isSet: true}
}

func (v NullableSharedDataIdRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedDataIdRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


