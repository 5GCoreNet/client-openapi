/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// ClassCriterion Indicates the dispersion class criterion for fixed, camper and/or traveller UE, and/or the  top-heavy UE dispersion class criterion. 
type ClassCriterion struct {
	DisperClass DispersionClass `json:"disperClass"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	ClassThreshold int32 `json:"classThreshold"`
	ThresMatch MatchingDirection `json:"thresMatch"`
}

// NewClassCriterion instantiates a new ClassCriterion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassCriterion(disperClass DispersionClass, classThreshold int32, thresMatch MatchingDirection) *ClassCriterion {
	this := ClassCriterion{}
	this.DisperClass = disperClass
	this.ClassThreshold = classThreshold
	this.ThresMatch = thresMatch
	return &this
}

// NewClassCriterionWithDefaults instantiates a new ClassCriterion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassCriterionWithDefaults() *ClassCriterion {
	this := ClassCriterion{}
	return &this
}

// GetDisperClass returns the DisperClass field value
func (o *ClassCriterion) GetDisperClass() DispersionClass {
	if o == nil {
		var ret DispersionClass
		return ret
	}

	return o.DisperClass
}

// GetDisperClassOk returns a tuple with the DisperClass field value
// and a boolean to check if the value has been set.
func (o *ClassCriterion) GetDisperClassOk() (*DispersionClass, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DisperClass, true
}

// SetDisperClass sets field value
func (o *ClassCriterion) SetDisperClass(v DispersionClass) {
	o.DisperClass = v
}

// GetClassThreshold returns the ClassThreshold field value
func (o *ClassCriterion) GetClassThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClassThreshold
}

// GetClassThresholdOk returns a tuple with the ClassThreshold field value
// and a boolean to check if the value has been set.
func (o *ClassCriterion) GetClassThresholdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ClassThreshold, true
}

// SetClassThreshold sets field value
func (o *ClassCriterion) SetClassThreshold(v int32) {
	o.ClassThreshold = v
}

// GetThresMatch returns the ThresMatch field value
func (o *ClassCriterion) GetThresMatch() MatchingDirection {
	if o == nil {
		var ret MatchingDirection
		return ret
	}

	return o.ThresMatch
}

// GetThresMatchOk returns a tuple with the ThresMatch field value
// and a boolean to check if the value has been set.
func (o *ClassCriterion) GetThresMatchOk() (*MatchingDirection, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ThresMatch, true
}

// SetThresMatch sets field value
func (o *ClassCriterion) SetThresMatch(v MatchingDirection) {
	o.ThresMatch = v
}

func (o ClassCriterion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["disperClass"] = o.DisperClass
	}
	if true {
		toSerialize["classThreshold"] = o.ClassThreshold
	}
	if true {
		toSerialize["thresMatch"] = o.ThresMatch
	}
	return json.Marshal(toSerialize)
}

type NullableClassCriterion struct {
	value *ClassCriterion
	isSet bool
}

func (v NullableClassCriterion) Get() *ClassCriterion {
	return v.value
}

func (v *NullableClassCriterion) Set(val *ClassCriterion) {
	v.value = val
	v.isSet = true
}

func (v NullableClassCriterion) IsSet() bool {
	return v.isSet
}

func (v *NullableClassCriterion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassCriterion(val *ClassCriterion) *NullableClassCriterion {
	return &NullableClassCriterion{value: val, isSet: true}
}

func (v NullableClassCriterion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassCriterion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


