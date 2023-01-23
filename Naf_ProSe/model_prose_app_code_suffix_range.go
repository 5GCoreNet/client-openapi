/*
Naf_ProSe API

Naf_ProSe Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naf_ProSe

import (
	"encoding/json"
)

// ProseAppCodeSuffixRange Contains a range of the Prose Application Code Suffixes which are consecutive.
type ProseAppCodeSuffixRange struct {
	// Contains the ProSe Application Code Suffix.
	BeginningSuffix string `json:"beginningSuffix"`
	// Contains the ProSe Application Code Suffix.
	EndingSuffix string `json:"endingSuffix"`
}

// NewProseAppCodeSuffixRange instantiates a new ProseAppCodeSuffixRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProseAppCodeSuffixRange(beginningSuffix string, endingSuffix string) *ProseAppCodeSuffixRange {
	this := ProseAppCodeSuffixRange{}
	this.BeginningSuffix = beginningSuffix
	this.EndingSuffix = endingSuffix
	return &this
}

// NewProseAppCodeSuffixRangeWithDefaults instantiates a new ProseAppCodeSuffixRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProseAppCodeSuffixRangeWithDefaults() *ProseAppCodeSuffixRange {
	this := ProseAppCodeSuffixRange{}
	return &this
}

// GetBeginningSuffix returns the BeginningSuffix field value
func (o *ProseAppCodeSuffixRange) GetBeginningSuffix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BeginningSuffix
}

// GetBeginningSuffixOk returns a tuple with the BeginningSuffix field value
// and a boolean to check if the value has been set.
func (o *ProseAppCodeSuffixRange) GetBeginningSuffixOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BeginningSuffix, true
}

// SetBeginningSuffix sets field value
func (o *ProseAppCodeSuffixRange) SetBeginningSuffix(v string) {
	o.BeginningSuffix = v
}

// GetEndingSuffix returns the EndingSuffix field value
func (o *ProseAppCodeSuffixRange) GetEndingSuffix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndingSuffix
}

// GetEndingSuffixOk returns a tuple with the EndingSuffix field value
// and a boolean to check if the value has been set.
func (o *ProseAppCodeSuffixRange) GetEndingSuffixOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EndingSuffix, true
}

// SetEndingSuffix sets field value
func (o *ProseAppCodeSuffixRange) SetEndingSuffix(v string) {
	o.EndingSuffix = v
}

func (o ProseAppCodeSuffixRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["beginningSuffix"] = o.BeginningSuffix
	}
	if true {
		toSerialize["endingSuffix"] = o.EndingSuffix
	}
	return json.Marshal(toSerialize)
}

type NullableProseAppCodeSuffixRange struct {
	value *ProseAppCodeSuffixRange
	isSet bool
}

func (v NullableProseAppCodeSuffixRange) Get() *ProseAppCodeSuffixRange {
	return v.value
}

func (v *NullableProseAppCodeSuffixRange) Set(val *ProseAppCodeSuffixRange) {
	v.value = val
	v.isSet = true
}

func (v NullableProseAppCodeSuffixRange) IsSet() bool {
	return v.isSet
}

func (v *NullableProseAppCodeSuffixRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProseAppCodeSuffixRange(val *ProseAppCodeSuffixRange) *NullableProseAppCodeSuffixRange {
	return &NullableProseAppCodeSuffixRange{value: val, isSet: true}
}

func (v NullableProseAppCodeSuffixRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProseAppCodeSuffixRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


