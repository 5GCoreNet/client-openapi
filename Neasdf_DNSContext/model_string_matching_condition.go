/*
Neasdf_DNSContext

EASDF DNS Context Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Neasdf_DNSContext

import (
	"encoding/json"
)

// StringMatchingCondition A String with Matching Operator
type StringMatchingCondition struct {
	MatchingString *string `json:"matchingString,omitempty"`
	MatchingOperator MatchingOperator `json:"matchingOperator"`
}

// NewStringMatchingCondition instantiates a new StringMatchingCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringMatchingCondition(matchingOperator MatchingOperator) *StringMatchingCondition {
	this := StringMatchingCondition{}
	this.MatchingOperator = matchingOperator
	return &this
}

// NewStringMatchingConditionWithDefaults instantiates a new StringMatchingCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringMatchingConditionWithDefaults() *StringMatchingCondition {
	this := StringMatchingCondition{}
	return &this
}

// GetMatchingString returns the MatchingString field value if set, zero value otherwise.
func (o *StringMatchingCondition) GetMatchingString() string {
	if o == nil || isNil(o.MatchingString) {
		var ret string
		return ret
	}
	return *o.MatchingString
}

// GetMatchingStringOk returns a tuple with the MatchingString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMatchingCondition) GetMatchingStringOk() (*string, bool) {
	if o == nil || isNil(o.MatchingString) {
    return nil, false
	}
	return o.MatchingString, true
}

// HasMatchingString returns a boolean if a field has been set.
func (o *StringMatchingCondition) HasMatchingString() bool {
	if o != nil && !isNil(o.MatchingString) {
		return true
	}

	return false
}

// SetMatchingString gets a reference to the given string and assigns it to the MatchingString field.
func (o *StringMatchingCondition) SetMatchingString(v string) {
	o.MatchingString = &v
}

// GetMatchingOperator returns the MatchingOperator field value
func (o *StringMatchingCondition) GetMatchingOperator() MatchingOperator {
	if o == nil {
		var ret MatchingOperator
		return ret
	}

	return o.MatchingOperator
}

// GetMatchingOperatorOk returns a tuple with the MatchingOperator field value
// and a boolean to check if the value has been set.
func (o *StringMatchingCondition) GetMatchingOperatorOk() (*MatchingOperator, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MatchingOperator, true
}

// SetMatchingOperator sets field value
func (o *StringMatchingCondition) SetMatchingOperator(v MatchingOperator) {
	o.MatchingOperator = v
}

func (o StringMatchingCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MatchingString) {
		toSerialize["matchingString"] = o.MatchingString
	}
	if true {
		toSerialize["matchingOperator"] = o.MatchingOperator
	}
	return json.Marshal(toSerialize)
}

type NullableStringMatchingCondition struct {
	value *StringMatchingCondition
	isSet bool
}

func (v NullableStringMatchingCondition) Get() *StringMatchingCondition {
	return v.value
}

func (v *NullableStringMatchingCondition) Set(val *StringMatchingCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableStringMatchingCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableStringMatchingCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringMatchingCondition(val *StringMatchingCondition) *NullableStringMatchingCondition {
	return &NullableStringMatchingCondition{value: val, isSet: true}
}

func (v NullableStringMatchingCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringMatchingCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


